package terrahelpers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/vault/api/auth/approle"
)

func TestValidEmptyRoleId(t *testing.T) {
	var roleId string = ""
	var secretId *approle.SecretID = nil
	var vaultSecretPath = ""
	const ErrorMessage = "invalid role ID was provided"

	defer func() {
		if r := recover(); r != nil {
			s := fmt.Sprintf("Panic detected:%s", r)
			fmt.Println(s)
			if !strings.Contains(s, ErrorMessage) {
				t.Errorf("expected message containing: %s", ErrorMessage)
			}
		}
	}()

	_, err := GetSecretWithAppRole(roleId, secretId, vaultSecretPath)
	if err != nil {
		t.Errorf(fmt.Sprintf("Test Failed: %s", err.Error()))
	}
}

func TestValidInvalidSecretId(t *testing.T) {
	var roleId string = "someAppRoleId"
	var secretId *approle.SecretID = nil
	var vaultSecretPath = ""
	const ErrorMessage = "invalid secret ID was provided"

	defer func() {
		if r := recover(); r != nil {
			s := fmt.Sprintf("Panic detected:%s", r)
			fmt.Println(s)
			if !strings.Contains(s, ErrorMessage) {
				t.Errorf("expected message containing: %s", ErrorMessage)
			}
		}
	}()

	_, err := GetSecretWithAppRole(roleId, secretId, vaultSecretPath)
	if err != nil {
		t.Errorf(fmt.Sprintf("Test Failed: %s", err.Error()))
	}
}
