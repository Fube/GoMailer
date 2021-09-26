package mailer

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

func TestDialer_SendEmail(t *testing.T) {

	// Explicit port setting for full coverage
	dialer := CreateMailer("smtp.gmail.com", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), 587)
	mail := Mail{To: "test@exmaple.com", Message: "msg"}
	assert.Nil(t, dialer.SendEmail(&mail))
}
