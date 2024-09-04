package aws

import (
	"github.com/infracost/infracost/internal/logging"
	"github.com/infracost/infracost/internal/schema"
	"github.com/shopspring/decimal"
)

// Utility function to convert a string to a pointer
func strPtr(s string) *string {
	return &s
}

// Utility function to convert a decimal to a pointer
func decimalPtr(d decimal.Decimal) *decimal.Decimal {
	return &d
}

// Utility function to convert an int64 to a pointer
func intPtr(i int64) *int64 {
	return &i
}

// Utility function to convert a float64 to a pointer
func floatPtr(f float64) *float64 {
	return &f
}

// Utility function to look up the region for a given AWS resource
func lookupRegion(d *schema.ResourceData, parentResourceKeys []string) string {
	// Check if the region is directly set on the resource
	if d.Get("region").String() != "" {
		return d.Get("region").String()
	}

	// Check for region in forProvider block
	if d.Get("forProvider.region").String() != "" {
		return d.Get("forProvider.region").String()
	}

	// If region is not found, use the default region
	defaultRegion := d.Get("default_region").String()
	logging.Logger.Debug().Msgf("Using %s for resource %s as its 'region' property could not be found.", defaultRegion, d.Address)
	return defaultRegion
}

// Converts AWS region to a human-readable name
func convertRegion(region string) string {
	regionNameMap := map[string]string{
		"us-east-1":      "US East (N. Virginia)",
		"us-east-2":      "US East (Ohio)",
		"us-west-1":      "US West (N. California)",
		"us-west-2":      "US West (Oregon)",
		"ap-south-1":     "Asia Pacific (Mumbai)",
		"ap-northeast-3": "Asia Pacific (Osaka)",
		"ap-northeast-2": "Asia Pacific (Seoul)",
		"ap-southeast-1": "Asia Pacific (Singapore)",
		"ap-southeast-2": "Asia Pacific (Sydney)",
		"ap-northeast-1": "Asia Pacific (Tokyo)",
		"ca-central-1":   "Canada (Central)",
		"eu-central-1":   "EU (Frankfurt)",
		"eu-west-1":      "EU (Ireland)",
		"eu-west-2":      "EU (London)",
		"eu-west-3":      "EU (Paris)",
		"eu-north-1":     "EU (Stockholm)",
		"sa-east-1":      "South America (São Paulo)",
		"us-gov-west-1":  "AWS GovCloud (US-West)",
		"us-gov-east-1":  "AWS GovCloud (US-East)",
		"me-south-1":     "Middle East (Bahrain)",
		"af-south-1":     "Africa (Cape Town)",
	}

	if name, ok := regionNameMap[region]; ok {
		return name
	}

	return region
}

// regionToVPCZone maps AWS regions to their respective VPC zone
func regionToVPCZone(region string) string {
	zoneMap := map[string]string{
		"us-east-1":      "Zone 1",
		"us-east-2":      "Zone 1",
		"us-west-1":      "Zone 2",
		"us-west-2":      "Zone 2",
		"ap-south-1":     "Zone 3",
		"ap-northeast-3": "Zone 3",
		"ap-northeast-2": "Zone 3",
		"ap-southeast-1": "Zone 3",
		"ap-southeast-2": "Zone 3",
		"ap-northeast-1": "Zone 3",
		"ca-central-1":   "Zone 1",
		"eu-central-1":   "Zone 1",
		"eu-west-1":      "Zone 1",
		"eu-west-2":      "Zone 1",
		"eu-west-3":      "Zone 1",
		"eu-north-1":     "Zone 1",
		"sa-east-1":      "Zone 4",
		"us-gov-west-1":  "GovCloud",
		"us-gov-east-1":  "GovCloud",
		"me-south-1":     "Zone 3",
		"af-south-1":     "Zone 3",
	}

	if zone, ok := zoneMap[region]; ok {
		return zone
	}

	return "Global"
}

// Checks if a string exists within a slice of strings
func contains(arr []string, e string) bool {
	for _, a := range arr {
		if a == e {
			return true
		}
	}
	return false
}
