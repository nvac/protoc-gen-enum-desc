package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayOfWeek(t *testing.T) {
	status := NewStatusFromNumber(1)
	assert.Equal(t, "ACTIVE", status.GetDesc())

	status9 := NewStatusFromNumber(9)
	assert.Equal(t, "", status9.GetDesc())

	assert.Equal(t, "DELETED", status.GetDescFromNumber(2))
	assert.Equal(t, "", status.GetDescFromNumber(99))

	assert.Equal(t, "DELETED", status.GetDescFromName("Status_Deleted"))
	assert.Equal(t, "", status.GetDescFromName("Status_XX"))
}
