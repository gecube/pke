/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: master
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type UpdateClusterRequestV2 struct {
	Nodepools []PkeOnAzureNodePool `json:"nodepools,omitempty"`
}
