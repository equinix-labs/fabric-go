/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b> </br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

type RouteFiltersData struct {
	// Route Filter URI
	Href string `json:"href,omitempty"`
	// Route Filter type
	Type_ string `json:"type,omitempty"`
	// Route Filter identifier
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	// Customer-provided connection description
	Description          string                   `json:"description,omitempty"`
	State                *RouteFilterState        `json:"state,omitempty"`
	Change               *RouteFiltersChange      `json:"change,omitempty"`
	NotMatchedRuleAction string                   `json:"notMatchedRuleAction,omitempty"`
	ConnectionsCount     int32                    `json:"connectionsCount,omitempty"`
	RulesCount           int32                    `json:"rulesCount,omitempty"`
	Project              *RouteFiltersDataProject `json:"project,omitempty"`
	Changelog            *Changelog               `json:"changelog,omitempty"`
}
