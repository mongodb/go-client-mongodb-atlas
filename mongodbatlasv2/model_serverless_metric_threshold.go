/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// ServerlessMetricThreshold - Threshold for the metric that, when exceeded, triggers an alert. The metric threshold pertains to event types which reflects changes of measurements and metrics about the serverless database.
type ServerlessMetricThreshold struct {
	DataMetricThreshold *DataMetricThreshold
	RPUMetricThreshold *RPUMetricThreshold
	RawMetricThreshold *RawMetricThreshold
	TimeMetricThreshold *TimeMetricThreshold
}

// DataMetricThresholdAsServerlessMetricThreshold is a convenience function that returns DataMetricThreshold wrapped in ServerlessMetricThreshold
func DataMetricThresholdAsServerlessMetricThreshold(v *DataMetricThreshold) ServerlessMetricThreshold {
	return ServerlessMetricThreshold{
		DataMetricThreshold: v,
	}
}

// RPUMetricThresholdAsServerlessMetricThreshold is a convenience function that returns RPUMetricThreshold wrapped in ServerlessMetricThreshold
func RPUMetricThresholdAsServerlessMetricThreshold(v *RPUMetricThreshold) ServerlessMetricThreshold {
	return ServerlessMetricThreshold{
		RPUMetricThreshold: v,
	}
}

// RawMetricThresholdAsServerlessMetricThreshold is a convenience function that returns RawMetricThreshold wrapped in ServerlessMetricThreshold
func RawMetricThresholdAsServerlessMetricThreshold(v *RawMetricThreshold) ServerlessMetricThreshold {
	return ServerlessMetricThreshold{
		RawMetricThreshold: v,
	}
}

// TimeMetricThresholdAsServerlessMetricThreshold is a convenience function that returns TimeMetricThreshold wrapped in ServerlessMetricThreshold
func TimeMetricThresholdAsServerlessMetricThreshold(v *TimeMetricThreshold) ServerlessMetricThreshold {
	return ServerlessMetricThreshold{
		TimeMetricThreshold: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServerlessMetricThreshold) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'DataMetricThreshold'
	if jsonDict["metricName"] == "DataMetricThreshold" {
		// try to unmarshal JSON data into DataMetricThreshold
		err = json.Unmarshal(data, &dst.DataMetricThreshold)
		if err == nil {
			return nil // data stored in dst.DataMetricThreshold, return on the first match
		} else {
			dst.DataMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as DataMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RPUMetricThreshold'
	if jsonDict["metricName"] == "RPUMetricThreshold" {
		// try to unmarshal JSON data into RPUMetricThreshold
		err = json.Unmarshal(data, &dst.RPUMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RPUMetricThreshold, return on the first match
		} else {
			dst.RPUMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RPUMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RawMetricThreshold'
	if jsonDict["metricName"] == "RawMetricThreshold" {
		// try to unmarshal JSON data into RawMetricThreshold
		err = json.Unmarshal(data, &dst.RawMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RawMetricThreshold, return on the first match
		} else {
			dst.RawMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RawMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_AVG_COMMAND_EXECUTION_TIME'
	if jsonDict["metricName"] == "SERVERLESS_AVG_COMMAND_EXECUTION_TIME" {
		// try to unmarshal JSON data into TimeMetricThreshold
		err = json.Unmarshal(data, &dst.TimeMetricThreshold)
		if err == nil {
			return nil // data stored in dst.TimeMetricThreshold, return on the first match
		} else {
			dst.TimeMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as TimeMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_AVG_READ_EXECUTION_TIME'
	if jsonDict["metricName"] == "SERVERLESS_AVG_READ_EXECUTION_TIME" {
		// try to unmarshal JSON data into TimeMetricThreshold
		err = json.Unmarshal(data, &dst.TimeMetricThreshold)
		if err == nil {
			return nil // data stored in dst.TimeMetricThreshold, return on the first match
		} else {
			dst.TimeMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as TimeMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_AVG_WRITE_EXECUTION_TIME'
	if jsonDict["metricName"] == "SERVERLESS_AVG_WRITE_EXECUTION_TIME" {
		// try to unmarshal JSON data into TimeMetricThreshold
		err = json.Unmarshal(data, &dst.TimeMetricThreshold)
		if err == nil {
			return nil // data stored in dst.TimeMetricThreshold, return on the first match
		} else {
			dst.TimeMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as TimeMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_CONNECTIONS'
	if jsonDict["metricName"] == "SERVERLESS_CONNECTIONS" {
		// try to unmarshal JSON data into RawMetricThreshold
		err = json.Unmarshal(data, &dst.RawMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RawMetricThreshold, return on the first match
		} else {
			dst.RawMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RawMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_CONNECTIONS_PERCENT'
	if jsonDict["metricName"] == "SERVERLESS_CONNECTIONS_PERCENT" {
		// try to unmarshal JSON data into RawMetricThreshold
		err = json.Unmarshal(data, &dst.RawMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RawMetricThreshold, return on the first match
		} else {
			dst.RawMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RawMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_DATA_SIZE_TOTAL'
	if jsonDict["metricName"] == "SERVERLESS_DATA_SIZE_TOTAL" {
		// try to unmarshal JSON data into DataMetricThreshold
		err = json.Unmarshal(data, &dst.DataMetricThreshold)
		if err == nil {
			return nil // data stored in dst.DataMetricThreshold, return on the first match
		} else {
			dst.DataMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as DataMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_NETWORK_BYTES_IN'
	if jsonDict["metricName"] == "SERVERLESS_NETWORK_BYTES_IN" {
		// try to unmarshal JSON data into DataMetricThreshold
		err = json.Unmarshal(data, &dst.DataMetricThreshold)
		if err == nil {
			return nil // data stored in dst.DataMetricThreshold, return on the first match
		} else {
			dst.DataMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as DataMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_NETWORK_BYTES_OUT'
	if jsonDict["metricName"] == "SERVERLESS_NETWORK_BYTES_OUT" {
		// try to unmarshal JSON data into DataMetricThreshold
		err = json.Unmarshal(data, &dst.DataMetricThreshold)
		if err == nil {
			return nil // data stored in dst.DataMetricThreshold, return on the first match
		} else {
			dst.DataMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as DataMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_NETWORK_NUM_REQUESTS'
	if jsonDict["metricName"] == "SERVERLESS_NETWORK_NUM_REQUESTS" {
		// try to unmarshal JSON data into RawMetricThreshold
		err = json.Unmarshal(data, &dst.RawMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RawMetricThreshold, return on the first match
		} else {
			dst.RawMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RawMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_OPCOUNTER_CMD'
	if jsonDict["metricName"] == "SERVERLESS_OPCOUNTER_CMD" {
		// try to unmarshal JSON data into RawMetricThreshold
		err = json.Unmarshal(data, &dst.RawMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RawMetricThreshold, return on the first match
		} else {
			dst.RawMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RawMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_OPCOUNTER_DELETE'
	if jsonDict["metricName"] == "SERVERLESS_OPCOUNTER_DELETE" {
		// try to unmarshal JSON data into RawMetricThreshold
		err = json.Unmarshal(data, &dst.RawMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RawMetricThreshold, return on the first match
		} else {
			dst.RawMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RawMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_OPCOUNTER_GETMORE'
	if jsonDict["metricName"] == "SERVERLESS_OPCOUNTER_GETMORE" {
		// try to unmarshal JSON data into RawMetricThreshold
		err = json.Unmarshal(data, &dst.RawMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RawMetricThreshold, return on the first match
		} else {
			dst.RawMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RawMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_OPCOUNTER_INSERT'
	if jsonDict["metricName"] == "SERVERLESS_OPCOUNTER_INSERT" {
		// try to unmarshal JSON data into RawMetricThreshold
		err = json.Unmarshal(data, &dst.RawMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RawMetricThreshold, return on the first match
		} else {
			dst.RawMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RawMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_OPCOUNTER_QUERY'
	if jsonDict["metricName"] == "SERVERLESS_OPCOUNTER_QUERY" {
		// try to unmarshal JSON data into RawMetricThreshold
		err = json.Unmarshal(data, &dst.RawMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RawMetricThreshold, return on the first match
		} else {
			dst.RawMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RawMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_OPCOUNTER_UPDATE'
	if jsonDict["metricName"] == "SERVERLESS_OPCOUNTER_UPDATE" {
		// try to unmarshal JSON data into RawMetricThreshold
		err = json.Unmarshal(data, &dst.RawMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RawMetricThreshold, return on the first match
		} else {
			dst.RawMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RawMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_TOTAL_READ_UNITS'
	if jsonDict["metricName"] == "SERVERLESS_TOTAL_READ_UNITS" {
		// try to unmarshal JSON data into RPUMetricThreshold
		err = json.Unmarshal(data, &dst.RPUMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RPUMetricThreshold, return on the first match
		} else {
			dst.RPUMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RPUMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SERVERLESS_TOTAL_WRITE_UNITS'
	if jsonDict["metricName"] == "SERVERLESS_TOTAL_WRITE_UNITS" {
		// try to unmarshal JSON data into RPUMetricThreshold
		err = json.Unmarshal(data, &dst.RPUMetricThreshold)
		if err == nil {
			return nil // data stored in dst.RPUMetricThreshold, return on the first match
		} else {
			dst.RPUMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as RPUMetricThreshold: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TimeMetricThreshold'
	if jsonDict["metricName"] == "TimeMetricThreshold" {
		// try to unmarshal JSON data into TimeMetricThreshold
		err = json.Unmarshal(data, &dst.TimeMetricThreshold)
		if err == nil {
			return nil // data stored in dst.TimeMetricThreshold, return on the first match
		} else {
			dst.TimeMetricThreshold = nil
			return fmt.Errorf("failed to unmarshal ServerlessMetricThreshold as TimeMetricThreshold: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServerlessMetricThreshold) MarshalJSON() ([]byte, error) {
	if src.DataMetricThreshold != nil {
		return json.Marshal(&src.DataMetricThreshold)
	}

	if src.RPUMetricThreshold != nil {
		return json.Marshal(&src.RPUMetricThreshold)
	}

	if src.RawMetricThreshold != nil {
		return json.Marshal(&src.RawMetricThreshold)
	}

	if src.TimeMetricThreshold != nil {
		return json.Marshal(&src.TimeMetricThreshold)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServerlessMetricThreshold) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DataMetricThreshold != nil {
		return obj.DataMetricThreshold
	}

	if obj.RPUMetricThreshold != nil {
		return obj.RPUMetricThreshold
	}

	if obj.RawMetricThreshold != nil {
		return obj.RawMetricThreshold
	}

	if obj.TimeMetricThreshold != nil {
		return obj.TimeMetricThreshold
	}

	// all schemas are nil
	return nil
}

type NullableServerlessMetricThreshold struct {
	value *ServerlessMetricThreshold
	isSet bool
}

func (v NullableServerlessMetricThreshold) Get() *ServerlessMetricThreshold {
	return v.value
}

func (v *NullableServerlessMetricThreshold) Set(val *ServerlessMetricThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessMetricThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessMetricThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessMetricThreshold(val *ServerlessMetricThreshold) *NullableServerlessMetricThreshold {
	return &NullableServerlessMetricThreshold{value: val, isSet: true}
}

func (v NullableServerlessMetricThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessMetricThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

