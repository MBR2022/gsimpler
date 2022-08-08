package store

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/MBR2022/gosimpler/internal/model"
)

type (
	memStore struct {
		Todos map[string]*model.Todo
		*sync.RWMutex
	}

	MemStore interface {
		Add(t *model.Todo) error
		Delete(id string) error
		Get(id string) (*model.Todo, error)
		Update(t *model.Todo) error
	}
)

var (
	ErrorTodoExist    = errors.New("todo is exist")
	ErrorTodoNotfound = errors.New("todo is not found")
)

func NewMemStore() MemStore {
	return &memStore{
		Todos:   make(map[string]*model.Todo),
		RWMutex: &sync.RWMutex{},
	}
}

func (svc *memStore) Add(t *model.Todo) error {
	s := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(s)
	t.ID = fmt.Sprintf("%d", rd.Int63())
	if todo, _ := svc.Get(t.ID); todo != nil {
		return ErrorTodoExist
	}
	svc.Lock()
	svc.Todos[t.ID] = t
	svc.Unlock()
	return nil
}
func (svc *memStore) Delete(id string) error {
	todo, err := svc.Get(id)
	if err != nil {
		return err
	}
	svc.Lock()
	delete(svc.Todos, todo.ID)
	return nil
}
func (svc *memStore) Get(id string) (*model.Todo, error) {
	todo, ok := svc.Todos[id]
	if !ok {
		return nil, ErrorTodoNotfound
	}
	return todo, nil
}
func (svc *memStore) Update(t *model.Todo) error {
	_, err := svc.Get(t.ID)
	if err != nil {
		return err
	}
	svc.Lock()
	svc.Todos[t.ID] = t
	svc.Unlock()
	return nil
}
