package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const randomPassword = `{
  "block": {
    "attributes": {
      "bcrypt_hash": {
        "computed": true,
        "description": "A bcrypt hash of the generated random string. **NOTE**: If the generated random string is greater than 72 bytes in length, ` + "`" + `bcrypt_hash` + "`" + ` will contain a hash of the first 72 bytes.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "A static value used internally by Terraform, this should not be referenced in configurations.",
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
        "sensitive": true,
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
    "description": "Identical to [random_string](string.html) with the exception that the result is treated as sensitive and, thus, _not_ displayed in console output. Read more about sensitive data handling in the [Terraform documentation](https://www.terraform.io/docs/language/state/sensitive-data.html).\n\nThis resource *does* use a cryptographic random number generator.",
    "description_kind": "plain"
  },
  "version": 3
}`

func RandomPasswordSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(randomPassword), &result)
	return &result
}
