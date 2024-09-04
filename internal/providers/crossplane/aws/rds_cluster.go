package aws

import (
	"github.com/infracost/infracost/internal/resources/aws"
	"github.com/infracost/infracost/internal/schema"
)

func getRDSClusterRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "rds.aws.upbound.io/Cluster",
		CoreRFunc: NewRDSCluster,
	}
}

func NewRDSCluster(d *schema.ResourceData) schema.CoreResource {
	forProvider := d.Get("forProvider")
	
	engineMode := d.GetStringOrDefault("engine_mode", "provisioned")
	r := &aws.RDSCluster{
		Address:               d.Address,
		Region:                forProvider.Get("region").String(),
		Engine:                d.GetStringOrDefault("forProvider.engine", "aurora"),
		BackupRetentionPeriod: d.GetInt64OrDefault("forProvider.backupRetentionPeriod", 1),
		EngineMode:            engineMode,
		IOOptimized:           forProvider.Get("storageType").String() == "aurora-iopt1",
	}
	return r
}
