package reporter

import (
	"bytes"
	"github.com/Hranoprovod/shared"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReporter(t *testing.T) {
	Convey("Given reporter", t, func() {
		var b bytes.Buffer
		rp := NewReporter(&b)
		Convey("Prints list of API results", func() {
			nl := shared.APINodeList{
				shared.APINode{
					Name: "test1",
				},
				shared.APINode{
					Name: "test2",
				},
			}
			expected := `test1
test2
`
			err := rp.PrintAPISearchResult(nl)
			So(err, ShouldBeNil)
			So(b.String(), ShouldEqual, expected)
		})
	})
}
