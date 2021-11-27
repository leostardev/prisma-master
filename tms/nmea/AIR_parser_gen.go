// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strings"
import "strconv"

const (
	// PrefixAIR prefix
	PrefixAIR = "AIR"
)

// AIR represents fix data.
type CoreAIR struct {
	Destination1ID uint32

	Destination1IDValidity bool

	Message11ID uint32

	Message11IDValidity bool

	Message11SubID uint32

	Message11SubIDValidity bool

	Message12ID uint32

	Message12IDValidity bool

	Message12SubID uint32

	Message12SubIDValidity bool

	Destination2ID uint32

	Destination2IDValidity bool

	Message21ID uint32

	Message21IDValidity bool

	Message21SubID uint32

	Message21SubIDValidity bool

	Channel string

	ChannelValidity bool

	SlotOffset11 uint32

	SlotOffset11Validity bool

	SlotOffset12 uint32

	SlotOffset12Validity bool

	SlotOffset21 uint32

	SlotOffset21Validity bool
}

type AIR struct {
	BaseSentence
	CoreAIR
}

func NewAIR(sentence BaseSentence) *AIR {
	s := new(AIR)
	s.BaseSentence = sentence

	s.Destination1IDValidity = false

	s.Message11IDValidity = false

	s.Message11SubIDValidity = false

	s.Message12IDValidity = false

	s.Message12SubIDValidity = false

	s.Destination2IDValidity = false

	s.Message21IDValidity = false

	s.Message21SubIDValidity = false

	s.ChannelValidity = false

	s.SlotOffset11Validity = false

	s.SlotOffset12Validity = false

	s.SlotOffset21Validity = false

	return s
}

func (s *AIR) parse() error {
	var err error

	if s.Format != PrefixAIR {
		err = fmt.Errorf("%s is not a %s", s.Format, PrefixAIR)
		return err
	}

	if len(s.Fields) == 0 {
		return nil
	} else {
		if s.Fields[0] != "" {
			i, err := strconv.ParseUint(s.Fields[0], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[0])
			} else {
				s.CoreAIR.Destination1ID = uint32(i)
				s.CoreAIR.Destination1IDValidity = true
			}

		}
	}

	if len(s.Fields) == 1 {
		return nil
	} else {
		if s.Fields[1] != "" {
			i, err := strconv.ParseUint(s.Fields[1], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[1])
			} else {
				s.CoreAIR.Message11ID = uint32(i)
				s.CoreAIR.Message11IDValidity = true
			}

		}
	}

	if len(s.Fields) == 2 {
		return nil
	} else {
		if s.Fields[2] != "" {
			i, err := strconv.ParseUint(s.Fields[2], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[2])
			} else {
				s.CoreAIR.Message11SubID = uint32(i)
				s.CoreAIR.Message11SubIDValidity = true
			}

		}
	}

	if len(s.Fields) == 3 {
		return nil
	} else {
		if s.Fields[3] != "" {
			i, err := strconv.ParseUint(s.Fields[3], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[3])
			} else {
				s.CoreAIR.Message12ID = uint32(i)
				s.CoreAIR.Message12IDValidity = true
			}

		}
	}

	if len(s.Fields) == 4 {
		return nil
	} else {
		if s.Fields[4] != "" {
			i, err := strconv.ParseUint(s.Fields[4], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[4])
			} else {
				s.CoreAIR.Message12SubID = uint32(i)
				s.CoreAIR.Message12SubIDValidity = true
			}

		}
	}

	if len(s.Fields) == 5 {
		return nil
	} else {
		if s.Fields[5] != "" {
			i, err := strconv.ParseUint(s.Fields[5], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[5])
			} else {
				s.CoreAIR.Destination2ID = uint32(i)
				s.CoreAIR.Destination2IDValidity = true
			}

		}
	}

	if len(s.Fields) == 6 {
		return nil
	} else {
		if s.Fields[6] != "" {
			i, err := strconv.ParseUint(s.Fields[6], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[6])
			} else {
				s.CoreAIR.Message21ID = uint32(i)
				s.CoreAIR.Message21IDValidity = true
			}

		}
	}

	if len(s.Fields) == 7 {
		return nil
	} else {
		if s.Fields[7] != "" {
			i, err := strconv.ParseUint(s.Fields[7], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[7])
			} else {
				s.CoreAIR.Message21SubID = uint32(i)
				s.CoreAIR.Message21SubIDValidity = true
			}

		}
	}

	if len(s.Fields) == 8 {
		return nil
	} else {
		if s.Fields[8] != "" {
			s.Channel = s.Fields[8]
			s.ChannelValidity = true
		}
	}

	if len(s.Fields) == 9 {
		return nil
	} else {
		if s.Fields[9] != "" {
			i, err := strconv.ParseUint(s.Fields[9], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[9])
			} else {
				s.CoreAIR.SlotOffset11 = uint32(i)
				s.CoreAIR.SlotOffset11Validity = true
			}

		}
	}

	if len(s.Fields) == 10 {
		return nil
	} else {
		if s.Fields[10] != "" {
			i, err := strconv.ParseUint(s.Fields[10], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[10])
			} else {
				s.CoreAIR.SlotOffset12 = uint32(i)
				s.CoreAIR.SlotOffset12Validity = true
			}

		}
	}

	if len(s.Fields) == 11 {
		return nil
	} else {
		if s.Fields[11] != "" {
			i, err := strconv.ParseUint(s.Fields[11], 10, 32)
			if err != nil {
				return fmt.Errorf("AIR decode variation error: %s", s.Fields[11])
			} else {
				s.CoreAIR.SlotOffset21 = uint32(i)
				s.CoreAIR.SlotOffset21Validity = true
			}

		}
	}

	return nil
}

func (s *AIR) Encode() (string, error) {
	var Raw string

	if s.Format != PrefixAIR {
		err := fmt.Errorf("Sentence format %s is not a AIR sentence", s.Format)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format

	if s.Destination1IDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Destination1ID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Destination1ID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.Message11IDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message11ID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message11ID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.Message11SubIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message11SubID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message11SubID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.Message12IDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message12ID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message12ID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.Message12SubIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message12SubID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message12SubID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.Destination2IDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Destination2ID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Destination2ID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.Message21IDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message21ID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message21ID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.Message21SubIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message21SubID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.Message21SubID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ChannelValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreAIR.Channel

		} else {
			Raw = Raw + "," + s.CoreAIR.Channel
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.SlotOffset11Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.SlotOffset11), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.SlotOffset11), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.SlotOffset12Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.SlotOffset12), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.SlotOffset12), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.SlotOffset21Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.SlotOffset21), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreAIR.SlotOffset21), 10)
		}

	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}