package weather

import (
	"encoding/json"
	"fmt"
	"time"
)

// Module contains information about any additionnal modules
type Module struct {
	ID                   string                      `json:"_id"`             // uniq ID of the module (MAC address)
	Type                 ModuleType                  `json:"type"`            // type of module (see ModuleType const values)
	ModuleName           string                      `json:"module_name"`     // user set name of the module
	DataType             []ModuleDataType            `json:"data_type"`       // array of data measured by the device (see ModuleDataType const values)
	LastSetup            time.Time                   ``                       // date of the last installation
	Reachable            bool                        `json:"reachable"`       // true if the station connected to Netatmo cloud within the last 4 hours
	Firmware             int64                       `json:"firmware"`        // version of the software
	LastMessage          time.Time                   ``                       // date of the last measure update
	LastSeen             time.Time                   ``                       // date of the last status update
	RfStatus             RadioQuality                `json:"rf_status"`       // current radio status per module (see RadioQuality const values)
	BatteryVp            int64                       `json:"battery_vp"`      // current battery status per module (legacy, see BatteryPercent)
	BatteryPercent       int64                       `json:"battery_percent"` // percentage of battery remaining (10=low)
	DashboardDataOutdoor *OutdoorModuleDashboardData ``                       // values summary if module type is outdoor and is reachable
	DashboardDataWind    *WindModuleDashboardData    ``                       // values summary if module type is wind and is reachable
	DashboardDataRain    *RainModuleDashboardData    ``                       // values summary if module type is rain and is reachable
	DashboardDataIndoor  *IndoorModuleDashboardData  ``                       // values summary if module type is indoor and is reachable
	DashboardDataRaw     json.RawMessage             ``                       // in case type auto detect has failed, raw dashboard will be kept here (module must still be reachable)
}

// UnmarshalJSON allows to automatically convert data to go types
func (m *Module) UnmarshalJSON(data []byte) (err error) {
	// Add tmp type
	type OriginalUnmarshal Module
	tmp := struct {
		LastSetup        int64           `json:"last_setup"`
		LastMessage      int64           `json:"last_message"`
		LastSeen         int64           `json:"last_seen"`
		DashboardDataRaw json.RawMessage `json:"dashboard_data"`
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(m),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("failed to unmarshal data to the temporary Module struct: %w", err)
		return
	}
	// Convert
	m.LastSetup = time.Unix(tmp.LastSetup, 0)
	m.LastMessage = time.Unix(tmp.LastMessage, 0)
	m.LastSeen = time.Unix(tmp.LastSeen, 0)
	// Handle module type to select the right dashboard
	if m.Reachable {
		switch tmp.Type {
		case ModuleTypeOutdoor:
			m.DashboardDataOutdoor = new(OutdoorModuleDashboardData)
			if err = json.Unmarshal(tmp.DashboardDataRaw, m.DashboardDataOutdoor); err != nil {
				err = fmt.Errorf("failed to parse module name '%s' of type '%s' into Outdoor dashboard: %w",
					m.ModuleName, m.Type, err)
				return
			}
		case ModuleTypeAnemometer:
			m.DashboardDataWind = new(WindModuleDashboardData)
			if err = json.Unmarshal(tmp.DashboardDataRaw, m.DashboardDataWind); err != nil {
				err = fmt.Errorf("failed to parse module name '%s' of type '%s' into Wind dashboard: %w",
					m.ModuleName, m.Type, err)
				return
			}
		case ModuleTypeRainGauge:
			m.DashboardDataRain = new(RainModuleDashboardData)
			if err = json.Unmarshal(tmp.DashboardDataRaw, m.DashboardDataRain); err != nil {
				err = fmt.Errorf("failed to parse module name '%s' of type '%s' into Rain dashboard: %w",
					m.ModuleName, m.Type, err)
				return
			}
		case ModuleTypeIndoor:
			m.DashboardDataIndoor = new(IndoorModuleDashboardData)
			if err = json.Unmarshal(tmp.DashboardDataRaw, m.DashboardDataIndoor); err != nil {
				err = fmt.Errorf("failed to parse module name '%s' of type '%s' into Indoor dashboard: %w",
					m.ModuleName, m.Type, err)
				return
			}
		default:
			// keep the raw value for user to work around if necessary
			m.DashboardDataRaw = tmp.DashboardDataRaw
		}
	}
	return
}
