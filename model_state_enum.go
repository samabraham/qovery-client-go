/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"fmt"
)

// StateEnum the model 'StateEnum'
type StateEnum string

// List of StateEnum
const (
	STATEENUM_BUILDING          StateEnum = "BUILDING"
	STATEENUM_CANCELED          StateEnum = "CANCELED"
	STATEENUM_CANCELING         StateEnum = "CANCELING"
	STATEENUM_DELETED           StateEnum = "DELETED"
	STATEENUM_DELETE_ERROR      StateEnum = "DELETE_ERROR"
	STATEENUM_DELETE_QUEUED     StateEnum = "DELETE_QUEUED"
	STATEENUM_DELETING          StateEnum = "DELETING"
	STATEENUM_DEPLOYED          StateEnum = "DEPLOYED"
	STATEENUM_DEPLOYING         StateEnum = "DEPLOYING"
	STATEENUM_DEPLOYMENT_ERROR  StateEnum = "DEPLOYMENT_ERROR"
	STATEENUM_DEPLOYMENT_QUEUED StateEnum = "DEPLOYMENT_QUEUED"
	STATEENUM_QUEUED            StateEnum = "QUEUED"
	STATEENUM_READY             StateEnum = "READY"
	STATEENUM_RUNNING           StateEnum = "RUNNING"
	STATEENUM_STOPPED           StateEnum = "STOPPED"
	STATEENUM_STOPPING          StateEnum = "STOPPING"
	STATEENUM_STOP_ERROR        StateEnum = "STOP_ERROR"
	STATEENUM_STOP_QUEUED       StateEnum = "STOP_QUEUED"
)

// All allowed values of StateEnum enum
var AllowedStateEnumEnumValues = []StateEnum{
	"BUILDING",
	"CANCELED",
	"CANCELING",
	"DELETED",
	"DELETE_ERROR",
	"DELETE_QUEUED",
	"DELETING",
	"DEPLOYED",
	"DEPLOYING",
	"DEPLOYMENT_ERROR",
	"DEPLOYMENT_QUEUED",
	"QUEUED",
	"READY",
	"RUNNING",
	"STOPPED",
	"STOPPING",
	"STOP_ERROR",
	"STOP_QUEUED",
}

func (v *StateEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StateEnum(value)
	for _, existing := range AllowedStateEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StateEnum", value)
}

// NewStateEnumFromValue returns a pointer to a valid StateEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStateEnumFromValue(v string) (*StateEnum, error) {
	ev := StateEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StateEnum: valid values are %v", v, AllowedStateEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StateEnum) IsValid() bool {
	for _, existing := range AllowedStateEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StateEnum value
func (v StateEnum) Ptr() *StateEnum {
	return &v
}

type NullableStateEnum struct {
	value *StateEnum
	isSet bool
}

func (v NullableStateEnum) Get() *StateEnum {
	return v.value
}

func (v *NullableStateEnum) Set(val *StateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStateEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateEnum(val *StateEnum) *NullableStateEnum {
	return &NullableStateEnum{value: val, isSet: true}
}

func (v NullableStateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}