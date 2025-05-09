package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	
	"github.com/lonegunmanb/terraform-random-schema/v3/generated/resource"
	"github.com/lonegunmanb/terraform-random-schema/v3/generated/ephemeral"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema
var EphemeralResources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)
	ephemeralResources := make(map[string]*tfjson.Schema, 0)  
	resources["random_bytes"] = resource.RandomBytesSchema()  
	resources["random_id"] = resource.RandomIdSchema()  
	resources["random_integer"] = resource.RandomIntegerSchema()  
	resources["random_password"] = resource.RandomPasswordSchema()  
	resources["random_pet"] = resource.RandomPetSchema()  
	resources["random_shuffle"] = resource.RandomShuffleSchema()  
	resources["random_string"] = resource.RandomStringSchema()  
	resources["random_uuid"] = resource.RandomUuidSchema()  
	ephemeralResources["random_password"] = ephemeral.RandomPasswordSchema()  
	Resources = resources
	DataSources = dataSources
    EphemeralResources = ephemeralResources
}