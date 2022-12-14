// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/MBR2022/gosimpler/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.BasicAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/todo",
					Handler: CreateTodoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/todo/:id",
					Handler: GetTodoHandler(serverCtx),
				},
			}...,
		),
	)
}
