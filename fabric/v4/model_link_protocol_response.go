/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.7
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Link Protocol response
type LinkProtocolResponse struct {
	// Equinix-assigned network identifier
	Uuid           string                    `json:"uuid,omitempty"`
	State          *LinkProtocolState        `json:"state,omitempty"`
	Type_          *LinkProtocolRequestType  `json:"type,omitempty"`
	VlanTag        int32                     `json:"vlanTag,omitempty"`
	VlanTagMin     int32                     `json:"vlanTagMin,omitempty"`
	VlanTagMax     int32                     `json:"vlanTagMax,omitempty"`
	VlanSTag       int32                     `json:"vlanSTag,omitempty"`
	VlanCTag       int32                     `json:"vlanCTag,omitempty"`
	VlanCTagMin    int32                     `json:"vlanCTagMin,omitempty"`
	VlanCTagMax    int32                     `json:"vlanCTagMax,omitempty"`
	SubInterface   *SubInterface             `json:"subInterface,omitempty"`
	AdditionalInfo []AdditionalInfo          `json:"additionalInfo,omitempty"`
	Connection     *LinkProtocolConnection   `json:"connection,omitempty"`
	ServiceToken   *LinkProtocolServiceToken `json:"serviceToken,omitempty"`
	ChangeLog      *Changelog                `json:"changeLog,omitempty"`
}