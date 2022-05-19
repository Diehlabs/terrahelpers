package terra_helper

import (
	"context"
	"fmt"
	"os"

	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
)

type Secret struct {
	Data map[string]interface{}
}

func NewSecret(roleID string,
	wrappedToken string,
	vaultSecretPath string) *Secret {

	r := Secret{}
	r.Data = getVaultSecret(roleID, wrappedToken, vaultSecretPath)

	return &r
}

//------------------------------------------------------------------------------------------------
// Fetches a key-value secret (kv-v2) after authenticating via AppRole.
//------------------------------------------------------------------------------------------------
func getVaultSecret(
	roleID string,
	wrappedToken string,
	vaultSecretPath string,
) map[string]interface{} {
	config := vault.DefaultConfig() // modify for more granular configuration

	client, err := vault.NewClient(config)
	if err != nil {
		panic(fmt.Errorf("unable to initialize Vault client: %w", err))
	}

	//------------------------------------------------------------------------------------------------
	// The Secret ID is a value that needs to be protected, so instead of the
	// app having knowledge of the secret ID directly, we have a trusted orchestrator (https://learn.hashicorp.com/tutorials/vault/secure-introduction?in=vault/app-integration#trusted-orchestrator)
	// give the app access to a short-lived response-wrapping token (https://www.vaultproject.io/docs/concepts/response-wrapping).
	// Read more at: https://learn.hashicorp.com/tutorials/vault/approle-best-practices?in=vault/auth-methods#secretid-delivery-best-practices
	//------------------------------------------------------------------------------------------------
	secretID := &auth.SecretID{FromString: wrappedToken}

	appRoleAuth, err := auth.NewAppRoleAuth(
		roleID,
		secretID,
		auth.WithWrappingToken(),
	)
	if err != nil {
		panic(fmt.Errorf("unable to initialize AppRole auth method: %w", err))
	}

	authInfo, err := client.Auth().Login(context.TODO(), appRoleAuth)
	if err != nil {
		panic(fmt.Errorf("unable to login to AppRole auth method: %w", err))
	}
	if authInfo == nil {
		panic(fmt.Errorf("no auth info was returned after login"))
	}

	// get secret
	secret, err := client.Logical().Read(vaultSecretPath)
	if err != nil {
		panic(fmt.Errorf("unable to read secret: %w", err))
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("data type assertion failed: %T %#v", secret.Data["data"], secret.Data["data"]))
	}

	return data

}

func (s *Secret) MapData(dataMap map[string]string) map[string]string {
	for keyName, valueName := range dataMap {
		value, ok := s.Data[valueName].(string)
		if !ok {
			panic(fmt.Sprintf("lookup failed for key / value pair  %s / %s ", keyName, valueName))
		} else {
			dataMap[keyName] = value
		}
	}

	return dataMap
}

func (s *Secret) SetEnv(data map[string]string) {
	for key, value := range data {
		os.Setenv(key, value)
	}
}
