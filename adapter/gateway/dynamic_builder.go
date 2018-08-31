package gateway

import (
	"github.com/golang/protobuf/proto"
	"github.com/MOZGIII/evans/adapter/protobuf"
	"github.com/MOZGIII/evans/entity"
)

type DynamicBuilder struct{}

func NewDynamicBuilder() *DynamicBuilder {
	return &DynamicBuilder{}
}

func (b *DynamicBuilder) NewMessage(m entity.Message) proto.Message {
	return protobuf.NewDynamicMessage(m)
}
