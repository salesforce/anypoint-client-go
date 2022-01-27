# Urls4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignOn** | Pointer to **string** | The Anypoint Platform URL where users must sign in. | [optional] 
**SignOut** | Pointer to **string** | URL to redirect sign out requests, so users both sign out of the Anypoint Platform and have their SAML user&#39;s status set to signed out. | [optional] 

## Methods

### NewUrls4

`func NewUrls4() *Urls4`

NewUrls4 instantiates a new Urls4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrls4WithDefaults

`func NewUrls4WithDefaults() *Urls4`

NewUrls4WithDefaults instantiates a new Urls4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignOn

`func (o *Urls4) GetSignOn() string`

GetSignOn returns the SignOn field if non-nil, zero value otherwise.

### GetSignOnOk

`func (o *Urls4) GetSignOnOk() (*string, bool)`

GetSignOnOk returns a tuple with the SignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOn

`func (o *Urls4) SetSignOn(v string)`

SetSignOn sets SignOn field to given value.

### HasSignOn

`func (o *Urls4) HasSignOn() bool`

HasSignOn returns a boolean if a field has been set.

### GetSignOut

`func (o *Urls4) GetSignOut() string`

GetSignOut returns the SignOut field if non-nil, zero value otherwise.

### GetSignOutOk

`func (o *Urls4) GetSignOutOk() (*string, bool)`

GetSignOutOk returns a tuple with the SignOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOut

`func (o *Urls4) SetSignOut(v string)`

SetSignOut sets SignOut field to given value.

### HasSignOut

`func (o *Urls4) HasSignOut() bool`

HasSignOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


