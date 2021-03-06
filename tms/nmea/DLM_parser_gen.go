// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strings"
import "strconv"

const (
	// PrefixDLM prefix
	PrefixDLM = "DLM"
)

// DLM represents fix data.
type CoreDLM struct {
	RecordIndex uint32

	RecordIndexValidity bool

	Channel string

	ChannelValidity bool

	ReservationOwnership1 string

	ReservationOwnership1Validity bool

	ReservationStartSlot1 uint32

	ReservationStartSlot1Validity bool

	ReservationCount1 uint32

	ReservationCount1Validity bool

	ReservationTimeout1 uint32

	ReservationTimeout1Validity bool

	ReservationIncrement1 uint32

	ReservationIncrement1Validity bool

	ReservationOwnership2 string

	ReservationOwnership2Validity bool

	ReservationStartSlot2 uint32

	ReservationStartSlot2Validity bool

	ReservationCount2 uint32

	ReservationCount2Validity bool

	ReservationTimeout2 uint32

	ReservationTimeout2Validity bool

	ReservationIncrement2 uint32

	ReservationIncrement2Validity bool

	ReservationOwnership3 string

	ReservationOwnership3Validity bool

	ReservationStartSlot3 uint32

	ReservationStartSlot3Validity bool

	ReservationCount3 uint32

	ReservationCount3Validity bool

	ReservationTimeout3 uint32

	ReservationTimeout3Validity bool

	ReservationIncrement3 uint32

	ReservationIncrement3Validity bool

	ReservationOwnership4 string

	ReservationOwnership4Validity bool

	ReservationStartSlot4 uint32

	ReservationStartSlot4Validity bool

	ReservationCount4 uint32

	ReservationCount4Validity bool

	ReservationTimeout4 uint32

	ReservationTimeout4Validity bool

	ReservationIncrement4 uint32

	ReservationIncrement4Validity bool

	Status string

	StatusValidity bool
}

type DLM struct {
	BaseSentence
	CoreDLM
}

func NewDLM(sentence BaseSentence) *DLM {
	s := new(DLM)
	s.BaseSentence = sentence

	s.RecordIndexValidity = false

	s.ChannelValidity = false

	s.ReservationOwnership1Validity = false

	s.ReservationStartSlot1Validity = false

	s.ReservationCount1Validity = false

	s.ReservationTimeout1Validity = false

	s.ReservationIncrement1Validity = false

	s.ReservationOwnership2Validity = false

	s.ReservationStartSlot2Validity = false

	s.ReservationCount2Validity = false

	s.ReservationTimeout2Validity = false

	s.ReservationIncrement2Validity = false

	s.ReservationOwnership3Validity = false

	s.ReservationStartSlot3Validity = false

	s.ReservationCount3Validity = false

	s.ReservationTimeout3Validity = false

	s.ReservationIncrement3Validity = false

	s.ReservationOwnership4Validity = false

	s.ReservationStartSlot4Validity = false

	s.ReservationCount4Validity = false

	s.ReservationTimeout4Validity = false

	s.ReservationIncrement4Validity = false

	s.StatusValidity = false

	return s
}

func (s *DLM) parse() error {
	var err error

	if s.Format != PrefixDLM {
		err = fmt.Errorf("%s is not a %s", s.Format, PrefixDLM)
		return err
	}

	if len(s.Fields) == 0 {
		return nil
	} else {
		if s.Fields[0] != "" {
			i, err := strconv.ParseUint(s.Fields[0], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[0])
			} else {
				s.CoreDLM.RecordIndex = uint32(i)
				s.CoreDLM.RecordIndexValidity = true
			}

		}
	}

	if len(s.Fields) == 1 {
		return nil
	} else {
		if s.Fields[1] != "" {
			s.Channel = s.Fields[1]
			s.ChannelValidity = true
		}
	}

	if len(s.Fields) == 2 {
		return nil
	} else {
		if s.Fields[2] != "" {
			s.ReservationOwnership1 = s.Fields[2]
			s.ReservationOwnership1Validity = true
		}
	}

	if len(s.Fields) == 3 {
		return nil
	} else {
		if s.Fields[3] != "" {
			i, err := strconv.ParseUint(s.Fields[3], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[3])
			} else {
				s.CoreDLM.ReservationStartSlot1 = uint32(i)
				s.CoreDLM.ReservationStartSlot1Validity = true
			}

		}
	}

	if len(s.Fields) == 4 {
		return nil
	} else {
		if s.Fields[4] != "" {
			i, err := strconv.ParseUint(s.Fields[4], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[4])
			} else {
				s.CoreDLM.ReservationCount1 = uint32(i)
				s.CoreDLM.ReservationCount1Validity = true
			}

		}
	}

	if len(s.Fields) == 5 {
		return nil
	} else {
		if s.Fields[5] != "" {
			i, err := strconv.ParseUint(s.Fields[5], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[5])
			} else {
				s.CoreDLM.ReservationTimeout1 = uint32(i)
				s.CoreDLM.ReservationTimeout1Validity = true
			}

		}
	}

	if len(s.Fields) == 6 {
		return nil
	} else {
		if s.Fields[6] != "" {
			i, err := strconv.ParseUint(s.Fields[6], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[6])
			} else {
				s.CoreDLM.ReservationIncrement1 = uint32(i)
				s.CoreDLM.ReservationIncrement1Validity = true
			}

		}
	}

	if len(s.Fields) == 7 {
		return nil
	} else {
		if s.Fields[7] != "" {
			s.ReservationOwnership2 = s.Fields[7]
			s.ReservationOwnership2Validity = true
		}
	}

	if len(s.Fields) == 8 {
		return nil
	} else {
		if s.Fields[8] != "" {
			i, err := strconv.ParseUint(s.Fields[8], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[8])
			} else {
				s.CoreDLM.ReservationStartSlot2 = uint32(i)
				s.CoreDLM.ReservationStartSlot2Validity = true
			}

		}
	}

	if len(s.Fields) == 9 {
		return nil
	} else {
		if s.Fields[9] != "" {
			i, err := strconv.ParseUint(s.Fields[9], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[9])
			} else {
				s.CoreDLM.ReservationCount2 = uint32(i)
				s.CoreDLM.ReservationCount2Validity = true
			}

		}
	}

	if len(s.Fields) == 10 {
		return nil
	} else {
		if s.Fields[10] != "" {
			i, err := strconv.ParseUint(s.Fields[10], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[10])
			} else {
				s.CoreDLM.ReservationTimeout2 = uint32(i)
				s.CoreDLM.ReservationTimeout2Validity = true
			}

		}
	}

	if len(s.Fields) == 11 {
		return nil
	} else {
		if s.Fields[11] != "" {
			i, err := strconv.ParseUint(s.Fields[11], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[11])
			} else {
				s.CoreDLM.ReservationIncrement2 = uint32(i)
				s.CoreDLM.ReservationIncrement2Validity = true
			}

		}
	}

	if len(s.Fields) == 12 {
		return nil
	} else {
		if s.Fields[12] != "" {
			s.ReservationOwnership3 = s.Fields[12]
			s.ReservationOwnership3Validity = true
		}
	}

	if len(s.Fields) == 13 {
		return nil
	} else {
		if s.Fields[13] != "" {
			i, err := strconv.ParseUint(s.Fields[13], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[13])
			} else {
				s.CoreDLM.ReservationStartSlot3 = uint32(i)
				s.CoreDLM.ReservationStartSlot3Validity = true
			}

		}
	}

	if len(s.Fields) == 14 {
		return nil
	} else {
		if s.Fields[14] != "" {
			i, err := strconv.ParseUint(s.Fields[14], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[14])
			} else {
				s.CoreDLM.ReservationCount3 = uint32(i)
				s.CoreDLM.ReservationCount3Validity = true
			}

		}
	}

	if len(s.Fields) == 15 {
		return nil
	} else {
		if s.Fields[15] != "" {
			i, err := strconv.ParseUint(s.Fields[15], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[15])
			} else {
				s.CoreDLM.ReservationTimeout3 = uint32(i)
				s.CoreDLM.ReservationTimeout3Validity = true
			}

		}
	}

	if len(s.Fields) == 16 {
		return nil
	} else {
		if s.Fields[16] != "" {
			i, err := strconv.ParseUint(s.Fields[16], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[16])
			} else {
				s.CoreDLM.ReservationIncrement3 = uint32(i)
				s.CoreDLM.ReservationIncrement3Validity = true
			}

		}
	}

	if len(s.Fields) == 17 {
		return nil
	} else {
		if s.Fields[17] != "" {
			s.ReservationOwnership4 = s.Fields[17]
			s.ReservationOwnership4Validity = true
		}
	}

	if len(s.Fields) == 18 {
		return nil
	} else {
		if s.Fields[18] != "" {
			i, err := strconv.ParseUint(s.Fields[18], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[18])
			} else {
				s.CoreDLM.ReservationStartSlot4 = uint32(i)
				s.CoreDLM.ReservationStartSlot4Validity = true
			}

		}
	}

	if len(s.Fields) == 19 {
		return nil
	} else {
		if s.Fields[19] != "" {
			i, err := strconv.ParseUint(s.Fields[19], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[19])
			} else {
				s.CoreDLM.ReservationCount4 = uint32(i)
				s.CoreDLM.ReservationCount4Validity = true
			}

		}
	}

	if len(s.Fields) == 20 {
		return nil
	} else {
		if s.Fields[20] != "" {
			i, err := strconv.ParseUint(s.Fields[20], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[20])
			} else {
				s.CoreDLM.ReservationTimeout4 = uint32(i)
				s.CoreDLM.ReservationTimeout4Validity = true
			}

		}
	}

	if len(s.Fields) == 21 {
		return nil
	} else {
		if s.Fields[21] != "" {
			i, err := strconv.ParseUint(s.Fields[21], 10, 32)
			if err != nil {
				return fmt.Errorf("DLM decode variation error: %s", s.Fields[21])
			} else {
				s.CoreDLM.ReservationIncrement4 = uint32(i)
				s.CoreDLM.ReservationIncrement4Validity = true
			}

		}
	}

	if len(s.Fields) == 22 {
		return nil
	} else {
		if s.Fields[22] != "" {
			s.Status = s.Fields[22]
			s.StatusValidity = true
		}
	}

	return nil
}

func (s *DLM) Encode() (string, error) {
	var Raw string

	if s.Format != PrefixDLM {
		err := fmt.Errorf("Sentence format %s is not a DLM sentence", s.Format)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format

	if s.RecordIndexValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.RecordIndex), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.RecordIndex), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ChannelValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreDLM.Channel

		} else {
			Raw = Raw + "," + s.CoreDLM.Channel
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationOwnership1Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreDLM.ReservationOwnership1

		} else {
			Raw = Raw + "," + s.CoreDLM.ReservationOwnership1
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationStartSlot1Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationStartSlot1), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationStartSlot1), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationCount1Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationCount1), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationCount1), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationTimeout1Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationTimeout1), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationTimeout1), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationIncrement1Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationIncrement1), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationIncrement1), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationOwnership2Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreDLM.ReservationOwnership2

		} else {
			Raw = Raw + "," + s.CoreDLM.ReservationOwnership2
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationStartSlot2Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationStartSlot2), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationStartSlot2), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationCount2Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationCount2), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationCount2), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationTimeout2Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationTimeout2), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationTimeout2), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationIncrement2Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationIncrement2), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationIncrement2), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationOwnership3Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreDLM.ReservationOwnership3

		} else {
			Raw = Raw + "," + s.CoreDLM.ReservationOwnership3
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationStartSlot3Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationStartSlot3), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationStartSlot3), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationCount3Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationCount3), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationCount3), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationTimeout3Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationTimeout3), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationTimeout3), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationIncrement3Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationIncrement3), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationIncrement3), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationOwnership4Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreDLM.ReservationOwnership4

		} else {
			Raw = Raw + "," + s.CoreDLM.ReservationOwnership4
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationStartSlot4Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationStartSlot4), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationStartSlot4), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationCount4Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationCount4), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationCount4), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationTimeout4Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationTimeout4), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationTimeout4), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ReservationIncrement4Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationIncrement4), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CoreDLM.ReservationIncrement4), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.StatusValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreDLM.Status

		} else {
			Raw = Raw + "," + s.CoreDLM.Status
		}
	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}
