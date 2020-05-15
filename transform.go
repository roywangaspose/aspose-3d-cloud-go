/*
 * Aspose.3D Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose3dcloud

// A transform contains information that allow access to object's translate/scale/rotation or transform matrix at minimum cost This is used by local transform.             
type Transform struct {
	// Gets or sets the translation
	Translation *Vector3 `json:"Translation,omitempty"`
	// Gets or sets the scale
	Scale *Vector3 `json:"Scale,omitempty"`
	// Gets or sets the pre-rotation represented in degree
	PreRotation *Vector3 `json:"PreRotation,omitempty"`
	// Gets or sets the post-rotation represented in degree
	PostRotation *Vector3 `json:"PostRotation,omitempty"`
	// Gets or sets the rotation represented in euler angles, measured in degree             
	EulerAngles *Vector3 `json:"EulerAngles,omitempty"`
}