package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const randomUuid = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "The generated uuid presented in string format.",
        "description_kind": "plain",
        "type": "string"
      },
      "keepers": {
        "description": "Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider documentation](../index.html) for more information.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "result": {
        "computed": true,
        "description": "The generated uuid presented in string format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The resource ` + "`" + `random_uuid` + "`" + ` generates a random uuid string that is intended to be used as a unique identifier for other resources.\n\nThis resource uses [hashicorp/go-uuid](https://github.com/hashicorp/go-uuid) to generate a UUID-formatted string for use with services needing a unique string identifier.",
    "description_kind": "plain"
  },
  "version": 0
}`

func RandomUuidSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(randomUuid), &result)
	return &result
}
