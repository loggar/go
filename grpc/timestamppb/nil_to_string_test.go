package timestamppb_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNilTimeToString(t *testing.T) {
	var ts *timestamppb.Timestamp

	assert.Nil(t, ts, "Expected nil")

	str := ts.String()

	assert.Equal(t, "<nil>", str, "nil to string")
}
