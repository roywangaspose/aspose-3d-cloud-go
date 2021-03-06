/*
 * Aspose.3D Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose3dcloud

// Entity class             
type Entity struct {
	// Box/Cylinder/Sphere/Torus/Plane
	Type_ string `json:"Type,omitempty"`
	// Gets or sets Box entity
	Box *Box `json:"Box,omitempty"`
	// Gets or sets Cylinder entity
	Cylinder *Cylinder `json:"Cylinder,omitempty"`
	// Gets or sets Sphere entity
	Sphere *Sphere `json:"Sphere,omitempty"`
	// Gets or sets Torus entity
	Torus *Torus `json:"Torus,omitempty"`
	// Gets or sets Plane entity
	Plane *Plane `json:"Plane,omitempty"`
}
