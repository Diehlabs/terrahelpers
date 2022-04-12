#!/bin/sh
# This work can be done with Curl versus the Vault CLI, but it's easier this way and we probably have it installed already anyway.
TT_TIMEOUT="40m"
# export SKIP_cleanupModule=true

approle_name=$1
if [[ -z "$approle_name"]]; then
  echo 'Usage:'
  echo 'vault login -method=xxx username=xxx'
  echo './local-test.sh <approle name>'
fi

export VAULT_APPROLE_ID=$(vault read -format=json auth/approle/role/${approle_name}/role-id | jq -r .data.role_id)
export VAULT_WRAPPED_TOKEN=$(vault write -f -format=json -wrap-ttl=5m auth/approle/role/${approle_name}/secret-id | jq -r .wrap_info.token)
echo "Using approle ${approle_name}"
echo "Using approle ID ${VAULT_APPROLE_ID}"
echo "Using wrapped token ${VAULT_WRAPPED_TOKEN}"
go test -v -timeout $TT_TIMEOUT
