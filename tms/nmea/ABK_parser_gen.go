// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strings"
import "strconv"

const (
	// PrefixABK prefix
	PrefixABK = "ABK"
)

// ABK represents fix data.
type CoreABK struct {
	DestinationID uint32

	DestinationIDValidity bool

	ReceivedChannel string

	ReceivedChannelValidity bool

	MessageID uint32

	MessageIDValidity bool

	SeqMsgID uint32

	SeqMsgIDValidity bool

	AckType uint32

	AckTypeValidity bool
}

type ABK struct {
	BaseSentence
	CoreABK
}

func NewABK(sentence BaseSentence) *ABK {
	s := new(ABK)
	s.BaseSentence = sentence

	s.DestinationIDValidity = false

	s.ReceivedChannelValidity = false

	s.MessageIDValidity = false

	s.SeqMsgIDValidity = false

	s.AckTypeValidity = false

	return s
}

func (s *ABK) parse() error {
	var err error

	if s.Format != PrefixABK {
		err = fmt.Errorf("%s is not a %s", s.Format, PrefixABK)
		return err
	}

	if len(s.Fields) == 0 {
		return nil
	} else {
		if s.Fields[0] != "" {
			i, err := strconv.ParseUint(s.Fields[0], 10, 32)
			if err != nil {
				return fmt.Errorf("ABK decode variation error: %s", s.Fields[0])
			} else {
				s.CoreABK.DestinationID = uint32(i)
				s.CoreABK.DestinationIDValidity = true
			}

		}
	}

	if len(s.Fields) == 1 {
		return nil
	} else {
		if s.Fields[1] != "" {
			s.ReceivedChannel = s.Fields[1]
			s.ReceivedChannelValidity = true
		}
	}

	if len(s.Fields) == 2 {
		return nil
	} else {
		if s.Fields[2] != "" {
			i, err := strconv.ParseUint(s.Fields[2], 10, 32)
			if err != nil {
				return fmt.Errorf("ABK decode variation error: %s", s.Fields[2])
			} else {
				s.CoreABK.MessageID = uint32(i)
				s.CoreABK.MessageIDValidity = true
			}

		}
	}

	if len(s.Fields) == 3 {
		return nil
	} else {
		if s.Fields[3] != "" {
			i, err := strconv.ParseUint(s.Fields[3], 10, 32)
			if err != nil {
				return fmt.Errorf("ABK decode variation error: %s", s.Fields[3])
			} else {
				s.CoreABK.SeqMsgID = uint32(i)
				s.CoreABK.SeqMsgIDValidity = true
			}

		}
	}

	if len(s.Fields) == 4 {
		return nil
	} else {
		if s.Fields[4] != "" {
			i, err := strconv.ParseUint(s.Fields[4], 10, 32)
			if err != nil {
				return fmt.Errorf("ABK decode variation error: %s", s.Fields[4])
			} else {
				s.CoreABK.AckType = uint32(i)
				s.CoreABK.AckTypeValidity = true
			}

		}
	}

	return nil
}

func (s *ABK) Encode() (string, error) {
	var Raw string

	if s.Format != PrefixABK {
		err := fmt.Errorf("Sentence format %s is not a ABK sentence", s.Format)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format

	if s.DestinationIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreABK.DestinationID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreABK.DestinationID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReceivedChannelValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreABK.ReceivedChannel

		} else {
			Raw = Raw + "," + s.CoreABK.ReceivedChannel
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.MessageIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreABK.MessageID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreABK.MessageID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.SeqMsgIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreABK.SeqMsgID), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreABK.SeqMsgID), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.AckTypeValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreABK.AckType), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreABK.AckType), 10)
		}

	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}