package netatmo

/*
	Scopes
	https://dev.netatmo.com/apidocumentation/oauth#scopes
*/

const (
	// ScopeStationRead - to retrieve weather station data (Getstationsdata, Getmeasure)
	ScopeStationRead = "read_station"
	// ScopeThermostatRead - to retrieve thermostat data (Homestatus, Getroommeasure...)
	ScopeThermostatRead = "read_thermostat"
	// ScopeThermostatWrite - to set up the thermostat (Synchomeschedule, Setroomthermpoint...)
	ScopeThermostatWrite = "write_thermostat"
	// ScopeCameraRead - to retrieve Smart Indoor Cameradata (Gethomedata, Getcamerapicture...)
	ScopeCameraRead = "read_camera"
	// ScopeCameraWrite - to inform the Smart Indoor Camera that a specific person or everybody has left the Home (Setpersonsaway, Setpersonshome)
	ScopeCameraWrite = "write_camera"
	// ScopeCameraAccess - to access the camera, the videos and the live stream
	ScopeCameraAccess = "access_camera"
	// ScopePresenceRead - to retrieve Smart Outdoor Camera data (Gethomedata, Getcamerapicture...)
	ScopePresenceRead = "read_presence"
	// ScopePresenceAccess - to access the camera, the videos and the live stream
	ScopePresenceAccess = "access_presence"
	// ScopeSmokeDetectorRead - to retrieve the Smart Smoke Alarm informations and events (Gethomedata, Geteventsuntil...)
	ScopeSmokeDetectorRead = "read_smokedetector"
	// ScopeHomeCoachRead - to read data coming from Smart Indoor Air Quality Monitor (gethomecoachsdata)
	ScopeHomeCoachRead = "read_homecoach"
)
