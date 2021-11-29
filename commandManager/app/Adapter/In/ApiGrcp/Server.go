package ApiGrcp

import (
	"context"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
)

type TaskServer struct {
	task.UnimplementedTaskServiceServer
}

func (server TaskServer) CreateTask(ctx context.Context, request *task.CreateTaskRequest) (*task.CreateTaskResponse, error) {
	panic("implement me")
}

func (server TaskServer) ReadTask(ctx context.Context, request *task.ReadTaskRequest) (*task.ReadTaskResponse, error) {
	panic("implement me")
}

func (server TaskServer) UpdateTask(ctx context.Context, request *task.UpdateTaskRequest) (*task.UpdateTaskResponse, error) {
	panic("implement me")
}

func (server TaskServer) DeleteTask(ctx context.Context, request *task.DeleteTaskRequest) (*task.DeleteTaskResponse, error) {
	panic("implement me")
}

func (server TaskServer) ListTask(request *task.ListTaskRequest, server2 task.TaskService_ListTaskServer) error {
	panic("implement me")
}
