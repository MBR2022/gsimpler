package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MBR2022/gosimpler/internal/handler"
	"github.com/MBR2022/gosimpler/internal/model"
	"github.com/MBR2022/gosimpler/internal/store"
	"github.com/MBR2022/gosimpler/internal/svc"
	"github.com/MBR2022/gosimpler/internal/types"
	"github.com/MBR2022/gosimpler/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/rest/pathvar"
)

func TestShoudGetTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock.NewMockMemStore(ctrl)
	req := types.GetTodoRequest{
		ID: todoId,
	}
	todo := &model.Todo{
		ID:          todoId,
		Name:        "Name",
		Description: "Description",
		Status:      "Status",
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
			name:           "Test should handle success",
			req:            req,
			mockCall:       mockStore.EXPECT().Get(todoId).Return(todo, nil),
			wantData:       todo,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "Test should handle input invalid",
			req:            nil,
			mockCall:       nil,
			wantData:       nil,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "Test should handle store add error",
			req:            req,
			mockCall:       mockStore.EXPECT().Get(todoId).Return(nil, store.ErrorTodoNotfound),
			wantData:       nil,
			wantStatusCode: http.StatusBadRequest,
		},
	}

	t.Parallel()
	for i := range tt {
		tc := tt[i]
		t.Run(tc.name, HanlderGetTodo(svc, tc.req, tc.mockCall, tc.wantData, tc.wantStatusCode))
	}
}
func HanlderGetTodo(svct *svc.ServiceContext, req interface{}, mockCall *gomock.Call, wantData interface{}, wantStatus int) func(t *testing.T) {
	return func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		if rq, ok := req.(types.GetTodoRequest); ok {
			r = pathvar.WithVars(r, map[string]string{"id": rq.ID})
		}
		handler.GetTodoHandler(svct)(w, r)
		assert.Equal(t, wantStatus, w.Code)
		out := new(model.Todo)
		err := json.NewDecoder(w.Body).Decode(out)
		if wantData == nil {
			assert.Error(t, err)
			return
		}
		assert.Equal(t, wantData, out)
	}
}
