package exec

import (
	"encoding/json"
	"fmt"
	"log"

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

		bytes, err := json.Marshal(output)
		if err != nil {
			status = model.ResultStatusError
			errMsg = err.Error()
		}

		resultItem := &model.ResultItem{
			ResultID: result.ID,
			TaskID:   t.ID,
			Output:   string(bytes),
			Error:    errMsg,
			Status:   status,
		}
		if result.Status != model.ResultStatusError {
			result.Status = status
		}
		err = resultItem.Create(s.db)
		if err != nil {
			s.log.Println(err)
			result.Status = model.ResultStatusError
		}
	}

	result.Update(s.db)
}

func (s *Service) handleTask(result model.Result, task *model.Task) (map[string]watchful.InputValue, error) {
	executable, ok := s.executables[task.Executable]
	if !ok {
		return nil, fmt.Errorf("Could not find executable: %s", task.Executable)
	}

	commands := map[string]interface{}{}
	for _, i := range task.Inputs {
		val, err := s.handleInput(i.Value, result)
		if err != nil {
			return nil, err
		}
		commands[i.Name] = val
	}

	return executable.Execute(commands)
}

func (s *Service) handleInput(iv *watchful.InputValue, result model.Result) (interface{}, error) {
	switch iv.Type {
	case watchful.ParamDynamic:
		val, ok := iv.Val.(watchful.DynamicSource)
		if !ok {
			return nil, fmt.Errorf("Expected dynamic, but got %T", iv.Val)
		}

		return s.handleDynamic(val, result)
	default:
		return iv.Val, nil
	}
}

func (s *Service) handleDynamic(ds watchful.DynamicSource, result model.Result) (interface{}, error) {
	resultItems, err := model.ResultItemAllByResultID(s.db, result.ID)
	if err != nil {
		return nil, err
	}

	sourceIV := map[string]watchful.InputValue{}

	for _, r := range resultItems {
		if r.TaskID == ds.TaskID {
			err := json.Unmarshal([]byte(r.Output), &sourceIV)
			if err != nil {
				return nil, err
			}
			iv, ok := sourceIV[ds.OutputName]
			if !ok {
				break
			}
			return s.handleInput(&iv, result)
		}
	}

	return nil, fmt.Errorf("Could not find output with name %s from task %d", ds.OutputName, ds.TaskID)
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
