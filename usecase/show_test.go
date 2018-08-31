package usecase

import (
	"testing"

	"github.com/MOZGIII/evans/adapter/presenter"
	"github.com/MOZGIII/evans/entity"
	"github.com/MOZGIII/evans/tests/helper"
	"github.com/MOZGIII/evans/usecase/port"
	"github.com/stretchr/testify/require"
)

type showEnv struct {
	entity.Environment

	expected []*entity.Package
}

func (e *showEnv) Packages() []*entity.Package {
	return e.expected
}

func TestShow(t *testing.T) {
	expected := []*entity.Package{{Name: "example_package"}}
	params := &port.ShowParams{Type: port.ShowTypePackage}
	presenter := presenter.NewJSONCLIPresenter()
	env := &showEnv{expected: expected}

	res, err := Show(params, presenter, env)
	require.NoError(t, err)

	actual := helper.ReadAllAsStr(t, res)

	require.Equal(t, packages(expected).Show(), actual)
}
