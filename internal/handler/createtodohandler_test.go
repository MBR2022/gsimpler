package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MBR2022/gosimpler/internal/handler"
	"github.com/MBR2022/gosimpler/internal/model"
	"github.com/MBR2022/gosimpler/internal/svc"
	"github.com/MBR2022/gosimpler/internal/types"
	"github.com/MBR2022/gosimpler/mock"
	"github.com/golang/mock/gomock"
)

var (
	todoId = "generate-id-dummy"
)

func TestShoudCreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock.NewMockMemStore(ctrl)
	req := types.CreateTodoRequest{
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
	w := httptest.NewRecorder()
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRequest("POST", "/", body)
	r.Header.Set("Content-Type", "application/json")
	handler.CreateTodoHandler(&svc.ServiceContext{MemStore: mockStore})(w, r)
	if got, want := w.Code, http.StatusOK; got != want {
		t.Fatalf("Want status cod %d, but got: %d", want, got)
	}
	out := new(model.Todo)
	err = json.NewDecoder(w.Body).Decode(out)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := out.ID, todoId; got != want {
		t.Errorf("Want id: %s, but got: %s", want, got)
	}
}
