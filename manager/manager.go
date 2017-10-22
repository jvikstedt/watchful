package manager

import (
	"log"
	"time"

	"github.com/jvikstedt/watchful/model"
)

type Executor interface {
	Identifier() string
	Instruction() Instruction
	Execute(map[string]interface{}) (map[string]interface{}, error)
}

type Handler interface {
}

type Service struct {
	log       *log.Logger
	model     *model.Service
	executors map[string]Executor
	close     chan bool
}

func NewService(log *log.Logger, model *model.Service) *Service {
	return &Service{
		log:       log,
		model:     model,
		executors: make(map[string]Executor),
		close:     make(chan bool),
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

func (s *Service) Run() error {
	for {
		select {
		case <-s.close:
			s.log.Println("manager closed")
			return nil
		default:
			time.Sleep(time.Second * 2)
		}
	}
}
