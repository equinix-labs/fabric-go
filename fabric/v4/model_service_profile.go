/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b></br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Service Profile is a software definition for a named provider service and it's network connectivity requirements. This includes the basic marketing information and one or more sets of access points (a set per each access point type) fulfilling the provider service.
type ServiceProfile struct {
	// Service Profile URI response attribute
	Href  string                  `json:"href,omitempty"`
	Type_ *ServiceProfileTypeEnum `json:"type,omitempty"`
	// Customer-assigned service profile name
	Name string `json:"name,omitempty"`
	// Equinix-assigned service profile identifier
	Uuid string `json:"uuid,omitempty"`
	// User-provided service description should be of maximum length 375
	Description string `json:"description,omitempty"`
	// Recipients of notifications on service profile change
	Notifications          []SimplifiedNotification        `json:"notifications,omitempty"`
	Tags                   *[]string                       `json:"tags,omitempty"`
	Visibility             *ServiceProfileVisibilityEnum   `json:"visibility,omitempty"`
	AllowedEmails          []string                        `json:"allowedEmails,omitempty"`
	AccessPointTypeConfigs []ServiceProfileAccessPointType `json:"accessPointTypeConfigs,omitempty"`
	CustomFields           []CustomField                   `json:"customFields,omitempty"`
	MarketingInfo          *MarketingInfo                  `json:"marketingInfo,omitempty"`
	Ports                  []ServiceProfileAccessPointColo `json:"ports,omitempty"`
	VirtualDevices         []ServiceProfileAccessPointVd   `json:"virtualDevices,omitempty"`
	// Derived response attribute.
	Metros []ServiceMetro `json:"metros,omitempty"`
	// response attribute indicates whether the profile belongs to the same organization as the api-invoker.
	SelfProfile bool                     `json:"selfProfile,omitempty"`
	ProjectId   string                   `json:"projectId,omitempty"`
	State       *ServiceProfileStateEnum `json:"state,omitempty"`
	// Seller Account for Service Profile.
	Account *AllOfServiceProfileAccount `json:"account,omitempty"`
	Project *Project                    `json:"project,omitempty"`
	// Seller Account for Service Profile.
	ChangeLog *AllOfServiceProfileChangeLog `json:"changeLog,omitempty"`
}
