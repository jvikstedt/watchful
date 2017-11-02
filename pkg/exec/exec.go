package exec

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jvikstedt/watchful"
	"github.com/jvikstedt/watchful/pkg/model"
	uuid "github.com/satori/go.uuid"
)

type scheduledJob struct {
	id        string
	job       *model.Job
	isTestRun bool
}

type Service struct {
	log            *log.Logger
	model          *model.Service
	executables    map[string]watchful.Executable
	close          chan bool
	scheduledJobCh chan scheduledJob
}

func New(log *log.Logger, model *model.Service) *Service {
	return &Service{
		log:            log,
		model:          model,
		executables:    make(map[string]watchful.Executable),
		close:          make(chan bool),
		scheduledJobCh: make(chan scheduledJob, 10),
	}
}

func (s *Service) RegisterExecutable(e watchful.Executable) {
	s.executables[e.Identifier()] = e
}

func (s *Service) Executables() map[string]watchful.Executable {
	return s.executables
}

func (s *Service) Shutdown() error {
	s.close <- true
	return nil
}

func (s *Service) AddScheduledJob(job *model.Job, isTestRun bool) string {
	u1 := uuid.NewV1()
	s.scheduledJobCh <- scheduledJob{
		id:        u1.String(),
		job:       job,
		isTestRun: isTestRun,
	}
	return u1.String()
}

func (s *Service) Run() error {
	for {
		select {
		case sj := <-s.scheduledJobCh:
			result := model.Result{
				UUID:    sj.id,
				TestRun: sj.isTestRun,
				JobID:   sj.job.ID,
				Status:  model.ResultStatusWaiting,
			}
			err := s.model.DB().ResultCreate(&result)
			if err != nil {
				s.log.Println(err)
			}
			go s.executeJob(result)
		case <-s.close:
			s.log.Println("exec closed")
			return nil
		default:
			time.Sleep(time.Second * 2)
		}
	}
}

func (s *Service) executeJob(result model.Result) {
	job := model.Job{}
	err := s.model.DB().JobGetOne(result.JobID, &job)
	if err != nil {
		s.log.Println(err)
		return
	}
	tasks, err := s.model.TasksWithInputsByJobID(result.JobID)
	if err != nil {
		s.log.Println(err)
		return
	}

	for _, t := range tasks {
		err := s.handleTask(result, t)
		if err != nil {
			s.log.Println(err)
		}
	}

	result.Status = model.ResultStatusDone
	s.model.DB().ResultUpdate(&result)
}

func (s *Service) handleTask(result model.Result, task *model.Task) error {
	executable, ok := s.executables[task.Executable]
	if !ok {
		return fmt.Errorf("Could not find executable: %s", task.Executable)
	}

	commands := map[string]interface{}{}
	for _, i := range task.Inputs {
		commands[i.Name] = i.Value
	}

	output, err := executable.Execute(commands)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		return err
	}

	resultItem := model.ResultItem{
		ResultID: result.ID,
		TaskID:   task.ID,
		Output:   string(bytes),
	}

	return s.model.DB().ResultItemCreate(&resultItem)
}
