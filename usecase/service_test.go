package usecase

import (
	"testing"

	"github.com/MOZGIII/evans/adapter/presenter"
	"github.com/MOZGIII/evans/entity"
	"github.com/MOZGIII/evans/usecase/port"
	"github.com/stretchr/testify/require"
)

type serviceEnv struct {
	entity.Environment

	usedService string
}

func (e *serviceEnv) UseService(pkgName string) error {
	e.usedService = pkgName
	return nil
}

func TestService(t *testing.T) {
	expected := "example_service"
	params := &port.ServiceParams{SvcName: expected}
	presenter := &presenter.StubPresenter{}
	env := &serviceEnv{}

	_, err := Service(params, presenter, env)
	require.NoError(t, err)
	require.Equal(t, expected, env.usedService)
}
