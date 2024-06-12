/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b> </br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Metadata. Response attribute. Ignored on request payload.
type ServiceProfileMetadata struct {
	Props                  string `json:"props,omitempty"`
	RegEx                  string `json:"regEx,omitempty"`
	RegExMsg               string `json:"regExMsg,omitempty"`
	VlanRangeMaxValue      int32  `json:"vlanRangeMaxValue,omitempty"`
	VlanRangeMinValue      int32  `json:"vlanRangeMinValue,omitempty"`
	MaxQinq                string `json:"maxQinq,omitempty"`
	MaxDot1q               int32  `json:"maxDot1q,omitempty"`
	VariableBilling        bool   `json:"variableBilling,omitempty"`
	GlobalOrganization     string `json:"globalOrganization,omitempty"`
	LimitAuthKeyConn       bool   `json:"limitAuthKeyConn,omitempty"`
	AllowSecondaryLocation bool   `json:"allowSecondaryLocation,omitempty"`
	RedundantProfileId     string `json:"redundantProfileId,omitempty"`
	AllowVcMigration       bool   `json:"allowVcMigration,omitempty"`
	ConnectionEditable     bool   `json:"connectionEditable,omitempty"`
	ReleaseVlan            bool   `json:"releaseVlan,omitempty"`
	MaxConnectionsOnPort   int32  `json:"maxConnectionsOnPort,omitempty"`
	PortAssignmentStrategy string `json:"portAssignmentStrategy,omitempty"`
	EqxManagedPort         bool   `json:"eqxManagedPort,omitempty"`
	ConnectionNameEditable bool   `json:"connectionNameEditable,omitempty"`
}
