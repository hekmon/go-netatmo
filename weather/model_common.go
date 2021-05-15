package weather

// Place struct for Place
type Place struct {
	Timezone string        `json:"timezone"` // Timezone
	Country  string        `json:"country"`  // Country
	Altitude float32       `json:"altitude"` // Altitude
	Location []interface{} `json:"location"`
}

/*
	TODO organize
*/

// BodyMeasureGet struct for BodyMeasureGet
type BodyMeasureGet struct {
	Body []MeasureGet `json:"body,omitempty"`
}

// MeasureGet struct for MeasureGet
type MeasureGet struct {
	BegTime  float32   `json:"beg_time,omitempty"`
	StepTime float32   `json:"step_time,omitempty"`
	Value    []float32 `json:"value,omitempty"`
}

// InvalidServerResponse struct for InvalidServerResponse
type InvalidServerResponse struct {
	Error InvalidServerResponseError `json:"error"`
}

// InvalidServerResponseError struct for InvalidServerResponseError
type InvalidServerResponseError struct {
	Code    float32 `json:"code"`
	Message string  `json:"message"`
}

// Measure struct for Measure
type Measure struct {
	MacAddressNAMain    *MeasureMacAddressNAMain    `json:"mac_address_NAMain,omitempty"`
	MacAddressNAModule1 *MeasureMacAddressNAModule1 `json:"mac_address_NAModule1,omitempty"`
	MacAddressNAModule2 *MeasureMacAddressNAModule2 `json:"mac_address_NAModule2,omitempty"`
	MacAddressNAModule3 *MeasureMacAddressNAModule3 `json:"mac_address_NAModule3,omitempty"`
}

// MeasureMacAddressNAMain struct for MeasureMacAddressNAMain
type MeasureMacAddressNAMain struct {
	Res  *MeasureMacAddressNAMainRes `json:"res,omitempty"`
	Type *string                     `json:"type,omitempty"`
}

// MeasureMacAddressNAMainRes struct for MeasureMacAddressNAMainRes
type MeasureMacAddressNAMainRes struct {
	TimeStamp *[]float32 `json:"time_stamp,omitempty"`
}

// MeasureMacAddressNAModule1 struct for MeasureMacAddressNAModule1
type MeasureMacAddressNAModule1 struct {
	Res  *MeasureMacAddressNAModule1Res `json:"res,omitempty"`
	Type *[]string                      `json:"type,omitempty"`
}

// MeasureMacAddressNAModule1Res struct for MeasureMacAddressNAModule1Res
type MeasureMacAddressNAModule1Res struct {
	TimeStamp *[]float32 `json:"time_stamp,omitempty"`
}

// PublicData struct for PublicData
type PublicData struct {
	Body *[]PublicDataBody `json:"body,omitempty"`
}

// ServerResponse struct for ServerResponse
type ServerResponse struct {
	Status     *string `json:"status,omitempty"`
	TimeExec   *string `json:"time_exec,omitempty"`
	TimeServer *string `json:"time_server,omitempty"`
}

// StationData struct for StationData
type StationData struct {
	Body *StationDataBody `json:"body,omitempty"`
}

// StationDataBody struct for StationDataBody
type StationDataBody struct {
	Devices *[]StationDataBodyDevices `json:"devices,omitempty"`
	User    *UserWeather              `json:"user,omitempty"`
}
