package usecase

import (
	"testing"

	"github.com/MOZGIII/evans/adapter/presenter"
	"github.com/MOZGIII/evans/entity"
	"github.com/MOZGIII/evans/entity/testentity"
	"github.com/MOZGIII/evans/tests/helper"
	"github.com/MOZGIII/evans/usecase/port"
	"github.com/stretchr/testify/require"
)

type describeEnv struct {
	entity.Environment

	expected entity.Message
}

func (e *describeEnv) Message(msgName string) (entity.Message, error) {
	return e.expected, nil
}

func TestDescribe(t *testing.T) {
	var expected entity.Message = testentity.NewMsg()

	params := &port.DescribeParams{}
	presenter := presenter.NewJSONCLIPresenter()
	env := &describeEnv{expected: expected}

	res, err := Describe(params, presenter, env)
	require.NoError(t, err)

	actual := helper.ReadAllAsStr(t, res)
	m := &message{expected}
	require.Equal(t, m.Show(), actual)
}
