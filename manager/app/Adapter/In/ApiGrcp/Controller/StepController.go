package Controller

import (
	"context"
	"fmt"
	stepProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/step"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/CreateStep"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/DeleteStep"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/ListSteps"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/ReadStep"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Step/UpdateStep"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StepController struct {
	CreateStepUseCase CreateStep.UseCase
	ReadStepUseCase   ReadStep.UseCase
	UpdateStepUseCase UpdateStep.UseCase
	DeleteStepUseCase DeleteStep.UseCase
	ListStepsUseCase  ListSteps.UseCase
	stepProto.UnimplementedStepServiceServer
}

func (controller StepController) CreateStep(ctx context.Context, request *stepProto.CreateStepRequest) (*stepProto.CreateStepResponse, error) {
	protoStep := request.GetStepParams()
	var cmd CreateStep.Command
	cmd.TaskUuid = protoStep.GetTaskUuid()
	cmd.Sentence = protoStep.GetSentence()
	stepDomain, err := controller.CreateStepUseCase.Create(cmd)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	newStep := stepProto.Step{
		Uuid:     stepDomain.Uuid.String(),
		TaskUuid: stepDomain.TaskUuid.String(),
		Sentence: stepDomain.Sentence,
	}
	return &stepProto.CreateStepResponse{Step: &newStep}, nil
}

func (controller StepController) ReadStep(ctx context.Context, request *stepProto.ReadStepRequest) (*stepProto.ReadStepResponse, error) {
	var query = ReadStep.Query{Uuid: request.GetStepUuid()}
	step, err := controller.ReadStepUseCase.Read(query)
	if err != nil {
		return &stepProto.ReadStepResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}
	return &stepProto.ReadStepResponse{Step: &stepProto.Step{
		Uuid:     step.Uuid.String(),
		TaskUuid: step.TaskUuid.String(),
		Sentence: step.Sentence,
	}}, nil
}

func (controller StepController) UpdateStep(ctx context.Context, request *stepProto.UpdateStepRequest) (*stepProto.UpdateStepResponse, error) {
	cmd := UpdateStep.Command{}
	params := request.GetStepParams()
	cmd.Uuid = request.GetStepUuid()
	if params.GetTaskUuid() != nil {
		cmd.TaskUuid.Change = true
	}
	if params.GetSentence() != nil {
		cmd.Sentence.Change = true
	}
	cmd.TaskUuid.Value = params.GetTaskUuid().Value
	cmd.Sentence.Value = params.GetSentence().Value
	err := controller.UpdateStepUseCase.Update(cmd)
	return &stepProto.UpdateStepResponse{}, err
}

func (controller StepController) DeleteStep(ctx context.Context, request *stepProto.DeleteStepRequest) (*stepProto.DeleteStepResponse, error) {
	var cmd = DeleteStep.Command{Uuid: request.GetStepUuid()}
	err := controller.DeleteStepUseCase.Delete(cmd)
	if err != nil {
		return &stepProto.DeleteStepResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}

	return &stepProto.DeleteStepResponse{}, nil
}

func (controller StepController) ListSteps(ctx context.Context, request *stepProto.ListStepsRequest) (*stepProto.ListStepsResponse, error) {
	steps := controller.ListStepsUseCase.List(ListSteps.Query{})
	if steps == nil {
		return &stepProto.ListStepsResponse{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("error"),
		)
	}
	stepProtoArray := []*stepProto.Step{}
	for _, step := range steps {
		stepProtoArray = append(stepProtoArray, &stepProto.Step{
			Uuid:     step.Uuid.String(),
			TaskUuid: step.TaskUuid.String(),
			Sentence: step.Sentence,
		})
	}
	return &stepProto.ListStepsResponse{Steps: stepProtoArray}, nil
}
