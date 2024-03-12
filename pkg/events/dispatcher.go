package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (dispatcher *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := dispatcher.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (dispatcher *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := dispatcher.handlers[eventName]; ok {
		for _, h := range dispatcher.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	dispatcher.handlers[eventName] = append(dispatcher.handlers[eventName], handler)
	return nil
}

func (dispatcher *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := dispatcher.handlers[eventName]; ok {
		for _, h := range dispatcher.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (dispatcher *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := dispatcher.handlers[eventName]; ok {
		for i, h := range dispatcher.handlers[eventName] {
			if h == handler {
				dispatcher.handlers[eventName] = append(dispatcher.handlers[eventName][:i], dispatcher.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (dispatcher *EventDispatcher) Clear() {
	dispatcher.handlers = make(map[string][]EventHandlerInterface)
}
