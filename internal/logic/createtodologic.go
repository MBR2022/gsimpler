package logic

import (
	"context"

	"github.com/MBR2022/gosimpler/internal/model"
	"github.com/MBR2022/gosimpler/internal/svc"
	"github.com/MBR2022/gosimpler/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTodoLogic {
	return &CreateTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTodoLogic) CreateTodo(req *types.CreateTodoRequest) (resp *types.TodoResponse, err error) {
	todo := &model.Todo{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}
	err = l.svcCtx.MemStore.Add(todo)
	if err != nil {
		return
	}
	resp = &types.TodoResponse{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
	}
	return
}
