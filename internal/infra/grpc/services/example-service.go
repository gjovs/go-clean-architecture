package services

import (
	"context"

	"github.com/gjovs/clean/internal/infra/grpc/pb"
	usecase "github.com/gjovs/clean/internal/usecase/example"
)

type ExampleService struct {
	CreateExampleUc usecase.CreateExampleUseCase
	pb.UnimplementedExampleServiceServer
}

func NewExampleService(createExampleUc usecase.CreateExampleUseCase) *ExampleService {
	return &ExampleService{
		CreateExampleUc: createExampleUc,
	}
}

func (service *ExampleService) CreateExample(ctx *context.Context, in *pb.CreateExampleRequest) (*pb.CreateExampleResponse, error) {
	dto := usecase.ExampleInputDTO{
		ID:          in.Id,
		Name:        in.Name,
		Description: in.Description,
	}

	output, err := service.CreateExampleUc.Execute(dto)

	if err != nil {
		return nil, err
	}

	return &pb.CreateExampleResponse{
		Id:          output.ID,
		Name:        output.Name,
		Description: output.Description,
		CreatedAt:   output.CreatedAt.String(),
	}, nil
}
