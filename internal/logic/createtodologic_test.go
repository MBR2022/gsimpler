package logic_test

import (
	"context"
	"testing"

	"github.com/MBR2022/gosimpler/internal/logic"
	"github.com/MBR2022/gosimpler/internal/model"
	"github.com/MBR2022/gosimpler/internal/store"
	"github.com/MBR2022/gosimpler/internal/svc"
	"github.com/MBR2022/gosimpler/internal/types"
	"github.com/MBR2022/gosimpler/mock"
	"github.com/golang/mock/gomock"
)

var (
	todoId = "generate-id-dummy"
)

func Test_Logic_Create_Todo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := mock.NewMockMemStore(ctrl)
	req := &types.CreateTodoRequest{
		Name:        "Name",
		Description: "Description",
		Status:      "Status",
	}
	todo := &model.Todo{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}
	mockStore.EXPECT().Add(todo).DoAndReturn(func(todo *model.Todo) error {
		todo.ID = todoId
		return nil
	})
	l := logic.NewCreateTodoLogic(context.TODO(), &svc.ServiceContext{
		MemStore: mockStore,
	})
	t.Run("Test_Should_Create_Todo_Success", CreateTodoSuccess(l, req))

	mockStore.EXPECT().Add(todo).DoAndReturn(func(todo *model.Todo) error {
		todo.ID = todoId
		return store.ErrorTodoExist
	})
	t.Run("Test_Shoud_Create_Todo_Failed", CreateTodoFailed(l, req))
}

func CreateTodoSuccess(l *logic.CreateTodoLogic, req *types.CreateTodoRequest) func(t *testing.T) {
	return func(t *testing.T) {
		resp, err := l.CreateTodo(req)
		if err != nil {
			t.Error(err)
		}
		if resp == nil {
			t.Error("Response is nil")
			return
		}
		if want, got := todoId, resp.ID; want != got {
			t.Errorf("Want todo ID: %s, but got: %s", want, got)
		}
	}
}

func CreateTodoFailed(l *logic.CreateTodoLogic, req *types.CreateTodoRequest) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := l.CreateTodo(req)
		if err == nil {
			t.Error("Want error is not nil")
		}
		if want, got := store.ErrorTodoExist, err; want != got {
			t.Errorf("Want Error: %s, but got %s", want, got)
		}
	}
}
