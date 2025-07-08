package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-random-schema/v3/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestRandomPasswordSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.RandomPasswordSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
