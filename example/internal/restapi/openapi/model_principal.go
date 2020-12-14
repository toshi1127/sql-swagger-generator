/*
 * TestService
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// Principal Security principal for validating that a user is authorized to execute certain actions
type Principal struct {
	UserId *string `json:"userId,omitempty"`
	Permissions *[]string `json:"permissions,omitempty"`
}

// NewPrincipal instantiates a new Principal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipal() *Principal {
	this := Principal{}
	return &this
}

// NewPrincipalWithDefaults instantiates a new Principal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalWithDefaults() *Principal {
	this := Principal{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Principal) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Principal) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Principal) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Principal) SetUserId(v string) {
	o.UserId = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *Principal) GetPermissions() []string {
	if o == nil || o.Permissions == nil {
		var ret []string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Principal) GetPermissionsOk() (*[]string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *Principal) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *Principal) SetPermissions(v []string) {
	o.Permissions = &v
}

func (o Principal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullablePrincipal struct {
	value *Principal
	isSet bool
}

func (v NullablePrincipal) Get() *Principal {
	return v.value
}

func (v *NullablePrincipal) Set(val *Principal) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipal(val *Principal) *NullablePrincipal {
	return &NullablePrincipal{value: val, isSet: true}
}

func (v NullablePrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
