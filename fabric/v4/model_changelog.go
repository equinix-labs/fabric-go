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

// Change log
type Changelog struct {
	// Created by User Key
	CreatedBy string `json:"createdBy,omitempty"`
	// Created by User Full Name
	CreatedByFullName string `json:"createdByFullName,omitempty"`
	// Created by User Email Address
	CreatedByEmail string `json:"createdByEmail,omitempty"`
	// Created by Date and Time
	CreatedDateTime time.Time `json:"createdDateTime,omitempty"`
	// Updated by User Key
	UpdatedBy string `json:"updatedBy,omitempty"`
	// Updated by User Full Name
	UpdatedByFullName string `json:"updatedByFullName,omitempty"`
	// Updated by User Email Address
	UpdatedByEmail string `json:"updatedByEmail,omitempty"`
	// Updated by Date and Time
	UpdatedDateTime time.Time `json:"updatedDateTime,omitempty"`
	// Deleted by User Key
	DeletedBy string `json:"deletedBy,omitempty"`
	// Deleted by User Full Name
	DeletedByFullName string `json:"deletedByFullName,omitempty"`
	// Deleted by User Email Address
	DeletedByEmail string `json:"deletedByEmail,omitempty"`
	// Deleted by Date and Time
	DeletedDateTime time.Time `json:"deletedDateTime,omitempty"`
}
