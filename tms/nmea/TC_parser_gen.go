// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strings"
import "strconv"

const (
	// PrefixTC prefix
	PrefixTC = "TC"
)

// TC represents fix data.
type CoreTC struct {
	CurrentProfile uint32

	CurrentProfileValidity bool

	ClassifierState uint32

	ClassifierStateValidity bool

	ClassifierThreshold float64

	ClassifierThresholdValidity bool

	MinDirOfInterestInDegrees int32

	MinDirOfInterestInDegreesValidity bool

	MaxDirOfInterestInDegrees int32

	MaxDirOfInterestInDegreesValidity bool

	MinVelocityInKnots uint32

	MinVelocityInKnotsValidity bool

	MaxVelocityInKnots uint32

	MaxVelocityInKnotsValidity bool
}

type TC struct {
	BaseSentence
	CoreTC
}

func NewTC(sentence BaseSentence) *TC {
	s := new(TC)
	s.BaseSentence = sentence

	s.CurrentProfileValidity = false

	s.ClassifierStateValidity = false

	s.ClassifierThresholdValidity = false

	s.MinDirOfInterestInDegreesValidity = false

	s.MaxDirOfInterestInDegreesValidity = false

	s.MinVelocityInKnotsValidity = false

	s.MaxVelocityInKnotsValidity = false

	return s
}

func (s *TC) parse() error {
	var err error

	if s.Format != PrefixTC {
		err = fmt.Errorf("%s is not a %s", s.Format, PrefixTC)
		return err
	}

	if len(s.Fields) == 0 {
		return nil
	} else {
		if s.Fields[0] != "" {
			i, err := strconv.ParseUint(s.Fields[0], 10, 32)
			if err != nil {
				return fmt.Errorf("TC decode variation error: %s", s.Fields[0])
			} else {
				s.CoreTC.CurrentProfile = uint32(i)
				s.CoreTC.CurrentProfileValidity = true
			}

		}
	}

	if len(s.Fields) == 1 {
		return nil
	} else {
		if s.Fields[1] != "" {
			i, err := strconv.ParseUint(s.Fields[1], 10, 32)
			if err != nil {
				return fmt.Errorf("TC decode variation error: %s", s.Fields[1])
			} else {
				s.CoreTC.ClassifierState = uint32(i)
				s.CoreTC.ClassifierStateValidity = true
			}

		}
	}

	if len(s.Fields) == 2 {
		return nil
	} else {
		if s.Fields[2] != "" {
			i, err := strconv.ParseFloat(s.Fields[2], 64)
			if err != nil {
				return fmt.Errorf("TC decode variation error: %s", s.Fields[2])
			} else {
				s.CoreTC.ClassifierThreshold = float64(i)
				s.CoreTC.ClassifierThresholdValidity = true
			}

		}
	}

	if len(s.Fields) == 3 {
		return nil
	} else {
		if s.Fields[3] != "" {
			i, err := strconv.ParseInt(s.Fields[3], 10, 32)
			if err != nil {
				return fmt.Errorf("TC decode variation error: %s", s.Fields[3])
			} else {
				s.CoreTC.MinDirOfInterestInDegrees = int32(i)
				s.CoreTC.MinDirOfInterestInDegreesValidity = true
			}

		}
	}

	if len(s.Fields) == 4 {
		return nil
	} else {
		if s.Fields[4] != "" {
			i, err := strconv.ParseInt(s.Fields[4], 10, 32)
			if err != nil {
				return fmt.Errorf("TC decode variation error: %s", s.Fields[4])
			} else {
				s.CoreTC.MaxDirOfInterestInDegrees = int32(i)
				s.CoreTC.MaxDirOfInterestInDegreesValidity = true
			}

		}
	}

	if len(s.Fields) == 5 {
		return nil
	} else {
		if s.Fields[5] != "" {
			i, err := strconv.ParseUint(s.Fields[5], 10, 32)
			if err != nil {
				return fmt.Errorf("TC decode variation error: %s", s.Fields[5])
			} else {
				s.CoreTC.MinVelocityInKnots = uint32(i)
				s.CoreTC.MinVelocityInKnotsValidity = true
			}

		}
	}

	if len(s.Fields) == 6 {
		return nil
	} else {
		if s.Fields[6] != "" {
			i, err := strconv.ParseUint(s.Fields[6], 10, 32)
			if err != nil {
				return fmt.Errorf("TC decode variation error: %s", s.Fields[6])
			} else {
				s.CoreTC.MaxVelocityInKnots = uint32(i)
				s.CoreTC.MaxVelocityInKnotsValidity = true
			}

		}
	}

	return nil
}

func (s *TC) Encode() (string, error) {
	var Raw string

	if s.Format != PrefixTC {
		err := fmt.Errorf("Sentence format %s is not a TC sentence", s.Format)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format

	if s.CurrentProfileValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreTC.CurrentProfile), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreTC.CurrentProfile), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ClassifierStateValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreTC.ClassifierState), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreTC.ClassifierState), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ClassifierThresholdValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreTC.ClassifierThreshold, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreTC.ClassifierThreshold, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.MinDirOfInterestInDegreesValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatInt(int64(s.CoreTC.MinDirOfInterestInDegrees), 10)

		} else {
			Raw = Raw + "," + strconv.FormatInt(int64(s.CoreTC.MinDirOfInterestInDegrees), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.MaxDirOfInterestInDegreesValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatInt(int64(s.CoreTC.MaxDirOfInterestInDegrees), 10)

		} else {
			Raw = Raw + "," + strconv.FormatInt(int64(s.CoreTC.MaxDirOfInterestInDegrees), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.MinVelocityInKnotsValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreTC.MinVelocityInKnots), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreTC.MinVelocityInKnots), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.MaxVelocityInKnotsValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreTC.MaxVelocityInKnots), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreTC.MaxVelocityInKnots), 10)
		}

	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}