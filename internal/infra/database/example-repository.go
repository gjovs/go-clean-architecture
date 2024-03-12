package database

import (
	"database/sql"
	"time"

	"github.com/gjovs/clean/internal/entity"
)

type ExampleRepository struct {
	DB *sql.DB
}

func NewExampleRepository(db *sql.DB) *ExampleRepository {
	return &ExampleRepository{
		DB: db,
	}
}

func (repo *ExampleRepository) Save(example *entity.ExampleModel) (*entity.ExampleModel, error) {
	statement, err := repo.DB.Prepare("INSERT INTO [Example].[example].[tb_example] (id, name, description) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	result, err := statement.Exec(example.ID, example.Name, example.Description)

	if err != nil {
		return nil, err
	}

	pk, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	var created_at time.Time

	err = repo.DB.QueryRow("SELECT created_at FROM [Example].[example].[tb_example] WHERE id = $1", pk).Scan(&created_at)

	if err != nil {
		return nil, err
	}

	example_fullfilled := &entity.ExampleModel{
		ID:          example.ID,
		Name:        example.Name,
		Description: example.Description,
		CreatedAt:   created_at,
	}

	return example_fullfilled, nil
}
