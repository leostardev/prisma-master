// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strconv"
import "strings"

// M13713 represents fix data.
type CoreM13713 struct {
	MessageID uint8

	RepeatIndicator uint32

	Mmsi uint32

	NavigationalStatus uint32

	RateOfTurn int8

	SpeedOverGround uint32

	PositionAccuracy bool

	Longitude float64

	Latitude float64

	CourseOverGround uint32

	TrueHeading uint32

	TimeStamp uint32

	SpecialManoeuvre uint32

	Spare uint32 //Supposed to be Unknown

	RaimFlag bool

	CommState uint32 //Supposed to be Unknown
}
type M13713 struct {
	VDMO
	CoreM13713
}

func NewM13713(sentence VDMO) *M13713 {
	s := new(M13713)
	s.VDMO = sentence
	return s
}

func (s *M13713) parse() error {
	var err error

	if MessageType(s.EncapData) != 3 {
		err = fmt.Errorf("message %d is not a M13713", MessageType(s.EncapData))
		return err
	}

	data := []byte(s.EncapData)

	//if len(data)*6 > 168 {
	//	err = fmt.Errorf("Message lenght is larger than it should be [%d!=168]", len(data)*6)
	//	return err
	//}

	s.MessageID = MessageType(s.EncapData)

	s.CoreM13713.RepeatIndicator = BitsToInt(6, 7, data)

	s.CoreM13713.Mmsi = BitsToInt(8, 37, data)

	s.CoreM13713.NavigationalStatus = BitsToInt(38, 41, data)

	s.CoreM13713.RateOfTurn = int8(BitsToInt(42, 49, data))

	s.CoreM13713.SpeedOverGround = BitsToInt(50, 59, data)

	s.CoreM13713.PositionAccuracy = CbnBool(60, data)

	s.CoreM13713.Longitude = (float64(int32(BitsToInt(61, 88, data)) << 4)) / 16

	s.CoreM13713.Latitude = (float64(int32(BitsToInt(89, 115, data)) << 5)) / 32

	s.CoreM13713.CourseOverGround = BitsToInt(116, 127, data)

	s.CoreM13713.TrueHeading = BitsToInt(128, 136, data)

	s.CoreM13713.TimeStamp = BitsToInt(137, 142, data)

	s.CoreM13713.SpecialManoeuvre = BitsToInt(143, 144, data)

	s.CoreM13713.Spare = BitsToInt(145, 147, data)

	s.CoreM13713.RaimFlag = CbnBool(148, data)

	s.CoreM13713.CommState = BitsToInt(149, 6*(len(data)-1), data)

	return nil
}

func (s *M13713) Encode() (string, error) {
	var Raw string
	var Sbinary string

	if s.MessageID != 3 {
		err := fmt.Errorf("message %d is not a M13713", s.MessageID)
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

	str := strconv.FormatInt(int64(s.CoreM13713.MessageID), 2)
	for len(str) < 6 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.RepeatIndicator), 2)
	for len(str) < 2 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.Mmsi), 2)
	for len(str) < 30 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.NavigationalStatus), 2)
	for len(str) < 4 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.RateOfTurn), 2)
	for len(str) < 8 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.SpeedOverGround), 2)
	for len(str) < 10 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	if s.PositionAccuracy == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.Longitude), 2)
	for len(str) < 28 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.Latitude), 2)
	for len(str) < 27 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.CourseOverGround), 2)
	for len(str) < 12 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.TrueHeading), 2)
	for len(str) < 9 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.TimeStamp), 2)
	for len(str) < 6 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.SpecialManoeuvre), 2)
	for len(str) < 2 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.Spare), 2)
	for len(str) < 3 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	if s.RaimFlag == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM13713.CommState), 2)
	for len(str) < 19 {
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
