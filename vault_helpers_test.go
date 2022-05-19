package terra_helper

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidEmptyRoleId(t *testing.T) {
	var roleId string = ""
	var wrappedToken string = ""
	var vaultSecretPath = ""
	const ErrorMessage = "unable to initialize AppRole auth method:"

	defer func() {
		if r := recover(); r != nil {
			s := fmt.Sprintf("Panic detected:%s", r)
			fmt.Println(s)
			if !strings.Contains(s, ErrorMessage) {
				t.Errorf("expected message containing: %s", ErrorMessage)
			}
		}
	}()

	NewSecret(roleId, wrappedToken, vaultSecretPath)

}

func TestValidInvalidSecretId(t *testing.T) {
	var roleId string = "someAppRoleId"
	var wrappedToken string = ""
	var vaultSecretPath = ""
	const ErrorMessage = "unable to initialize AppRole auth method:"

	defer func() {
		if r := recover(); r != nil {
			s := fmt.Sprintf("Panic detected:%s", r)
			fmt.Println(s)
			if !strings.Contains(s, ErrorMessage) {
				t.Errorf("expected message containing: %s", ErrorMessage)
			}
		}
	}()

	NewSecret(roleId, wrappedToken, vaultSecretPath)

}

// func TestGetEnvSettings(t *testing.T) {
// 	var roleId string = os.Getenv("VAULT_APPROLE_ID")
// 	var wrappedToken string = os.Getenv("VAULT_WRAPPED_TOKEN")
// 	var vaultSecretPath = "cloudauto/data/terraform/nonprod/azure/svcazsp-cloudauto-terratest-devtest-terraform-managed"
// 	var dataMap = map[string]string{
// 		"ARM_CLIENT_ID":         "client_id",
// 		"ARM_CLIENT_SECRET":     "client_secret",
// 		"ARM_TENANT_ID":         "tenant_id",
// 		"ARM_SUBSCRIPTION_ID":   "subscription_id",
// 		"AZURE_CLIENT_ID":       "client_id",
// 		"AZURE_CLIENT_SECRET":   "client_secret",
// 		"AZURE_TENANT_ID":       "tenant_id",
// 		"AZURE_SUBSCRIPTION_ID": "subscription_id",
// 	}

// 	s := NewSecret(roleId, wrappedToken, vaultSecretPath)

// 	mapped := s.MapData(dataMap)
// 	for key, value := range mapped {
// 		os.Setenv(key, value)
// 		fmt.Println("Key:", key, "=>", "Value:", value)
// 	}
// }
