package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const randomShuffle = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "A static value used internally by Terraform, this should not be referenced in configurations.",
        "description_kind": "plain",
        "type": "string"
      },
      "input": {
        "description": "The list of strings to shuffle.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
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
        "description": "Random permutation of the list of strings given in ` + "`" + `input` + "`" + `. The number of elements is determined by ` + "`" + `result_count` + "`" + ` if set, or the number of elements in ` + "`" + `input` + "`" + `.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "result_count": {
        "description": "The number of results to return. Defaults to the number of items in the ` + "`" + `input` + "`" + ` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "seed": {
        "description": "Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.\n\n**Important:** Even with an identical seed, it is not guaranteed that the same permutation will be produced across different versions of Terraform. This argument causes the result to be *less volatile*, but not fixed for all time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The resource ` + "`" + `random_shuffle` + "`" + ` generates a random permutation of a list of strings given as an argument.",
    "description_kind": "plain"
  },
  "version": 0
}`

func RandomShuffleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(randomShuffle), &result)
	return &result
}
