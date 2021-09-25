package controller

import (
	"GoMailer/internal/projectpath"
	mailerM "GoMailer/mailer"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type MockMailer struct {
	mock.Mock
}

func (m MockMailer) SendEmail(mail *mailerM.Mail) error {
	args := m.Called()
	return args.Error(0)
}

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

func TestRegisterMailRouteWithValidInfo(t *testing.T) {
	// Just so that it doesn't actually send an email
	mockMailer := new(MockMailer)

	impl := MailerControllerImpl{}

	router := gin.Default()
	impl.Inject(mockMailer)
	err := impl.Routes(router)

	assert.Nil(t, err)

	marshal, _ := json.Marshal(mailerM.Mail{To: "test@test.test", Subject: "Subject", Message: "Message"})
	serial := string(marshal)

	req, err := http.NewRequest(http.MethodPost, "/mail", strings.NewReader(serial))

	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {

		assert.Equal(t, http.StatusOK, w.Code)
		return true
	})
}

func TestRegisterMailRouteWithInvalidInfo(t *testing.T) {

	// Just so that it doesn't actually send an email
	mockMailer := new(MockMailer)

	impl := MailerControllerImpl{}

	router := gin.Default()
	impl.Inject(mockMailer)
	err := impl.Routes(router)
	assert.Nil(t, err)

	marshal, _ := json.Marshal(mailerM.Mail{To: "", Subject: "Subject", Message: "Message"})
	serial := string(marshal)

	req, err := http.NewRequest(http.MethodPost, "/mail", strings.NewReader(serial))

	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {

		assert.Equal(t, http.StatusBadRequest, w.Code)
		return true
	})
}

func TestNilInject(t *testing.T) {
	router := gin.Default()
	impl := MailerControllerImpl{}
	assert.Error(t, impl.Routes(router))

	impl.Inject(nil)
	assert.Error(t, impl.Routes(router))
}

func TestHandleSendEmailNilMailInContext(t *testing.T) {

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	MailerControllerImpl{}.handleSendEmail(c)
	assert.Equal(t, "Unable to parse e-mail from body", recorder.Body)
}

func TestHandleSendEmailErrInMailerSendEmail(t *testing.T) {

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	testMail := mailerM.Mail{To: "test@test.test", Subject: "test-subject", Message: "test-message"}
	c.Set("mail", &testMail)

	impl := MailerControllerImpl{}
	mailer := MockMailer{}
	mailer.On("SendEmail", mock.Anything).Return(errors.New("Test error"))
	impl.mailer = mailer
	impl.handleSendEmail(c)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	assert.Equal(t, "\"Test error\"", recorder.Body.String())
}