package parsers

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/roots"
)

type application struct {
	managerBuilder ManagerBuilder
}

func createApplication(
	managerBuilder ManagerBuilder,
) Application {
	out := application{
		managerBuilder: managerBuilder,
	}

	return &out
}

// Execute executes the application and returns the parsed instance
func (app *application) Execute(root roots.Root, events map[string][]EventFunc) (interface{}, error) {
	manager, err := app.managerBuilder.Create().WithEvents(events).Now()
	if err != nil {
		return nil, err
	}

	result, err := manager.Execute(root)
	if err != nil {
		return nil, err
	}

	return result.First(), nil
}
