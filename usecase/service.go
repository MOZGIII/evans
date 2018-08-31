package usecase

import (
	"io"

	"github.com/MOZGIII/evans/entity"
	"github.com/MOZGIII/evans/usecase/port"
)

func Service(params *port.ServiceParams, outputPort port.OutputPort, env entity.Environment) (io.Reader, error) {
	if err := env.UseService(params.SvcName); err != nil {
		return nil, err
	}
	return outputPort.Service()
}
