package parsers

import (
	"errors"
	"fmt"
	"log"

	"github.com/steve-care-software/steve/domain/languages/lexers/roots"
)

type manager struct {
	resultBuilder ResultBuilder
	events        map[string][]EventFunc
}

func createManager(
	resultBuilder ResultBuilder,
	events map[string][]EventFunc,
) Manager {
	out := manager{
		resultBuilder: resultBuilder,
		events:        events,
	}

	return &out
}

// Execute executes an event on the given root instance and returns the result
func (app *manager) Execute(root roots.Root) (Result, error) {
	name := root.Name()
	list, ins, err := app.executeEvents(name, root)
	if err != nil {
		return nil, err
	}

	builder := app.resultBuilder.Create()
	if list != nil {
		builder.WithList(list)
	}

	if ins != nil {
		builder.WithElement(ins)
	}

	return builder.Now()
}

// Trigger triggers an event on the root instance and returns the result
func (app *manager) Trigger(name string, input roots.Root) (Result, error) {
	root, node, container, content, err := input.Nodes().Scan(name)
	if err != nil {
		return nil, err
	}

	builder := app.resultBuilder.Create()
	if root != nil {
		list, ins, err := app.executeEvents(name, root)
		if err != nil {
			return nil, err
		}

		if list != nil {
			builder.WithList(list)
		}

		if ins != nil {
			builder.WithElement(ins)
		}
	}

	if node != nil {
		list, err := app.executeEventList(name, node)
		if err != nil {
			return nil, err
		}

		builder.WithList(list)
	}

	if container != nil {
		list, ins, err := app.fetchContentOrExecuteEvent(name, container)
		if err != nil {
			return nil, err
		}

		if list != nil {
			builder.WithList(list)
		}

		if ins != nil {
			builder.WithElement(ins)
		}
	}

	if content != "" {
		builder.WithElement(content)
	}

	return builder.Now()
}

func (app *manager) executeEvents(name string, root roots.Root) ([]interface{}, interface{}, error) {
	if list, ok := app.events[name]; ok {
		for index, oneEventFunc := range list {
			ins, err := oneEventFunc(root, app)
			if err != nil {
				str := fmt.Sprintf("%s.%d: %s", name, index, err.Error())
				log.Printf(str)
				continue
			}

			if res, ok := ins.(Result); ok {
				if res.IsSingle() {
					return nil, res.First(), nil
				}

				return res.List(), nil, nil
			}

			return nil, ins, nil
		}

		str := fmt.Sprintf("here was %d event(s) associated to the root token (name: %s), but none was successfully executed", len(list), name)
		return nil, nil, errors.New(str)
	}

	str := fmt.Sprintf("there is no Event associated with the given root token (name: %s)", name)
	return nil, nil, errors.New(str)
}

func (app *manager) executeEventList(name string, node roots.Node) ([]interface{}, error) {
	out := []interface{}{}
	containers := node.Containers().List()
	for _, oneContainer := range containers {
		list, ins, err := app.fetchContentOrExecuteEvent(name, oneContainer)
		if err != nil {
			return nil, err
		}

		if list != nil {
			out = append(out, list)
			continue
		}

		out = append(out, ins)
	}

	return out, nil
}

func (app *manager) fetchContentOrExecuteEvent(name string, container roots.Container) ([]interface{}, interface{}, error) {
	if !container.HasRoot() {
		return nil, container.Content(), nil
	}

	root := container.Root()
	return app.executeEvents(name, root)
}
