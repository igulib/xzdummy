package xzdummy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var expect = require.True

func TestSayHello(t *testing.T) {
	SayHello()
	expect(t, true, "TestSayHello failed")
}
