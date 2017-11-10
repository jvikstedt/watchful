package exec

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jvikstedt/watchful"
	"github.com/jvikstedt/watchful/pkg/model"
	uuid "github.com/satori/go.uuid"
)

type work struct {
	id        string
	job       *model.Job
	isTestRun bool
}

type Service struct {
	log         *log.Logger
	db          *sqlx.DB
	executables map[string]watchful.Executable
	close       chan bool
	workCh      chan work
}

func New(log *log.Logger, db *sqlx.DB) *Service {
	return &Service{
		log:         log,
		db:          db,
		executables: make(map[string]watchful.Executable),
		close:       make(chan bool),
		workCh:      make(chan work, 10),
	}
}

func (s *Service) RegisterExecutable(e watchful.Executable) {
	s.executables[e.Identifier()] = e
}

func (s *Service) Executables() map[string]watchful.Executable {
	return s.executables
}

// Shutdown blocks until it's done working
func (s *Service) Shutdown() {
	s.close <- true
}

func (s *Service) AddJob(job *model.Job, isTestRun bool) string {
	u1 := uuid.NewV1()
	s.workCh <- work{
		id:        u1.String(),
		job:       job,
		isTestRun: isTestRun,
	}
	return u1.String()
}

func (s *Service) Run() error {
	for {
		select {
		case sj := <-s.workCh:
			result := model.Result{
				UUID:    sj.id,
				TestRun: sj.isTestRun,
				JobID:   sj.job.ID,
				Status:  model.ResultStatusWaiting,
			}
			err := result.Create(s.db)
			if err != nil {
				s.log.Println(err)
			}
			go s.executeJob(result)
		case <-s.close:
			return nil
		}
	}
}

func (s *Service) executeJob(result model.Result) {
	job := model.Job{}
	err := model.JobGetOne(s.db, result.JobID, &job)
	if err != nil {
		s.log.Println(err)
		return
	}
	tasks, err := model.TasksWithInputsByJobID(s.db, result.JobID)
	if err != nil {
		s.log.Println(err)
		return
	}

	for _, t := range tasks {
		var status model.ResultStatus = model.ResultStatusSuccess
		var errMsg string
		output, err := s.handleTask(result, t)
		if err != nil {
			status = model.ResultStatusError
			errMsg = err.Error()
		}

		resultItem := &model.ResultItem{
			ResultID: result.ID,
			TaskID:   t.ID,
			Output:   output,
			Error:    errMsg,
			Status:   status,
		}

		result.Status = status
		err = resultItem.Create(s.db)
		if err != nil {
			result.Status = model.ResultStatusError
		}
	}

	result.Update(s.db)
}

func (s *Service) handleTask(result model.Result, task *model.Task) (string, error) {
	executable, ok := s.executables[task.Executable]
	if !ok {
		return "", fmt.Errorf("Could not find executable: %s", task.Executable)
	}

	resultItems, err := model.ResultItemAllByResultID(s.db, result.ID)
	if err != nil {
		return "", err
	}

	commands := map[string]interface{}{}
	for _, i := range task.Inputs {
		if i.Dynamic {
			for _, r := range resultItems {
				if r.TaskID == *i.SourceTaskID {
					output := map[string]interface{}{}

					d := json.NewDecoder(strings.NewReader(r.Output))
					d.UseNumber()
					if err = d.Decode(&output); err != nil {
						return "", err
					}

					val, err := s.handleDynamicInput(output[i.SourceName], i.Type)
					if err != nil {
						return "", err
					}
					commands[i.Name] = val
				}
			}
		} else {
			val, err := s.handleStaticInput(i.Value, i.Type)
			if err != nil {
				return "", err
			}
			commands[i.Name] = val
		}
	}

	output, err := executable.Execute(commands)
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (s *Service) handleDynamicInput(i interface{}, desiredType watchful.ParamType) (interface{}, error) {
	switch v := i.(type) {
	case json.Number:
		switch desiredType {
		case watchful.ParamString:
			return v.String(), nil
		case watchful.ParamInt:
			return v.Int64()
		case watchful.ParamFloat:
			return v.Float64()
		default:
			return nil, fmt.Errorf("desiredType was not any of the expected values %d", desiredType)
		}
	default:
		return i, nil
	}
}

func (s *Service) handleStaticInput(str string, desiredType watchful.ParamType) (interface{}, error) {
	switch desiredType {
	case watchful.ParamString:
		return str, nil
	case watchful.ParamInt:
		return strconv.ParseInt(str, 0, 64)
	case watchful.ParamFloat:
		return strconv.ParseFloat(str, 64)
	default:
		return nil, fmt.Errorf("desiredType was not any of the expected values %d", desiredType)
	}
}
