package obtime

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNow(t *testing.T) {
	Convey("now", t, func() {
		So(Now().Unix(), ShouldEqual, time.Now().Unix())

		Convey("set local", func() {
			now := time.Now()
			localTime := time.Date(
				now.Year()-1,
				now.Month(),
				now.Day(),
				now.Hour(),
				now.Minute(),
				now.Second(),
				0,
				now.Location(),
			)
			SetLocal(localTime)
			So(Now().Unix(), ShouldEqual, localTime.Unix())

			Convey("reset local", func() {
				ResetLocal()
				So(Now().Unix(), ShouldEqual, time.Now().Unix())
			})
		})
	})
}
