package usecase

import (
	"time"

	"github.com/gjovs/clean/internal/entity"
)

type ExampleInputDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExampleOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateExampleUseCase struct {
	ExampleRepository entity.ExampleRepositoryInterface
}

func NewCreateExampleUseCase(exampleRepository entity.ExampleRepositoryInterface) *CreateExampleUseCase {
	return &CreateExampleUseCase{
		ExampleRepository: exampleRepository,
	}
}

func (usecase *CreateExampleUseCase) Execute(input ExampleInputDTO) (*ExampleOutputDTO, error) {
	example := entity.ExampleModel{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	}

	if err := usecase.ExampleRepository.Save(&example); err != nil {
		return nil, err
	}

	dto := &ExampleOutputDTO{
		ID:          example.ID,
		Name:        example.Name,
		Description: example.Description,
		CreatedAt:   time.Now(),
	}

	return dto, nil
}
