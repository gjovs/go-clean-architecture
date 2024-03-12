package entity

type ExampleRepositoryInterface interface {
	Save(example *ExampleModel) (*ExampleModel, error)
}
