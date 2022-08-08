package handler

import (
	"net/http"

	"github.com/MBR2022/gosimpler/internal/logic"
	"github.com/MBR2022/gosimpler/internal/svc"
	"github.com/MBR2022/gosimpler/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateTodoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateTodoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateTodoLogic(r.Context(), svcCtx)
		resp, err := l.CreateTodo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
