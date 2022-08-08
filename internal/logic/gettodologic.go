package logic

import (
	"context"

	"github.com/MBR2022/gosimpler/internal/svc"
	"github.com/MBR2022/gosimpler/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTodoLogic {
	return &GetTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTodoLogic) GetTodo(req *types.GetTodoRequest) (resp *types.TodoResponse, err error) {
	todo, err := l.svcCtx.MemStore.Get(req.ID)
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
