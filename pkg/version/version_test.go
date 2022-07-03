package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	assert.True(t, len(Version) > 0)
	assert.False(t, len(CommitHash) > 0)
	assert.False(t, len(Timestamp) > 0)
}
