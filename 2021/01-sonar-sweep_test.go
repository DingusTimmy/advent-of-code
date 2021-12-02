package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSonarSweep(t *testing.T) {
	Convey("Given a previous string and a current string", t, func() {
		previous := "0"
		current := "0"

		Convey("When the current string's value is higher", func() {
			current = "1"

			Convey("Then sonarSweep should return true and the expected string", func() {
				increased, printString := sonarSweep(previous, current)
				So(increased, ShouldBeTrue)
				So(printString, ShouldEqual, "1 (increased)")
			})
		})
		Convey("When the current string's value is lower", func() {
			current = "-1"

			Convey("Then sonarSweep should return false and the expected string", func() {
				increased, printString := sonarSweep(previous, current)
				So(increased, ShouldBeFalse)
				So(printString, ShouldEqual, "-1 (decreased)")
			})
		})
		Convey("When the both strings' values are equal", func() {
			current = "0"

			Convey("Then sonarSweep should return false and the expected string", func() {
				increased, printString := sonarSweep(previous, current)
				So(increased, ShouldBeFalse)
				So(printString, ShouldEqual, "0 (equal)")
			})
		})
	})
}
