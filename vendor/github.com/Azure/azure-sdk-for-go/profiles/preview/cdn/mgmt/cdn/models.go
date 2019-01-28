// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package cdn

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2017-10-12/cdn"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CacheBehavior = original.CacheBehavior

const (
	BypassCache  CacheBehavior = original.BypassCache
	Override     CacheBehavior = original.Override
	SetIfMissing CacheBehavior = original.SetIfMissing
)

type CustomDomainResourceState = original.CustomDomainResourceState

const (
	Active   CustomDomainResourceState = original.Active
	Creating CustomDomainResourceState = original.Creating
	Deleting CustomDomainResourceState = original.Deleting
)

type CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningState

const (
	Disabled  CustomHTTPSProvisioningState = original.Disabled
	Disabling CustomHTTPSProvisioningState = original.Disabling
	Enabled   CustomHTTPSProvisioningState = original.Enabled
	Enabling  CustomHTTPSProvisioningState = original.Enabling
	Failed    CustomHTTPSProvisioningState = original.Failed
)

type CustomHTTPSProvisioningSubstate = original.CustomHTTPSProvisioningSubstate

const (
	CertificateDeleted                            CustomHTTPSProvisioningSubstate = original.CertificateDeleted
	CertificateDeployed                           CustomHTTPSProvisioningSubstate = original.CertificateDeployed
	DeletingCertificate                           CustomHTTPSProvisioningSubstate = original.DeletingCertificate
	DeployingCertificate                          CustomHTTPSProvisioningSubstate = original.DeployingCertificate
	DomainControlValidationRequestApproved        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestApproved
	DomainControlValidationRequestRejected        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestRejected
	DomainControlValidationRequestTimedOut        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestTimedOut
	IssuingCertificate                            CustomHTTPSProvisioningSubstate = original.IssuingCertificate
	PendingDomainControlValidationREquestApproval CustomHTTPSProvisioningSubstate = original.PendingDomainControlValidationREquestApproval
	SubmittingDomainControlValidationRequest      CustomHTTPSProvisioningSubstate = original.SubmittingDomainControlValidationRequest
)

type EndpointResourceState = original.EndpointResourceState

const (
	EndpointResourceStateCreating EndpointResourceState = original.EndpointResourceStateCreating
	EndpointResourceStateDeleting EndpointResourceState = original.EndpointResourceStateDeleting
	EndpointResourceStateRunning  EndpointResourceState = original.EndpointResourceStateRunning
	EndpointResourceStateStarting EndpointResourceState = original.EndpointResourceStateStarting
	EndpointResourceStateStopped  EndpointResourceState = original.EndpointResourceStateStopped
	EndpointResourceStateStopping EndpointResourceState = original.EndpointResourceStateStopping
)

type GeoFilterActions = original.GeoFilterActions

const (
	Allow GeoFilterActions = original.Allow
	Block GeoFilterActions = original.Block
)

type MatchType = original.MatchType

const (
	Literal  MatchType = original.Literal
	Wildcard MatchType = original.Wildcard
)

type Name = original.Name

const (
	NameCacheExpiration    Name = original.NameCacheExpiration
	NameDeliveryRuleAction Name = original.NameDeliveryRuleAction
)

type NameBasicDeliveryRuleCondition = original.NameBasicDeliveryRuleCondition

const (
	NameDeliveryRuleCondition NameBasicDeliveryRuleCondition = original.NameDeliveryRuleCondition
	NameURLFileExtension      NameBasicDeliveryRuleCondition = original.NameURLFileExtension
	NameURLPath               NameBasicDeliveryRuleCondition = original.NameURLPath
)

type OptimizationType = original.OptimizationType

const (
	DynamicSiteAcceleration     OptimizationType = original.DynamicSiteAcceleration
	GeneralMediaStreaming       OptimizationType = original.GeneralMediaStreaming
	GeneralWebDelivery          OptimizationType = original.GeneralWebDelivery
	LargeFileDownload           OptimizationType = original.LargeFileDownload
	VideoOnDemandMediaStreaming OptimizationType = original.VideoOnDemandMediaStreaming
)

type OriginResourceState = original.OriginResourceState

const (
	OriginResourceStateActive   OriginResourceState = original.OriginResourceStateActive
	OriginResourceStateCreating OriginResourceState = original.OriginResourceStateCreating
	OriginResourceStateDeleting OriginResourceState = original.OriginResourceStateDeleting
)

type ProfileResourceState = original.ProfileResourceState

const (
	ProfileResourceStateActive   ProfileResourceState = original.ProfileResourceStateActive
	ProfileResourceStateCreating ProfileResourceState = original.ProfileResourceStateCreating
	ProfileResourceStateDeleting ProfileResourceState = original.ProfileResourceStateDeleting
	ProfileResourceStateDisabled ProfileResourceState = original.ProfileResourceStateDisabled
)

type QueryStringCachingBehavior = original.QueryStringCachingBehavior

const (
	BypassCaching     QueryStringCachingBehavior = original.BypassCaching
	IgnoreQueryString QueryStringCachingBehavior = original.IgnoreQueryString
	NotSet            QueryStringCachingBehavior = original.NotSet
	UseQueryString    QueryStringCachingBehavior = original.UseQueryString
)

type ResourceType = original.ResourceType

const (
	MicrosoftCdnProfilesEndpoints ResourceType = original.MicrosoftCdnProfilesEndpoints
)

type SkuName = original.SkuName

const (
	CustomVerizon     SkuName = original.CustomVerizon
	PremiumVerizon    SkuName = original.PremiumVerizon
	StandardAkamai    SkuName = original.StandardAkamai
	StandardChinaCdn  SkuName = original.StandardChinaCdn
	StandardMicrosoft SkuName = original.StandardMicrosoft
	StandardVerizon   SkuName = original.StandardVerizon
)

type BaseClient = original.BaseClient
type BasicDeliveryRuleAction = original.BasicDeliveryRuleAction
type BasicDeliveryRuleCondition = original.BasicDeliveryRuleCondition
type CacheExpirationActionParameters = original.CacheExpirationActionParameters
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CidrIPAddress = original.CidrIPAddress
type CustomDomain = original.CustomDomain
type CustomDomainListResult = original.CustomDomainListResult
type CustomDomainListResultIterator = original.CustomDomainListResultIterator
type CustomDomainListResultPage = original.CustomDomainListResultPage
type CustomDomainParameters = original.CustomDomainParameters
type CustomDomainProperties = original.CustomDomainProperties
type CustomDomainPropertiesParameters = original.CustomDomainPropertiesParameters
type CustomDomainsClient = original.CustomDomainsClient
type CustomDomainsCreateFuture = original.CustomDomainsCreateFuture
type CustomDomainsDeleteFuture = original.CustomDomainsDeleteFuture
type DeepCreatedOrigin = original.DeepCreatedOrigin
type DeepCreatedOriginProperties = original.DeepCreatedOriginProperties
type DeliveryRule = original.DeliveryRule
type DeliveryRuleAction = original.DeliveryRuleAction
type DeliveryRuleCacheExpirationAction = original.DeliveryRuleCacheExpirationAction
type DeliveryRuleCondition = original.DeliveryRuleCondition
type DeliveryRuleURLFileExtensionCondition = original.DeliveryRuleURLFileExtensionCondition
type DeliveryRuleURLPathCondition = original.DeliveryRuleURLPathCondition
type EdgeNode = original.EdgeNode
type EdgeNodeProperties = original.EdgeNodeProperties
type EdgeNodesClient = original.EdgeNodesClient
type EdgenodeResult = original.EdgenodeResult
type EdgenodeResultIterator = original.EdgenodeResultIterator
type EdgenodeResultPage = original.EdgenodeResultPage
type Endpoint = original.Endpoint
type EndpointListResult = original.EndpointListResult
type EndpointListResultIterator = original.EndpointListResultIterator
type EndpointListResultPage = original.EndpointListResultPage
type EndpointProperties = original.EndpointProperties
type EndpointPropertiesUpdateParameters = original.EndpointPropertiesUpdateParameters
type EndpointPropertiesUpdateParametersDeliveryPolicy = original.EndpointPropertiesUpdateParametersDeliveryPolicy
type EndpointUpdateParameters = original.EndpointUpdateParameters
type EndpointsClient = original.EndpointsClient
type EndpointsCreateFuture = original.EndpointsCreateFuture
type EndpointsDeleteFuture = original.EndpointsDeleteFuture
type EndpointsLoadContentFuture = original.EndpointsLoadContentFuture
type EndpointsPurgeContentFuture = original.EndpointsPurgeContentFuture
type EndpointsStartFuture = original.EndpointsStartFuture
type EndpointsStopFuture = original.EndpointsStopFuture
type EndpointsUpdateFuture = original.EndpointsUpdateFuture
type ErrorResponse = original.ErrorResponse
type GeoFilter = original.GeoFilter
type IPAddressGroup = original.IPAddressGroup
type LoadParameters = original.LoadParameters
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsListResult = original.OperationsListResult
type OperationsListResultIterator = original.OperationsListResultIterator
type OperationsListResultPage = original.OperationsListResultPage
type Origin = original.Origin
type OriginListResult = original.OriginListResult
type OriginListResultIterator = original.OriginListResultIterator
type OriginListResultPage = original.OriginListResultPage
type OriginProperties = original.OriginProperties
type OriginPropertiesParameters = original.OriginPropertiesParameters
type OriginUpdateParameters = original.OriginUpdateParameters
type OriginsClient = original.OriginsClient
type OriginsUpdateFuture = original.OriginsUpdateFuture
type Profile = original.Profile
type ProfileListResult = original.ProfileListResult
type ProfileListResultIterator = original.ProfileListResultIterator
type ProfileListResultPage = original.ProfileListResultPage
type ProfileProperties = original.ProfileProperties
type ProfileUpdateParameters = original.ProfileUpdateParameters
type ProfilesClient = original.ProfilesClient
type ProfilesCreateFuture = original.ProfilesCreateFuture
type ProfilesDeleteFuture = original.ProfilesDeleteFuture
type ProfilesUpdateFuture = original.ProfilesUpdateFuture
type ProxyResource = original.ProxyResource
type PurgeParameters = original.PurgeParameters
type Resource = original.Resource
type ResourceUsage = original.ResourceUsage
type ResourceUsageClient = original.ResourceUsageClient
type ResourceUsageListResult = original.ResourceUsageListResult
type ResourceUsageListResultIterator = original.ResourceUsageListResultIterator
type ResourceUsageListResultPage = original.ResourceUsageListResultPage
type Sku = original.Sku
type SsoURI = original.SsoURI
type SupportedOptimizationTypesListResult = original.SupportedOptimizationTypesListResult
type TrackedResource = original.TrackedResource
type URLFileExtensionConditionParameters = original.URLFileExtensionConditionParameters
type URLPathConditionParameters = original.URLPathConditionParameters
type ValidateCustomDomainInput = original.ValidateCustomDomainInput
type ValidateCustomDomainOutput = original.ValidateCustomDomainOutput
type ValidateProbeInput = original.ValidateProbeInput
type ValidateProbeOutput = original.ValidateProbeOutput

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCustomDomainListResultIterator(page CustomDomainListResultPage) CustomDomainListResultIterator {
	return original.NewCustomDomainListResultIterator(page)
}
func NewCustomDomainListResultPage(getNextPage func(context.Context, CustomDomainListResult) (CustomDomainListResult, error)) CustomDomainListResultPage {
	return original.NewCustomDomainListResultPage(getNextPage)
}
func NewCustomDomainsClient(subscriptionID string) CustomDomainsClient {
	return original.NewCustomDomainsClient(subscriptionID)
}
func NewCustomDomainsClientWithBaseURI(baseURI string, subscriptionID string) CustomDomainsClient {
	return original.NewCustomDomainsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEdgeNodesClient(subscriptionID string) EdgeNodesClient {
	return original.NewEdgeNodesClient(subscriptionID)
}
func NewEdgeNodesClientWithBaseURI(baseURI string, subscriptionID string) EdgeNodesClient {
	return original.NewEdgeNodesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEdgenodeResultIterator(page EdgenodeResultPage) EdgenodeResultIterator {
	return original.NewEdgenodeResultIterator(page)
}
func NewEdgenodeResultPage(getNextPage func(context.Context, EdgenodeResult) (EdgenodeResult, error)) EdgenodeResultPage {
	return original.NewEdgenodeResultPage(getNextPage)
}
func NewEndpointListResultIterator(page EndpointListResultPage) EndpointListResultIterator {
	return original.NewEndpointListResultIterator(page)
}
func NewEndpointListResultPage(getNextPage func(context.Context, EndpointListResult) (EndpointListResult, error)) EndpointListResultPage {
	return original.NewEndpointListResultPage(getNextPage)
}
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsListResultIterator(page OperationsListResultPage) OperationsListResultIterator {
	return original.NewOperationsListResultIterator(page)
}
func NewOperationsListResultPage(getNextPage func(context.Context, OperationsListResult) (OperationsListResult, error)) OperationsListResultPage {
	return original.NewOperationsListResultPage(getNextPage)
}
func NewOriginListResultIterator(page OriginListResultPage) OriginListResultIterator {
	return original.NewOriginListResultIterator(page)
}
func NewOriginListResultPage(getNextPage func(context.Context, OriginListResult) (OriginListResult, error)) OriginListResultPage {
	return original.NewOriginListResultPage(getNextPage)
}
func NewOriginsClient(subscriptionID string) OriginsClient {
	return original.NewOriginsClient(subscriptionID)
}
func NewOriginsClientWithBaseURI(baseURI string, subscriptionID string) OriginsClient {
	return original.NewOriginsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProfileListResultIterator(page ProfileListResultPage) ProfileListResultIterator {
	return original.NewProfileListResultIterator(page)
}
func NewProfileListResultPage(getNextPage func(context.Context, ProfileListResult) (ProfileListResult, error)) ProfileListResultPage {
	return original.NewProfileListResultPage(getNextPage)
}
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return original.NewProfilesClient(subscriptionID)
}
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return original.NewProfilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceUsageClient(subscriptionID string) ResourceUsageClient {
	return original.NewResourceUsageClient(subscriptionID)
}
func NewResourceUsageClientWithBaseURI(baseURI string, subscriptionID string) ResourceUsageClient {
	return original.NewResourceUsageClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceUsageListResultIterator(page ResourceUsageListResultPage) ResourceUsageListResultIterator {
	return original.NewResourceUsageListResultIterator(page)
}
func NewResourceUsageListResultPage(getNextPage func(context.Context, ResourceUsageListResult) (ResourceUsageListResult, error)) ResourceUsageListResultPage {
	return original.NewResourceUsageListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCacheBehaviorValues() []CacheBehavior {
	return original.PossibleCacheBehaviorValues()
}
func PossibleCustomDomainResourceStateValues() []CustomDomainResourceState {
	return original.PossibleCustomDomainResourceStateValues()
}
func PossibleCustomHTTPSProvisioningStateValues() []CustomHTTPSProvisioningState {
	return original.PossibleCustomHTTPSProvisioningStateValues()
}
func PossibleCustomHTTPSProvisioningSubstateValues() []CustomHTTPSProvisioningSubstate {
	return original.PossibleCustomHTTPSProvisioningSubstateValues()
}
func PossibleEndpointResourceStateValues() []EndpointResourceState {
	return original.PossibleEndpointResourceStateValues()
}
func PossibleGeoFilterActionsValues() []GeoFilterActions {
	return original.PossibleGeoFilterActionsValues()
}
func PossibleMatchTypeValues() []MatchType {
	return original.PossibleMatchTypeValues()
}
func PossibleNameBasicDeliveryRuleConditionValues() []NameBasicDeliveryRuleCondition {
	return original.PossibleNameBasicDeliveryRuleConditionValues()
}
func PossibleNameValues() []Name {
	return original.PossibleNameValues()
}
func PossibleOptimizationTypeValues() []OptimizationType {
	return original.PossibleOptimizationTypeValues()
}
func PossibleOriginResourceStateValues() []OriginResourceState {
	return original.PossibleOriginResourceStateValues()
}
func PossibleProfileResourceStateValues() []ProfileResourceState {
	return original.PossibleProfileResourceStateValues()
}
func PossibleQueryStringCachingBehaviorValues() []QueryStringCachingBehavior {
	return original.PossibleQueryStringCachingBehaviorValues()
}
func PossibleResourceTypeValues() []ResourceType {
	return original.PossibleResourceTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
