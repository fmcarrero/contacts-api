package integration

import (
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	if !strings.EqualFold(os.Getenv("ENV"), "") {
		m.Run()
	}
}
