/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b> </br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// SortBy : Possible field names to use on sorting
type SortBy string

// List of SortBy
const (
	NAME_SortBy                                      SortBy = "/name"
	DIRECTION_SortBy                                 SortBy = "/direction"
	A_SIDEACCESS_POINTNAME_SortBy                    SortBy = "/aSide/accessPoint/name"
	A_SIDEACCESS_POINTTYPE_SortBy                    SortBy = "/aSide/accessPoint/type"
	A_SIDEACCESS_POINTACCOUNTACCOUNT_NAME_SortBy     SortBy = "/aSide/accessPoint/account/accountName"
	A_SIDEACCESS_POINTLOCATIONMETRO_NAME_SortBy      SortBy = "/aSide/accessPoint/location/metroName"
	A_SIDEACCESS_POINTLOCATIONMETRO_CODE_SortBy      SortBy = "/aSide/accessPoint/location/metroCode"
	A_SIDEACCESS_POINTLINK_PROTOCOLVLAN_C_TAG_SortBy SortBy = "/aSide/accessPoint/linkProtocol/vlanCTag"
	A_SIDEACCESS_POINTLINK_PROTOCOLVLAN_S_TAG_SortBy SortBy = "/aSide/accessPoint/linkProtocol/vlanSTag"
	Z_SIDEACCESS_POINTNAME_SortBy                    SortBy = "/zSide/accessPoint/name"
	Z_SIDEACCESS_POINTTYPE_SortBy                    SortBy = "/zSide/accessPoint/type"
	Z_SIDEACCESS_POINTACCOUNTACCOUNT_NAME_SortBy     SortBy = "/zSide/accessPoint/account/accountName"
	Z_SIDEACCESS_POINTLOCATIONMETRO_NAME_SortBy      SortBy = "/zSide/accessPoint/location/metroName"
	Z_SIDEACCESS_POINTLOCATIONMETRO_CODE_SortBy      SortBy = "/zSide/accessPoint/location/metroCode"
	Z_SIDEACCESS_POINTLINK_PROTOCOLVLAN_C_TAG_SortBy SortBy = "/zSide/accessPoint/linkProtocol/vlanCTag"
	Z_SIDEACCESS_POINTLINK_PROTOCOLVLAN_S_TAG_SortBy SortBy = "/zSide/accessPoint/linkProtocol/vlanSTag"
	Z_SIDEACCESS_POINTAUTHENTICATION_KEY_SortBy      SortBy = "/zSide/accessPoint/authenticationKey"
	BANDWIDTH_SortBy                                 SortBy = "/bandwidth"
	GEO_SCOPE_SortBy                                 SortBy = "/geoScope"
	UUID_SortBy                                      SortBy = "/uuid"
	CHANGE_LOGCREATED_DATE_TIME_SortBy               SortBy = "/changeLog/createdDateTime"
	CHANGE_LOGUPDATED_DATE_TIME_SortBy               SortBy = "/changeLog/updatedDateTime"
	OPERATIONEQUINIX_STATUS_SortBy                   SortBy = "/operation/equinixStatus"
	OPERATIONPROVIDER_STATUS_SortBy                  SortBy = "/operation/providerStatus"
	REDUNDANCYPRIORITY_SortBy                        SortBy = "/redundancy/priority"
)
