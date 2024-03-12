package entity

type ExampleModel struct {
	ID          string
	Name        string
	Description string
}

func NewExampleModel(id, name, description string) (*ExampleModel, error) {
	example := &ExampleModel{
		ID:          id,
		Name:        name,
		Description: description,
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
