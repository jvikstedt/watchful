package manager

import (
	"log"
	"time"

	"github.com/jvikstedt/watchful/model"
	uuid "github.com/satori/go.uuid"
)

type Executor interface {
	Identifier() string
	Instruction() Instruction
	Execute(map[string]interface{}) (map[string]interface{}, error)
}

type Handler interface {
}

type scheduledJob struct {
	id        string
	job       *model.Job
	isTestRun bool
}

type Service struct {
	log            *log.Logger
	model          *model.Service
	executors      map[string]Executor
	close          chan bool
	scheduledJobCh chan scheduledJob
}

func NewService(log *log.Logger, model *model.Service) *Service {
	return &Service{
		log:            log,
		model:          model,
		executors:      make(map[string]Executor),
		close:          make(chan bool),
		scheduledJobCh: make(chan scheduledJob, 10),
	}
}

func (s *Service) RegisterExecutor(e Executor) {
	s.executors[e.Identifier()] = e
}

func (s *Service) Executors() map[string]Executor {
	return s.executors
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
			err := s.model.ResultCreate(&result)
			if err != nil {
				s.log.Println(err)
			}
			go s.executeJob(result)
		case <-s.close:
			s.log.Println("manager closed")
			return nil
		default:
			time.Sleep(time.Second * 2)
		}
	}
}

func (s *Service) executeJob(result model.Result) {
	// Working
	time.Sleep(time.Second * 5)

	result.Status = model.ResultStatusDone
	s.model.ResultUpdate(&result)
}
