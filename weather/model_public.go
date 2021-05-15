package weather

// PublicDataBody struct for PublicDataBody
type PublicDataBody struct {
	ID          string                   `json:"_id"`
	Place       Place                    `json:"place"`
	Mark        float32                  `json:"mark"`
	Measures    Measure                  `json:"measures"`
	Modules     []string                 `json:"modules"`
	ModuleTypes []map[string]interface{} `json:"module_types"`
}
