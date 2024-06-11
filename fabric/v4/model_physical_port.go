/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b></br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Physical Port specification
type PhysicalPort struct {
	Type_ *PhysicalPortType `json:"type,omitempty"`
	// Equinix assigned response attribute for Physical Port Id
	Id int32 `json:"id,omitempty"`
	// Equinix assigned response attribute for an absolute URL that is the subject of the link's context.
	Href    string             `json:"href,omitempty"`
	State   *PortState         `json:"state,omitempty"`
	Account *SimplifiedAccount `json:"account,omitempty"`
	// Physical Port Speed in Mbps
	InterfaceSpeed int32 `json:"interfaceSpeed,omitempty"`
	// Physical Port Interface Type
	InterfaceType    string                `json:"interfaceType,omitempty"`
	Tether           *PortTether           `json:"tether,omitempty"`
	DemarcationPoint *PortDemarcationPoint `json:"demarcationPoint,omitempty"`
	// Physical Port additional information
	AdditionalInfo []PortAdditionalInfo `json:"additionalInfo,omitempty"`
	Order          *PortOrder           `json:"order,omitempty"`
	Operation      *PortOperation       `json:"operation,omitempty"`
	// Port Loas
	Loas []PortLoa `json:"loas,omitempty"`
}
