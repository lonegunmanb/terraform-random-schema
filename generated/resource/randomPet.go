package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const randomPet = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "The random pet name.",
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
      "length": {
        "computed": true,
        "description": "The length (in words) of the pet name. Defaults to 2",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "prefix": {
        "description": "A string to prefix the name with.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "separator": {
        "computed": true,
        "description": "The character to separate words in the pet name. Defaults to \"-\"",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The resource ` + "`" + `random_pet` + "`" + ` generates random pet names that are intended to be used as unique identifiers for other resources.\n\nThis resource can be used in conjunction with resources that have the ` + "`" + `create_before_destroy` + "`" + ` lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.",
    "description_kind": "plain"
  },
  "version": 0
}`

func RandomPetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(randomPet), &result)
	return &result
}
