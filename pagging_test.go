package mygo_pagination

import (
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMyPagging(t *testing.T) {

	Convey("test page gen case 1", t, func() {
		papping := Pagging{
			Limit: 10,
			Page:  1,
		}
		pg := papping.PageInfoGenerator(200)
		So(pg.Page, ShouldEqual, 1)
		So(pg.TotalRecord, ShouldEqual, 200)
		So(pg.TotalPage, ShouldEqual, 20)
		So(pg.QLimit, ShouldEqual, 9)
		So(pg.QOffset, ShouldEqual, 0)
	})

	Convey("test page gen case 2", t, func() {
		papping := Pagging{
			Limit: 10,
			Page:  1,
		}
		pg := papping.PageInfoGenerator(207)
		log.Infof("%v", pg)
		So(pg.TotalRecord, ShouldEqual, 207)
		So(pg.TotalPage, ShouldEqual, 21)
	})

	Convey("test page gen case 2", t, func() {
		papping := Pagging{
			Limit: 20,
			Page:  2,
		}
		pg := papping.PageInfoGenerator(207)
		log.Infof("%v", pg)
		So(pg.TotalRecord, ShouldEqual, 207)
		So(pg.TotalPage, ShouldEqual, 11)
		So(pg.QOffset, ShouldEqual, 20)
		So(pg.QLimit, ShouldEqual, 39)
	})

}
