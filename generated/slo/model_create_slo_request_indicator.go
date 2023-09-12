/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
	"fmt"
)

// CreateSloRequestIndicator - struct for CreateSloRequestIndicator
type CreateSloRequestIndicator struct {
	IndicatorPropertiesApmAvailability *IndicatorPropertiesApmAvailability
	IndicatorPropertiesApmLatency      *IndicatorPropertiesApmLatency
	IndicatorPropertiesCustomKql       *IndicatorPropertiesCustomKql
	IndicatorPropertiesCustomMetric    *IndicatorPropertiesCustomMetric
	IndicatorPropertiesHistogram       *IndicatorPropertiesHistogram
}

// IndicatorPropertiesApmAvailabilityAsCreateSloRequestIndicator is a convenience function that returns IndicatorPropertiesApmAvailability wrapped in CreateSloRequestIndicator
func IndicatorPropertiesApmAvailabilityAsCreateSloRequestIndicator(v *IndicatorPropertiesApmAvailability) CreateSloRequestIndicator {
	return CreateSloRequestIndicator{
		IndicatorPropertiesApmAvailability: v,
	}
}

// IndicatorPropertiesApmLatencyAsCreateSloRequestIndicator is a convenience function that returns IndicatorPropertiesApmLatency wrapped in CreateSloRequestIndicator
func IndicatorPropertiesApmLatencyAsCreateSloRequestIndicator(v *IndicatorPropertiesApmLatency) CreateSloRequestIndicator {
	return CreateSloRequestIndicator{
		IndicatorPropertiesApmLatency: v,
	}
}

// IndicatorPropertiesCustomKqlAsCreateSloRequestIndicator is a convenience function that returns IndicatorPropertiesCustomKql wrapped in CreateSloRequestIndicator
func IndicatorPropertiesCustomKqlAsCreateSloRequestIndicator(v *IndicatorPropertiesCustomKql) CreateSloRequestIndicator {
	return CreateSloRequestIndicator{
		IndicatorPropertiesCustomKql: v,
	}
}

// IndicatorPropertiesCustomMetricAsCreateSloRequestIndicator is a convenience function that returns IndicatorPropertiesCustomMetric wrapped in CreateSloRequestIndicator
func IndicatorPropertiesCustomMetricAsCreateSloRequestIndicator(v *IndicatorPropertiesCustomMetric) CreateSloRequestIndicator {
	return CreateSloRequestIndicator{
		IndicatorPropertiesCustomMetric: v,
	}
}

// IndicatorPropertiesHistogramAsCreateSloRequestIndicator is a convenience function that returns IndicatorPropertiesHistogram wrapped in CreateSloRequestIndicator
func IndicatorPropertiesHistogramAsCreateSloRequestIndicator(v *IndicatorPropertiesHistogram) CreateSloRequestIndicator {
	return CreateSloRequestIndicator{
		IndicatorPropertiesHistogram: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateSloRequestIndicator) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IndicatorPropertiesApmAvailability
	err = json.Unmarshal(data, &dst.IndicatorPropertiesApmAvailability)
	if err == nil {
		jsonIndicatorPropertiesApmAvailability, _ := json.Marshal(dst.IndicatorPropertiesApmAvailability)
		if string(jsonIndicatorPropertiesApmAvailability) == "{}" { // empty struct
			dst.IndicatorPropertiesApmAvailability = nil
		} else {
			match++
		}
	} else {
		dst.IndicatorPropertiesApmAvailability = nil
	}

	// try to unmarshal data into IndicatorPropertiesApmLatency
	err = json.Unmarshal(data, &dst.IndicatorPropertiesApmLatency)
	if err == nil {
		jsonIndicatorPropertiesApmLatency, _ := json.Marshal(dst.IndicatorPropertiesApmLatency)
		if string(jsonIndicatorPropertiesApmLatency) == "{}" { // empty struct
			dst.IndicatorPropertiesApmLatency = nil
		} else {
			match++
		}
	} else {
		dst.IndicatorPropertiesApmLatency = nil
	}

	// try to unmarshal data into IndicatorPropertiesCustomKql
	err = json.Unmarshal(data, &dst.IndicatorPropertiesCustomKql)
	if err == nil {
		jsonIndicatorPropertiesCustomKql, _ := json.Marshal(dst.IndicatorPropertiesCustomKql)
		if string(jsonIndicatorPropertiesCustomKql) == "{}" { // empty struct
			dst.IndicatorPropertiesCustomKql = nil
		} else {
			match++
		}
	} else {
		dst.IndicatorPropertiesCustomKql = nil
	}

	// try to unmarshal data into IndicatorPropertiesCustomMetric
	err = json.Unmarshal(data, &dst.IndicatorPropertiesCustomMetric)
	if err == nil {
		jsonIndicatorPropertiesCustomMetric, _ := json.Marshal(dst.IndicatorPropertiesCustomMetric)
		if string(jsonIndicatorPropertiesCustomMetric) == "{}" { // empty struct
			dst.IndicatorPropertiesCustomMetric = nil
		} else {
			match++
		}
	} else {
		dst.IndicatorPropertiesCustomMetric = nil
	}

	// try to unmarshal data into IndicatorPropertiesHistogram
	err = json.Unmarshal(data, &dst.IndicatorPropertiesHistogram)
	if err == nil {
		jsonIndicatorPropertiesHistogram, _ := json.Marshal(dst.IndicatorPropertiesHistogram)
		if string(jsonIndicatorPropertiesHistogram) == "{}" { // empty struct
			dst.IndicatorPropertiesHistogram = nil
		} else {
			match++
		}
	} else {
		dst.IndicatorPropertiesHistogram = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IndicatorPropertiesApmAvailability = nil
		dst.IndicatorPropertiesApmLatency = nil
		dst.IndicatorPropertiesCustomKql = nil
		dst.IndicatorPropertiesCustomMetric = nil
		dst.IndicatorPropertiesHistogram = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateSloRequestIndicator)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateSloRequestIndicator)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateSloRequestIndicator) MarshalJSON() ([]byte, error) {
	if src.IndicatorPropertiesApmAvailability != nil {
		return json.Marshal(&src.IndicatorPropertiesApmAvailability)
	}

	if src.IndicatorPropertiesApmLatency != nil {
		return json.Marshal(&src.IndicatorPropertiesApmLatency)
	}

	if src.IndicatorPropertiesCustomKql != nil {
		return json.Marshal(&src.IndicatorPropertiesCustomKql)
	}

	if src.IndicatorPropertiesCustomMetric != nil {
		return json.Marshal(&src.IndicatorPropertiesCustomMetric)
	}

	if src.IndicatorPropertiesHistogram != nil {
		return json.Marshal(&src.IndicatorPropertiesHistogram)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateSloRequestIndicator) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IndicatorPropertiesApmAvailability != nil {
		return obj.IndicatorPropertiesApmAvailability
	}

	if obj.IndicatorPropertiesApmLatency != nil {
		return obj.IndicatorPropertiesApmLatency
	}

	if obj.IndicatorPropertiesCustomKql != nil {
		return obj.IndicatorPropertiesCustomKql
	}

	if obj.IndicatorPropertiesCustomMetric != nil {
		return obj.IndicatorPropertiesCustomMetric
	}

	if obj.IndicatorPropertiesHistogram != nil {
		return obj.IndicatorPropertiesHistogram
	}

	// all schemas are nil
	return nil
}

type NullableCreateSloRequestIndicator struct {
	value *CreateSloRequestIndicator
	isSet bool
}

func (v NullableCreateSloRequestIndicator) Get() *CreateSloRequestIndicator {
	return v.value
}

func (v *NullableCreateSloRequestIndicator) Set(val *CreateSloRequestIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSloRequestIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSloRequestIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSloRequestIndicator(val *CreateSloRequestIndicator) *NullableCreateSloRequestIndicator {
	return &NullableCreateSloRequestIndicator{value: val, isSet: true}
}

func (v NullableCreateSloRequestIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSloRequestIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}