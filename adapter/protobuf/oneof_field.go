package protobuf

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/MOZGIII/evans/entity"
)

type oneOfField struct {
	d       *desc.OneOfDescriptor
	choices []entity.Field
}

func (o *oneOfField) FieldName() string {
	return o.d.GetName()
}

func (o *oneOfField) FQRN() string {
	return o.d.GetFullyQualifiedName()
}

func (o *oneOfField) Type() entity.FieldType {
	return entity.FieldTypeOneOf
}

func (o *oneOfField) IsRepeated() bool {
	return false
}

// oneof hasn't any type
func (o *oneOfField) PBType() string {
	return "oneof"
}

func (o *oneOfField) Choices() []entity.Field {
	return o.choices
}
