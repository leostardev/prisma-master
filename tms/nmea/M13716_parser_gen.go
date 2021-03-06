// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strconv"
import "strings"

// M13716 represents fix data.
type CoreM13716 struct {
	MessageID uint8

	RepeatIndicator uint32

	Mmsi uint32

	SeqNumber uint32

	DestinationID uint32

	RetransmitFlag bool

	Spare uint32 //Supposed to be Unknown

	DesignatedAreaCode uint32

	FunctionID uint32

	ApplicationData uint32 //Supposed to be Unknown
}
type M13716 struct {
	VDMO
	CoreM13716
}

func NewM13716(sentence VDMO) *M13716 {
	s := new(M13716)
	s.VDMO = sentence
	return s
}

func (s *M13716) parse() error {
	var err error

	if MessageType(s.EncapData) != 6 {
		err = fmt.Errorf("message %d is not a M13716", MessageType(s.EncapData))
		return err
	}

	data := []byte(s.EncapData)

	//if len(data)*6 > 1008 {
	//	err = fmt.Errorf("Message lenght is larger than it should be [%d!=1008]", len(data)*6)
	//	return err
	//}

	s.MessageID = MessageType(s.EncapData)

	s.CoreM13716.RepeatIndicator = BitsToInt(6, 7, data)

	s.CoreM13716.Mmsi = BitsToInt(8, 37, data)

	s.CoreM13716.SeqNumber = BitsToInt(38, 39, data)

	s.CoreM13716.DestinationID = BitsToInt(40, 69, data)

	s.CoreM13716.RetransmitFlag = CbnBool(70, data)

	s.CoreM13716.Spare = BitsToInt(71, 71, data)

	s.CoreM13716.DesignatedAreaCode = BitsToInt(72, 81, data)

	s.CoreM13716.FunctionID = BitsToInt(82, 87, data)

	s.CoreM13716.ApplicationData = BitsToInt(88, 6*(len(data)-1), data)

	return nil
}

func (s *M13716) Encode() (string, error) {
	var Raw string
	var Sbinary string

	if s.MessageID != 6 {
		err := fmt.Errorf("message %d is not a M13716", s.MessageID)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format + ","

	if s.SentenceCountValidity == true {
		Raw = Raw + strconv.FormatInt(int64(s.SentenceCount), 10) + ","
	} else {
		Raw = Raw + ","
	}

	if s.SentenceIndexValidity == true {
		Raw = Raw + strconv.FormatInt(int64(s.SentenceIndex), 10) + ","
	} else {
		Raw = Raw + ","
	}

	if s.SeqMsgIDValidity == true {
		Raw = Raw + strconv.FormatInt(int64(s.SeqMsgID), 10) + ","
	} else {
		Raw = Raw + ","
	}

	if s.ChannelValidity == true {
		Raw = Raw + s.Channel
	}

	str := strconv.FormatInt(int64(s.CoreM13716.MessageID), 2)
	for len(str) < 6 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13716.RepeatIndicator), 2)
	for len(str) < 2 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13716.Mmsi), 2)
	for len(str) < 30 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13716.SeqNumber), 2)
	for len(str) < 2 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13716.DestinationID), 2)
	for len(str) < 30 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	if s.RetransmitFlag == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13716.Spare), 2)
	for len(str) < 1 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13716.DesignatedAreaCode), 2)
	for len(str) < 10 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13716.FunctionID), 2)
	for len(str) < 6 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13716.ApplicationData), 2)
	for len(str) < 920 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	field := strings.SplitN(Sbinary, "", len(Sbinary))

	var encdata = make([]string, int((len(Sbinary)+int(s.FillBits))/6))

	j := 0
	for i := 0; i < int((len(Sbinary)+int(s.FillBits))/6); i++ {

		if i == (int((len(Sbinary)+int(s.FillBits))/6) - 1) {
			for j < len(Sbinary) {
				encdata[i] = encdata[i] + field[j]
				j = j + 1
			}
			for h := 0; h < int(s.FillBits); h++ {
				encdata[i] = encdata[i] + "0" // fill bits
			}
		} else {
			encdata[i] = field[j] + field[j+1] + field[j+2] + field[j+3] + field[j+4] + field[j+5]
			j = j + 6
		}
	}

	var data string
	for j := 0; j < int((len(Sbinary)+int(s.FillBits))/6); j++ {
		i, _ := strconv.ParseInt(encdata[j], 2, 8)
		if i < 40 {
			i = i + 48
		} else {
			i = i + 8 + 48
		}
		data = data + string(rune(i))
	}

	Raw = Raw + "," + data + ","

	if s.FillBitsValidity == true {
		Raw = Raw + strconv.FormatInt(int64(s.FillBits), 10)
	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}
