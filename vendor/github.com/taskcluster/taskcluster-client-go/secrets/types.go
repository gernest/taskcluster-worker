// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package secrets

import (
	"encoding/json"

	tcclient "github.com/taskcluster/taskcluster-client-go"
)

type (
	// Message containing a TaskCluster Secret
	//
	// See http://schemas.taskcluster.net/secrets/v1/secret.json#
	Secret struct {

		// An expiration date for this secret.
		//
		// See http://schemas.taskcluster.net/secrets/v1/secret.json#/properties/expires
		Expires tcclient.Time `json:"expires"`

		// The secret value to be encrypted.
		//
		// See http://schemas.taskcluster.net/secrets/v1/secret.json#/properties/secret
		Secret json.RawMessage `json:"secret"`
	}

	// Message containing a list of secret names
	//
	// See http://schemas.taskcluster.net/secrets/v1/secret-list.json#
	SecretsList struct {

		// Secret names
		//
		// See http://schemas.taskcluster.net/secrets/v1/secret-list.json#/properties/secrets
		Secrets []string `json:"secrets"`
	}
)
