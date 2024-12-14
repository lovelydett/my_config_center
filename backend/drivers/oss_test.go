package drivers

import (
	"testing"
)

func TestOSSConnection(t *testing.T) {
	// Test OSS connection
	_, err := Bucket.ListObjects()

	if err != nil {
		t.Errorf("Failed to connect to OSS: %v", err)
	}
}
