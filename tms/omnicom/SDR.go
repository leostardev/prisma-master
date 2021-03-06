//Split diagonostic request(0x41), Iridium
// Code generated by omnicom; DO NOT EDIT
package omnicom

import (
	"errors"
	"strconv"
)

type SDR struct {
	Header       byte
	Date         Date
	Split_Msg_ID uint64
	Padding      uint32
	CRC          uint32
}

func (SDR *SDR) Parse(input string) error {
	var err error
	if len(input) > 112 {
		err = errors.New("Input message is longer than limit")
		return err
	}
	var count int = 0
	if count+8 > len(input) {
		err = errors.New("Input message length of SDR is shorter than required")
		return err
	}
	count = count + 8
	if count+27 > len(input) {
		err = errors.New("Input message length of SDR is shorter than required")
		return err
	}
	SDR.Date = *new(Date)
	err = SDR.Date.parse(input[count : count+27])
	count = count + 27
	if err != nil {
		return err
	}
	var num uint64
	if count+64 > len(input) {
		err = errors.New("Input message length of SDR is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+64], 2, 32)
	SDR.Split_Msg_ID = num*1 - 0
	count = count + 64
	if err != nil {
		return err
	}
	if count+(8-count%8) > len(input) {
		err = errors.New("Input message length of SDR is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+(8-count%8)], 2, 32)
	SDR.Padding = uint32(num*1 - 0)
	count = count + (8 - count%8)
	if err != nil {
		return err
	}
	if count+8 > len(input) {
		err = errors.New("Input message length of SDR is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+8], 2, 32)
	SDR.CRC = uint32(num*1 - 0)
	count = count + 8
	if err != nil {
		return err
	}
	return err
}
func (SDR *SDR) Encode() ([]byte, error) {
	var str string
	var s string
	var err error
	str += pad(strconv.FormatUint(uint64(SDR.Header), 2), 8)
	s, err = SDR.Date.encode()
	if err != nil {
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(((SDR.Split_Msg_ID+0)/1), 2), 64)
	if len(s) > 64 {
		err = errors.New("Value assigned for SDR.Split_Msg_ID exceeds limit")
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((SDR.Padding+0)/1), 2), 8-len(str)%8)
	if len(s) > (8 - len(str)%8) {
		err = errors.New("Value assigned for SDR.Padding exceeds limit")
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((SDR.CRC+0)/1), 2), 8)
	if len(s) > 8 {
		err = errors.New("Value assigned for SDR.CRC exceeds limit")
		return []byte{}, err
	}
	str += s

	str = attachCRC(str)
	byteList, err := decToByte(str)
	return byteList, err
}
func (SDR *SDR) getHeader() byte {
	return SDR.Header
}
func (SDR *SDR) getCRC() uint32 {
	return SDR.CRC
}
func (SDR *SDR) setCRC(crc uint32) {
	SDR.CRC = crc
}
