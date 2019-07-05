package alerts_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListAlerts(t *testing.T) {
	underTest, err := setupAlertsTest()

	if assert.NoError(t, err) {
		_, err = underTest.ListAlerts()
	} else {
		t.Fatalf("%q should not have raised an error: %v", "ListAlerts", err)
	}
}
