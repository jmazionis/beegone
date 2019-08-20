package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"path/filepath"
	"runtime"
	"testing"

	"github.com/ICanHaz/beegone/internal/api/models"
	_ "github.com/ICanHaz/beegone/internal/api/routers"
	"github.com/ICanHaz/beegone/internal/api/storages"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestAddCarplate(t *testing.T) {

	Convey("Subject: Test AddCarplate Endpoint\n", t, func() {
		Convey("Given valid request", func() {
			rBody, err := json.Marshal(map[string]interface{}{
				"plateId":   "AAA-200",
				"modelName": "Subaru Outback",
				"modelYear": 1990,
				"owner":     "Driver 2",
			})

			if err != nil {
				beego.Error(err)
			}

			r, _ := http.NewRequest("POST", "/api/carplates", bytes.NewBuffer(rBody))
			w := httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)

			beego.Trace("test", "TestAddCarplate", "Code[%d]\n%s", w.Code, w.Body.String())

			Convey("Status code should Be 201", func() {
				So(w.Code, ShouldEqual, 201)
			})

			Convey("Should include location header", func() {
				So(len(w.Header().Get("location")), ShouldBeGreaterThan, 0)
			})

			Convey("The body should be empty", func() {
				So(w.Body.Len(), ShouldBeZeroValue)
			})
		})

		Convey("Given invalid request", func() {
			rBody, err := json.Marshal(map[string]interface{}{
				"plateId":   "AAA-20",
				"modelName": "Subaru Impreza",
				"modelYear": 1990,
			})

			if err != nil {
				beego.Error(err)
			}

			r, _ := http.NewRequest("POST", "/api/carplates", bytes.NewBuffer(rBody))
			w := httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)

			beego.Trace("test", "TestAddCarplate", "Code[%d]\n%s", w.Code, w.Body.String())

			Convey("Status code should Be 400", func() {
				So(w.Code, ShouldEqual, 400)
			})
		})

		Reset(resetStorage)
	})

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
	Convey("Subject: Test GetCarplate Endpoint\n", t, func() {
		Convey("Given non existing carplate", func() {
			r, _ := http.NewRequest("GET", "/api/carplates/1", nil)
			w := httptest.NewRecorder()
			beego.BeeApp.Handlers.ServeHTTP(w, r)

			beego.Trace("test", "TestGetCarplate", "Code[%d]\n%s", w.Code, w.Body.String())
			Convey("Status code should be 404", func() {
				So(w.Code, ShouldEqual, 404)
			})
			Convey("The body should be empty", func() {
				So(w.Body.Len(), ShouldBeZeroValue)
			})
		})

		Convey("Given existing carplate", func() {
			const carplateId = "1"
			storages.CarPlateDb().Add(&models.CarPlate{
				ID:        carplateId,
				PlateID:   "AAA-000",
				ModelName: "Toyoda",
				ModelYear: 2010,
				Owner:     "Person",
			})
			r, _ := http.NewRequest("GET", "/api/carplates/1", nil)
			w := httptest.NewRecorder()

			beego.BeeApp.Handlers.ServeHTTP(w, r)

			beego.Trace("test", "TestGetCarplate", "Code[%d]\n%s", w.Code, w.Body.String())
			Convey("Status code should be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})

			Convey("The body should contain retrieved item", func() {
				carplate := models.CarPlate{}
				_ = json.Unmarshal(w.Body.Bytes(), &carplate)
				So(carplate.ID, ShouldEqual, carplateId)
			})
		})

		Reset(resetStorage)
	})
}

func resetStorage() {
	storages.CarPlateDb().Reset()
}
