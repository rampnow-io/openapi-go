package openapi3_test

import (
	"testing"

	"github.com/rampnow-io/openapi-go/openapi3"
	"github.com/stretchr/testify/require"
)

func TestSpec_MarshalYAML(t *testing.T) {
	var s openapi3.Spec

	spec := `openapi: 3.0.3
info:
  description: description
  license:
    name: Apache-2.0
    url: https://www.apache.org/licenses/LICENSE-2.0.html
  title: title
  version: 2.0.0
servers:
  - url: /v2
paths:
  /user:
    put:
      summary: updates the user by id
      operationId: UpdateUser
      requestBody:
        content:
          application/json:
            schema:
              type: string
        description: Updated user object
        required: true
      responses:
        "404":
          description: User not found
components:
  securitySchemes:
    api_key:
      in: header
      name: x-api-key
      type: apiKey
    bearer_auth:
      type: http
      scheme: bearer
      bearerFormat: JWT`

	require.NoError(t, s.UnmarshalYAML([]byte(spec)))
}

func TestSpec_MarshalYAML_2(t *testing.T) {
	var s openapi3.Spec

	spec := `openapi: 3.0.0
info:
  title: MyProject
  description: "My Project Description"
  version: v1.0.0
# 1) Define the security scheme type (HTTP bearer)
components:
  securitySchemes:
    bearerAuth: # arbitrary name for the security scheme
      type: http
      scheme: bearer
      bearerFormat: JWT # optional, arbitrary value for documentation purposes
# 2) Apply the security globally to all operations
security:
  - bearerAuth: [] # use the same name as above
paths:
`

	require.NoError(t, s.UnmarshalYAML([]byte(spec)))
}

func TestSpec_MarshalYAML_3(t *testing.T) {
	var s openapi3.Spec

	spec := `openapi: 3.0.3
info:
  title: MyProject
  description: "My Project Description"
  version: v1.0.0
components:
  securitySchemes:
    basicAuth: # <-- arbitrary name for the security scheme
      type: http
      scheme: basic
security:
  - basicAuth: [] # <-- use the same name here  
paths:
`

	require.NoError(t, s.UnmarshalYAML([]byte(spec)))
}
