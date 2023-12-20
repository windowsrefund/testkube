/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

// TestTriggerConcurrencyPolicies : supported concurrency policies for test triggers
type TestTriggerConcurrencyPolicies string

// List of TestTriggerConcurrencyPolicies
const (
	ALLOW_TestTriggerConcurrencyPolicies   TestTriggerConcurrencyPolicies = "allow"
	FORBID_TestTriggerConcurrencyPolicies  TestTriggerConcurrencyPolicies = "forbid"
	REPLACE_TestTriggerConcurrencyPolicies TestTriggerConcurrencyPolicies = "replace"
)