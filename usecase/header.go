package usecase

import (
	"io"

	"github.com/MOZGIII/evans/entity"
	"github.com/MOZGIII/evans/usecase/port"
)

func Header(
	params *port.HeaderParams,
	outputPort port.OutputPort,
	env entity.Environment,
) (io.Reader, error) {
	for _, h := range params.Headers {
		if h.NeedToRemove {
			env.RemoveHeader(h.Key)
		} else {
			env.AddHeader(h)
		}
	}
	return outputPort.Header()
}
