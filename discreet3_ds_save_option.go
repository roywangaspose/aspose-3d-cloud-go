/*
 * Aspose.3D Cloud API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aspose3dcloud

type Discreet3DsSaveOption struct {
	// Gets or sets  of the SaveFormat.
	SaveFormat SaveFormat `json:"SaveFormat,omitempty"`
	// Some files like OBJ depends on external file, the lookup paths will allows Aspose.3D to look for external file to load
	LookupPaths []string `json:"LookupPaths,omitempty"`
	// The file name of the exporting/importing scene. This is optional, but useful when serialize external assets like OBJ's material.
	FileName string `json:"FileName,omitempty"`
	// The file format like FBX,U3D,PDF ....
	FileFormat string `json:"FileFormat,omitempty"`
	// Gets or sets whether export all lights in the scene.
	ExportLight bool `json:"ExportLight,omitempty"`
	// Gets or sets whether export all cameras in the scene
	ExportCamera bool `json:"ExportCamera,omitempty"`
	// The separator between object's name and the duplicated counter, default value is \"_\". When scene contains objects that use the same name, Aspose.3D 3DS exporter will generate a different name for the object. For example there's two nodes named \"Box\", the first node will have a name \"Box\", and the second node will get a new name \"Box_2\" using the default configuration
	DuplicatedNameSeparator string `json:"DuplicatedNameSeparator,omitempty"`
	// The counter used by generating new name for duplicated names
	DuplicatedNameCounterBase int32 `json:"DuplicatedNameCounterBase,omitempty"`
	// The format of the duplicated counter, default value is empty string.
	DuplicatedNameCounterFormat string `json:"DuplicatedNameCounterFormat,omitempty"`
	// Gets or sets the master scale used in exporting.
	MasterScale float64 `json:"MasterScale,omitempty"`
	// Gets or sets the GammaCorrectedColor.
	GammaCorrectedColor bool `json:"GammaCorrectedColor,omitempty"`
	// Gets or sets flip coordinate system of control points/normal during importing/exporting..
	FlipCoordinateSystem bool `json:"FlipCoordinateSystem,omitempty"`
	// Gets or sets the HighPreciseColor.
	HighPreciseColor bool `json:"HighPreciseColor,omitempty"`
}
