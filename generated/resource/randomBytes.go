package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const randomBytes = `{
  "block": {
    "attributes": {
      "base64": {
        "computed": true,
        "description": "The generated bytes presented in base64 string format.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "hex": {
        "computed": true,
        "description": "The generated bytes presented in lowercase hexadecimal string format. The length of the encoded string is exactly twice the ` + "`" + `length` + "`" + ` parameter.",
        "description_kind": "plain",
        "sensitive": true,
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
      "length": {
        "description": "The number of bytes requested. The minimum value for length is 1.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description": "The resource ` + "`" + `random_bytes` + "`" + ` generates random bytes that are intended to be used as a secret, or key. Use this in preference to ` + "`" + `random_id` + "`" + ` when the output is considered sensitive, and should not be displayed in the CLI.\n\nThis resource *does* use a cryptographic random number generator.",
    "description_kind": "plain"
  },
  "version": 0
}`

func RandomBytesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(randomBytes), &result)
	return &result
}
