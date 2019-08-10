package integration_tests

import (
	"net/http"
	"net/http/httptest"

	"path/filepath"
	"runtime"
	"testing"
	_ "do/internal/api/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestGetCarplates(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/carplates", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("test", "TestGetCarplates", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test GetCarplates Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGetCarplate(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/carplates/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("test", "TestGetCarplate", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test GetCarplate Endpoint\n", t, func() {
		Convey("Status Code Should Be 404 if resource does not exist", func() {
			So(w.Code, ShouldEqual, 404)
		})
		Convey("The Result Should Be Empty", func() {
			So(w.Body.Len(), ShouldBeZeroValue)
		})
	})
}
