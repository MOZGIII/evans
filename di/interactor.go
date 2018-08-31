package di

import (
	"io"

	"github.com/MOZGIII/evans/config"
	"github.com/MOZGIII/evans/usecase"
)

func NewCLIInteractorParams(cfg *config.Config, in io.Reader) (*usecase.InteractorParams, error) {
	if err := initDependencies(cfg, in); err != nil {
		return nil, err
	}

	return &usecase.InteractorParams{
		Env:            env,
		OutputPort:     jsonCLIPresenter,
		InputterPort:   jsonFileInputter,
		GRPCClient:     gRPCClient,
		DynamicBuilder: dynamicBuilder,
	}, nil
}

func NewREPLInteractorParams(cfg *config.Config, in io.Reader) (param *usecase.InteractorParams, err error) {
	if err := initDependencies(cfg, in); err != nil {
		return nil, err
	}
	return &usecase.InteractorParams{
		Env:            env,
		OutputPort:     jsonCLIPresenter,
		InputterPort:   promptInputter,
		GRPCClient:     gRPCClient,
		DynamicBuilder: dynamicBuilder,
	}, nil
}
