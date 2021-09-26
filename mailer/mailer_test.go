package mailer

import (
	"GoMailer/internal/projectpath"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {

	// Perhaps the hachiest solution one could come up with, but it works so
	godotenv.Load(filepath.Join(projectpath.Root, "/.env"))

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
