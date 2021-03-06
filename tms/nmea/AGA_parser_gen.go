// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strings"
import "strconv"

const (
	// PrefixAGA prefix
	PrefixAGA = "AGA"
)

// AGA represents fix data.
type CoreAGA struct {
	StationID string

	StationIDValidity bool

	StationType uint32

	StationTypeValidity bool

	Latitude float64

	LatitudeValidity bool

	LatitudeDirection string

	LatitudeDirectionValidity bool

	Longitude float64

	LongitudeValidity bool

	LongitudeDirection string

	LongitudeDirectionValidity bool

	ReportInterval uint32

	ReportIntervalValidity bool

	TxRxMode uint32

	TxRxModeValidity bool

	QuiteTime uint32

	QuiteTimeValidity bool

	Status string

	StatusValidity bool
}

type AGA struct {
	BaseSentence
	CoreAGA
}

func NewAGA(sentence BaseSentence) *AGA {
	s := new(AGA)
	s.BaseSentence = sentence

	s.StationIDValidity = false

	s.StationTypeValidity = false

	s.LatitudeValidity = false

	s.LatitudeDirectionValidity = false

	s.LongitudeValidity = false

	s.LongitudeDirectionValidity = false

	s.ReportIntervalValidity = false

	s.TxRxModeValidity = false

	s.QuiteTimeValidity = false

	s.StatusValidity = false

	return s
}

func (s *AGA) parse() error {
	var err error

	if s.Format != PrefixAGA {
		err = fmt.Errorf("%s is not a %s", s.Format, PrefixAGA)
		return err
	}

	if len(s.Fields) == 0 {
		return nil
	} else {
		if s.Fields[0] != "" {
			s.StationID = s.Fields[0]
			s.StationIDValidity = true
		}
	}

	if len(s.Fields) == 1 {
		return nil
	} else {
		if s.Fields[1] != "" {
			i, err := strconv.ParseUint(s.Fields[1], 10, 32)
			if err != nil {
				return fmt.Errorf("AGA decode variation error: %s", s.Fields[1])
			} else {
				s.CoreAGA.StationType = uint32(i)
				s.CoreAGA.StationTypeValidity = true
			}

		}
	}

	if len(s.Fields) == 2 {
		return nil
	} else {
		if s.Fields[2] != "" {
			i, err := strconv.ParseFloat(s.Fields[2], 64)
			if err != nil {
				return fmt.Errorf("AGA decode variation error: %s", s.Fields[2])
			} else {
				s.CoreAGA.Latitude = float64(i)
				s.CoreAGA.LatitudeValidity = true
			}

		}
	}

	if len(s.Fields) == 3 {
		return nil
	} else {
		if s.Fields[3] != "" {
			s.LatitudeDirection = s.Fields[3]
			s.LatitudeDirectionValidity = true
		}
	}

	if len(s.Fields) == 4 {
		return nil
	} else {
		if s.Fields[4] != "" {
			i, err := strconv.ParseFloat(s.Fields[4], 64)
			if err != nil {
				return fmt.Errorf("AGA decode variation error: %s", s.Fields[4])
			} else {
				s.CoreAGA.Longitude = float64(i)
				s.CoreAGA.LongitudeValidity = true
			}

		}
	}

	if len(s.Fields) == 5 {
		return nil
	} else {
		if s.Fields[5] != "" {
			s.LongitudeDirection = s.Fields[5]
			s.LongitudeDirectionValidity = true
		}
	}

	if len(s.Fields) == 6 {
		return nil
	} else {
		if s.Fields[6] != "" {
			i, err := strconv.ParseUint(s.Fields[6], 10, 32)
			if err != nil {
				return fmt.Errorf("AGA decode variation error: %s", s.Fields[6])
			} else {
				s.CoreAGA.ReportInterval = uint32(i)
				s.CoreAGA.ReportIntervalValidity = true
			}

		}
	}

	if len(s.Fields) == 7 {
		return nil
	} else {
		if s.Fields[7] != "" {
			i, err := strconv.ParseUint(s.Fields[7], 10, 32)
			if err != nil {
				return fmt.Errorf("AGA decode variation error: %s", s.Fields[7])
			} else {
				s.CoreAGA.TxRxMode = uint32(i)
				s.CoreAGA.TxRxModeValidity = true
			}

		}
	}

	if len(s.Fields) == 8 {
		return nil
	} else {
		if s.Fields[8] != "" {
			i, err := strconv.ParseUint(s.Fields[8], 10, 32)
			if err != nil {
				return fmt.Errorf("AGA decode variation error: %s", s.Fields[8])
			} else {
				s.CoreAGA.QuiteTime = uint32(i)
				s.CoreAGA.QuiteTimeValidity = true
			}

		}
	}

	if len(s.Fields) == 9 {
		return nil
	} else {
		if s.Fields[9] != "" {
			s.Status = s.Fields[9]
			s.StatusValidity = true
		}
	}

	return nil
}

func (s *AGA) Encode() (string, error) {
	var Raw string

	if s.Format != PrefixAGA {
		err := fmt.Errorf("Sentence format %s is not a AGA sentence", s.Format)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format

	if s.StationIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreAGA.StationID

		} else {
			Raw = Raw + "," + s.CoreAGA.StationID
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.StationTypeValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAGA.StationType), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAGA.StationType), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.LatitudeValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreAGA.Latitude, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreAGA.Latitude, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.LatitudeDirectionValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreAGA.LatitudeDirection

		} else {
			Raw = Raw + "," + s.CoreAGA.LatitudeDirection
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.LongitudeValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreAGA.Longitude, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreAGA.Longitude, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.LongitudeDirectionValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreAGA.LongitudeDirection

		} else {
			Raw = Raw + "," + s.CoreAGA.LongitudeDirection
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReportIntervalValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAGA.ReportInterval), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAGA.ReportInterval), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.TxRxModeValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAGA.TxRxMode), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAGA.TxRxMode), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.QuiteTimeValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAGA.QuiteTime), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAGA.QuiteTime), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.StatusValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreAGA.Status

		} else {
			Raw = Raw + "," + s.CoreAGA.Status
		}
	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}
