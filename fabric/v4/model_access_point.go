/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b> </br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Access point object
type AccessPoint struct {
	Type_         *AccessPointType          `json:"type,omitempty"`
	Account       *SimplifiedAccount        `json:"account,omitempty"`
	Location      *SimplifiedLocation       `json:"location,omitempty"`
	Port          *SimplifiedPort           `json:"port,omitempty"`
	Profile       *SimplifiedServiceProfile `json:"profile,omitempty"`
	Router        *CloudRouter              `json:"router,omitempty"`
	LinkProtocol  *SimplifiedLinkProtocol   `json:"linkProtocol,omitempty"`
	VirtualDevice *VirtualDevice            `json:"virtualDevice,omitempty"`
	Interface_    *ModelInterface           `json:"interface,omitempty"`
	Network       *SimplifiedNetwork        `json:"network,omitempty"`
	// Access point seller region
	SellerRegion string       `json:"sellerRegion,omitempty"`
	PeeringType  *PeeringType `json:"peeringType,omitempty"`
	// Access point authentication key
	AuthenticationKey string `json:"authenticationKey,omitempty"`
	// Provider assigned Connection Id
	ProviderConnectionId string                `json:"providerConnectionId,omitempty"`
	VirtualNetwork       *VirtualNetwork       `json:"virtualNetwork,omitempty"`
	Interconnection      *MetalInterconnection `json:"interconnection,omitempty"`
	VpicInterface        *VpicInterface        `json:"vpic_interface,omitempty"`
}
