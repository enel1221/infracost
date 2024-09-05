package aws

import (
	"github.com/infracost/infracost/internal/resources/aws"
	"github.com/infracost/infracost/internal/schema"
)

func getCloudFormationStackRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "cloudformation.aws.upbound.io/Stack", // Updated name for Crossplane
		CoreRFunc: NewCloudFormationStack,
	}
}

func NewCloudFormationStack(d *schema.ResourceData) schema.CoreResource {
	forProvider := d.Get("forProvider")   // Get the forProvider field, which holds the resource configuration in Crossplane
	region := lookupRegion(d, []string{}) // Custom function to lookup the region, could be adjusted based on your environment
	templateBody := forProvider.Get("templateBody").String()

	r := &aws.CloudFormationStack{
		Address:      d.Address,
		Region:       region,
		TemplateBody: templateBody,
	}
	return r
}
