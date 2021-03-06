// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strings"
import "strconv"

const (
	// PrefixCT prefix
	PrefixCT = "CT"
)

// CT represents fix data.
type CoreCT struct {
	TimeInUtcFormat uint32

	TimeInUtcFormatValidity bool

	TrackID uint32

	TrackIDValidity bool

	TrackPositionOffsetWtoEinM int32

	TrackPositionOffsetWtoEinMValidity bool

	TrackPositionOffsetNtoSinM int32

	TrackPositionOffsetNtoSinMValidity bool

	TrackLocationWtoEinDeg float64

	TrackLocationWtoEinDegValidity bool

	TrackLocationNtoSinDeg float64

	TrackLocationNtoSinDegValidity bool
}

type CT struct {
	BaseSentence
	CoreCT
}

func NewCT(sentence BaseSentence) *CT {
	s := new(CT)
	s.BaseSentence = sentence

	s.TimeInUtcFormatValidity = false

	s.TrackIDValidity = false

	s.TrackPositionOffsetWtoEinMValidity = false

	s.TrackPositionOffsetNtoSinMValidity = false

	s.TrackLocationWtoEinDegValidity = false

	s.TrackLocationNtoSinDegValidity = false

	return s
}

func (s *CT) parse() error {
	var err error

	if s.Format != PrefixCT {
		err = fmt.Errorf("%s is not a %s", s.Format, PrefixCT)
		return err
	}

	if len(s.Fields) == 0 {
		return nil
	} else {
		if s.Fields[0] != "" {
			i, err := strconv.ParseUint(s.Fields[0], 10, 32)
			if err != nil {
				return fmt.Errorf("CT decode variation error: %s", s.Fields[0])
			} else {
				s.CoreCT.TimeInUtcFormat = uint32(i)
				s.CoreCT.TimeInUtcFormatValidity = true
			}

		}
	}

	if len(s.Fields) == 1 {
		return nil
	} else {
		if s.Fields[1] != "" {
			i, err := strconv.ParseUint(s.Fields[1], 10, 32)
			if err != nil {
				return fmt.Errorf("CT decode variation error: %s", s.Fields[1])
			} else {
				s.CoreCT.TrackID = uint32(i)
				s.CoreCT.TrackIDValidity = true
			}

		}
	}

	if len(s.Fields) == 2 {
		return nil
	} else {
		if s.Fields[2] != "" {
			i, err := strconv.ParseInt(s.Fields[2], 10, 32)
			if err != nil {
				return fmt.Errorf("CT decode variation error: %s", s.Fields[2])
			} else {
				s.CoreCT.TrackPositionOffsetWtoEinM = int32(i)
				s.CoreCT.TrackPositionOffsetWtoEinMValidity = true
			}

		}
	}

	if len(s.Fields) == 3 {
		return nil
	} else {
		if s.Fields[3] != "" {
			i, err := strconv.ParseInt(s.Fields[3], 10, 32)
			if err != nil {
				return fmt.Errorf("CT decode variation error: %s", s.Fields[3])
			} else {
				s.CoreCT.TrackPositionOffsetNtoSinM = int32(i)
				s.CoreCT.TrackPositionOffsetNtoSinMValidity = true
			}

		}
	}

	if len(s.Fields) == 4 {
		return nil
	} else {
		if s.Fields[4] != "" {
			i, err := strconv.ParseFloat(s.Fields[4], 64)
			if err != nil {
				return fmt.Errorf("CT decode variation error: %s", s.Fields[4])
			} else {
				s.CoreCT.TrackLocationWtoEinDeg = float64(i)
				s.CoreCT.TrackLocationWtoEinDegValidity = true
			}

		}
	}

	if len(s.Fields) == 5 {
		return nil
	} else {
		if s.Fields[5] != "" {
			i, err := strconv.ParseFloat(s.Fields[5], 64)
			if err != nil {
				return fmt.Errorf("CT decode variation error: %s", s.Fields[5])
			} else {
				s.CoreCT.TrackLocationNtoSinDeg = float64(i)
				s.CoreCT.TrackLocationNtoSinDegValidity = true
			}

		}
	}

	return nil
}

func (s *CT) Encode() (string, error) {
	var Raw string

	if s.Format != PrefixCT {
		err := fmt.Errorf("Sentence format %s is not a CT sentence", s.Format)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format

	if s.TimeInUtcFormatValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreCT.TimeInUtcFormat), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreCT.TimeInUtcFormat), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.TrackIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreCT.TrackID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreCT.TrackID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.TrackPositionOffsetWtoEinMValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatInt(int64(s.CoreCT.TrackPositionOffsetWtoEinM), 10)

		} else {
			Raw = Raw + "," + strconv.FormatInt(int64(s.CoreCT.TrackPositionOffsetWtoEinM), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.TrackPositionOffsetNtoSinMValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatInt(int64(s.CoreCT.TrackPositionOffsetNtoSinM), 10)

		} else {
			Raw = Raw + "," + strconv.FormatInt(int64(s.CoreCT.TrackPositionOffsetNtoSinM), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.TrackLocationWtoEinDegValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreCT.TrackLocationWtoEinDeg, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreCT.TrackLocationWtoEinDeg, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.TrackLocationNtoSinDegValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreCT.TrackLocationNtoSinDeg, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreCT.TrackLocationNtoSinDeg, 'f', -1, 64)
		}

	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}
