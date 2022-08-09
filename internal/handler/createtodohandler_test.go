package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/MBR2022/gosimpler/internal/handler"
	"github.com/MBR2022/gosimpler/internal/model"
	"github.com/MBR2022/gosimpler/internal/store"
	"github.com/MBR2022/gosimpler/internal/svc"
	"github.com/MBR2022/gosimpler/internal/types"
	"github.com/MBR2022/gosimpler/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
	svc := &svc.ServiceContext{MemStore: mockStore}
	tt := []struct {
		name           string
		req            interface{}
		mockCall       *gomock.Call
		wantData       interface{}
		wantStatusCode int
	}{
		{
			name: "Test should handle success",
			req:  req,
			mockCall: mockStore.EXPECT().Add(todo).DoAndReturn(func(todo *model.Todo) error {
				todo.ID = todoId
				return nil
			}),
			wantData: &model.Todo{
				ID:          todoId,
				Name:        req.Name,
				Description: req.Description,
				Status:      req.Status,
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "Test should handle input invalid",
			req:            map[string]interface{}{},
			mockCall:       nil,
			wantData:       nil,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "Test should handle store add error",
			req:  req,
			mockCall: mockStore.EXPECT().Add(todo).DoAndReturn(func(_ *model.Todo) error {
				return store.ErrorTodoExist
			}),
			wantData:       nil,
			wantStatusCode: http.StatusBadRequest,
		},
	}

	t.Parallel()
	for i := range tt {
		tc := tt[i]
		t.Run(tc.name, HanlderAddTodo(svc, tc.req, tc.mockCall, tc.wantData, tc.wantStatusCode))
	}
}
func HanlderAddTodo(svct *svc.ServiceContext, req interface{}, mockCall *gomock.Call, wantData interface{}, wantStatus int) func(t *testing.T) {
	return func(t *testing.T) {
		w := httptest.NewRecorder()
		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(req)
		if err != nil {
			t.Fatal(err)
		}
		r := httptest.NewRequest("POST", "/", body)
		r.Header.Set("Content-Type", "application/json")
		handler.CreateTodoHandler(svct)(w, r)
		assert.Equal(t, wantStatus, w.Code)
		out := new(model.Todo)
		err = json.NewDecoder(w.Body).Decode(out)
		if wantData == nil {
			assert.Error(t, err)
			return
		}
		if got, want := out, wantData; !reflect.DeepEqual(want, got) {
			t.Fatalf("Want data %v, but got %v", want, got)
		}
	}
}
