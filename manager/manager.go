package manager

import (
	"log"
	"time"

	"github.com/jvikstedt/watchful/storage"
)

type Executor interface {
	Name() string
	Instructions() Instruction
	Execute(map[string]interface{}) (map[string]interface{}, error)
}

type Checker interface {
	Name() string
	Check(string, interface{}) error
}

type Handler interface {
}

type Service struct {
	log       *log.Logger
	storage   storage.Service
	executors map[string]Executor
	checkers  map[string]Checker
	close     chan bool
}

func NewService(log *log.Logger, storage storage.Service) *Service {
	return &Service{
		log:       log,
		storage:   storage,
		executors: make(map[string]Executor),
		checkers:  make(map[string]Checker),
		close:     make(chan bool),
	}
}

func (s *Service) RegisterExecutor(e Executor) {
	s.executors[e.Name()] = e
}

func (s *Service) Executors() map[string]Executor {
	return s.executors
}

func (s *Service) RegisterChecker(c Checker) {
	s.checkers[c.Name()] = c
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
