package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"time"
)


var inputTimeLayout = []string{
	TimeLayout,
	TimeLayoutHourColon,
	TimeLayoutSecond,
	TimeLayoutMinute,
	TimeLayoutHour,
}

type Time struct {
	string    `json:"string,omitempty"`
	time.Time `json:"time,omitempty"`
	Error     error
}


// UnmarshalJSON - Convert JSON to value
func (dt *Time) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}

		// Store result from string
		err = json.Unmarshal(data, &dt.string)
		if err == nil {
			dt.SetString(dt.string)
			break
		}

		// Store result from time
		err = json.Unmarshal(data, &dt.Time)
		if err == nil {
			dt.SetValue(dt.Time)
			break
		}
	}

	return err
}

// MarshalJSON - Convert value to JSON
func (dt Time) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		// data, err = json.Marshal(dt.string)
		// if err != nil {
		// 	break
		// }
		data = []byte("\"" + dt.Time.Format(TimeLayout) + "\"")
	}

	return data, err
}

func (dt Time) Value() time.Time {
	return dt.Time
}

func (dt Time) String() string {
	return dt.string
}

func (dt *Time) SetString(value string) *Time {
	for range Only.Once {
		dt.string = value
		dt.Time = time.Time{}

		if value == "" {
			break
		}

		if value == "--" {
			value = ""
			break
		}

		for _, f := range inputTimeLayout {
			dt.Time, dt.Error = time.Parse(f, value)
			if dt.Error == nil {
				dt.string = dt.Time.Format(TimeLayout)
				break
			}
		}
	}

	return dt
}

func (dt *Time) SetValue(value time.Time) *Time {
	for range Only.Once {
		dt.string = ""
		dt.Time = value

		if value.IsZero() {
			break
		}

		dt.string = value.Format(TimeLayout)
	}

	return dt
}

func SetTimeString(value string) *Time {
	var t Time
	return t.SetString(value)
}

func SetTimeValue(value time.Time) *Time {
	var t Time
	return t.SetValue(value)
}
