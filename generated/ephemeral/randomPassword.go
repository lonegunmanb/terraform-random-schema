package ephemeral

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
    "description": "-\u003e If the managed resource doesn't have a write-only argument available for the password (first introduced in Terraform 1.11), then the password can only be created with the managed resource variant of [` + "`" + `random_password` + "`" + `](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/password).\n\nGenerates an ephemeral password string using a cryptographic random number generator.\n\nThe primary use-case for generating an ephemeral random password is to be used in combination with a write-only argument in a managed resource, which will avoid Terraform storing the password string in the plan or state file.",
    "description_kind": "plain"
  },
  "version": 0
}`

func RandomPasswordSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(randomPassword), &result)
	return &result
}
