package integration

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	time.Local = time.UTC
	if !strings.EqualFold(os.Getenv("ENV"), "") {
		m.Run()
	}
}
