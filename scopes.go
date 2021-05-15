package netatmo

// Scope is used to hold all netatmo API scopes.
// https://dev.netatmo.com/apidocumentation/oauth#scopes
type Scope string

const (
	// ScopeStationRead: to retrieve weather station data (Getstationsdata, Getmeasure)
	ScopeStationRead Scope = "read_station"
	// ScopeThermostatRead: to retrieve thermostat data (Homestatus, Getroommeasure...)
	ScopeThermostatRead Scope = "read_thermostat"
	// ScopeThermostatWrite: to set up the thermostat (Synchomeschedule, Setroomthermpoint...)
	ScopeThermostatWrite Scope = "write_thermostat"
	// ScopeCameraRead: to retrieve Smart Indoor Cameradata (Gethomedata, Getcamerapicture...)
	ScopeCameraRead Scope = "read_camera"
	// ScopeCameraWrite: to inform the Smart Indoor Camera that a specific person or everybody has left the Home (Setpersonsaway, Setpersonshome)
	ScopeCameraWrite Scope = "write_camera"
	// ScopeCameraAccess: to access the camera, the videos and the live stream
	ScopeCameraAccess Scope = "access_camera"
	// ScopePresenceRead: to retrieve Smart Outdoor Camera data (Gethomedata, Getcamerapicture...)
	ScopePresenceRead Scope = "read_presence"
	// ScopePresenceAccess: to access the camera, the videos and the live stream
	ScopePresenceAccess Scope = "access_presence"
	// ScopeSmokeDetectorRead: to retrieve the Smart Smoke Alarm informations and events (Gethomedata, Geteventsuntil...)
	ScopeSmokeDetectorRead Scope = "read_smokedetector"
	// read_homecoach: to read data coming from Smart Indoor Air Quality Monitor (gethomecoachsdata)
	ScopeHomeCoachRead Scope = "read_homecoach"
)

type Scopes []Scope

func (sc Scopes) toStrSlice() (output []string) {
	output = make([]string, len(sc))
	for index, scope := range sc {
		output[index] = string(scope)
	}
	return
}
