package todo

import (
	"errors"

	"github.com/go-spirit/spirit/component"
	"github.com/go-spirit/spirit/mail"
	"github.com/go-spirit/spirit/worker"
	"github.com/pborman/uuid"
)

type Task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TaskId struct {
	Id string `json:"id"`
}

type Todo struct {
	tasks map[string]Task
}

func init() {
	component.RegisterComponent("examples-todo", NewTodo)
}

func NewTodo(opts ...component.Option) (comp component.Component, err error) {

	return &Todo{
		tasks: make(map[string]Task),
	}, nil
}

func (p *Todo) Start() error {

	return nil
}

func (p *Todo) Stop() error {

	return nil
}

func (p *Todo) NewTask(session mail.Session) (err error) {

	task := Task{}
	err = session.Payload().Content().ToObject(&task)
	if err != nil {
		return
	}

	id := uuid.New()

	task.Id = id

	err = session.Payload().Content().SetBody(&TaskId{Id: id})
	if err != nil {
		return
	}

	p.tasks[task.Id] = task

	return
}

func (p *Todo) GetTask(session mail.Session) (err error) {

	taskId := TaskId{}
	err = session.Payload().Content().ToObject(&taskId)
	if err != nil {
		return
	}

	task, exist := p.tasks[taskId.Id]
	if !exist {
		err = errors.New("task not exist")
		return
	}

	err = session.Payload().Content().SetBody(task)
	if err != nil {
		return
	}

	return
}

func (p *Todo) Route(session mail.Session) worker.HandlerFunc {

	switch session.Query("action") {
	case "new":
		{
			return p.NewTask
		}
	case "get":
		{
			return p.GetTask
		}
	}

	return nil
}
