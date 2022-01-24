package chains

import (
	"errors"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	id          *uuid.UUID
	root        *hash.Hash
	createdOn   *time.Time
	head        Block
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		id:          nil,
		root:        nil,
		createdOn:   nil,
		head:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithID adds an id to the builder
func (app *builder) WithID(id uuid.UUID) Builder {
	app.id = &id
	return app
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root hash.Hash) Builder {
	app.root = &root
	return app
}

// WithHead adds a head to the builder
func (app *builder) WithHead(head Block) Builder {
	app.head = head
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.id == nil {
		return nil, errors.New("the id is mandatory in order to build a Chain instance")
	}

	if app.root == nil {
		return nil, errors.New("the root hash is mandatory in order to build a Chain instance")
	}

	if app.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Chain instance")
	}

	data := [][]byte{
		app.id.Bytes(),
		app.root.Bytes(),
		[]byte(strconv.Itoa(int(app.createdOn.Unix()))),
	}

	if app.head != nil {
		data = append(data, app.head.Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.head != nil {
		return createChainWithHead(*app.id, *hash, *app.root, *app.createdOn, app.head), nil
	}

	return createChain(*app.id, *hash, *app.root, *app.createdOn), nil
}
