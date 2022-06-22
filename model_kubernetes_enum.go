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

// KubernetesEnum the model 'KubernetesEnum'
type KubernetesEnum string

// List of KubernetesEnum
const (
	KUBERNETESENUM_K3_S    KubernetesEnum = "K3S"
	KUBERNETESENUM_MANAGED KubernetesEnum = "MANAGED"
)

// All allowed values of KubernetesEnum enum
var AllowedKubernetesEnumEnumValues = []KubernetesEnum{
	"K3S",
	"MANAGED",
}

func (v *KubernetesEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KubernetesEnum(value)
	for _, existing := range AllowedKubernetesEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KubernetesEnum", value)
}

// NewKubernetesEnumFromValue returns a pointer to a valid KubernetesEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKubernetesEnumFromValue(v string) (*KubernetesEnum, error) {
	ev := KubernetesEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KubernetesEnum: valid values are %v", v, AllowedKubernetesEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KubernetesEnum) IsValid() bool {
	for _, existing := range AllowedKubernetesEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KubernetesEnum value
func (v KubernetesEnum) Ptr() *KubernetesEnum {
	return &v
}

type NullableKubernetesEnum struct {
	value *KubernetesEnum
	isSet bool
}

func (v NullableKubernetesEnum) Get() *KubernetesEnum {
	return v.value
}

func (v *NullableKubernetesEnum) Set(val *KubernetesEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesEnum(val *KubernetesEnum) *NullableKubernetesEnum {
	return &NullableKubernetesEnum{value: val, isSet: true}
}

func (v NullableKubernetesEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}