package aws

import (
	"github.com/infracost/infracost/internal/resources/aws"
	"github.com/infracost/infracost/internal/schema"
)

func getEC2HostRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "ec2.aws.upbound.io/Instance",
		CoreRFunc: newEC2Host,
	}
}

func newEC2Host(d *schema.ResourceData) schema.CoreResource {
	
	forProvider := d.Get("forProvider")
	region := lookupRegion(d, []string{})
	instanceType := forProvider.Get("instanceType").String()

	r := &aws.EC2Host{
		Address:        d.Address,
		Region:         region,
		InstanceType:   instanceType,
		// InstanceFamily: d.Get("instance_family").String(),
	}
	return r
}
