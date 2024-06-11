/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b></br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// GET Metros retrieves all Equinix® Fabric™ metros, as well as latency data for each location.This performance data helps network planning engineers and administrators make strategic decisions about port locations and traffic routes.
type Metro struct {
	// The Canonical URL at which the resource resides.
	Href string `json:"href,omitempty"`
	// Indicator of a Fabric Metro
	Type_ string `json:"type,omitempty"`
	// Code Assigned to an Equinix IBX data center in a specified metropolitan area.
	Code string `json:"code,omitempty"`
	// Board geographic area in which the data center is located
	Region string `json:"region,omitempty"`
	// Name of the region in which the data center is located.
	Name string `json:"name,omitempty"`
	// Autonomous system number (ASN) for a specified Fabric metro. The ASN is a unique identifier that carries the network routing protocol and exchanges that data with other internal systems via border gateway protocol.
	EquinixAsn int64 `json:"equinixAsn,omitempty"`
	// This field holds Max Connection speed with in the metro
	LocalVCBandwidthMax int64            `json:"localVCBandwidthMax,omitempty"`
	GeoCoordinates      *GeoCoordinates  `json:"geoCoordinates,omitempty"`
	ConnectedMetros     []ConnectedMetro `json:"connectedMetros,omitempty"`
	// List of supported geographic boundaries of a Fabric Metro.
	GeoScopes []GeoScopeType `json:"geoScopes,omitempty"`
}
