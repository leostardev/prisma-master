//Update global parameters(0x34), Iridium/3G
// Code generated by omnicom; DO NOT EDIT
package omnicom

import (
	"errors"
	"strconv"
)

type UGP struct {
	Header                         byte
	ID_Msg                         uint32
	Date                           Date
	Position_Reporting_Interval    Position_Reporting_Interval_StoV
	Geofencing_Enable              Geofencing_Enable_StoV
	Geofence_Status_Check_Interval Geofence_Status_Check_Interval_StoV
	Password                       Password_StoV
	Routing                        Routing_StoV
	Padding                        uint32
	CRC                            uint32
}

func (UGP *UGP) Parse(input string) error {
	var err error
	if len(input) > 208 {
		err = errors.New("Input message is longer than limit")
		return err
	}
	var count int = 0
	if count+8 > len(input) {
		err = errors.New("Input message length of UGP is shorter than required")
		return err
	}
	count = count + 8
	var num uint64
	if count+12 > len(input) {
		err = errors.New("Input message length of UGP is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+12], 2, 32)
	UGP.ID_Msg = uint32(num*1 - 0)
	count = count + 12
	if err != nil {
		return err
	}
	if count+27 > len(input) {
		err = errors.New("Input message length of UGP is shorter than required")
		return err
	}
	UGP.Date = *new(Date)
	err = UGP.Date.parse(input[count : count+27])
	count = count + 27
	if err != nil {
		return err
	}
	if count+10 > len(input) {
		err = errors.New("Input message length of UGP is shorter than required")
		return err
	}
	UGP.Position_Reporting_Interval = *new(Position_Reporting_Interval_StoV)
	err = UGP.Position_Reporting_Interval.parse(input[count : count+10])
	count = count + 10
	if err != nil {
		return err
	}
	if count+2 > len(input) {
		err = errors.New("Input message length of UGP is shorter than required")
		return err
	}
	UGP.Geofencing_Enable = *new(Geofencing_Enable_StoV)
	err = UGP.Geofencing_Enable.parse(input[count : count+2])
	count = count + 2
	if err != nil {
		return err
	}
	if count+10 > len(input) {
		err = errors.New("Input message length of UGP is shorter than required")
		return err
	}
	UGP.Geofence_Status_Check_Interval = *new(Geofence_Status_Check_Interval_StoV)
	err = UGP.Geofence_Status_Check_Interval.parse(input[count : count+10])
	count = count + 10
	if err != nil {
		return err
	}
	if count+121 > len(input) {
		err = errors.New("Input message length of UGP is shorter than required")
		return err
	}
	UGP.Password = *new(Password_StoV)
	err = UGP.Password.parse(input[count : count+121])
	count = count + 121
	if err != nil {
		return err
	}
	if count+3 > len(input) {
		err = errors.New("Input message length of UGP is shorter than required")
		return err
	}
	UGP.Routing = *new(Routing_StoV)
	err = UGP.Routing.parse(input[count : count+3])
	count = count + 3
	if err != nil {
		return err
	}
	if count+(8-count%8) > len(input) {
		err = errors.New("Input message length of UGP is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+(8-count%8)], 2, 32)
	UGP.Padding = uint32(num*1 - 0)
	count = count + (8 - count%8)
	if err != nil {
		return err
	}
	if count+8 > len(input) {
		err = errors.New("Input message length of UGP is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+8], 2, 32)
	UGP.CRC = uint32(num*1 - 0)
	count = count + 8
	if err != nil {
		return err
	}
	return err
}
func (UGP *UGP) Encode() ([]byte, error) {
	var str string
	var s string
	var err error
	str += pad(strconv.FormatUint(uint64(UGP.Header), 2), 8)
	s = pad(strconv.FormatUint(uint64((UGP.ID_Msg+0)/1), 2), 12)
	if len(s) > 12 {
		err = errors.New("Value assigned for UGP.ID_Msg exceeds limit")
		return []byte{}, err
	}
	str += s

	s, err = UGP.Date.encode()
	if err != nil {
		return []byte{}, err
	}
	str += s

	s, err = UGP.Position_Reporting_Interval.encode()
	if err != nil {
		return []byte{}, err
	}
	str += s

	s, err = UGP.Geofencing_Enable.encode()
	if err != nil {
		return []byte{}, err
	}
	str += s

	s, err = UGP.Geofence_Status_Check_Interval.encode()
	if err != nil {
		return []byte{}, err
	}
	str += s

	s, err = UGP.Password.encode()
	if err != nil {
		return []byte{}, err
	}
	str += s

	s, err = UGP.Routing.encode()
	if err != nil {
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((UGP.Padding+0)/1), 2), 8-len(str)%8)
	if len(s) > (8 - len(str)%8) {
		err = errors.New("Value assigned for UGP.Padding exceeds limit")
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((UGP.CRC+0)/1), 2), 8)
	if len(s) > 8 {
		err = errors.New("Value assigned for UGP.CRC exceeds limit")
		return []byte{}, err
	}
	str += s

	str = attachCRC(str)
	byteList, err := decToByte(str)
	return byteList, err
}
func (UGP *UGP) getHeader() byte {
	return UGP.Header
}
func (UGP *UGP) getCRC() uint32 {
	return UGP.CRC
}
func (UGP *UGP) setCRC(crc uint32) {
	UGP.CRC = crc
}
