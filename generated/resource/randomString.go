package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const randomString = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "The generated random string.",
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
        "description": "The length of the string desired. The minimum value for length is 1 and, length must also be \u003e= (` + "`" + `min_upper` + "`" + ` + ` + "`" + `min_lower` + "`" + ` + ` + "`" + `min_numeric` + "`" + ` + ` + "`" + `min_special` + "`" + `).",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "lower": {
        "computed": true,
        "description": "Include lowercase alphabet characters in the result. Default value is ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "min_lower": {
        "computed": true,
        "description": "Minimum number of lowercase alphabet characters in the result. Default value is ` + "`" + `0` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_numeric": {
        "computed": true,
        "description": "Minimum number of numeric characters in the result. Default value is ` + "`" + `0` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_special": {
        "computed": true,
        "description": "Minimum number of special characters in the result. Default value is ` + "`" + `0` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_upper": {
        "computed": true,
        "description": "Minimum number of uppercase alphabet characters in the result. Default value is ` + "`" + `0` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "number": {
        "computed": true,
        "deprecated": true,
        "description": "Include numeric characters in the result. Default value is ` + "`" + `true` + "`" + `. If ` + "`" + `number` + "`" + `, ` + "`" + `upper` + "`" + `, ` + "`" + `lower` + "`" + `, and ` + "`" + `special` + "`" + ` are all configured, at least one of them must be set to ` + "`" + `true` + "`" + `. **NOTE**: This is deprecated, use ` + "`" + `numeric` + "`" + ` instead.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "numeric": {
        "computed": true,
        "description": "Include numeric characters in the result. Default value is ` + "`" + `true` + "`" + `. If ` + "`" + `numeric` + "`" + `, ` + "`" + `upper` + "`" + `, ` + "`" + `lower` + "`" + `, and ` + "`" + `special` + "`" + ` are all configured, at least one of them must be set to ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "override_special": {
        "description": "Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The ` + "`" + `special` + "`" + ` argument must still be set to true for any overwritten characters to be used in generation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "result": {
        "computed": true,
        "description": "The generated random string.",
        "description_kind": "plain",
        "type": "string"
      },
      "special": {
        "computed": true,
        "description": "Include special characters in the result. These are ` + "`" + `!@#$%\u0026*()-_=+[]{}\u003c\u003e:?` + "`" + `. Default value is ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "upper": {
        "computed": true,
        "description": "Include uppercase alphabet characters in the result. Default value is ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "The resource ` + "`" + `random_string` + "`" + ` generates a random permutation of alphanumeric characters and optionally special characters.\n\nThis resource *does* use a cryptographic random number generator.\n\nHistorically this resource's intended usage has been ambiguous as the original example used it in a password. For backwards compatibility it will continue to exist. For unique ids please use [random_id](id.html), for sensitive random values please use [random_password](password.html).",
    "description_kind": "plain"
  },
  "version": 2
}`

func RandomStringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(randomString), &result)
	return &result
}
