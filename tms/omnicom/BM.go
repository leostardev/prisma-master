//Binary Message(0x07), Iridium/3G
// Code generated by omnicom; DO NOT EDIT
package omnicom

import (
	"errors"
	"strconv"
)

type BM struct {
	Header             byte
	Date               Date
	ID_Msg             uint32
	Length_Msg_Content uint32
	Msg_Content        []byte
	CRC                uint32
}

func (BM *BM) Parse(input string) error {
	var err error
	if len(input) > 16777288 {
		err = errors.New("Input message is longer than limit")
		return err
	}
	var count int = 0
	if count+8 > len(input) {
		err = errors.New("Input message length of BM is shorter than required")
		return err
	}
	count = count + 8
	if count+27 > len(input) {
		err = errors.New("Input message length of BM is shorter than required")
		return err
	}
	BM.Date = *new(Date)
	err = BM.Date.parse(input[count : count+27])
	count = count + 27
	if err != nil {
		return err
	}
	var num uint64
	if count+16 > len(input) {
		err = errors.New("Input message length of BM is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+16], 2, 32)
	BM.ID_Msg = uint32(num*1 - 0)
	count = count + 16
	if err != nil {
		return err
	}
	if count+21 > len(input) {
		err = errors.New("Input message length of BM is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+21], 2, 32)
	BM.Length_Msg_Content = uint32(num*1 - 0)
	count = count + 21
	if err != nil {
		return err
	}
	if count+int(BM.Length_Msg_Content)*8 > len(input) {
		err = errors.New("Input message length of BM is shorter than required")
		return err
	}
	if 16777208 < int(BM.Length_Msg_Content)*8 {
		err = errors.New("BM.Length_Msg_Content is larger than limit")
		return err
	}
	BM.Msg_Content, err = decodeToByte(input[count : count+int(BM.Length_Msg_Content)*8])
	count = count + int(BM.Length_Msg_Content)*8
	if err != nil {
		return err
	}
	if count+8 > len(input) {
		err = errors.New("Input message length of BM is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+8], 2, 32)
	BM.CRC = uint32(num*1 - 0)
	count = count + 8
	if err != nil {
		return err
	}
	return err
}
func (BM *BM) Encode() ([]byte, error) {
	var str string
	var s string
	var err error
	str += pad(strconv.FormatUint(uint64(BM.Header), 2), 8)
	s, err = BM.Date.encode()
	if err != nil {
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((BM.ID_Msg+0)/1), 2), 16)
	if len(s) > 16 {
		err = errors.New("Value assigned for BM.ID_Msg exceeds limit")
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((BM.Length_Msg_Content+0)/1), 2), 21)
	if len(s) > 21 {
		err = errors.New("Value assigned for BM.Length_Msg_Content exceeds limit")
		return []byte{}, err
	}
	str += s

	if 16777208 < int(BM.Length_Msg_Content)*8 {
		err = errors.New("BM.Length_Msg_Content is larger than limit")
		return []byte{}, err
	}
	s, err = encodeFromByte(BM.Msg_Content, int(BM.Length_Msg_Content)*8)
	if err != nil {
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((BM.CRC+0)/1), 2), 8)
	if len(s) > 8 {
		err = errors.New("Value assigned for BM.CRC exceeds limit")
		return []byte{}, err
	}
	str += s

	str = attachCRC(str)
	byteList, err := decToByte(str)
	return byteList, err
}
func (BM *BM) getHeader() byte {
	return BM.Header
}
func (BM *BM) getCRC() uint32 {
	return BM.CRC
}
func (BM *BM) setCRC(crc uint32) {
	BM.CRC = crc
}