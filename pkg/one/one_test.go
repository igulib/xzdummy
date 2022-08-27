package one

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var expect = require.True

func TestPackageOneHello(t *testing.T) {
	PackageOneHello()
	expect(t, true, "PackageOneHello failed")
}
