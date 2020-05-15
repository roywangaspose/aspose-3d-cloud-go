/*
 * Aspose.3D Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose3dcloud

// The Sphere Entity
type Sphere struct {
	// Gets or sets the Name of Sphere.
	Name string `json:"Name,omitempty"`
	// Gets or sets the width segments.
	WidthSegments int32 `json:"WidthSegments"`
	// Gets or sets the height segments.             
	HeightSegments int32 `json:"HeightSegments"`
	// Gets or sets the phi start.             
	PhiStart float64 `json:"PhiStart"`
	// Gets or sets the length of the phi.
	PhiLength float64 `json:"PhiLength"`
	// Gets or sets the theta start.             
	ThetaStart float64 `json:"ThetaStart"`
	// Gets or sets the length of the theta.
	ThetaLength float64 `json:"ThetaLength"`
	// Gets or sets the radius 
	Radius float64 `json:"Radius"`
}