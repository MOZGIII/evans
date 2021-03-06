package usecase

import (
	"io"

	"github.com/MOZGIII/evans/entity"
	"github.com/MOZGIII/evans/usecase/port"
)

func Describe(params *port.DescribeParams, outputPort port.OutputPort, env entity.Environment) (io.Reader, error) {
	msg, err := env.Message(params.MsgName)
	if err != nil {
		return nil, err
	}
	return outputPort.Describe(&message{msg})
}
