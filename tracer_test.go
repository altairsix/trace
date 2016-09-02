package trace_test

import (
	"context"
	"testing"

	"io"

	"github.com/altairsix/trace"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFrom(t *testing.T) {
	ctx := context.Background()

	Convey("Verify simple log message", t, func() {
		tracer := trace.FromContext(ctx)
		err := tracer.Wrap(io.EOF, "test", trace.String("hello", "world"))
		So(err, ShouldNotBeNil)
	})
}
