/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"time"
)

// EventViewForGroup struct for EventViewForGroup
type EventViewForGroup struct {
	// (**For audit only**), Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **userId** parameter.
	ApiKeyId *string `json:"apiKeyId,omitempty"`
	// Date and time when this event occurred. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	EventTypeName *UserEventTypeForGroup `json:"eventTypeName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the event.
	Id *string `json:"id,omitempty"`
	// (**For audit only**), Flag that indicates whether a MongoDB employee triggered the specified event.
	IsGlobalAdmin *bool `json:"isGlobalAdmin,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization to which these events apply.
	OrgId *string `json:"orgId,omitempty"`
	// (**For audit only**), Public part of the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **username** parameter.
	PublicKey *string `json:"publicKey,omitempty"`
	Raw *Raw `json:"raw,omitempty"`
	// (**For audit only**), IPv4 or IPv6 address from which the user triggered this event.
	RemoteAddress *string `json:"remoteAddress,omitempty"`
	// (**For audit only**), Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn't return the **apiKeyId** parameter.
	UserId *string `json:"userId,omitempty"`
	// (**For audit only**), Email address for the user who triggered this event. If this resource returns this parameter, it doesn't return the **publicApiKey** parameter.
	Username *string `json:"username,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the alert associated with the event.
	AlertId *string `json:"alertId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the alert configuration associated with the **alertId**.
	AlertConfigId *string `json:"alertConfigId,omitempty"`
	// Human-readable label of the replica set associated with the event.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// IANA port on which the MongoDB process listens for requests.
	Port *int32 `json:"port,omitempty"`
	// Unique 24-hexadecimal digit string that identifies of the invoice associated with the event.
	InvoiceId *string `json:"invoiceId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the invoice payment associated with this event.
	PaymentId *string `json:"paymentId,omitempty"`
	// Human-readable label of the shard associated with the event.
	ShardName *string `json:"shardName,omitempty"`
	// (**For audit only**), Human-readable label of the collection on which the event occurred. The resource returns this parameter when the **eventTypeName** includes `DATA_EXPLORER`.
	Collection *string `json:"collection,omitempty"`
	// (**For audit only**), Human-readable label of the database on which this incident occurred. The resource returns this parameter when `\"eventTypeName\" : \"DATA_EXPLORER\"` or `\"eventTypeName\" : \"DATA_EXPLORER_CRUD\"`.
	Database *string `json:"database,omitempty"`
	// (**For audit only**), Action that the database attempted to execute when the event triggered. The response returns this parameter when `eventTypeName\" : \"DATA_EXPLORER\"`.
	OpType *string `json:"opType,omitempty"`
	CurrentValue *HostMetricValueView `json:"currentValue,omitempty"`
	// Human-readable label of the metric associated with the **alertId**. This field may change type of **currentValue** field.
	MetricName *string `json:"metricName,omitempty"`
	// (**For audit only**), Entry in the list of source host addresses that the API key accepts and this event targets.
	WhitelistEntry *string `json:"whitelistEntry,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the endpoint associated with this event.
	EndpointId *string `json:"endpointId,omitempty"`
	// Unique identification string that the cloud provider uses to identify the private endpoint.
	ProviderEndpointId *string `json:"providerEndpointId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization team associated with this event.
	TeamId *string `json:"teamId,omitempty"`
	// Email address for the console user that this event targets. The resource returns this parameter when `\"eventTypeName\" : \"USER\"`.
	TargetUsername *string `json:"targetUsername,omitempty"`
}

// NewEventViewForGroup instantiates a new EventViewForGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventViewForGroup() *EventViewForGroup {
	this := EventViewForGroup{}
	return &this
}

// NewEventViewForGroupWithDefaults instantiates a new EventViewForGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventViewForGroupWithDefaults() *EventViewForGroup {
	this := EventViewForGroup{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetApiKeyId() string {
	if o == nil || o.ApiKeyId == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetApiKeyIdOk() (*string, bool) {
	if o == nil || o.ApiKeyId == nil {
		return nil, false
	}
	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasApiKeyId() bool {
	if o != nil && o.ApiKeyId != nil {
		return true
	}

	return false
}

// SetApiKeyId gets a reference to the given string and assigns it to the ApiKeyId field.
func (o *EventViewForGroup) SetApiKeyId(v string) {
	o.ApiKeyId = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *EventViewForGroup) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *EventViewForGroup) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *EventViewForGroup) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEventTypeName returns the EventTypeName field value if set, zero value otherwise.
func (o *EventViewForGroup) GetEventTypeName() UserEventTypeForGroup {
	if o == nil || o.EventTypeName == nil {
		var ret UserEventTypeForGroup
		return ret
	}
	return *o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetEventTypeNameOk() (*UserEventTypeForGroup, bool) {
	if o == nil || o.EventTypeName == nil {
		return nil, false
	}
	return o.EventTypeName, true
}

// HasEventTypeName returns a boolean if a field has been set.
func (o *EventViewForGroup) HasEventTypeName() bool {
	if o != nil && o.EventTypeName != nil {
		return true
	}

	return false
}

// SetEventTypeName gets a reference to the given UserEventTypeForGroup and assigns it to the EventTypeName field.
func (o *EventViewForGroup) SetEventTypeName(v UserEventTypeForGroup) {
	o.EventTypeName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *EventViewForGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventViewForGroup) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventViewForGroup) SetId(v string) {
	o.Id = &v
}

// GetIsGlobalAdmin returns the IsGlobalAdmin field value if set, zero value otherwise.
func (o *EventViewForGroup) GetIsGlobalAdmin() bool {
	if o == nil || o.IsGlobalAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsGlobalAdmin
}

// GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetIsGlobalAdminOk() (*bool, bool) {
	if o == nil || o.IsGlobalAdmin == nil {
		return nil, false
	}
	return o.IsGlobalAdmin, true
}

// HasIsGlobalAdmin returns a boolean if a field has been set.
func (o *EventViewForGroup) HasIsGlobalAdmin() bool {
	if o != nil && o.IsGlobalAdmin != nil {
		return true
	}

	return false
}

// SetIsGlobalAdmin gets a reference to the given bool and assigns it to the IsGlobalAdmin field.
func (o *EventViewForGroup) SetIsGlobalAdmin(v bool) {
	o.IsGlobalAdmin = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EventViewForGroup) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EventViewForGroup) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *EventViewForGroup) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *EventViewForGroup) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *EventViewForGroup) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *EventViewForGroup) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *EventViewForGroup) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *EventViewForGroup) GetRaw() Raw {
	if o == nil || o.Raw == nil {
		var ret Raw
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetRawOk() (*Raw, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *EventViewForGroup) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given Raw and assigns it to the Raw field.
func (o *EventViewForGroup) SetRaw(v Raw) {
	o.Raw = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *EventViewForGroup) GetRemoteAddress() string {
	if o == nil || o.RemoteAddress == nil {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetRemoteAddressOk() (*string, bool) {
	if o == nil || o.RemoteAddress == nil {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *EventViewForGroup) HasRemoteAddress() bool {
	if o != nil && o.RemoteAddress != nil {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *EventViewForGroup) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EventViewForGroup) SetUserId(v string) {
	o.UserId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *EventViewForGroup) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EventViewForGroup) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EventViewForGroup) SetUsername(v string) {
	o.Username = &v
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetAlertId() string {
	if o == nil || o.AlertId == nil {
		var ret string
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetAlertIdOk() (*string, bool) {
	if o == nil || o.AlertId == nil {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasAlertId() bool {
	if o != nil && o.AlertId != nil {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given string and assigns it to the AlertId field.
func (o *EventViewForGroup) SetAlertId(v string) {
	o.AlertId = &v
}

// GetAlertConfigId returns the AlertConfigId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetAlertConfigId() string {
	if o == nil || o.AlertConfigId == nil {
		var ret string
		return ret
	}
	return *o.AlertConfigId
}

// GetAlertConfigIdOk returns a tuple with the AlertConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetAlertConfigIdOk() (*string, bool) {
	if o == nil || o.AlertConfigId == nil {
		return nil, false
	}
	return o.AlertConfigId, true
}

// HasAlertConfigId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasAlertConfigId() bool {
	if o != nil && o.AlertConfigId != nil {
		return true
	}

	return false
}

// SetAlertConfigId gets a reference to the given string and assigns it to the AlertConfigId field.
func (o *EventViewForGroup) SetAlertConfigId(v string) {
	o.AlertConfigId = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise.
func (o *EventViewForGroup) GetReplicaSetName() string {
	if o == nil || o.ReplicaSetName == nil {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || o.ReplicaSetName == nil {
		return nil, false
	}
	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *EventViewForGroup) HasReplicaSetName() bool {
	if o != nil && o.ReplicaSetName != nil {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *EventViewForGroup) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *EventViewForGroup) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *EventViewForGroup) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *EventViewForGroup) SetPort(v int32) {
	o.Port = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetInvoiceId() string {
	if o == nil || o.InvoiceId == nil {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetInvoiceIdOk() (*string, bool) {
	if o == nil || o.InvoiceId == nil {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasInvoiceId() bool {
	if o != nil && o.InvoiceId != nil {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *EventViewForGroup) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *EventViewForGroup) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetShardName returns the ShardName field value if set, zero value otherwise.
func (o *EventViewForGroup) GetShardName() string {
	if o == nil || o.ShardName == nil {
		var ret string
		return ret
	}
	return *o.ShardName
}

// GetShardNameOk returns a tuple with the ShardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetShardNameOk() (*string, bool) {
	if o == nil || o.ShardName == nil {
		return nil, false
	}
	return o.ShardName, true
}

// HasShardName returns a boolean if a field has been set.
func (o *EventViewForGroup) HasShardName() bool {
	if o != nil && o.ShardName != nil {
		return true
	}

	return false
}

// SetShardName gets a reference to the given string and assigns it to the ShardName field.
func (o *EventViewForGroup) SetShardName(v string) {
	o.ShardName = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *EventViewForGroup) GetCollection() string {
	if o == nil || o.Collection == nil {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetCollectionOk() (*string, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *EventViewForGroup) HasCollection() bool {
	if o != nil && o.Collection != nil {
		return true
	}

	return false
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *EventViewForGroup) SetCollection(v string) {
	o.Collection = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *EventViewForGroup) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *EventViewForGroup) HasDatabase() bool {
	if o != nil && o.Database != nil {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *EventViewForGroup) SetDatabase(v string) {
	o.Database = &v
}

// GetOpType returns the OpType field value if set, zero value otherwise.
func (o *EventViewForGroup) GetOpType() string {
	if o == nil || o.OpType == nil {
		var ret string
		return ret
	}
	return *o.OpType
}

// GetOpTypeOk returns a tuple with the OpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetOpTypeOk() (*string, bool) {
	if o == nil || o.OpType == nil {
		return nil, false
	}
	return o.OpType, true
}

// HasOpType returns a boolean if a field has been set.
func (o *EventViewForGroup) HasOpType() bool {
	if o != nil && o.OpType != nil {
		return true
	}

	return false
}

// SetOpType gets a reference to the given string and assigns it to the OpType field.
func (o *EventViewForGroup) SetOpType(v string) {
	o.OpType = &v
}

// GetCurrentValue returns the CurrentValue field value if set, zero value otherwise.
func (o *EventViewForGroup) GetCurrentValue() HostMetricValueView {
	if o == nil || o.CurrentValue == nil {
		var ret HostMetricValueView
		return ret
	}
	return *o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetCurrentValueOk() (*HostMetricValueView, bool) {
	if o == nil || o.CurrentValue == nil {
		return nil, false
	}
	return o.CurrentValue, true
}

// HasCurrentValue returns a boolean if a field has been set.
func (o *EventViewForGroup) HasCurrentValue() bool {
	if o != nil && o.CurrentValue != nil {
		return true
	}

	return false
}

// SetCurrentValue gets a reference to the given HostMetricValueView and assigns it to the CurrentValue field.
func (o *EventViewForGroup) SetCurrentValue(v HostMetricValueView) {
	o.CurrentValue = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *EventViewForGroup) GetMetricName() string {
	if o == nil || o.MetricName == nil {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetMetricNameOk() (*string, bool) {
	if o == nil || o.MetricName == nil {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *EventViewForGroup) HasMetricName() bool {
	if o != nil && o.MetricName != nil {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *EventViewForGroup) SetMetricName(v string) {
	o.MetricName = &v
}

// GetWhitelistEntry returns the WhitelistEntry field value if set, zero value otherwise.
func (o *EventViewForGroup) GetWhitelistEntry() string {
	if o == nil || o.WhitelistEntry == nil {
		var ret string
		return ret
	}
	return *o.WhitelistEntry
}

// GetWhitelistEntryOk returns a tuple with the WhitelistEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetWhitelistEntryOk() (*string, bool) {
	if o == nil || o.WhitelistEntry == nil {
		return nil, false
	}
	return o.WhitelistEntry, true
}

// HasWhitelistEntry returns a boolean if a field has been set.
func (o *EventViewForGroup) HasWhitelistEntry() bool {
	if o != nil && o.WhitelistEntry != nil {
		return true
	}

	return false
}

// SetWhitelistEntry gets a reference to the given string and assigns it to the WhitelistEntry field.
func (o *EventViewForGroup) SetWhitelistEntry(v string) {
	o.WhitelistEntry = &v
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetEndpointId() string {
	if o == nil || o.EndpointId == nil {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetEndpointIdOk() (*string, bool) {
	if o == nil || o.EndpointId == nil {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasEndpointId() bool {
	if o != nil && o.EndpointId != nil {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *EventViewForGroup) SetEndpointId(v string) {
	o.EndpointId = &v
}

// GetProviderEndpointId returns the ProviderEndpointId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetProviderEndpointId() string {
	if o == nil || o.ProviderEndpointId == nil {
		var ret string
		return ret
	}
	return *o.ProviderEndpointId
}

// GetProviderEndpointIdOk returns a tuple with the ProviderEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetProviderEndpointIdOk() (*string, bool) {
	if o == nil || o.ProviderEndpointId == nil {
		return nil, false
	}
	return o.ProviderEndpointId, true
}

// HasProviderEndpointId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasProviderEndpointId() bool {
	if o != nil && o.ProviderEndpointId != nil {
		return true
	}

	return false
}

// SetProviderEndpointId gets a reference to the given string and assigns it to the ProviderEndpointId field.
func (o *EventViewForGroup) SetProviderEndpointId(v string) {
	o.ProviderEndpointId = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *EventViewForGroup) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *EventViewForGroup) HasTeamId() bool {
	if o != nil && o.TeamId != nil {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *EventViewForGroup) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTargetUsername returns the TargetUsername field value if set, zero value otherwise.
func (o *EventViewForGroup) GetTargetUsername() string {
	if o == nil || o.TargetUsername == nil {
		var ret string
		return ret
	}
	return *o.TargetUsername
}

// GetTargetUsernameOk returns a tuple with the TargetUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForGroup) GetTargetUsernameOk() (*string, bool) {
	if o == nil || o.TargetUsername == nil {
		return nil, false
	}
	return o.TargetUsername, true
}

// HasTargetUsername returns a boolean if a field has been set.
func (o *EventViewForGroup) HasTargetUsername() bool {
	if o != nil && o.TargetUsername != nil {
		return true
	}

	return false
}

// SetTargetUsername gets a reference to the given string and assigns it to the TargetUsername field.
func (o *EventViewForGroup) SetTargetUsername(v string) {
	o.TargetUsername = &v
}

func (o EventViewForGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKeyId != nil {
		toSerialize["apiKeyId"] = o.ApiKeyId
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.EventTypeName != nil {
		toSerialize["eventTypeName"] = o.EventTypeName
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsGlobalAdmin != nil {
		toSerialize["isGlobalAdmin"] = o.IsGlobalAdmin
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.OrgId != nil {
		toSerialize["orgId"] = o.OrgId
	}
	if o.PublicKey != nil {
		toSerialize["publicKey"] = o.PublicKey
	}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.RemoteAddress != nil {
		toSerialize["remoteAddress"] = o.RemoteAddress
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.AlertId != nil {
		toSerialize["alertId"] = o.AlertId
	}
	if o.AlertConfigId != nil {
		toSerialize["alertConfigId"] = o.AlertConfigId
	}
	if o.ReplicaSetName != nil {
		toSerialize["replicaSetName"] = o.ReplicaSetName
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.InvoiceId != nil {
		toSerialize["invoiceId"] = o.InvoiceId
	}
	if o.PaymentId != nil {
		toSerialize["paymentId"] = o.PaymentId
	}
	if o.ShardName != nil {
		toSerialize["shardName"] = o.ShardName
	}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.OpType != nil {
		toSerialize["opType"] = o.OpType
	}
	if o.CurrentValue != nil {
		toSerialize["currentValue"] = o.CurrentValue
	}
	if o.MetricName != nil {
		toSerialize["metricName"] = o.MetricName
	}
	if o.WhitelistEntry != nil {
		toSerialize["whitelistEntry"] = o.WhitelistEntry
	}
	if o.EndpointId != nil {
		toSerialize["endpointId"] = o.EndpointId
	}
	if o.ProviderEndpointId != nil {
		toSerialize["providerEndpointId"] = o.ProviderEndpointId
	}
	if o.TeamId != nil {
		toSerialize["teamId"] = o.TeamId
	}
	if o.TargetUsername != nil {
		toSerialize["targetUsername"] = o.TargetUsername
	}
	return json.Marshal(toSerialize)
}

type NullableEventViewForGroup struct {
	value *EventViewForGroup
	isSet bool
}

func (v NullableEventViewForGroup) Get() *EventViewForGroup {
	return v.value
}

func (v *NullableEventViewForGroup) Set(val *EventViewForGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableEventViewForGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableEventViewForGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventViewForGroup(val *EventViewForGroup) *NullableEventViewForGroup {
	return &NullableEventViewForGroup{value: val, isSet: true}
}

func (v NullableEventViewForGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventViewForGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

