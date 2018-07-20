package providers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: overwrite the exec.Command stuff to we aren't actually executing
// 	     shell commands. Need mocks.

func Test_ReadAccounts(t *testing.T) {
	gcloud := GCloudProvider{}
	accs, err := gcloud.ReadAccounts()
	assert.NotNil(t, accs)
	assert.NoError(t, err)
}

func Test_SelectAccount(t *testing.T) {
	gcloud := GCloudProvider{}
	err := gcloud.SelectAccount("")
	// Even an empty string would come back exit status 0.
	assert.NoError(t, err)
}
