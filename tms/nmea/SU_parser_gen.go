// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strings"
import "strconv"

const (
	// PrefixSU prefix
	PrefixSU = "SU"
)

// SU represents fix data.
type CoreSU struct {
	ResponseTypeID uint32

	ResponseTypeIDValidity bool

	ResponseStr string

	ResponseStrValidity bool
}

type SU struct {
	BaseSentence
	CoreSU
}

func NewSU(sentence BaseSentence) *SU {
	s := new(SU)
	s.BaseSentence = sentence

	s.ResponseTypeIDValidity = false

	s.ResponseStrValidity = false

	return s
}

func (s *SU) parse() error {
	var err error

	if s.Format != PrefixSU {
		err = fmt.Errorf("%s is not a %s", s.Format, PrefixSU)
		return err
	}

	if len(s.Fields) == 0 {
		return nil
	} else {
		if s.Fields[0] != "" {
			i, err := strconv.ParseUint(s.Fields[0], 10, 32)
			if err != nil {
				return fmt.Errorf("SU decode variation error: %s", s.Fields[0])
			} else {
				s.CoreSU.ResponseTypeID = uint32(i)
				s.CoreSU.ResponseTypeIDValidity = true
			}

		}
	}

	if len(s.Fields) == 1 {
		return nil
	} else {
		if s.Fields[1] != "" {
			s.ResponseStr = s.Fields[1]
			s.ResponseStrValidity = true
		}
	}

	return nil
}

func (s *SU) Encode() (string, error) {
	var Raw string

	if s.Format != PrefixSU {
		err := fmt.Errorf("Sentence format %s is not a SU sentence", s.Format)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format

	if s.ResponseTypeIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreSU.ResponseTypeID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreSU.ResponseTypeID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ResponseStrValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreSU.ResponseStr

		} else {
			Raw = Raw + "," + s.CoreSU.ResponseStr
		}
	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}