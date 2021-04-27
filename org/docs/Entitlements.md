# Entitlements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AngGovernance** | [**AngGovernance**](AngGovernance.md) |  | 
**AnypointSecurityEdgePolicies** | [**AnypointSecurityEdgePolicies**](AnypointSecurityEdgePolicies.md) |  | 
**AnypointSecurityTokenization** | [**AnypointSecurityTokenization**](AnypointSecurityTokenization.md) |  | 
**ApiCommunityManager** | [**ApiCommunityManager**](ApiCommunityManager.md) |  | 
**ApiMonitoring** | [**ApiMonitoring**](ApiMonitoring.md) |  | 
**ApiQuery** | [**ApiQuery**](ApiQuery.md) |  | 
**ApiQueryC360** | [**ApiQueryC360**](ApiQueryC360.md) |  | 
**Apis** | [**Apis**](Apis.md) |  | 
**AppViz** | **bool** | An explanation about the purpose of this instance. | [default to false]
**ArmAlerts** | **bool** | An explanation about the purpose of this instance. | [default to false]
**Autoscaling** | **bool** | An explanation about the purpose of this instance. | [default to false]
**Cam** | [**Cam**](Cam.md) |  | 
**CreateEnvironments** | **bool** | An explanation about the purpose of this instance. | [default to false]
**CreateSubOrgs** | **bool** | An explanation about the purpose of this instance. | [default to false]
**Crowd** | [**Crowd**](Crowd.md) |  | 
**CrowdSelfServiceMigration** | [**CrowdSelfServiceMigration**](CrowdSelfServiceMigration.md) |  | 
**DesignCenter** | [**TheDesignCenterSchema**](TheDesignCenterSchema.md) |  | 
**Exchange2** | [**Exchange2**](Exchange2.md) |  | 
**ExternalIdentity** | **bool** | An explanation about the purpose of this instance. | [default to false]
**Gateways** | [**Gateways**](Gateways.md) |  | 
**GlobalDeployment** | **bool** | An explanation about the purpose of this instance. | [default to false]
**Hybrid** | [**TheHybridSchema**](TheHybridSchema.md) |  | 
**HybridAutoDiscoverProperties** | **bool** | An explanation about the purpose of this instance. | [default to false]
**HybridInsight** | **bool** | An explanation about the purpose of this instance. | [default to false]
**KpiDashboard** | [**KpiDashboard**](KpiDashboard.md) |  | 
**LoadBalancer** | [**LoadBalancer**](LoadBalancer.md) |  | 
**Messaging** | [**Messaging**](Messaging.md) |  | 
**MonitoringCenter** | [**MonitoringCenter**](MonitoringCenter.md) |  | 
**MqAdvancedFeatures** | [**MqAdvancedFeatures**](MqAdvancedFeatures.md) |  | 
**MqMessages** | [**MqMessages**](MqMessages.md) |  | 
**MqRequests** | [**MqRequests**](MqRequests.md) |  | 
**ObjectStoreKeys** | [**ObjectStoreKeys**](ObjectStoreKeys.md) |  | 
**ObjectStoreRequestUnits** | [**ObjectStoreRequestUnits**](ObjectStoreRequestUnits.md) |  | 
**PartnersProduction** | [**PartnersProduction**](PartnersProduction.md) |  | 
**PartnersSandbox** | [**PartnersSandbox**](PartnersSandbox.md) |  | 
**Pcf** | **bool** | An explanation about the purpose of this instance. | [default to false]
**RuntimeFabric** | **bool** | An explanation about the purpose of this instance. | [default to false]
**RuntimeFabricCloud** | [**RuntimeFabricCloud**](RuntimeFabricCloud.md) |  | 
**ServiceMesh** | [**ServiceMesh**](ServiceMesh.md) |  | 
**StaticIps** | [**StaticIps**](StaticIps.md) |  | 
**TradingPartnersProduction** | [**TradingPartnersProduction**](TradingPartnersProduction.md) |  | 
**TradingPartnersSandbox** | [**TradingPartnersSandbox**](TradingPartnersSandbox.md) |  | 
**VCoresDesign** | [**VCoresDesign**](VCoresDesign.md) |  | 
**VCoresProduction** | [**VCoresProduction**](VCoresProduction.md) |  | 
**VCoresSandbox** | [**VCoresSandbox**](VCoresSandbox.md) |  | 
**Vpcs** | [**Vpcs**](Vpcs.md) |  | 
**Vpns** | [**Vpns**](Vpns.md) |  | 
**WorkerClouds** | [**WorkerClouds**](WorkerClouds.md) |  | 
**WorkerLoggingOverride** | [**WorkerLoggingOverride**](WorkerLoggingOverride.md) |  | 

## Methods

### NewEntitlements

`func NewEntitlements(angGovernance AngGovernance, anypointSecurityEdgePolicies AnypointSecurityEdgePolicies, anypointSecurityTokenization AnypointSecurityTokenization, apiCommunityManager ApiCommunityManager, apiMonitoring ApiMonitoring, apiQuery ApiQuery, apiQueryC360 ApiQueryC360, apis Apis, appViz bool, armAlerts bool, autoscaling bool, cam Cam, createEnvironments bool, createSubOrgs bool, crowd Crowd, crowdSelfServiceMigration CrowdSelfServiceMigration, designCenter TheDesignCenterSchema, exchange2 Exchange2, externalIdentity bool, gateways Gateways, globalDeployment bool, hybrid TheHybridSchema, hybridAutoDiscoverProperties bool, hybridInsight bool, kpiDashboard KpiDashboard, loadBalancer LoadBalancer, messaging Messaging, monitoringCenter MonitoringCenter, mqAdvancedFeatures MqAdvancedFeatures, mqMessages MqMessages, mqRequests MqRequests, objectStoreKeys ObjectStoreKeys, objectStoreRequestUnits ObjectStoreRequestUnits, partnersProduction PartnersProduction, partnersSandbox PartnersSandbox, pcf bool, runtimeFabric bool, runtimeFabricCloud RuntimeFabricCloud, serviceMesh ServiceMesh, staticIps StaticIps, tradingPartnersProduction TradingPartnersProduction, tradingPartnersSandbox TradingPartnersSandbox, vCoresDesign VCoresDesign, vCoresProduction VCoresProduction, vCoresSandbox VCoresSandbox, vpcs Vpcs, vpns Vpns, workerClouds WorkerClouds, workerLoggingOverride WorkerLoggingOverride, ) *Entitlements`

NewEntitlements instantiates a new Entitlements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsWithDefaults

`func NewEntitlementsWithDefaults() *Entitlements`

NewEntitlementsWithDefaults instantiates a new Entitlements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAngGovernance

`func (o *Entitlements) GetAngGovernance() AngGovernance`

GetAngGovernance returns the AngGovernance field if non-nil, zero value otherwise.

### GetAngGovernanceOk

`func (o *Entitlements) GetAngGovernanceOk() (*AngGovernance, bool)`

GetAngGovernanceOk returns a tuple with the AngGovernance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngGovernance

`func (o *Entitlements) SetAngGovernance(v AngGovernance)`

SetAngGovernance sets AngGovernance field to given value.


### GetAnypointSecurityEdgePolicies

`func (o *Entitlements) GetAnypointSecurityEdgePolicies() AnypointSecurityEdgePolicies`

GetAnypointSecurityEdgePolicies returns the AnypointSecurityEdgePolicies field if non-nil, zero value otherwise.

### GetAnypointSecurityEdgePoliciesOk

`func (o *Entitlements) GetAnypointSecurityEdgePoliciesOk() (*AnypointSecurityEdgePolicies, bool)`

GetAnypointSecurityEdgePoliciesOk returns a tuple with the AnypointSecurityEdgePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnypointSecurityEdgePolicies

`func (o *Entitlements) SetAnypointSecurityEdgePolicies(v AnypointSecurityEdgePolicies)`

SetAnypointSecurityEdgePolicies sets AnypointSecurityEdgePolicies field to given value.


### GetAnypointSecurityTokenization

`func (o *Entitlements) GetAnypointSecurityTokenization() AnypointSecurityTokenization`

GetAnypointSecurityTokenization returns the AnypointSecurityTokenization field if non-nil, zero value otherwise.

### GetAnypointSecurityTokenizationOk

`func (o *Entitlements) GetAnypointSecurityTokenizationOk() (*AnypointSecurityTokenization, bool)`

GetAnypointSecurityTokenizationOk returns a tuple with the AnypointSecurityTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnypointSecurityTokenization

`func (o *Entitlements) SetAnypointSecurityTokenization(v AnypointSecurityTokenization)`

SetAnypointSecurityTokenization sets AnypointSecurityTokenization field to given value.


### GetApiCommunityManager

`func (o *Entitlements) GetApiCommunityManager() ApiCommunityManager`

GetApiCommunityManager returns the ApiCommunityManager field if non-nil, zero value otherwise.

### GetApiCommunityManagerOk

`func (o *Entitlements) GetApiCommunityManagerOk() (*ApiCommunityManager, bool)`

GetApiCommunityManagerOk returns a tuple with the ApiCommunityManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCommunityManager

`func (o *Entitlements) SetApiCommunityManager(v ApiCommunityManager)`

SetApiCommunityManager sets ApiCommunityManager field to given value.


### GetApiMonitoring

`func (o *Entitlements) GetApiMonitoring() ApiMonitoring`

GetApiMonitoring returns the ApiMonitoring field if non-nil, zero value otherwise.

### GetApiMonitoringOk

`func (o *Entitlements) GetApiMonitoringOk() (*ApiMonitoring, bool)`

GetApiMonitoringOk returns a tuple with the ApiMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMonitoring

`func (o *Entitlements) SetApiMonitoring(v ApiMonitoring)`

SetApiMonitoring sets ApiMonitoring field to given value.


### GetApiQuery

`func (o *Entitlements) GetApiQuery() ApiQuery`

GetApiQuery returns the ApiQuery field if non-nil, zero value otherwise.

### GetApiQueryOk

`func (o *Entitlements) GetApiQueryOk() (*ApiQuery, bool)`

GetApiQueryOk returns a tuple with the ApiQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiQuery

`func (o *Entitlements) SetApiQuery(v ApiQuery)`

SetApiQuery sets ApiQuery field to given value.


### GetApiQueryC360

`func (o *Entitlements) GetApiQueryC360() ApiQueryC360`

GetApiQueryC360 returns the ApiQueryC360 field if non-nil, zero value otherwise.

### GetApiQueryC360Ok

`func (o *Entitlements) GetApiQueryC360Ok() (*ApiQueryC360, bool)`

GetApiQueryC360Ok returns a tuple with the ApiQueryC360 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiQueryC360

`func (o *Entitlements) SetApiQueryC360(v ApiQueryC360)`

SetApiQueryC360 sets ApiQueryC360 field to given value.


### GetApis

`func (o *Entitlements) GetApis() Apis`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *Entitlements) GetApisOk() (*Apis, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *Entitlements) SetApis(v Apis)`

SetApis sets Apis field to given value.


### GetAppViz

`func (o *Entitlements) GetAppViz() bool`

GetAppViz returns the AppViz field if non-nil, zero value otherwise.

### GetAppVizOk

`func (o *Entitlements) GetAppVizOk() (*bool, bool)`

GetAppVizOk returns a tuple with the AppViz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppViz

`func (o *Entitlements) SetAppViz(v bool)`

SetAppViz sets AppViz field to given value.


### GetArmAlerts

`func (o *Entitlements) GetArmAlerts() bool`

GetArmAlerts returns the ArmAlerts field if non-nil, zero value otherwise.

### GetArmAlertsOk

`func (o *Entitlements) GetArmAlertsOk() (*bool, bool)`

GetArmAlertsOk returns a tuple with the ArmAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmAlerts

`func (o *Entitlements) SetArmAlerts(v bool)`

SetArmAlerts sets ArmAlerts field to given value.


### GetAutoscaling

`func (o *Entitlements) GetAutoscaling() bool`

GetAutoscaling returns the Autoscaling field if non-nil, zero value otherwise.

### GetAutoscalingOk

`func (o *Entitlements) GetAutoscalingOk() (*bool, bool)`

GetAutoscalingOk returns a tuple with the Autoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaling

`func (o *Entitlements) SetAutoscaling(v bool)`

SetAutoscaling sets Autoscaling field to given value.


### GetCam

`func (o *Entitlements) GetCam() Cam`

GetCam returns the Cam field if non-nil, zero value otherwise.

### GetCamOk

`func (o *Entitlements) GetCamOk() (*Cam, bool)`

GetCamOk returns a tuple with the Cam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCam

`func (o *Entitlements) SetCam(v Cam)`

SetCam sets Cam field to given value.


### GetCreateEnvironments

`func (o *Entitlements) GetCreateEnvironments() bool`

GetCreateEnvironments returns the CreateEnvironments field if non-nil, zero value otherwise.

### GetCreateEnvironmentsOk

`func (o *Entitlements) GetCreateEnvironmentsOk() (*bool, bool)`

GetCreateEnvironmentsOk returns a tuple with the CreateEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnvironments

`func (o *Entitlements) SetCreateEnvironments(v bool)`

SetCreateEnvironments sets CreateEnvironments field to given value.


### GetCreateSubOrgs

`func (o *Entitlements) GetCreateSubOrgs() bool`

GetCreateSubOrgs returns the CreateSubOrgs field if non-nil, zero value otherwise.

### GetCreateSubOrgsOk

`func (o *Entitlements) GetCreateSubOrgsOk() (*bool, bool)`

GetCreateSubOrgsOk returns a tuple with the CreateSubOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSubOrgs

`func (o *Entitlements) SetCreateSubOrgs(v bool)`

SetCreateSubOrgs sets CreateSubOrgs field to given value.


### GetCrowd

`func (o *Entitlements) GetCrowd() Crowd`

GetCrowd returns the Crowd field if non-nil, zero value otherwise.

### GetCrowdOk

`func (o *Entitlements) GetCrowdOk() (*Crowd, bool)`

GetCrowdOk returns a tuple with the Crowd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrowd

`func (o *Entitlements) SetCrowd(v Crowd)`

SetCrowd sets Crowd field to given value.


### GetCrowdSelfServiceMigration

`func (o *Entitlements) GetCrowdSelfServiceMigration() CrowdSelfServiceMigration`

GetCrowdSelfServiceMigration returns the CrowdSelfServiceMigration field if non-nil, zero value otherwise.

### GetCrowdSelfServiceMigrationOk

`func (o *Entitlements) GetCrowdSelfServiceMigrationOk() (*CrowdSelfServiceMigration, bool)`

GetCrowdSelfServiceMigrationOk returns a tuple with the CrowdSelfServiceMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrowdSelfServiceMigration

`func (o *Entitlements) SetCrowdSelfServiceMigration(v CrowdSelfServiceMigration)`

SetCrowdSelfServiceMigration sets CrowdSelfServiceMigration field to given value.


### GetDesignCenter

`func (o *Entitlements) GetDesignCenter() TheDesignCenterSchema`

GetDesignCenter returns the DesignCenter field if non-nil, zero value otherwise.

### GetDesignCenterOk

`func (o *Entitlements) GetDesignCenterOk() (*TheDesignCenterSchema, bool)`

GetDesignCenterOk returns a tuple with the DesignCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignCenter

`func (o *Entitlements) SetDesignCenter(v TheDesignCenterSchema)`

SetDesignCenter sets DesignCenter field to given value.


### GetExchange2

`func (o *Entitlements) GetExchange2() Exchange2`

GetExchange2 returns the Exchange2 field if non-nil, zero value otherwise.

### GetExchange2Ok

`func (o *Entitlements) GetExchange2Ok() (*Exchange2, bool)`

GetExchange2Ok returns a tuple with the Exchange2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange2

`func (o *Entitlements) SetExchange2(v Exchange2)`

SetExchange2 sets Exchange2 field to given value.


### GetExternalIdentity

`func (o *Entitlements) GetExternalIdentity() bool`

GetExternalIdentity returns the ExternalIdentity field if non-nil, zero value otherwise.

### GetExternalIdentityOk

`func (o *Entitlements) GetExternalIdentityOk() (*bool, bool)`

GetExternalIdentityOk returns a tuple with the ExternalIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentity

`func (o *Entitlements) SetExternalIdentity(v bool)`

SetExternalIdentity sets ExternalIdentity field to given value.


### GetGateways

`func (o *Entitlements) GetGateways() Gateways`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *Entitlements) GetGatewaysOk() (*Gateways, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *Entitlements) SetGateways(v Gateways)`

SetGateways sets Gateways field to given value.


### GetGlobalDeployment

`func (o *Entitlements) GetGlobalDeployment() bool`

GetGlobalDeployment returns the GlobalDeployment field if non-nil, zero value otherwise.

### GetGlobalDeploymentOk

`func (o *Entitlements) GetGlobalDeploymentOk() (*bool, bool)`

GetGlobalDeploymentOk returns a tuple with the GlobalDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDeployment

`func (o *Entitlements) SetGlobalDeployment(v bool)`

SetGlobalDeployment sets GlobalDeployment field to given value.


### GetHybrid

`func (o *Entitlements) GetHybrid() TheHybridSchema`

GetHybrid returns the Hybrid field if non-nil, zero value otherwise.

### GetHybridOk

`func (o *Entitlements) GetHybridOk() (*TheHybridSchema, bool)`

GetHybridOk returns a tuple with the Hybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrid

`func (o *Entitlements) SetHybrid(v TheHybridSchema)`

SetHybrid sets Hybrid field to given value.


### GetHybridAutoDiscoverProperties

`func (o *Entitlements) GetHybridAutoDiscoverProperties() bool`

GetHybridAutoDiscoverProperties returns the HybridAutoDiscoverProperties field if non-nil, zero value otherwise.

### GetHybridAutoDiscoverPropertiesOk

`func (o *Entitlements) GetHybridAutoDiscoverPropertiesOk() (*bool, bool)`

GetHybridAutoDiscoverPropertiesOk returns a tuple with the HybridAutoDiscoverProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybridAutoDiscoverProperties

`func (o *Entitlements) SetHybridAutoDiscoverProperties(v bool)`

SetHybridAutoDiscoverProperties sets HybridAutoDiscoverProperties field to given value.


### GetHybridInsight

`func (o *Entitlements) GetHybridInsight() bool`

GetHybridInsight returns the HybridInsight field if non-nil, zero value otherwise.

### GetHybridInsightOk

`func (o *Entitlements) GetHybridInsightOk() (*bool, bool)`

GetHybridInsightOk returns a tuple with the HybridInsight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybridInsight

`func (o *Entitlements) SetHybridInsight(v bool)`

SetHybridInsight sets HybridInsight field to given value.


### GetKpiDashboard

`func (o *Entitlements) GetKpiDashboard() KpiDashboard`

GetKpiDashboard returns the KpiDashboard field if non-nil, zero value otherwise.

### GetKpiDashboardOk

`func (o *Entitlements) GetKpiDashboardOk() (*KpiDashboard, bool)`

GetKpiDashboardOk returns a tuple with the KpiDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpiDashboard

`func (o *Entitlements) SetKpiDashboard(v KpiDashboard)`

SetKpiDashboard sets KpiDashboard field to given value.


### GetLoadBalancer

`func (o *Entitlements) GetLoadBalancer() LoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *Entitlements) GetLoadBalancerOk() (*LoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *Entitlements) SetLoadBalancer(v LoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.


### GetMessaging

`func (o *Entitlements) GetMessaging() Messaging`

GetMessaging returns the Messaging field if non-nil, zero value otherwise.

### GetMessagingOk

`func (o *Entitlements) GetMessagingOk() (*Messaging, bool)`

GetMessagingOk returns a tuple with the Messaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessaging

`func (o *Entitlements) SetMessaging(v Messaging)`

SetMessaging sets Messaging field to given value.


### GetMonitoringCenter

`func (o *Entitlements) GetMonitoringCenter() MonitoringCenter`

GetMonitoringCenter returns the MonitoringCenter field if non-nil, zero value otherwise.

### GetMonitoringCenterOk

`func (o *Entitlements) GetMonitoringCenterOk() (*MonitoringCenter, bool)`

GetMonitoringCenterOk returns a tuple with the MonitoringCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringCenter

`func (o *Entitlements) SetMonitoringCenter(v MonitoringCenter)`

SetMonitoringCenter sets MonitoringCenter field to given value.


### GetMqAdvancedFeatures

`func (o *Entitlements) GetMqAdvancedFeatures() MqAdvancedFeatures`

GetMqAdvancedFeatures returns the MqAdvancedFeatures field if non-nil, zero value otherwise.

### GetMqAdvancedFeaturesOk

`func (o *Entitlements) GetMqAdvancedFeaturesOk() (*MqAdvancedFeatures, bool)`

GetMqAdvancedFeaturesOk returns a tuple with the MqAdvancedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqAdvancedFeatures

`func (o *Entitlements) SetMqAdvancedFeatures(v MqAdvancedFeatures)`

SetMqAdvancedFeatures sets MqAdvancedFeatures field to given value.


### GetMqMessages

`func (o *Entitlements) GetMqMessages() MqMessages`

GetMqMessages returns the MqMessages field if non-nil, zero value otherwise.

### GetMqMessagesOk

`func (o *Entitlements) GetMqMessagesOk() (*MqMessages, bool)`

GetMqMessagesOk returns a tuple with the MqMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqMessages

`func (o *Entitlements) SetMqMessages(v MqMessages)`

SetMqMessages sets MqMessages field to given value.


### GetMqRequests

`func (o *Entitlements) GetMqRequests() MqRequests`

GetMqRequests returns the MqRequests field if non-nil, zero value otherwise.

### GetMqRequestsOk

`func (o *Entitlements) GetMqRequestsOk() (*MqRequests, bool)`

GetMqRequestsOk returns a tuple with the MqRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqRequests

`func (o *Entitlements) SetMqRequests(v MqRequests)`

SetMqRequests sets MqRequests field to given value.


### GetObjectStoreKeys

`func (o *Entitlements) GetObjectStoreKeys() ObjectStoreKeys`

GetObjectStoreKeys returns the ObjectStoreKeys field if non-nil, zero value otherwise.

### GetObjectStoreKeysOk

`func (o *Entitlements) GetObjectStoreKeysOk() (*ObjectStoreKeys, bool)`

GetObjectStoreKeysOk returns a tuple with the ObjectStoreKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreKeys

`func (o *Entitlements) SetObjectStoreKeys(v ObjectStoreKeys)`

SetObjectStoreKeys sets ObjectStoreKeys field to given value.


### GetObjectStoreRequestUnits

`func (o *Entitlements) GetObjectStoreRequestUnits() ObjectStoreRequestUnits`

GetObjectStoreRequestUnits returns the ObjectStoreRequestUnits field if non-nil, zero value otherwise.

### GetObjectStoreRequestUnitsOk

`func (o *Entitlements) GetObjectStoreRequestUnitsOk() (*ObjectStoreRequestUnits, bool)`

GetObjectStoreRequestUnitsOk returns a tuple with the ObjectStoreRequestUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreRequestUnits

`func (o *Entitlements) SetObjectStoreRequestUnits(v ObjectStoreRequestUnits)`

SetObjectStoreRequestUnits sets ObjectStoreRequestUnits field to given value.


### GetPartnersProduction

`func (o *Entitlements) GetPartnersProduction() PartnersProduction`

GetPartnersProduction returns the PartnersProduction field if non-nil, zero value otherwise.

### GetPartnersProductionOk

`func (o *Entitlements) GetPartnersProductionOk() (*PartnersProduction, bool)`

GetPartnersProductionOk returns a tuple with the PartnersProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnersProduction

`func (o *Entitlements) SetPartnersProduction(v PartnersProduction)`

SetPartnersProduction sets PartnersProduction field to given value.


### GetPartnersSandbox

`func (o *Entitlements) GetPartnersSandbox() PartnersSandbox`

GetPartnersSandbox returns the PartnersSandbox field if non-nil, zero value otherwise.

### GetPartnersSandboxOk

`func (o *Entitlements) GetPartnersSandboxOk() (*PartnersSandbox, bool)`

GetPartnersSandboxOk returns a tuple with the PartnersSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnersSandbox

`func (o *Entitlements) SetPartnersSandbox(v PartnersSandbox)`

SetPartnersSandbox sets PartnersSandbox field to given value.


### GetPcf

`func (o *Entitlements) GetPcf() bool`

GetPcf returns the Pcf field if non-nil, zero value otherwise.

### GetPcfOk

`func (o *Entitlements) GetPcfOk() (*bool, bool)`

GetPcfOk returns a tuple with the Pcf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcf

`func (o *Entitlements) SetPcf(v bool)`

SetPcf sets Pcf field to given value.


### GetRuntimeFabric

`func (o *Entitlements) GetRuntimeFabric() bool`

GetRuntimeFabric returns the RuntimeFabric field if non-nil, zero value otherwise.

### GetRuntimeFabricOk

`func (o *Entitlements) GetRuntimeFabricOk() (*bool, bool)`

GetRuntimeFabricOk returns a tuple with the RuntimeFabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeFabric

`func (o *Entitlements) SetRuntimeFabric(v bool)`

SetRuntimeFabric sets RuntimeFabric field to given value.


### GetRuntimeFabricCloud

`func (o *Entitlements) GetRuntimeFabricCloud() RuntimeFabricCloud`

GetRuntimeFabricCloud returns the RuntimeFabricCloud field if non-nil, zero value otherwise.

### GetRuntimeFabricCloudOk

`func (o *Entitlements) GetRuntimeFabricCloudOk() (*RuntimeFabricCloud, bool)`

GetRuntimeFabricCloudOk returns a tuple with the RuntimeFabricCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeFabricCloud

`func (o *Entitlements) SetRuntimeFabricCloud(v RuntimeFabricCloud)`

SetRuntimeFabricCloud sets RuntimeFabricCloud field to given value.


### GetServiceMesh

`func (o *Entitlements) GetServiceMesh() ServiceMesh`

GetServiceMesh returns the ServiceMesh field if non-nil, zero value otherwise.

### GetServiceMeshOk

`func (o *Entitlements) GetServiceMeshOk() (*ServiceMesh, bool)`

GetServiceMeshOk returns a tuple with the ServiceMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMesh

`func (o *Entitlements) SetServiceMesh(v ServiceMesh)`

SetServiceMesh sets ServiceMesh field to given value.


### GetStaticIps

`func (o *Entitlements) GetStaticIps() StaticIps`

GetStaticIps returns the StaticIps field if non-nil, zero value otherwise.

### GetStaticIpsOk

`func (o *Entitlements) GetStaticIpsOk() (*StaticIps, bool)`

GetStaticIpsOk returns a tuple with the StaticIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIps

`func (o *Entitlements) SetStaticIps(v StaticIps)`

SetStaticIps sets StaticIps field to given value.


### GetTradingPartnersProduction

`func (o *Entitlements) GetTradingPartnersProduction() TradingPartnersProduction`

GetTradingPartnersProduction returns the TradingPartnersProduction field if non-nil, zero value otherwise.

### GetTradingPartnersProductionOk

`func (o *Entitlements) GetTradingPartnersProductionOk() (*TradingPartnersProduction, bool)`

GetTradingPartnersProductionOk returns a tuple with the TradingPartnersProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingPartnersProduction

`func (o *Entitlements) SetTradingPartnersProduction(v TradingPartnersProduction)`

SetTradingPartnersProduction sets TradingPartnersProduction field to given value.


### GetTradingPartnersSandbox

`func (o *Entitlements) GetTradingPartnersSandbox() TradingPartnersSandbox`

GetTradingPartnersSandbox returns the TradingPartnersSandbox field if non-nil, zero value otherwise.

### GetTradingPartnersSandboxOk

`func (o *Entitlements) GetTradingPartnersSandboxOk() (*TradingPartnersSandbox, bool)`

GetTradingPartnersSandboxOk returns a tuple with the TradingPartnersSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingPartnersSandbox

`func (o *Entitlements) SetTradingPartnersSandbox(v TradingPartnersSandbox)`

SetTradingPartnersSandbox sets TradingPartnersSandbox field to given value.


### GetVCoresDesign

`func (o *Entitlements) GetVCoresDesign() VCoresDesign`

GetVCoresDesign returns the VCoresDesign field if non-nil, zero value otherwise.

### GetVCoresDesignOk

`func (o *Entitlements) GetVCoresDesignOk() (*VCoresDesign, bool)`

GetVCoresDesignOk returns a tuple with the VCoresDesign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCoresDesign

`func (o *Entitlements) SetVCoresDesign(v VCoresDesign)`

SetVCoresDesign sets VCoresDesign field to given value.


### GetVCoresProduction

`func (o *Entitlements) GetVCoresProduction() VCoresProduction`

GetVCoresProduction returns the VCoresProduction field if non-nil, zero value otherwise.

### GetVCoresProductionOk

`func (o *Entitlements) GetVCoresProductionOk() (*VCoresProduction, bool)`

GetVCoresProductionOk returns a tuple with the VCoresProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCoresProduction

`func (o *Entitlements) SetVCoresProduction(v VCoresProduction)`

SetVCoresProduction sets VCoresProduction field to given value.


### GetVCoresSandbox

`func (o *Entitlements) GetVCoresSandbox() VCoresSandbox`

GetVCoresSandbox returns the VCoresSandbox field if non-nil, zero value otherwise.

### GetVCoresSandboxOk

`func (o *Entitlements) GetVCoresSandboxOk() (*VCoresSandbox, bool)`

GetVCoresSandboxOk returns a tuple with the VCoresSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCoresSandbox

`func (o *Entitlements) SetVCoresSandbox(v VCoresSandbox)`

SetVCoresSandbox sets VCoresSandbox field to given value.


### GetVpcs

`func (o *Entitlements) GetVpcs() Vpcs`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *Entitlements) GetVpcsOk() (*Vpcs, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *Entitlements) SetVpcs(v Vpcs)`

SetVpcs sets Vpcs field to given value.


### GetVpns

`func (o *Entitlements) GetVpns() Vpns`

GetVpns returns the Vpns field if non-nil, zero value otherwise.

### GetVpnsOk

`func (o *Entitlements) GetVpnsOk() (*Vpns, bool)`

GetVpnsOk returns a tuple with the Vpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpns

`func (o *Entitlements) SetVpns(v Vpns)`

SetVpns sets Vpns field to given value.


### GetWorkerClouds

`func (o *Entitlements) GetWorkerClouds() WorkerClouds`

GetWorkerClouds returns the WorkerClouds field if non-nil, zero value otherwise.

### GetWorkerCloudsOk

`func (o *Entitlements) GetWorkerCloudsOk() (*WorkerClouds, bool)`

GetWorkerCloudsOk returns a tuple with the WorkerClouds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerClouds

`func (o *Entitlements) SetWorkerClouds(v WorkerClouds)`

SetWorkerClouds sets WorkerClouds field to given value.


### GetWorkerLoggingOverride

`func (o *Entitlements) GetWorkerLoggingOverride() WorkerLoggingOverride`

GetWorkerLoggingOverride returns the WorkerLoggingOverride field if non-nil, zero value otherwise.

### GetWorkerLoggingOverrideOk

`func (o *Entitlements) GetWorkerLoggingOverrideOk() (*WorkerLoggingOverride, bool)`

GetWorkerLoggingOverrideOk returns a tuple with the WorkerLoggingOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerLoggingOverride

`func (o *Entitlements) SetWorkerLoggingOverride(v WorkerLoggingOverride)`

SetWorkerLoggingOverride sets WorkerLoggingOverride field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


