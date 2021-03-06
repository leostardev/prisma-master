// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strconv"
import "strings"

// M137118 represents fix data.
type CoreM137118 struct {
	MessageID uint8

	RepeatIndicator uint32

	Mmsi uint32

	Spare uint32

	SpeedOverGround uint32

	PositionAccuracy bool

	Longitude float64

	Latitude float64

	CourseOverGround uint32

	TrueHeading uint32

	TimeStamp uint32

	Spare2 uint32 //Supposed to be Unknown

	UnitFlag bool

	DisplayFlag bool

	DscFlag bool

	BandFlag bool

	Message22Flag bool

	ModeFlag bool

	RaimFlag bool

	CommStateSelectFlag bool

	CommState uint32 //Supposed to be Unknown
}
type M137118 struct {
	VDMO
	CoreM137118
}

func NewM137118(sentence VDMO) *M137118 {
	s := new(M137118)
	s.VDMO = sentence
	return s
}

func (s *M137118) parse() error {
	var err error

	if MessageType(s.EncapData) != 18 {
		err = fmt.Errorf("message %d is not a M137118", MessageType(s.EncapData))
		return err
	}

	data := []byte(s.EncapData)

	//if len(data)*6 > 168 {
	//	err = fmt.Errorf("Message lenght is larger than it should be [%d!=168]", len(data)*6)
	//	return err
	//}

	s.MessageID = MessageType(s.EncapData)

	s.CoreM137118.RepeatIndicator = BitsToInt(6, 7, data)

	s.CoreM137118.Mmsi = BitsToInt(8, 37, data)

	s.CoreM137118.Spare = BitsToInt(38, 45, data)

	s.CoreM137118.SpeedOverGround = BitsToInt(46, 55, data)

	s.CoreM137118.PositionAccuracy = CbnBool(56, data)

	s.CoreM137118.Longitude = (float64(int32(BitsToInt(57, 84, data)) << 4)) / 16

	s.CoreM137118.Latitude = (float64(int32(BitsToInt(85, 111, data)) << 5)) / 32

	s.CoreM137118.CourseOverGround = BitsToInt(112, 123, data)

	s.CoreM137118.TrueHeading = BitsToInt(124, 132, data)

	s.CoreM137118.TimeStamp = BitsToInt(133, 138, data)

	s.CoreM137118.Spare2 = BitsToInt(139, 140, data)

	s.CoreM137118.UnitFlag = CbnBool(141, data)

	s.CoreM137118.DisplayFlag = CbnBool(142, data)

	s.CoreM137118.DscFlag = CbnBool(143, data)

	s.CoreM137118.BandFlag = CbnBool(144, data)

	s.CoreM137118.Message22Flag = CbnBool(145, data)

	s.CoreM137118.ModeFlag = CbnBool(146, data)

	s.CoreM137118.RaimFlag = CbnBool(147, data)

	s.CoreM137118.CommStateSelectFlag = CbnBool(148, data)

	s.CoreM137118.CommState = BitsToInt(149, 6*(len(data)-1), data)

	return nil
}

func (s *M137118) Encode() (string, error) {
	var Raw string
	var Sbinary string

	if s.MessageID != 18 {
		err := fmt.Errorf("message %d is not a M137118", s.MessageID)
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

	str := strconv.FormatInt(int64(s.CoreM137118.MessageID), 2)
	for len(str) < 6 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137118.RepeatIndicator), 2)
	for len(str) < 2 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137118.Mmsi), 2)
	for len(str) < 30 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137118.Spare), 2)
	for len(str) < 8 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137118.SpeedOverGround), 2)
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

	str = strconv.FormatInt(int64(s.CoreM137118.Longitude), 2)
	for len(str) < 28 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137118.Latitude), 2)
	for len(str) < 27 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137118.CourseOverGround), 2)
	for len(str) < 12 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137118.TrueHeading), 2)
	for len(str) < 9 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137118.TimeStamp), 2)
	for len(str) < 6 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137118.Spare2), 2)
	for len(str) < 2 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	if s.UnitFlag == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	if s.DisplayFlag == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	if s.DscFlag == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	if s.BandFlag == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	if s.Message22Flag == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	if s.ModeFlag == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	if s.RaimFlag == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	if s.CommStateSelectFlag == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137118.CommState), 2)
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
