/*
Task Execution Service

## Executive Summary The Task Execution Service (TES) API is a standardized schema and API for describing and executing batch execution tasks. A task defines a set of input files, a set of containers and commands to run, a set of output files and some other logging and metadata.  TES servers accept task documents and execute them asynchronously on available compute resources. A TES server could be built on top of a traditional HPC queuing system, such as Grid Engine, Slurm or cloud style compute systems such as AWS Batch or Kubernetes. ## Introduction This document describes the TES API and provides details on the specific endpoints, request formats, and responses. It is intended to provide key information for developers of TES-compatible services as well as clients that will call these TES services. Use cases include:    - Deploying existing workflow engines on new infrastructure. Workflow engines   such as CWL-Tes and Cromwell have extentions for using TES. This will allow   a system engineer to deploy them onto a new infrastructure using a job scheduling   system not previously supported by the engine.    - Developing a custom workflow management system. This API provides a common   interface to asynchronous batch processing capabilities. A developer can write   new tools against this interface and expect them to work using a variety of   backend solutions that all support the same specification.   ## Standards The TES API specification is written in OpenAPI and embodies a RESTful service philosophy. It uses JSON in requests and responses and standard HTTP/HTTPS for information transport. HTTPS should be used rather than plain HTTP except for testing or internal-only purposes. ### Authentication and Authorization Is is envisaged that most TES API instances will require users to authenticate to use the endpoints. However, the decision if authentication is required should be taken by TES API implementers.  If authentication is required, we recommend that TES implementations use an OAuth2  bearer token, although they can choose other mechanisms if appropriate.  Checking that a user is authorized to submit TES requests is a responsibility of TES implementations. ### CORS If TES API implementation is to be used by another website or domain it must implement Cross Origin Resource Sharing (CORS). Please refer to https://w3id.org/ga4gh/product-approval-support/cors for more information about GA4GH’s recommendations and how to implement CORS. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tes

import (
	"encoding/json"
	"fmt"
)

// TesFileType Define if input/output element is a file or a directory. It is not required that the user provide this value, but it is required that the server fill in the value once the information is avalible at run time.
type TesFileType string

// List of tesFileType
const (
	FILE TesFileType = "FILE"
	DIRECTORY TesFileType = "DIRECTORY"
)

// All allowed values of TesFileType enum
var AllowedTesFileTypeEnumValues = []TesFileType{
	"FILE",
	"DIRECTORY",
}

func (v *TesFileType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TesFileType(value)
	for _, existing := range AllowedTesFileTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TesFileType", value)
}

// NewTesFileTypeFromValue returns a pointer to a valid TesFileType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTesFileTypeFromValue(v string) (*TesFileType, error) {
	ev := TesFileType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TesFileType: valid values are %v", v, AllowedTesFileTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TesFileType) IsValid() bool {
	for _, existing := range AllowedTesFileTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tesFileType value
func (v TesFileType) Ptr() *TesFileType {
	return &v
}

type NullableTesFileType struct {
	value *TesFileType
	isSet bool
}

func (v NullableTesFileType) Get() *TesFileType {
	return v.value
}

func (v *NullableTesFileType) Set(val *TesFileType) {
	v.value = val
	v.isSet = true
}

func (v NullableTesFileType) IsSet() bool {
	return v.isSet
}

func (v *NullableTesFileType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTesFileType(val *TesFileType) *NullableTesFileType {
	return &NullableTesFileType{value: val, isSet: true}
}

func (v NullableTesFileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTesFileType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

