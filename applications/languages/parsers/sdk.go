package parsers

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/roots"
)

// EventFunc represents an event func
type EventFunc func(root roots.Root, manager Manager) (interface{}, error)

// NewApplication creates a new application instance
func NewApplication() Application {
	managerBuilder := NewManagerBuilder()
	return createApplication(managerBuilder)
}

// NewManagerBuilder creates a new manager builder
func NewManagerBuilder() ManagerBuilder {
	resultBuilder := NewResultBuilder()
	return createManagerBuilder(resultBuilder)
}

// NewResultBuilder creates a new result builder
func NewResultBuilder() ResultBuilder {
	return createResultBuilder()
}

// Application represents a parser application
type Application interface {
	Execute(root roots.Root, events map[string][]EventFunc) (interface{}, error)
}

// ManagerBuilder represents a manager builder
type ManagerBuilder interface {
	Create() ManagerBuilder
	WithEvents(events map[string][]EventFunc) ManagerBuilder
	Now() (Manager, error)
}

// Manager presents the event manager
type Manager interface {
	Execute(root roots.Root) (Result, error)
	Trigger(name string, root roots.Root) (Result, error)
}

// ResultBuilder represents a result builder
type ResultBuilder interface {
	Create() ResultBuilder
	WithElement(element interface{}) ResultBuilder
	WithList(list []interface{}) ResultBuilder
	Now() (Result, error)
}

// Result represents a result
type Result interface {
	First() interface{}
	List() []interface{}
	IsSingle() bool
}
