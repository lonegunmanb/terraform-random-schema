package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const randomId = `{
  "block": {
    "attributes": {
      "b64_std": {
        "computed": true,
        "description": "The generated id presented in base64 without additional transformations.",
        "description_kind": "plain",
        "type": "string"
      },
      "b64_url": {
        "computed": true,
        "description": "The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters ` + "`" + `_` + "`" + ` and ` + "`" + `-` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "byte_length": {
        "description": "The number of random bytes to produce. The minimum value is 1, which produces eight bits of randomness.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "dec": {
        "computed": true,
        "description": "The generated id presented in non-padded decimal digits.",
        "description_kind": "plain",
        "type": "string"
      },
      "hex": {
        "computed": true,
        "description": "The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The generated id presented in base64 without additional transformations or prefix.",
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
      "prefix": {
        "description": "Arbitrary string to prefix the output value with. This string is supplied as-is, meaning it is not guaranteed to be URL-safe or base64 encoded.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "\nThe resource ` + "`" + `random_id` + "`" + ` generates random numbers that are intended to be\nused as unique identifiers for other resources. If the output is considered \nsensitive, and should not be displayed in the CLI, use ` + "`" + `random_bytes` + "`" + `\ninstead.\n\nThis resource *does* use a cryptographic random number generator in order\nto minimize the chance of collisions, making the results of this resource\nwhen a 16-byte identifier is requested of equivalent uniqueness to a\ntype-4 UUID.\n\nThis resource can be used in conjunction with resources that have\nthe ` + "`" + `create_before_destroy` + "`" + ` lifecycle flag set to avoid conflicts with\nunique names during the brief period where both the old and new resources\nexist concurrently.\n",
    "description_kind": "plain"
  },
  "version": 0
}`

func RandomIdSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(randomId), &result)
	return &result
}
