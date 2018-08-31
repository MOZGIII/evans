package usecase

import (
	"testing"

	"github.com/MOZGIII/evans/adapter/presenter"
	"github.com/MOZGIII/evans/entity"
	"github.com/MOZGIII/evans/usecase/port"
	"github.com/stretchr/testify/require"
)

type packageEnv struct {
	entity.Environment

	usedPackage string
}

func (e *packageEnv) UsePackage(pkgName string) error {
	e.usedPackage = pkgName
	return nil
}

func TestPackage(t *testing.T) {
	expected := "example_package"
	params := &port.PackageParams{PkgName: expected}
	presenter := &presenter.StubPresenter{}
	env := &packageEnv{}

	_, err := Package(params, presenter, env)
	require.NoError(t, err)
	require.Equal(t, expected, env.usedPackage)
}
