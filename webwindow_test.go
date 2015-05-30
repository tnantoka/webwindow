package webwindow

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig(8080, "Config")

	if !(config.Port == 8080 &&
		config.Host == "localhost" &&
		config.Root == "index.html" &&
		config.Width == 640 &&
		config.Height == 480 &&
		config.Title == "Config") {
		t.FailNow()
	}
}
