package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const randomInteger = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "The string representation of the integer result.",
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
      "max": {
        "description": "The maximum inclusive value of the range.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "min": {
        "description": "The minimum inclusive value of the range.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "result": {
        "computed": true,
        "description": "The random integer result.",
        "description_kind": "plain",
        "type": "number"
      },
      "seed": {
        "description": "A custom seed to always produce the same value.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The resource ` + "`" + `random_integer` + "`" + ` generates random values from a given range, described by the ` + "`" + `min` + "`" + ` and ` + "`" + `max` + "`" + ` attributes of a given resource.\n\nThis resource can be used in conjunction with resources that have the ` + "`" + `create_before_destroy` + "`" + ` lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.",
    "description_kind": "plain"
  },
  "version": 0
}`

func RandomIntegerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(randomInteger), &result)
	return &result
}
