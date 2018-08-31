package e2e

import (
	"log"
	"testing"

	"github.com/MOZGIII/evans/cache"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	if err := cache.Clear(); err != nil {
		log.Fatal(err)
	}

	goleak.VerifyTestMain(
		m,
		// TODO: invest these leaks
		goleak.IgnoreTopFunction("github.com/MOZGIII/evans/vendor/google.golang.org/grpc"),
		goleak.IgnoreTopFunction("github.com/MOZGIII/evans/vendor/google.golang.org/grpc.(*ccBalancerWrapper).watcher"),
		goleak.IgnoreTopFunction("github.com/MOZGIII/evans/vendor/google.golang.org/grpc.(*ccResolverWrapper).watcher"),
		goleak.IgnoreTopFunction("github.com/MOZGIII/evans/vendor/google.golang.org/grpc.(*addrConn).createTransport"),

		// ref. controller.(*executor).execute comments
		goleak.IgnoreTopFunction("time.Sleep"),
	)
}
