package utils

import (
	"testing"
)

func TestNotifier (t *testing.T) {
    err := SendNotification("karthikeya", "Hello from go!");
    if err != nil {
        t.Errorf("Error: %s", err);
    }
}
