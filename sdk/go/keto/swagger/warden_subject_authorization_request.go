/*
 * Package main ORY Keto
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type WardenSubjectAuthorizationRequest struct {

	// Action is the action that is requested on the resource.
	Action string `json:"action,omitempty"`

	// Context is the request's environmental context.
	Context map[string]interface{} `json:"context,omitempty"`

	// Resource is the resource that access is requested to.
	Resource string `json:"resource,omitempty"`

	// Subejct is the subject that is requesting access.
	Subject string `json:"subject,omitempty"`
}
