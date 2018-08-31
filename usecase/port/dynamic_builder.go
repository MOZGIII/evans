package port

import (
	"github.com/golang/protobuf/proto"
	"github.com/MOZGIII/evans/entity"
)

type DynamicBuilder interface {
	NewMessage(m entity.Message) proto.Message
}
