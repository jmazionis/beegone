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

func TestGetCarplate(t *testing.T) {
	Convey("Subject: Test GetCarplate endpoint", t, func() {
		Convey("Given non existing carplate", func() {
			r, _ := http.NewRequest("GET", "/api/carplates/1", nil)
			w := httptest.NewRecorder()

			beego.BeeApp.Handlers.ServeHTTP(w, r)

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

func TestGetAllCarplates(t *testing.T) {
	Convey("Subject: Test GetCarplates endpoint", t, func() {
		Convey("Given valid request", func() {
			storages.CarPlateDb().Add(&models.CarPlate{
				ID: "A",
			})
			storages.CarPlateDb().Add(&models.CarPlate{
				ID: "B",
			})
			r, _ := http.NewRequest("GET", "/api/carplates", nil)
			w := httptest.NewRecorder()

			beego.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status code should be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The body should contain retrieved items", func() {
				carplates := []*models.CarPlate{}
				_ = json.Unmarshal(w.Body.Bytes(), &carplates)
				So(len(carplates), ShouldEqual, 2)
			})
			Reset(resetStorage)
		})
	})
}

func TestAddCarplate(t *testing.T) {
	Convey("Subject: Test AddCarplate endpoint", t, func() {
		Convey("Given valid request", func() {
			carplateId := "carplate1"
			rBody, _ := json.Marshal(map[string]interface{}{
				"id":        carplateId,
				"plateId":   "AAA-200",
				"modelName": "Subaru Outback",
				"modelYear": "1990",
				"owner":     "Driver 2",
			})
			r, _ := http.NewRequest("POST", "/api/carplates", bytes.NewBuffer(rBody))
			w := httptest.NewRecorder()

			beego.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status code should Be 201", func() {
				So(w.Code, ShouldEqual, 201)
			})
			Convey("Should include location header with location to new carplate", func() {
				So(len(w.Header().Get("location")), ShouldBeGreaterThan, 0)
			})
			Convey("Should return id of newly created carplate", func() {
				newCarplateResp := &models.AddCarplateResponse{}
				_ = json.Unmarshal(w.Body.Bytes(), newCarplateResp)

				So(len(newCarplateResp.ID), ShouldBeGreaterThan, 0)
			})
		})

		Convey("Given invalid request", func() {
			rBody, _ := json.Marshal(map[string]interface{}{
				"plateId":   "AAA-20",
				"modelName": "Subaru Impreza",
				"modelYear": "1990",
			})
			r, _ := http.NewRequest("POST", "/api/carplates", bytes.NewBuffer(rBody))
			w := httptest.NewRecorder()

			beego.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status code should Be 400", func() {
				So(w.Code, ShouldEqual, 400)
			})
			Convey("The body should contain validation errors summary", func() {
				errors := struct {
					Errors []string `json:"errors"`
				}{}
				err := json.Unmarshal(w.Body.Bytes(), &errors)
				if err != nil {
					beego.Error(err)
				}
				So(len(errors.Errors), ShouldEqual, 2)
			})
		})
		Reset(resetStorage)
	})
}

func TestUpdateCarplate(t *testing.T) {
	Convey("Subject: Test UpdateCarplate endpoint", t, func() {
		Convey("Given valid request", func() {
			storages.CarPlateDb().Add(&models.CarPlate{
				ID: "A",
			})
			rBody, _ := json.Marshal(map[string]interface{}{
				"id":        "A",
				"plateId":   "AAA-200",
				"modelName": "Subaru Outback",
				"modelYear": "1990",
				"owner":     "Driver 2",
			})
			r, _ := http.NewRequest("PUT", "/api/carplates", bytes.NewBuffer(rBody))
			w := httptest.NewRecorder()

			beego.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status code should Be 204", func() {
				So(w.Code, ShouldEqual, 204)
			})
			Convey("The body should be empty", func() {
				So(w.Body.Len(), ShouldBeZeroValue)
			})
		})

		Convey("Given non existing resource", func() {
			storages.CarPlateDb().Add(&models.CarPlate{
				ID: "A",
			})
			rBody, _ := json.Marshal(map[string]interface{}{
				"id":        "B",
				"plateId":   "AAA-200",
				"modelName": "Subaru Outback",
				"modelYear": "1990",
				"owner":     "Driver 2",
			})
			r, _ := http.NewRequest("PUT", "/api/carplates", bytes.NewBuffer(rBody))
			w := httptest.NewRecorder()

			beego.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status code should Be 404", func() {
				So(w.Code, ShouldEqual, 404)
			})
			Convey("The body should be empty", func() {
				So(w.Body.Len(), ShouldBeZeroValue)
			})
		})

		Convey("Given invalid request", func() {
			storages.CarPlateDb().Add(&models.CarPlate{
				ID: "A",
			})
			rBody, _ := json.Marshal(map[string]interface{}{
				"plateId":   "AA-200",
				"modelName": "Subaru Outback",
				"modelYear": "1700",
				"owner":     "Driver 2",
			})
			r, _ := http.NewRequest("PUT", "/api/carplates", bytes.NewBuffer(rBody))
			w := httptest.NewRecorder()

			beego.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status code should Be 400", func() {
				So(w.Code, ShouldEqual, 400)
			})
			Convey("The body should contain validation errors summary", func() {
				errors := struct {
					Errors []string `json:"errors"`
				}{}
				err := json.Unmarshal(w.Body.Bytes(), &errors)
				if err != nil {
					beego.Error(err)
				}
				So(len(errors.Errors), ShouldEqual, 2)
			})
		})
		Reset(resetStorage)
	})
}

func TestDeleteCarplate(t *testing.T) {
	Convey("Subject: Test DeleteCarplate endpoint", t, func() {
		Convey("Given existing resource", func() {
			storages.CarPlateDb().Add(&models.CarPlate{
				ID: "A",
			})
			r, _ := http.NewRequest("DELETE", "/api/carplates/A", nil)
			w := httptest.NewRecorder()

			beego.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status code should Be 204", func() {
				So(w.Code, ShouldEqual, 204)
			})
			Convey("The body should be empty", func() {
				So(w.Body.Len(), ShouldBeZeroValue)
			})
		})
		Convey("Given non existing resource", func() {
			storages.CarPlateDb().Add(&models.CarPlate{
				ID: "A",
			})
			r, _ := http.NewRequest("DELETE", "/api/carplates/B", nil)
			w := httptest.NewRecorder()

			beego.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status code should Be 204", func() {
				So(w.Code, ShouldEqual, 204)
			})
			Convey("The body should be empty", func() {
				So(w.Body.Len(), ShouldBeZeroValue)
			})
		})
		Reset(resetStorage)
	})
}

func resetStorage() {
	storages.CarPlateDb().Reset()
}
