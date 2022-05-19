package terra_helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpnEnvVars(t *testing.T) {
	mymap := GetSpnEnvVars("1", "secret", "sub", "tenant")
	assert.Equal(t, mymap["ARM_CLIENT_ID"], "1", "The Client Id should be the same.")
	assert.Equal(t, mymap["ARM_CLIENT_SECRET"], "secret", "The secret should be the same.")
	assert.Equal(t, mymap["ARM_TENANT_ID"], "tenant", "The tenant should be the same.")
	assert.Equal(t, mymap["ARM_SUBSCRIPTION_ID"], "sub", "The subscription should be the same.")
}

func TestGetResourceTags(t *testing.T) {
	mymap := GetResourceTags("product", "region", "environment", "owner", "tech contact", "cost center")
	assert.Equal(t, mymap["product"], "product", "The product should be the same.")
	assert.Equal(t, mymap["region"], "region", "The region should be the same.")
	assert.Equal(t, mymap["environment"], "environment", "The environment should be the same.")
	assert.Equal(t, mymap["owner"], "owner", "The owner should be the same.")
	assert.Equal(t, mymap["technical_contact"], "tech contact", "The technical contact should be the same.")
	assert.Equal(t, mymap["cost_center"], "cost center", "The cost center should be the same.")
}
