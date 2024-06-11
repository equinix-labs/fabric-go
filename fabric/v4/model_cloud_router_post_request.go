/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b></br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Create Cloud Router
type CloudRouterPostRequest struct {
	Type_ string `json:"type,omitempty"`
	// Customer-provided Cloud Router name
	Name     string                         `json:"name,omitempty"`
	Location *SimplifiedLocationWithoutIbx  `json:"location,omitempty"`
	Package_ *CloudRouterPostRequestPackage `json:"package,omitempty"`
	Order    *Order                         `json:"order,omitempty"`
	Project  *Project                       `json:"project,omitempty"`
	Account  *SimplifiedAccount             `json:"account,omitempty"`
	// Preferences for notifications on connection configuration or status changes
	Notifications []SimplifiedNotification `json:"notifications,omitempty"`
}
