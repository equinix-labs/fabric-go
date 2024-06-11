/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b></br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

import (
	"time"
)

// Current state of latest Routing Protocol change
type RoutingProtocolChangeData struct {
	// Uniquely identifies a change
	Uuid string `json:"uuid"`
	// Type of change
	Type_ string `json:"type"`
	// Routing Protocol Change URI
	Href string `json:"href,omitempty"`
	// Current outcome of the change flow
	Status string `json:"status,omitempty"`
	// Created by User Key
	CreatedBy string `json:"createdBy,omitempty"`
	// Set when change flow starts
	CreatedDateTime time.Time `json:"createdDateTime,omitempty"`
	// Updated by User Key
	UpdatedBy string `json:"updatedBy,omitempty"`
	// Set when change object is updated
	UpdatedDateTime time.Time `json:"updatedDateTime,omitempty"`
	// Additional information
	Information string                          `json:"information,omitempty"`
	Data        *RoutingProtocolChangeOperation `json:"data,omitempty"`
}
