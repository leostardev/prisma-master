// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strconv"
import "strings"

// M137122 represents fix data.
type CoreM137122 struct {
	MessageID uint8

	RepeatIndicator uint32

	Mmsi uint32

	Spare1 uint32 //Supposed to be Unknown

	ChannelA uint32

	ChannelB uint32

	TxRxMode uint32

	Power bool

	Longitude1 float64

	Latitude1 float64

	Longitude2 float64

	Latitude2 float64

	AddrBcast bool

	ChannelABandwidth bool

	ChannelBBandwidth bool

	TransitionalZoneSize uint32

	Spare2 uint32 //Supposed to be Unknown
}
type M137122 struct {
	VDMO
	CoreM137122
}

func NewM137122(sentence VDMO) *M137122 {
	s := new(M137122)
	s.VDMO = sentence
	return s
}

func (s *M137122) parse() error {
	var err error

	if MessageType(s.EncapData) != 22 {
		err = fmt.Errorf("message %d is not a M137122", MessageType(s.EncapData))
		return err
	}

	data := []byte(s.EncapData)

	//if len(data)*6 > 168 {
	//	err = fmt.Errorf("Message lenght is larger than it should be [%d!=168]", len(data)*6)
	//	return err
	//}

	s.MessageID = MessageType(s.EncapData)

	s.CoreM137122.RepeatIndicator = BitsToInt(6, 7, data)

	s.CoreM137122.Mmsi = BitsToInt(8, 37, data)

	s.CoreM137122.Spare1 = BitsToInt(38, 39, data)

	s.CoreM137122.ChannelA = BitsToInt(40, 51, data)

	s.CoreM137122.ChannelB = BitsToInt(52, 63, data)

	s.CoreM137122.TxRxMode = BitsToInt(64, 67, data)

	s.CoreM137122.Power = CbnBool(68, data)

	s.CoreM137122.Longitude1 = float64(BitsToInt(69, 86, data))

	s.CoreM137122.Latitude1 = float64(BitsToInt(87, 103, data))

	s.CoreM137122.Longitude2 = float64(BitsToInt(104, 121, data))

	s.CoreM137122.Latitude2 = float64(BitsToInt(122, 138, data))

	s.CoreM137122.AddrBcast = CbnBool(139, data)

	s.CoreM137122.ChannelABandwidth = CbnBool(140, data)

	s.CoreM137122.ChannelBBandwidth = CbnBool(141, data)

	s.CoreM137122.TransitionalZoneSize = BitsToInt(142, 144, data)

	s.CoreM137122.Spare2 = BitsToInt(145, 6*(len(data)-1), data)

	return nil
}

func (s *M137122) Encode() (string, error) {
	var Raw string
	var Sbinary string

	if s.MessageID != 22 {
		err := fmt.Errorf("message %d is not a M137122", s.MessageID)
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

	str := strconv.FormatInt(int64(s.CoreM137122.MessageID), 2)
	for len(str) < 6 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.RepeatIndicator), 2)
	for len(str) < 2 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.Mmsi), 2)
	for len(str) < 30 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.Spare1), 2)
	for len(str) < 2 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.ChannelA), 2)
	for len(str) < 12 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.ChannelB), 2)
	for len(str) < 12 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.TxRxMode), 2)
	for len(str) < 4 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	if s.Power == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.Longitude1), 2)
	for len(str) < 18 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.Latitude1), 2)
	for len(str) < 17 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.Longitude2), 2)
	for len(str) < 18 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.Latitude2), 2)
	for len(str) < 17 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	if s.AddrBcast == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	if s.ChannelABandwidth == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	if s.ChannelBBandwidth == true {
		str = "1"
	} else {
		str = "0"
	}

	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.TransitionalZoneSize), 2)
	for len(str) < 3 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137122.Spare2), 2)
	for len(str) < 23 {
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
