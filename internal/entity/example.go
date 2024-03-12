package entity

import "time"

type ExampleModel struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
}

func NewExampleModel(id, name, description string, createdAt time.Time) (*ExampleModel, error) {
	example := &ExampleModel{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
	}

	if err := example.IsValid(); err != nil {
		return nil, err
	}

	return example, nil
}

func (model *ExampleModel) IsValid() error {
	// TODO: add validation here
	return nil
}
