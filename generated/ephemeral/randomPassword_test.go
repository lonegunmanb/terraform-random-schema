package ephemeral_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-random-schema/v3/generated/ephemeral"
	"github.com/stretchr/testify/assert"
)

func TestRandomPasswordSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := ephemeral.RandomPasswordSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
