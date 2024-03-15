package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gjovs/clean/internal/entity"
	usecase "github.com/gjovs/clean/internal/usecase/example"
	"github.com/gjovs/clean/pkg/events"
)

type WebExampleHandler struct {
	EventDispatcher     events.EventDispatcherInterface
	ExampleRepository   entity.ExampleRepositoryInterface
	ExampleCreatedEvent events.EventHandlerInterface
}

func NewWebExampleHandler(
	eventDispatcher events.EventDispatcherInterface,
	exampleRepository entity.ExampleRepositoryInterface,
	exampleCreatedEvent events.EventHandlerInterface,
) *WebExampleHandler {
	return &WebExampleHandler{
		EventDispatcher:     eventDispatcher,
		ExampleRepository:   exampleRepository,
		ExampleCreatedEvent: exampleCreatedEvent,
	}
}

func (handler *WebExampleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.ExampleInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	create_example_usecase := usecase.NewCreateExampleUseCase(handler.ExampleRepository)

	output, err := create_example_usecase.Execute(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
