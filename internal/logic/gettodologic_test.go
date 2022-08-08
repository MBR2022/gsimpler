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

func Test_Logic_Get_Todo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := mock.NewMockMemStore(ctrl)
	req := &types.GetTodoRequest{ID: todoId}
	todo := &model.Todo{
		ID:          todoId,
		Name:        "Name",
		Description: "Description",
		Status:      "Status",
	}
	mockStore.EXPECT().Get(todoId).DoAndReturn(func(id string) (*model.Todo, error) {
		return todo, nil
	})
	l := logic.NewGetTodoLogic(context.TODO(), &svc.ServiceContext{
		MemStore: mockStore,
	})
	t.Run("Test_Should_Create_Todo_Success", GetTodoSuccess(l, req))

	mockStore.EXPECT().Get(gomock.Any()).DoAndReturn(func(id string) (*model.Todo, error) {
		return nil, store.ErrorTodoNotfound
	})
	t.Run("Test_Shoud_Create_Todo_Failed", GetTodoFailed(l, req))
}

func GetTodoSuccess(l *logic.GetTodoLogic, req *types.GetTodoRequest) func(t *testing.T) {
	return func(t *testing.T) {
		resp, err := l.GetTodo(req)
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

func GetTodoFailed(l *logic.GetTodoLogic, req *types.GetTodoRequest) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := l.GetTodo(req)
		if err == nil {
			t.Error("Want error is not nil")
		}
		if want, got := store.ErrorTodoNotfound, err; want != got {
			t.Errorf("Want Error: %s, but got %s", want, got)
		}
	}
}
