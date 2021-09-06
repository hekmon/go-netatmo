package weather

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// PublicStationData represents the public data of a station
type PublicStationData struct {
	ID       string                          `json:"_id"`
	Place    Place                           `json:"place"`
	Mark     float32                         `json:"mark"`
	Pressure PublicStationDataPressureValues `json:"pressure"` // not in this form on the orignal payload
	Outdoor  []PublicStationDataOutdoor      `json:"outdoor"`  // not in this form on the orignal payload
	Wind     *WindMeasures                   `json:"wind"`     // not in this form on the orignal payload
	Rain     *RainMeasures                   `json:"rain"`     // not in this form on the orignal payload
}

// PublicStationDataPressureValues is a collection of PublicStationDataPressure
type PublicStationDataPressureValues []PublicStationDataPressure

// Len returns the number of values (implements the https://pkg.go.dev/sort#Interface)
func (psdpv PublicStationDataPressureValues) Len() int {
	return len(psdpv)
}

// Less returns true if the timestamp of i is before j (implements the https://pkg.go.dev/sort#Interface)
func (psdpv PublicStationDataPressureValues) Less(i, j int) bool {
	return psdpv[i].Time.Before(psdpv[j].Time)
}

// Swap changes the values places between them within the array (implements the https://pkg.go.dev/sort#Interface)
func (psdpv PublicStationDataPressureValues) Swap(i, j int) {
	psdpv[i], psdpv[j] = psdpv[j], psdpv[i]
}

// PublicStationDataPressure represents a given presure at a given time
type PublicStationDataPressure struct {
	Time     time.Time
	Pressure float64
}

type PublicStationDataOutdoorValues []PublicStationDataOutdoor

// Len returns the number of values (implements the https://pkg.go.dev/sort#Interface)
func (psdov PublicStationDataOutdoorValues) Len() int {
	return len(psdov)
}

// Less returns true if the timestamp of i is before j (implements the https://pkg.go.dev/sort#Interface)
func (psdov PublicStationDataOutdoorValues) Less(i, j int) bool {
	return psdov[i].Time.Before(psdov[j].Time)
}

// Swap changes the values places between them within the array (implements the https://pkg.go.dev/sort#Interface)
func (psdov PublicStationDataOutdoorValues) Swap(i, j int) {
	psdov[i], psdov[j] = psdov[j], psdov[i]
}

// PublicStationDataOutdoor represents a given outdoor temperature and humidity measures at a given time
type PublicStationDataOutdoor struct {
	Time        time.Time
	Temperature float64
	Humidity    int
}

// UnmarshalJSON allows to create a proper payloade on the fly during JSON unmarshaling
func (pdb *PublicStationData) UnmarshalJSON(data []byte) (err error) {
	/*
		original payload has a shitty JSON schema, its going to be hard to make it pretty
	*/
	// Add tmp type
	type OriginalUnmarshal PublicStationData
	tmp := struct {
		Measures    map[string]json.RawMessage `json:"measures"`     // key is MAC addr (IDs), value is dynamic payload given module type
		Modules     []string                   `json:"modules"`      // list of modules MAC addr (IDs)
		ModuleTypes map[string]ModuleType      `json:"module_types"` // key is MAC addr (IDs) value is module type
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(pdb),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("failed to unmarshal data to the temporary PublicStationData struct: %w", err)
		return
	}
	// Now detect modules types
	var (
		mtype ModuleType
		found bool
	)
	for moduleID, data := range tmp.Measures {
		// Which type is module ?
		if moduleID == tmp.ID {
			mtype = ModuleTypeStation
		} else {
			if mtype, found = tmp.ModuleTypes[moduleID]; !found {
				return fmt.Errorf("module ID %s is not listed on ModuleTypes", moduleID)
			}
		}
		// Prepare unmarshall in dedicated struct given type
		switch mtype {
		case ModuleTypeStation:
			if pdb.Pressure, err = unmarshalPublicStationData(data); err != nil {
				err = fmt.Errorf("failed to parse module ID %s payload as station data: %v", moduleID, err)
				return
			}
		case ModuleTypeOutdoor:
			if pdb.Outdoor, err = unmarshalPublicOutdoorData(data); err != nil {
				err = fmt.Errorf("failed to parse module ID %s payload as outdoor data: %v", moduleID, err)
				return
			}
		case ModuleTypeAnemometer:
			pdb.Wind = new(WindMeasures)
			if err = json.Unmarshal(data, pdb.Wind); err != nil {
				err = fmt.Errorf("failed to parse module ID %s payload as wind data: %w. Payload: %s", moduleID, err, string(data))
				return
			}
		case ModuleTypeRainGauge:
			pdb.Rain = new(RainMeasures)
			if err = json.Unmarshal(data, pdb.Rain); err != nil {
				err = fmt.Errorf("failed to parse module ID %s payload as rain data: %w. Payload: %s", moduleID, err, string(data))
				return
			}
		default:
			return fmt.Errorf("module ID %s has an unknown type: %s", moduleID, mtype)
		}
	}
	return
}

type publicStationDataMeasures struct {
	Measures map[string][]float64 `json:"res"`
	Types    []string             `json:"type"`
}

func unmarshalPublicStationData(data json.RawMessage) (pressure PublicStationDataPressureValues, err error) {
	// Unmarshall station data
	var mainDataTmp publicStationDataMeasures
	if err = json.Unmarshal(data, &mainDataTmp); err != nil {
		err = fmt.Errorf("failed to unmarshall as JSON: %w. Payload: %s", err, string(data))
		return
	}
	// verify
	if len(mainDataTmp.Types) != 1 {
		err = fmt.Errorf("station data should only have one type, payload has %d: %v", len(mainDataTmp.Types), mainDataTmp.Types)
		return
	}
	if mainDataTmp.Types[0] != strings.ToLower(string(ModuleDataTypePressure)) {
		err = fmt.Errorf("station data measure type is unknown: %v", mainDataTmp.Types[0])
		return
	}
	// Extract data
	var timestamp int64
	pressure = make(PublicStationDataPressureValues, len(mainDataTmp.Measures))
	index := 0
	for timestampStr, values := range mainDataTmp.Measures {
		if len(values) != 1 {
			err = fmt.Errorf("multiple station data pressure values encountered for timestamp %s: %v", timestampStr, values)
			return
		}
		if timestamp, err = strconv.ParseInt(timestampStr, 10, 64); err != nil {
			err = fmt.Errorf("can not convert '%s' timestamp as int64: %w", timestampStr, err)
			return
		}
		pressure[index] = PublicStationDataPressure{
			Time:     time.Unix(timestamp, 0),
			Pressure: values[0],
		}
		index++
	}
	// sort values by timestamps
	sort.Sort(pressure)
	return
}

func unmarshalPublicOutdoorData(data json.RawMessage) (outdoorData PublicStationDataOutdoorValues, err error) {
	// Unmarshall station data
	var mainDataTmp publicStationDataMeasures
	if err = json.Unmarshal(data, &mainDataTmp); err != nil {
		err = fmt.Errorf("failed to unmarshall as JSON: %w. Payload: %s", err, string(data))
		return
	}
	// verify
	if len(mainDataTmp.Types) != 2 {
		err = fmt.Errorf("outdoor data should only have two types, payload has %d: %v", len(mainDataTmp.Types), mainDataTmp.Types)
		return
	}
	tempIndex := -1
	humidityIndex := -1
	for index, value := range mainDataTmp.Types {
		switch value {
		case strings.ToLower(string(ModuleDataTypeTemperature)):
			tempIndex = index
		case strings.ToLower(string(ModuleDataTypeHumidity)):
			humidityIndex = index
		default:
			err = fmt.Errorf("outdoor data types at index %d is unknown: %s", index, value)
			return
		}
	}
	if tempIndex == -1 {
		err = fmt.Errorf("can not find index for outdoor '%s' data type: %v", strings.ToLower(string(ModuleDataTypeTemperature)), mainDataTmp.Types)
		return
	}
	if humidityIndex == -1 {
		err = fmt.Errorf("can not find index for outdoor '%s' data type: %v", strings.ToLower(string(ModuleDataTypeHumidity)), mainDataTmp.Types)
		return
	}
	// Extract data
	var timestamp int64
	outdoorData = make(PublicStationDataOutdoorValues, len(mainDataTmp.Measures))
	index := 0
	for timestampStr, values := range mainDataTmp.Measures {
		if len(values) != len(mainDataTmp.Types) {
			err = fmt.Errorf("unexpected number of values (%d, expecting %d) for timestamp %s: %v", len(values), len(mainDataTmp.Types), timestampStr, values)
			return
		}
		if timestamp, err = strconv.ParseInt(timestampStr, 10, 64); err != nil {
			err = fmt.Errorf("can not convert '%s' timestamp as integer: %w", timestampStr, err)
			return
		}
		outdoorData[index] = PublicStationDataOutdoor{
			Time:        time.Unix(timestamp, 0),
			Temperature: values[tempIndex],
			Humidity:    int(values[humidityIndex]),
		}
		index++
	}
	// sort values by timestamps
	sort.Sort(outdoorData)
	return
}
