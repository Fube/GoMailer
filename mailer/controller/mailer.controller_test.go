package controller

import (
	"GoMailer/internal/projectpath"
	"GoMailer/mailer"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func TestMain(m *testing.M) {

	// Perhaps the hachiest solution one could come up with, but it works so
	godotenv.Load(filepath.Join(projectpath.Root, "/.env"))

	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

func TestRegisterMailRoute(t *testing.T) {
	router := gin.Default()
	Routes(router)

	marshal, _ := json.Marshal(mailer.Mail{To: "test@test.test", Subject: "Subject", Message: "Message"})
	json := string(marshal)

	req, err := http.NewRequest(http.MethodPost, "/mail", strings.NewReader(json))

	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {

		statusOk := w.Code == http.StatusOK
		return statusOk
	})
}
