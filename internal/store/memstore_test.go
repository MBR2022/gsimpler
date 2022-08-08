package store_test

import (
	"testing"

	"github.com/MBR2022/gosimpler/internal/model"
	"github.com/MBR2022/gosimpler/internal/store"
)

func TestMemStore(t *testing.T) {
	st := store.NewMemStore()
	todo := &model.Todo{
		Name:        "Todo test",
		Description: "Test test test",
		Status:      "Done",
	}
	t.Run("Test_Should_Add_Todo", Should_Add_Todo(st, todo))
	t.Run("Test_Should_Get_Todo", Should_Get_Todo(st, todo.ID))
	t.Run("Test_If_NotFound", If_Todo_Notfound(st, "test-notfound-id"))
	t.Run("Test_Should_Update_Todo", Should_Update_Todo(st, todo))
	t.Run("Test_Should_Delete_Todo", Should_Delete_Todo(st, todo.ID))

}

func If_Todo_Notfound(st store.MemStore, s string) func(t *testing.T) {
	return func(t *testing.T) {
		todo, err := st.Get(s)
		if want, got := store.ErrorTodoNotfound, err; want != got {
			t.Errorf("Want error: %v, but got: %v", want.Error(), got.Error())
		}
		if todo != nil {
			t.Error("Want todo is nil")
		}
	}
}

func Should_Add_Todo(svc store.MemStore, todo *model.Todo) func(t *testing.T) {
	return func(t *testing.T) {
		if err := svc.Add(todo); err != nil {
			t.Error(err)
		}
	}
}

func Should_Delete_Todo(svc store.MemStore, id string) func(t *testing.T) {
	return func(t *testing.T) {
		if err := svc.Delete(id); err != nil {
			t.Error(err)
		}
		t.Run("Test_If_NotFound", If_Todo_Notfound(svc, id))
	}
}

func Should_Get_Todo(svc store.MemStore, id string) func(t *testing.T) {
	return func(t *testing.T) {
		if _, err := svc.Get(id); err != nil {
			t.Error(err)
		}
	}
}

func Should_Update_Todo(svc store.MemStore, todo *model.Todo) func(t *testing.T) {
	return func(t *testing.T) {
		todo.Name = "Updated"
		if err := svc.Update(todo); err != nil {
			t.Error(err)
		}
		ntodo, err := svc.Get(todo.ID)
		if err != nil {
			t.Error(err)
		}

		if want, got := "Updated", ntodo.Name; got != want {
			t.Errorf("Want %s, but got %s ", want, got)
		}
	}
}
