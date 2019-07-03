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

type DeploymentClusterGroupDeployment struct {
	Atomic bool `json:"atomic,omitempty"`
	Dryrun bool `json:"dryrun,omitempty"`
	Name string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	Package []int32 `json:"package,omitempty"`
	ReleaseName string `json:"releaseName,omitempty"`
	ReuseValues bool `json:"reuseValues,omitempty"`
	RollingMode bool `json:"rollingMode,omitempty"`
	ValueOverrides map[string]interface{} `json:"valueOverrides,omitempty"`
	Values map[string]interface{} `json:"values,omitempty"`
	Version string `json:"version,omitempty"`
}
