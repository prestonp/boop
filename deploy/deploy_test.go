package deploy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeployer(t *testing.T) {
	d, err := New("./echo.sh", "/tmp/deploy-logs")
	require.NoError(t, err)
	assert.NotNil(t, d)

	// Deploy for the first time
	err = d.Deploy()
	assert.NoError(t, err)

	deployments, err := d.List()
	assert.NoError(t, err)
	assert.Len(t, deployments, 1)

	// Deploy again, should create second record
	err = d.Deploy()
	assert.NoError(t, err)

	deployments, err = d.List()
	assert.NoError(t, err)
	assert.Len(t, deployments, 2)

	// Fetch
	time.Sleep(100 * time.Millisecond)
	dep, err := d.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, StatusSuccess, dep.Status)
}
