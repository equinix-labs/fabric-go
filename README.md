# Go API client for v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 4.13
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://equinix.com/about/](https://equinix.com/about/)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "github.com/equinix-labs/fabric-go/fabric/v4"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.equinix.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CloudRoutersApi* | [**CreateCloudRouter**](docs/CloudRoutersApi.md#createcloudrouter) | **Post** /fabric/v4/routers | Create Routers
*CloudRoutersApi* | [**CreateCloudRouterAction**](docs/CloudRoutersApi.md#createcloudrouteraction) | **Post** /fabric/v4/routers/{routerId}/actions | Route table actions
*CloudRoutersApi* | [**DeleteCloudRouterByUuid**](docs/CloudRoutersApi.md#deletecloudrouterbyuuid) | **Delete** /fabric/v4/routers/{routerId} | Delete Routers
*CloudRoutersApi* | [**GetCloudRouterActions**](docs/CloudRoutersApi.md#getcloudrouteractions) | **Get** /fabric/v4/routers/{routerId}/actions | Get actions
*CloudRoutersApi* | [**GetCloudRouterByUuid**](docs/CloudRoutersApi.md#getcloudrouterbyuuid) | **Get** /fabric/v4/routers/{routerId} | Get Routers
*CloudRoutersApi* | [**GetCloudRouterPackageByCode**](docs/CloudRoutersApi.md#getcloudrouterpackagebycode) | **Get** /fabric/v4/routerPackages/{routerPackageCode} | Get Package Details
*CloudRoutersApi* | [**GetCloudRouterPackages**](docs/CloudRoutersApi.md#getcloudrouterpackages) | **Get** /fabric/v4/routerPackages | List Packages
*CloudRoutersApi* | [**SearchCloudRouterRoutes**](docs/CloudRoutersApi.md#searchcloudrouterroutes) | **Post** /fabric/v4/routers/{routerId}/routes/search | Search Route Table
*CloudRoutersApi* | [**SearchCloudRouters**](docs/CloudRoutersApi.md#searchcloudrouters) | **Post** /fabric/v4/routers/search | Search Routers
*CloudRoutersApi* | [**UpdateCloudRouterByUuid**](docs/CloudRoutersApi.md#updatecloudrouterbyuuid) | **Patch** /fabric/v4/routers/{routerId} | Update Routers
*ConnectionsApi* | [**CreateConnection**](docs/ConnectionsApi.md#createconnection) | **Post** /fabric/v4/connections | Create Connection
*ConnectionsApi* | [**CreateConnectionAction**](docs/ConnectionsApi.md#createconnectionaction) | **Post** /fabric/v4/connections/{connectionId}/actions | Connection Actions
*ConnectionsApi* | [**DeleteConnectionByUuid**](docs/ConnectionsApi.md#deleteconnectionbyuuid) | **Delete** /fabric/v4/connections/{connectionId} | Delete by ID
*ConnectionsApi* | [**GetConnectionByUuid**](docs/ConnectionsApi.md#getconnectionbyuuid) | **Get** /fabric/v4/connections/{connectionId} | Get Connection by ID
*ConnectionsApi* | [**SearchConnections**](docs/ConnectionsApi.md#searchconnections) | **Post** /fabric/v4/connections/search | Search connections
*ConnectionsApi* | [**UpdateConnectionByUuid**](docs/ConnectionsApi.md#updateconnectionbyuuid) | **Patch** /fabric/v4/connections/{connectionId} | Update by ID
*ConnectionsApi* | [**ValidateConnections**](docs/ConnectionsApi.md#validateconnections) | **Post** /fabric/v4/connections/validate | Validate Connection
*HealthApi* | [**GetStatus**](docs/HealthApi.md#getstatus) | **Get** /fabric/v4/health | Get service status
*MetrosApi* | [**GetMetroByCode**](docs/MetrosApi.md#getmetrobycode) | **Get** /fabric/v4/metros/{metroCode} | Get Metro by Code
*MetrosApi* | [**GetMetros**](docs/MetrosApi.md#getmetros) | **Get** /fabric/v4/metros | Get all Metros
*NetworksApi* | [**CreateNetwork**](docs/NetworksApi.md#createnetwork) | **Post** /fabric/v4/networks | Create Network
*NetworksApi* | [**DeleteNetworkByUuid**](docs/NetworksApi.md#deletenetworkbyuuid) | **Delete** /fabric/v4/networks/{networkId} | Delete Network By ID
*NetworksApi* | [**GetConnectionsByNetworkUuid**](docs/NetworksApi.md#getconnectionsbynetworkuuid) | **Get** /fabric/v4/networks/{networkId}/connections | Get Connections
*NetworksApi* | [**GetNetworkByUuid**](docs/NetworksApi.md#getnetworkbyuuid) | **Get** /fabric/v4/networks/{networkId} | Get Network By ID
*NetworksApi* | [**GetNetworkChangeByUuid**](docs/NetworksApi.md#getnetworkchangebyuuid) | **Get** /fabric/v4/networks/{networkId}/changes/{changeId} | Get Change By ID
*NetworksApi* | [**GetNetworkChanges**](docs/NetworksApi.md#getnetworkchanges) | **Get** /fabric/v4/networks/{networkId}/changes | Get Network Changes
*NetworksApi* | [**SearchNetworks**](docs/NetworksApi.md#searchnetworks) | **Post** /fabric/v4/networks/search | Search Network
*NetworksApi* | [**UpdateNetworkByUuid**](docs/NetworksApi.md#updatenetworkbyuuid) | **Patch** /fabric/v4/networks/{networkId} | Update Network By ID
*PortsApi* | [**AddToLag**](docs/PortsApi.md#addtolag) | **Post** /fabric/v4/ports/{portId}/physicalPorts/bulk | Add to Lag
*PortsApi* | [**CreateBulkPort**](docs/PortsApi.md#createbulkport) | **Post** /fabric/v4/ports/bulk | Create Port
*PortsApi* | [**CreatePort**](docs/PortsApi.md#createport) | **Post** /fabric/v4/ports | Create Port
*PortsApi* | [**GetPortByUuid**](docs/PortsApi.md#getportbyuuid) | **Get** /fabric/v4/ports/{portId} | Get Port by uuid
*PortsApi* | [**GetPorts**](docs/PortsApi.md#getports) | **Get** /fabric/v4/ports | Get All Ports
*PortsApi* | [**GetVlans**](docs/PortsApi.md#getvlans) | **Get** /fabric/v4/ports/{portUuid}/linkProtocols | Get Vlans
*PortsApi* | [**SearchPorts**](docs/PortsApi.md#searchports) | **Post** /fabric/v4/ports/search | Search ports
*PrecisionTimeApi* | [**CreateTimeServices**](docs/PrecisionTimeApi.md#createtimeservices) | **Post** /fabric/v4/timeServices | Create Time Service
*PrecisionTimeApi* | [**DeleteTimeServiceById**](docs/PrecisionTimeApi.md#deletetimeservicebyid) | **Delete** /fabric/v4/timeServices/{serviceId} | Delete time service
*PrecisionTimeApi* | [**GetTimeServicesById**](docs/PrecisionTimeApi.md#gettimeservicesbyid) | **Get** /fabric/v4/timeServices/{serviceId} | Get Time Service
*PrecisionTimeApi* | [**GetTimeServicesConnectionsByServiceId**](docs/PrecisionTimeApi.md#gettimeservicesconnectionsbyserviceid) | **Get** /fabric/v4/timeServices/{serviceId}/connections | Get Connection Links
*PrecisionTimeApi* | [**GetTimeServicesPackageByCode**](docs/PrecisionTimeApi.md#gettimeservicespackagebycode) | **Get** /fabric/v4/timeServicePackages/{packageCode} | Get Package By Code
*PrecisionTimeApi* | [**GetTimeServicesPackages**](docs/PrecisionTimeApi.md#gettimeservicespackages) | **Get** /fabric/v4/timeServicePackages | Get Packages
*PrecisionTimeApi* | [**UpdateTimeServicesById**](docs/PrecisionTimeApi.md#updatetimeservicesbyid) | **Patch** /fabric/v4/timeServices/{serviceId} | Patch time service
*PricesApi* | [**SearchPrices**](docs/PricesApi.md#searchprices) | **Post** /fabric/v4/prices/search | Get Prices
*RouteFilterRulesApi* | [**CreateRouteFilterRule**](docs/RouteFilterRulesApi.md#createroutefilterrule) | **Post** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules | Create RFRule
*RouteFilterRulesApi* | [**CreateRouteFilterRulesInBulk**](docs/RouteFilterRulesApi.md#createroutefilterrulesinbulk) | **Post** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/bulk | Bulk RFRules
*RouteFilterRulesApi* | [**DeleteRouteFilterRuleByUuid**](docs/RouteFilterRulesApi.md#deleteroutefilterrulebyuuid) | **Delete** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | DeleteRFRule
*RouteFilterRulesApi* | [**GetRouteFilterRuleByUuid**](docs/RouteFilterRulesApi.md#getroutefilterrulebyuuid) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | GetRFRule By UUID
*RouteFilterRulesApi* | [**GetRouteFilterRuleChangeByUuid**](docs/RouteFilterRulesApi.md#getroutefilterrulechangebyuuid) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId}/changes/{changeId} | Get Change By ID
*RouteFilterRulesApi* | [**GetRouteFilterRuleChanges**](docs/RouteFilterRulesApi.md#getroutefilterrulechanges) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId}/changes | Get All Changes
*RouteFilterRulesApi* | [**GetRouteFilterRules**](docs/RouteFilterRulesApi.md#getroutefilterrules) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules | GetRFRules
*RouteFilterRulesApi* | [**PatchRouteFilterRuleByUuid**](docs/RouteFilterRulesApi.md#patchroutefilterrulebyuuid) | **Patch** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | PatchRFilterRule
*RouteFilterRulesApi* | [**ReplaceRouteFilterRuleByUuid**](docs/RouteFilterRulesApi.md#replaceroutefilterrulebyuuid) | **Put** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | ReplaceRFRule
*RouteFiltersApi* | [**AttachConnectionRouteFilter**](docs/RouteFiltersApi.md#attachconnectionroutefilter) | **Put** /fabric/v4/connections/{connectionId}/routeFilters/{routeFilterId} | Attach Route Filter
*RouteFiltersApi* | [**CreateRouteFilter**](docs/RouteFiltersApi.md#createroutefilter) | **Post** /fabric/v4/routeFilters | Create Route Filters
*RouteFiltersApi* | [**DeleteRouteFilterByUuid**](docs/RouteFiltersApi.md#deleteroutefilterbyuuid) | **Delete** /fabric/v4/routeFilters/{routeFilterId} | Delete Route Filter
*RouteFiltersApi* | [**DetachConnectionRouteFilter**](docs/RouteFiltersApi.md#detachconnectionroutefilter) | **Delete** /fabric/v4/connections/{connectionId}/routeFilters/{routeFilterId} | Detach Route Filter
*RouteFiltersApi* | [**GetConnectionRouteFilterByUuid**](docs/RouteFiltersApi.md#getconnectionroutefilterbyuuid) | **Get** /fabric/v4/connections/{connectionId}/routeFilters/{routeFilterId} | Get Route Filter
*RouteFiltersApi* | [**GetConnectionRouteFilters**](docs/RouteFiltersApi.md#getconnectionroutefilters) | **Get** /fabric/v4/connections/{connectionId}/routeFilters | Get All RouteFilters
*RouteFiltersApi* | [**GetRouteFilterByUuid**](docs/RouteFiltersApi.md#getroutefilterbyuuid) | **Get** /fabric/v4/routeFilters/{routeFilterId} | Get Filter By UUID
*RouteFiltersApi* | [**GetRouteFilterChangeByUuid**](docs/RouteFiltersApi.md#getroutefilterchangebyuuid) | **Get** /fabric/v4/routeFilters/{routeFilterId}/changes/{changeId} | Get Change By ID
*RouteFiltersApi* | [**GetRouteFilterChanges**](docs/RouteFiltersApi.md#getroutefilterchanges) | **Get** /fabric/v4/routeFilters/{routeFilterId}/changes | Get All Changes
*RouteFiltersApi* | [**GetRouteFilterConnections**](docs/RouteFiltersApi.md#getroutefilterconnections) | **Get** /fabric/v4/routeFilters/{routeFilterId}/connections | Get Connections
*RouteFiltersApi* | [**PatchRouteFilterByUuid**](docs/RouteFiltersApi.md#patchroutefilterbyuuid) | **Patch** /fabric/v4/routeFilters/{routeFilterId} | Patch Route Filter
*RouteFiltersApi* | [**SearchRouteFilters**](docs/RouteFiltersApi.md#searchroutefilters) | **Post** /fabric/v4/routeFilters/search | Search Route Filters
*RoutingProtocolsApi* | [**CreateConnectionRoutingProtocol**](docs/RoutingProtocolsApi.md#createconnectionroutingprotocol) | **Post** /fabric/v4/connections/{connectionId}/routingProtocols | Create Protocol
*RoutingProtocolsApi* | [**CreateConnectionRoutingProtocolsInBulk**](docs/RoutingProtocolsApi.md#createconnectionroutingprotocolsinbulk) | **Post** /fabric/v4/connections/{connectionId}/routingProtocols/bulk | Bulk Create Protocol
*RoutingProtocolsApi* | [**DeleteConnectionRoutingProtocolByUuid**](docs/RoutingProtocolsApi.md#deleteconnectionroutingprotocolbyuuid) | **Delete** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Delete Protocol
*RoutingProtocolsApi* | [**GetConnectionRoutingProtocolAllBgpActions**](docs/RoutingProtocolsApi.md#getconnectionroutingprotocolallbgpactions) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/actions | Get BGP Actions
*RoutingProtocolsApi* | [**GetConnectionRoutingProtocolByUuid**](docs/RoutingProtocolsApi.md#getconnectionroutingprotocolbyuuid) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Get Protocol
*RoutingProtocolsApi* | [**GetConnectionRoutingProtocols**](docs/RoutingProtocolsApi.md#getconnectionroutingprotocols) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols | GetRoutingProtocols
*RoutingProtocolsApi* | [**GetConnectionRoutingProtocolsBgpActionByUuid**](docs/RoutingProtocolsApi.md#getconnectionroutingprotocolsbgpactionbyuuid) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/actions/{actionId} | Get BGP Action
*RoutingProtocolsApi* | [**GetConnectionRoutingProtocolsChangeByUuid**](docs/RoutingProtocolsApi.md#getconnectionroutingprotocolschangebyuuid) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/changes/{changeId} | Get Change By ID
*RoutingProtocolsApi* | [**GetConnectionRoutingProtocolsChanges**](docs/RoutingProtocolsApi.md#getconnectionroutingprotocolschanges) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/changes | Get Changes
*RoutingProtocolsApi* | [**PatchConnectionRoutingProtocolByUuid**](docs/RoutingProtocolsApi.md#patchconnectionroutingprotocolbyuuid) | **Patch** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Patch Protocol
*RoutingProtocolsApi* | [**PostConnectionRoutingProtocolBgpActionByUuid**](docs/RoutingProtocolsApi.md#postconnectionroutingprotocolbgpactionbyuuid) | **Post** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/actions | Clear/Reset BGP
*RoutingProtocolsApi* | [**ReplaceConnectionRoutingProtocolByUuid**](docs/RoutingProtocolsApi.md#replaceconnectionroutingprotocolbyuuid) | **Put** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Replace Protocol
*RoutingProtocolsApi* | [**ValidateRoutingProtocol**](docs/RoutingProtocolsApi.md#validateroutingprotocol) | **Post** /fabric/v4/routers/{routerId}/validate | Validate Subnet
*ServiceProfilesApi* | [**CreateServiceProfile**](docs/ServiceProfilesApi.md#createserviceprofile) | **Post** /fabric/v4/serviceProfiles | Create Profile
*ServiceProfilesApi* | [**DeleteServiceProfileByUuid**](docs/ServiceProfilesApi.md#deleteserviceprofilebyuuid) | **Delete** /fabric/v4/serviceProfiles/{serviceProfileId} | Delete Profile
*ServiceProfilesApi* | [**GetServiceProfileByUuid**](docs/ServiceProfilesApi.md#getserviceprofilebyuuid) | **Get** /fabric/v4/serviceProfiles/{serviceProfileId} | Get Profile
*ServiceProfilesApi* | [**GetServiceProfileMetrosByUuid**](docs/ServiceProfilesApi.md#getserviceprofilemetrosbyuuid) | **Get** /fabric/v4/serviceProfiles/{serviceProfileId}/metros | Get Profile Metros
*ServiceProfilesApi* | [**GetServiceProfiles**](docs/ServiceProfilesApi.md#getserviceprofiles) | **Get** /fabric/v4/serviceProfiles | Get all Profiles
*ServiceProfilesApi* | [**PutServiceProfileByUuid**](docs/ServiceProfilesApi.md#putserviceprofilebyuuid) | **Put** /fabric/v4/serviceProfiles/{serviceProfileId} | Replace Profile
*ServiceProfilesApi* | [**SearchServiceProfiles**](docs/ServiceProfilesApi.md#searchserviceprofiles) | **Post** /fabric/v4/serviceProfiles/search | Profile Search
*ServiceProfilesApi* | [**UpdateServiceProfileByUuid**](docs/ServiceProfilesApi.md#updateserviceprofilebyuuid) | **Patch** /fabric/v4/serviceProfiles/{serviceProfileId} | Update Profile
*ServiceTokensApi* | [**CreateServiceToken**](docs/ServiceTokensApi.md#createservicetoken) | **Post** /fabric/v4/serviceTokens | Create Service Token
*ServiceTokensApi* | [**CreateServiceTokenAction**](docs/ServiceTokensApi.md#createservicetokenaction) | **Post** /fabric/v4/serviceTokens/{serviceTokenId}/actions | ServiceToken Actions
*ServiceTokensApi* | [**DeleteServiceTokenByUuid**](docs/ServiceTokensApi.md#deleteservicetokenbyuuid) | **Delete** /fabric/v4/serviceTokens/{serviceTokenId} | Delete Token by uuid
*ServiceTokensApi* | [**GetServiceTokenByUuid**](docs/ServiceTokensApi.md#getservicetokenbyuuid) | **Get** /fabric/v4/serviceTokens/{serviceTokenId} | Get Token by uuid
*ServiceTokensApi* | [**GetServiceTokens**](docs/ServiceTokensApi.md#getservicetokens) | **Get** /fabric/v4/serviceTokens | Get All Tokens
*ServiceTokensApi* | [**SearchServiceTokens**](docs/ServiceTokensApi.md#searchservicetokens) | **Post** /fabric/v4/serviceTokens/search | Search servicetokens
*ServiceTokensApi* | [**UpdateServiceTokenByUuid**](docs/ServiceTokensApi.md#updateservicetokenbyuuid) | **Patch** /fabric/v4/serviceTokens/{serviceTokenId} | Update Token By ID
*StatisticsApi* | [**GetConnectionStatsByPortUuid**](docs/StatisticsApi.md#getconnectionstatsbyportuuid) | **Get** /fabric/v4/connections/{connectionId}/stats | Get Stats by uuid
*StatisticsApi* | [**GetPortStats**](docs/StatisticsApi.md#getportstats) | **Get** /fabric/v4/ports/stats | Top Port Statistics
*StatisticsApi* | [**GetPortStatsByPortUuid**](docs/StatisticsApi.md#getportstatsbyportuuid) | **Get** /fabric/v4/ports/{portId}/stats | Get Stats by uuid

## Documentation For Models

 - [AccessPoint](docs/AccessPoint.md)
 - [AccessPointSelector](docs/AccessPointSelector.md)
 - [AccessPointType](docs/AccessPointType.md)
 - [Account](docs/Account.md)
 - [Actions](docs/Actions.md)
 - [AddOperation](docs/AddOperation.md)
 - [AdvanceConfiguration](docs/AdvanceConfiguration.md)
 - [AllOfServiceProfileAccount](docs/AllOfServiceProfileAccount.md)
 - [AllOfServiceProfileChangeLog](docs/AllOfServiceProfileChangeLog.md)
 - [AllPhysicalPortsResponse](docs/AllPhysicalPortsResponse.md)
 - [AllPortsResponse](docs/AllPortsResponse.md)
 - [AnyOfCloudRouterFilter](docs/AnyOfCloudRouterFilter.md)
 - [AnyOfRouteTableEntryFilter](docs/AnyOfRouteTableEntryFilter.md)
 - [AnyOfServiceProfileFilter](docs/AnyOfServiceProfileFilter.md)
 - [ApiConfig](docs/ApiConfig.md)
 - [ApiServices](docs/ApiServices.md)
 - [AuthenticationKey](docs/AuthenticationKey.md)
 - [BandwidthUtilization](docs/BandwidthUtilization.md)
 - [BgpActionData](docs/BgpActionData.md)
 - [BgpActionRequest](docs/BgpActionRequest.md)
 - [BgpActionStates](docs/BgpActionStates.md)
 - [BgpActions](docs/BgpActions.md)
 - [BgpActionsBulkData](docs/BgpActionsBulkData.md)
 - [BgpConnectionIpv4](docs/BgpConnectionIpv4.md)
 - [BgpConnectionIpv6](docs/BgpConnectionIpv6.md)
 - [BgpConnectionOperation](docs/BgpConnectionOperation.md)
 - [BulkPhysicalPort](docs/BulkPhysicalPort.md)
 - [BulkPort](docs/BulkPort.md)
 - [Change](docs/Change.md)
 - [Changelog](docs/Changelog.md)
 - [CloudRouter](docs/CloudRouter.md)
 - [CloudRouterAccessPointState](docs/CloudRouterAccessPointState.md)
 - [CloudRouterActionRequest](docs/CloudRouterActionRequest.md)
 - [CloudRouterActionResponse](docs/CloudRouterActionResponse.md)
 - [CloudRouterActionState](docs/CloudRouterActionState.md)
 - [CloudRouterActionType](docs/CloudRouterActionType.md)
 - [CloudRouterChange](docs/CloudRouterChange.md)
 - [CloudRouterChangeOperation](docs/CloudRouterChangeOperation.md)
 - [CloudRouterFilter](docs/CloudRouterFilter.md)
 - [CloudRouterFilters](docs/CloudRouterFilters.md)
 - [CloudRouterOrFilter](docs/CloudRouterOrFilter.md)
 - [CloudRouterPackage](docs/CloudRouterPackage.md)
 - [CloudRouterPostRequest](docs/CloudRouterPostRequest.md)
 - [CloudRouterPostRequestPackage](docs/CloudRouterPostRequestPackage.md)
 - [CloudRouterSearchRequest](docs/CloudRouterSearchRequest.md)
 - [CloudRouterSimpleExpression](docs/CloudRouterSimpleExpression.md)
 - [CloudRouterSortBy](docs/CloudRouterSortBy.md)
 - [CloudRouterSortCriteria](docs/CloudRouterSortCriteria.md)
 - [CloudRouterSortDirection](docs/CloudRouterSortDirection.md)
 - [Code](docs/Code.md)
 - [ConnectedMetro](docs/ConnectedMetro.md)
 - [Connection](docs/Connection.md)
 - [ConnectionAcceptanceData](docs/ConnectionAcceptanceData.md)
 - [ConnectionAction](docs/ConnectionAction.md)
 - [ConnectionActionRequest](docs/ConnectionActionRequest.md)
 - [ConnectionChangeOperation](docs/ConnectionChangeOperation.md)
 - [ConnectionCompanyProfile](docs/ConnectionCompanyProfile.md)
 - [ConnectionDirection](docs/ConnectionDirection.md)
 - [ConnectionInvitation](docs/ConnectionInvitation.md)
 - [ConnectionLink](docs/ConnectionLink.md)
 - [ConnectionOperation](docs/ConnectionOperation.md)
 - [ConnectionPostRequest](docs/ConnectionPostRequest.md)
 - [ConnectionPriority](docs/ConnectionPriority.md)
 - [ConnectionRedundancy](docs/ConnectionRedundancy.md)
 - [ConnectionResponse](docs/ConnectionResponse.md)
 - [ConnectionRouteFilterData](docs/ConnectionRouteFilterData.md)
 - [ConnectionRouteFiltersBase](docs/ConnectionRouteFiltersBase.md)
 - [ConnectionRoutingProtocolPostRequest](docs/ConnectionRoutingProtocolPostRequest.md)
 - [ConnectionSearchResponse](docs/ConnectionSearchResponse.md)
 - [ConnectionSide](docs/ConnectionSide.md)
 - [ConnectionSideAdditionalInfo](docs/ConnectionSideAdditionalInfo.md)
 - [ConnectionState](docs/ConnectionState.md)
 - [ConnectionType](docs/ConnectionType.md)
 - [ConnectivitySource](docs/ConnectivitySource.md)
 - [ConnectivitySourceType](docs/ConnectivitySourceType.md)
 - [CustomField](docs/CustomField.md)
 - [DirectConnectionIpv4](docs/DirectConnectionIpv4.md)
 - [DirectConnectionIpv6](docs/DirectConnectionIpv6.md)
 - [Direction](docs/Direction.md)
 - [Duration](docs/Duration.md)
 - [EquinixStatus](docs/EquinixStatus.md)
 - [Expression](docs/Expression.md)
 - [FabricCloudRouterCode](docs/FabricCloudRouterCode.md)
 - [FabricCloudRouterPackages](docs/FabricCloudRouterPackages.md)
 - [FabricCloudRouterPrice](docs/FabricCloudRouterPrice.md)
 - [FabricConnectionUuid](docs/FabricConnectionUuid.md)
 - [FilterBody](docs/FilterBody.md)
 - [GeoCoordinates](docs/GeoCoordinates.md)
 - [GeoScopeType](docs/GeoScopeType.md)
 - [GetAllConnectionRouteFiltersResponse](docs/GetAllConnectionRouteFiltersResponse.md)
 - [GetResponse](docs/GetResponse.md)
 - [GetRouteFilterGetConnectionsResponse](docs/GetRouteFilterGetConnectionsResponse.md)
 - [GetRouteFilterRulesResponse](docs/GetRouteFilterRulesResponse.md)
 - [HealthResponse](docs/HealthResponse.md)
 - [IpBlockPrice](docs/IpBlockPrice.md)
 - [IpBlockType](docs/IpBlockType.md)
 - [Ipv4](docs/Ipv4.md)
 - [JsonPatchOperation](docs/JsonPatchOperation.md)
 - [Link](docs/Link.md)
 - [LinkAggregationGroup](docs/LinkAggregationGroup.md)
 - [LinkProtocol](docs/LinkProtocol.md)
 - [LinkProtocolConnection](docs/LinkProtocolConnection.md)
 - [LinkProtocolDot1q](docs/LinkProtocolDot1q.md)
 - [LinkProtocolEvpnVxlan](docs/LinkProtocolEvpnVxlan.md)
 - [LinkProtocolGetResponse](docs/LinkProtocolGetResponse.md)
 - [LinkProtocolIpv4Ipv6Config](docs/LinkProtocolIpv4Ipv6Config.md)
 - [LinkProtocolQinq](docs/LinkProtocolQinq.md)
 - [LinkProtocolRequestType](docs/LinkProtocolRequestType.md)
 - [LinkProtocolResponse](docs/LinkProtocolResponse.md)
 - [LinkProtocolServiceToken](docs/LinkProtocolServiceToken.md)
 - [LinkProtocolState](docs/LinkProtocolState.md)
 - [LinkProtocolType](docs/LinkProtocolType.md)
 - [LinkProtocolUntagged](docs/LinkProtocolUntagged.md)
 - [LinkProtocolVxlan](docs/LinkProtocolVxlan.md)
 - [MarketingInfo](docs/MarketingInfo.md)
 - [Md5](docs/Md5.md)
 - [MetalInterconnection](docs/MetalInterconnection.md)
 - [MetricInterval](docs/MetricInterval.md)
 - [Metrics](docs/Metrics.md)
 - [Metro](docs/Metro.md)
 - [MetroError](docs/MetroError.md)
 - [MetroResponse](docs/MetroResponse.md)
 - [ModelError](docs/ModelError.md)
 - [ModelInterface](docs/ModelInterface.md)
 - [Network](docs/Network.md)
 - [NetworkChange](docs/NetworkChange.md)
 - [NetworkChangeOperation](docs/NetworkChangeOperation.md)
 - [NetworkChangeResponse](docs/NetworkChangeResponse.md)
 - [NetworkChangeStatus](docs/NetworkChangeStatus.md)
 - [NetworkChangeType](docs/NetworkChangeType.md)
 - [NetworkConnections](docs/NetworkConnections.md)
 - [NetworkEquinixStatus](docs/NetworkEquinixStatus.md)
 - [NetworkFilter](docs/NetworkFilter.md)
 - [NetworkOperation](docs/NetworkOperation.md)
 - [NetworkPostRequest](docs/NetworkPostRequest.md)
 - [NetworkScope](docs/NetworkScope.md)
 - [NetworkSearchFieldName](docs/NetworkSearchFieldName.md)
 - [NetworkSearchRequest](docs/NetworkSearchRequest.md)
 - [NetworkSearchResponse](docs/NetworkSearchResponse.md)
 - [NetworkSortBy](docs/NetworkSortBy.md)
 - [NetworkSortByResponse](docs/NetworkSortByResponse.md)
 - [NetworkSortCriteria](docs/NetworkSortCriteria.md)
 - [NetworkSortCriteriaResponse](docs/NetworkSortCriteriaResponse.md)
 - [NetworkSortDirection](docs/NetworkSortDirection.md)
 - [NetworkSortDirectionResponse](docs/NetworkSortDirectionResponse.md)
 - [NetworkState](docs/NetworkState.md)
 - [NetworkType](docs/NetworkType.md)
 - [OneOfJsonPatchOperation](docs/OneOfJsonPatchOperation.md)
 - [OneOfLinkProtocol](docs/OneOfLinkProtocol.md)
 - [OneOfRoutingProtocolBase](docs/OneOfRoutingProtocolBase.md)
 - [OneOfRoutingProtocolData](docs/OneOfRoutingProtocolData.md)
 - [OneOfServiceProfileAccessPointType](docs/OneOfServiceProfileAccessPointType.md)
 - [OpEnum](docs/OpEnum.md)
 - [Order](docs/Order.md)
 - [PackageChangeLog](docs/PackageChangeLog.md)
 - [PackageResponse](docs/PackageResponse.md)
 - [Pagination](docs/Pagination.md)
 - [PaginationRequest](docs/PaginationRequest.md)
 - [PeeringType](docs/PeeringType.md)
 - [PhysicalPort](docs/PhysicalPort.md)
 - [PhysicalPortType](docs/PhysicalPortType.md)
 - [Port](docs/Port.md)
 - [PortAdditionalInfo](docs/PortAdditionalInfo.md)
 - [PortDemarcationPoint](docs/PortDemarcationPoint.md)
 - [PortDevice](docs/PortDevice.md)
 - [PortDeviceRedundancy](docs/PortDeviceRedundancy.md)
 - [PortEncapsulation](docs/PortEncapsulation.md)
 - [PortExpression](docs/PortExpression.md)
 - [PortInterface](docs/PortInterface.md)
 - [PortLag](docs/PortLag.md)
 - [PortLoa](docs/PortLoa.md)
 - [PortNotification](docs/PortNotification.md)
 - [PortOperation](docs/PortOperation.md)
 - [PortOrder](docs/PortOrder.md)
 - [PortOrderPurchaseOrder](docs/PortOrderPurchaseOrder.md)
 - [PortOrderSignature](docs/PortOrderSignature.md)
 - [PortOrderSignatureDelegate](docs/PortOrderSignatureDelegate.md)
 - [PortPriority](docs/PortPriority.md)
 - [PortRedundancy](docs/PortRedundancy.md)
 - [PortSearchFieldName](docs/PortSearchFieldName.md)
 - [PortSettings](docs/PortSettings.md)
 - [PortSortBy](docs/PortSortBy.md)
 - [PortSortCriteria](docs/PortSortCriteria.md)
 - [PortSortDirection](docs/PortSortDirection.md)
 - [PortState](docs/PortState.md)
 - [PortTether](docs/PortTether.md)
 - [PortType](docs/PortType.md)
 - [PortV4SearchRequest](docs/PortV4SearchRequest.md)
 - [PrecisionTimeChangeOperation](docs/PrecisionTimeChangeOperation.md)
 - [PrecisionTimePackageRequest](docs/PrecisionTimePackageRequest.md)
 - [PrecisionTimePackageResponse](docs/PrecisionTimePackageResponse.md)
 - [PrecisionTimeServiceConnectionsResponse](docs/PrecisionTimeServiceConnectionsResponse.md)
 - [PrecisionTimeServiceCreateResponse](docs/PrecisionTimeServiceCreateResponse.md)
 - [PrecisionTimeServicePackagesResponse](docs/PrecisionTimeServicePackagesResponse.md)
 - [PrecisionTimeServiceRequest](docs/PrecisionTimeServiceRequest.md)
 - [Presence](docs/Presence.md)
 - [Price](docs/Price.md)
 - [PriceCategory](docs/PriceCategory.md)
 - [PriceCharge](docs/PriceCharge.md)
 - [PriceError](docs/PriceError.md)
 - [PriceErrorAdditionalInfo](docs/PriceErrorAdditionalInfo.md)
 - [PriceLocation](docs/PriceLocation.md)
 - [PriceSearchResponse](docs/PriceSearchResponse.md)
 - [ProcessStep](docs/ProcessStep.md)
 - [ProductType](docs/ProductType.md)
 - [Project](docs/Project.md)
 - [ProviderStatus](docs/ProviderStatus.md)
 - [PtpAdvanceConfiguration](docs/PtpAdvanceConfiguration.md)
 - [QueryDirection](docs/QueryDirection.md)
 - [RemoveOperation](docs/RemoveOperation.md)
 - [ReplaceOperation](docs/ReplaceOperation.md)
 - [RouteFilterChangeData](docs/RouteFilterChangeData.md)
 - [RouteFilterChangeDataResponse](docs/RouteFilterChangeDataResponse.md)
 - [RouteFilterConnectionsData](docs/RouteFilterConnectionsData.md)
 - [RouteFilterRuleState](docs/RouteFilterRuleState.md)
 - [RouteFilterRulesBase](docs/RouteFilterRulesBase.md)
 - [RouteFilterRulesChange](docs/RouteFilterRulesChange.md)
 - [RouteFilterRulesChangeData](docs/RouteFilterRulesChangeData.md)
 - [RouteFilterRulesChangeDataResponse](docs/RouteFilterRulesChangeDataResponse.md)
 - [RouteFilterRulesChangeOperation](docs/RouteFilterRulesChangeOperation.md)
 - [RouteFilterRulesData](docs/RouteFilterRulesData.md)
 - [RouteFilterRulesPatchRequestItem](docs/RouteFilterRulesPatchRequestItem.md)
 - [RouteFilterRulesPostRequest](docs/RouteFilterRulesPostRequest.md)
 - [RouteFilterState](docs/RouteFilterState.md)
 - [RouteFiltersBase](docs/RouteFiltersBase.md)
 - [RouteFiltersChange](docs/RouteFiltersChange.md)
 - [RouteFiltersChangeOperation](docs/RouteFiltersChangeOperation.md)
 - [RouteFiltersData](docs/RouteFiltersData.md)
 - [RouteFiltersDataProject](docs/RouteFiltersDataProject.md)
 - [RouteFiltersPatchRequestItem](docs/RouteFiltersPatchRequestItem.md)
 - [RouteFiltersSearchBase](docs/RouteFiltersSearchBase.md)
 - [RouteFiltersSearchBaseFilter](docs/RouteFiltersSearchBaseFilter.md)
 - [RouteFiltersSearchFilterItem](docs/RouteFiltersSearchFilterItem.md)
 - [RouteFiltersSearchResponse](docs/RouteFiltersSearchResponse.md)
 - [RouteTableEntry](docs/RouteTableEntry.md)
 - [RouteTableEntryConnection](docs/RouteTableEntryConnection.md)
 - [RouteTableEntryFilter](docs/RouteTableEntryFilter.md)
 - [RouteTableEntryFilters](docs/RouteTableEntryFilters.md)
 - [RouteTableEntryOrFilter](docs/RouteTableEntryOrFilter.md)
 - [RouteTableEntryProtocolType](docs/RouteTableEntryProtocolType.md)
 - [RouteTableEntrySearchRequest](docs/RouteTableEntrySearchRequest.md)
 - [RouteTableEntrySearchResponse](docs/RouteTableEntrySearchResponse.md)
 - [RouteTableEntrySimpleExpression](docs/RouteTableEntrySimpleExpression.md)
 - [RouteTableEntrySortBy](docs/RouteTableEntrySortBy.md)
 - [RouteTableEntrySortCriteria](docs/RouteTableEntrySortCriteria.md)
 - [RouteTableEntrySortDirection](docs/RouteTableEntrySortDirection.md)
 - [RouteTableEntryState](docs/RouteTableEntryState.md)
 - [RouteTableEntryType](docs/RouteTableEntryType.md)
 - [RouterPackageCode](docs/RouterPackageCode.md)
 - [RoutingProtocolBase](docs/RoutingProtocolBase.md)
 - [RoutingProtocolBfd](docs/RoutingProtocolBfd.md)
 - [RoutingProtocolBgpData](docs/RoutingProtocolBgpData.md)
 - [RoutingProtocolBgpType](docs/RoutingProtocolBgpType.md)
 - [RoutingProtocolChange](docs/RoutingProtocolChange.md)
 - [RoutingProtocolChangeData](docs/RoutingProtocolChangeData.md)
 - [RoutingProtocolChangeDataResponse](docs/RoutingProtocolChangeDataResponse.md)
 - [RoutingProtocolChangeOperation](docs/RoutingProtocolChangeOperation.md)
 - [RoutingProtocolData](docs/RoutingProtocolData.md)
 - [RoutingProtocolDirectData](docs/RoutingProtocolDirectData.md)
 - [RoutingProtocolDirectType](docs/RoutingProtocolDirectType.md)
 - [RoutingProtocolOperation](docs/RoutingProtocolOperation.md)
 - [SearchExpression](docs/SearchExpression.md)
 - [SearchFieldName](docs/SearchFieldName.md)
 - [SearchRequest](docs/SearchRequest.md)
 - [SearchResponse](docs/SearchResponse.md)
 - [ServiceMetro](docs/ServiceMetro.md)
 - [ServiceMetros](docs/ServiceMetros.md)
 - [ServiceProfile](docs/ServiceProfile.md)
 - [ServiceProfileAccessPointColo](docs/ServiceProfileAccessPointColo.md)
 - [ServiceProfileAccessPointType](docs/ServiceProfileAccessPointType.md)
 - [ServiceProfileAccessPointTypeColo](docs/ServiceProfileAccessPointTypeColo.md)
 - [ServiceProfileAccessPointTypeEnum](docs/ServiceProfileAccessPointTypeEnum.md)
 - [ServiceProfileAccessPointTypeVd](docs/ServiceProfileAccessPointTypeVd.md)
 - [ServiceProfileAccessPointVd](docs/ServiceProfileAccessPointVd.md)
 - [ServiceProfileAndFilter](docs/ServiceProfileAndFilter.md)
 - [ServiceProfileFilter](docs/ServiceProfileFilter.md)
 - [ServiceProfileLinkProtocolConfig](docs/ServiceProfileLinkProtocolConfig.md)
 - [ServiceProfileMetadata](docs/ServiceProfileMetadata.md)
 - [ServiceProfileRequest](docs/ServiceProfileRequest.md)
 - [ServiceProfileSearchRequest](docs/ServiceProfileSearchRequest.md)
 - [ServiceProfileSimpleExpression](docs/ServiceProfileSimpleExpression.md)
 - [ServiceProfileSortBy](docs/ServiceProfileSortBy.md)
 - [ServiceProfileSortCriteria](docs/ServiceProfileSortCriteria.md)
 - [ServiceProfileSortDirection](docs/ServiceProfileSortDirection.md)
 - [ServiceProfileStateEnum](docs/ServiceProfileStateEnum.md)
 - [ServiceProfileTypeEnum](docs/ServiceProfileTypeEnum.md)
 - [ServiceProfileVisibilityEnum](docs/ServiceProfileVisibilityEnum.md)
 - [ServiceProfiles](docs/ServiceProfiles.md)
 - [ServiceToken](docs/ServiceToken.md)
 - [ServiceTokenActionRequest](docs/ServiceTokenActionRequest.md)
 - [ServiceTokenActions](docs/ServiceTokenActions.md)
 - [ServiceTokenChangeOperation](docs/ServiceTokenChangeOperation.md)
 - [ServiceTokenConnection](docs/ServiceTokenConnection.md)
 - [ServiceTokenSearchExpression](docs/ServiceTokenSearchExpression.md)
 - [ServiceTokenSearchFieldName](docs/ServiceTokenSearchFieldName.md)
 - [ServiceTokenSearchRequest](docs/ServiceTokenSearchRequest.md)
 - [ServiceTokenSide](docs/ServiceTokenSide.md)
 - [ServiceTokenState](docs/ServiceTokenState.md)
 - [ServiceTokenType](docs/ServiceTokenType.md)
 - [ServiceTokens](docs/ServiceTokens.md)
 - [SimplifiedAccount](docs/SimplifiedAccount.md)
 - [SimplifiedLinkProtocol](docs/SimplifiedLinkProtocol.md)
 - [SimplifiedLocation](docs/SimplifiedLocation.md)
 - [SimplifiedLocationWithoutIbx](docs/SimplifiedLocationWithoutIbx.md)
 - [SimplifiedMetadataEntity](docs/SimplifiedMetadataEntity.md)
 - [SimplifiedNetwork](docs/SimplifiedNetwork.md)
 - [SimplifiedNetworkChange](docs/SimplifiedNetworkChange.md)
 - [SimplifiedNotification](docs/SimplifiedNotification.md)
 - [SimplifiedPort](docs/SimplifiedPort.md)
 - [SimplifiedServiceProfile](docs/SimplifiedServiceProfile.md)
 - [Sort](docs/Sort.md)
 - [SortBy](docs/SortBy.md)
 - [SortCriteria](docs/SortCriteria.md)
 - [SortCriteriaResponse](docs/SortCriteriaResponse.md)
 - [SortDirection](docs/SortDirection.md)
 - [SortItem](docs/SortItem.md)
 - [Statistics](docs/Statistics.md)
 - [SubInterface](docs/SubInterface.md)
 - [TopUtilizedStatistics](docs/TopUtilizedStatistics.md)
 - [ValidateRequest](docs/ValidateRequest.md)
 - [ValidateRequestFilter](docs/ValidateRequestFilter.md)
 - [ValidateRequestFilterAnd](docs/ValidateRequestFilterAnd.md)
 - [ValidateSubnetResponse](docs/ValidateSubnetResponse.md)
 - [ViewPoint](docs/ViewPoint.md)
 - [VirtualConnectionBridgePackageCode](docs/VirtualConnectionBridgePackageCode.md)
 - [VirtualConnectionPrice](docs/VirtualConnectionPrice.md)
 - [VirtualConnectionPriceASide](docs/VirtualConnectionPriceASide.md)
 - [VirtualConnectionPriceASideAccessPoint](docs/VirtualConnectionPriceASideAccessPoint.md)
 - [VirtualConnectionPriceASideAccessPointPort](docs/VirtualConnectionPriceASideAccessPointPort.md)
 - [VirtualConnectionPriceASideAccessPointPortSettings](docs/VirtualConnectionPriceASideAccessPointPortSettings.md)
 - [VirtualConnectionPriceAccessPointType](docs/VirtualConnectionPriceAccessPointType.md)
 - [VirtualConnectionPriceConnectionType](docs/VirtualConnectionPriceConnectionType.md)
 - [VirtualConnectionPriceZSide](docs/VirtualConnectionPriceZSide.md)
 - [VirtualConnectionPriceZSideAccessPoint](docs/VirtualConnectionPriceZSideAccessPoint.md)
 - [VirtualConnectionPriceZSideAccessPointBridge](docs/VirtualConnectionPriceZSideAccessPointBridge.md)
 - [VirtualConnectionPriceZSideAccessPointBridgePackage](docs/VirtualConnectionPriceZSideAccessPointBridgePackage.md)
 - [VirtualConnectionPriceZSideAccessPointProfile](docs/VirtualConnectionPriceZSideAccessPointProfile.md)
 - [VirtualDevice](docs/VirtualDevice.md)
 - [VirtualNetwork](docs/VirtualNetwork.md)
 - [VirtualPortConfiguration](docs/VirtualPortConfiguration.md)
 - [VirtualPortLocation](docs/VirtualPortLocation.md)
 - [VirtualPortPrice](docs/VirtualPortPrice.md)
 - [VirtualPortRedundancy](docs/VirtualPortRedundancy.md)
 - [VirtualPortServiceType](docs/VirtualPortServiceType.md)
 - [VirtualPortType](docs/VirtualPortType.md)

## Documentation For Authorization

## BearerAuth

## Author

api-support@equinix.com
