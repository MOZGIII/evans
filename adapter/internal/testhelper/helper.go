package testhelper

import (
	"path/filepath"
	"testing"

	"github.com/jhump/protoreflect/desc"
	"github.com/MOZGIII/evans/adapter/internal/protoparser"
	"github.com/MOZGIII/evans/entity"
	"github.com/MOZGIII/evans/tests/helper"
	"github.com/stretchr/testify/require"
)

func ReadProtoAsFileDescriptors(t *testing.T, fpath ...string) []*desc.FileDescriptor {
	for i := range fpath {
		fpath[i] = filepath.Join("testdata", fpath[i])
	}
	set, err := protoparser.ParseFile(fpath, nil)
	require.NoError(t, err)
	require.Len(t, set, len(fpath))
	return set
}

func FindMessage(t *testing.T, name string, set []*desc.FileDescriptor) *desc.MessageDescriptor {
	for _, f := range set {
		for _, msg := range f.GetMessageTypes() {
			if msg.GetName() == name {
				return msg
			}
		}
	}
	require.Fail(t, "message not found: %s", name)
	return nil
}

func SetupEnv(t *testing.T, fpath, pkgName, svcName string) *entity.Env {
	t.Helper()

	set := helper.ReadProto(t, fpath)

	env := entity.NewEnv(set, helper.TestConfig())

	err := env.UsePackage(pkgName)
	require.NoError(t, err)

	if svcName != "" {
		err = env.UseService(svcName)
		require.NoError(t, err)
	}

	return env
}
