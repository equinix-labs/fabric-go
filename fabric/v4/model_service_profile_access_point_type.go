/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b> </br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Access Point Type
type ServiceProfileAccessPointType struct {
	SupportedBandwidths *[]int32 `json:"supportedBandwidths,omitempty"`
	// Setting to allow or prohibit remote connections to the service profile.
	AllowRemoteConnections bool `json:"allowRemoteConnections,omitempty"`
	// Setting to enable or disable the ability of the buyer to customize the bandwidth.
	AllowCustomBandwidth bool `json:"allowCustomBandwidth,omitempty"`
	// percentage of port bandwidth at which an allocation alert is generated - missing on wiki.
	BandwidthAlertThreshold float64 `json:"bandwidthAlertThreshold,omitempty"`
	// Setting to enable or disable the ability of the buyer to change connection bandwidth without approval of the seller.
	AllowBandwidthAutoApproval bool `json:"allowBandwidthAutoApproval,omitempty"`
	// Availability of a bandwidth upgrade. The default is false.
	AllowBandwidthUpgrade bool                              `json:"allowBandwidthUpgrade,omitempty"`
	LinkProtocolConfig    *ServiceProfileLinkProtocolConfig `json:"linkProtocolConfig,omitempty"`
	// for verizon only.
	EnableAutoGenerateServiceKey bool `json:"enableAutoGenerateServiceKey,omitempty"`
	// Mandate redundant connections
	ConnectionRedundancyRequired bool       `json:"connectionRedundancyRequired,omitempty"`
	ApiConfig                    *ApiConfig `json:"apiConfig,omitempty"`
	// custom name for \"Connection\"
	ConnectionLabel   string                             `json:"connectionLabel,omitempty"`
	AuthenticationKey *AuthenticationKey                 `json:"authenticationKey,omitempty"`
	Metadata          *ServiceProfileMetadata            `json:"metadata,omitempty"`
	Type_             *ServiceProfileAccessPointTypeEnum `json:"type"`
	Uuid              string                             `json:"uuid,omitempty"`
}
