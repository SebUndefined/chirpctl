package shared

import "strings"

type SecondChAckOffset string

const (
	SecondChAckOffsetUnknown SecondChAckOffset = "UNKNOWN"
	SecondChAckOffsetKhz0    SecondChAckOffset = "KHZ_0"
	SecondChAckOffsetKhz200  SecondChAckOffset = "KHZ_200"
	SecondChAckOffsetKhz400  SecondChAckOffset = "KHZ_400"
	SecondChAckOffsetKhz800  SecondChAckOffset = "KHZ_800"
	SecondChAckOffsetKhz1600 SecondChAckOffset = "KHZ_1600"
	SecondChAckOffsetKhz3200 SecondChAckOffset = "KHZ_3200"
)

func (c SecondChAckOffset) String() string {
	if len(strings.TrimSpace(string(c))) == 0 {
		return SecondChAckOffsetUnknown.String()
	}
	return string(c)
}

func ParseSecondChAckOffsetSourceFromInt32(r int32) SecondChAckOffset {
	switch r {
	case 0:
		return SecondChAckOffsetKhz0
	case 1:
		return SecondChAckOffsetKhz200
	case 2:
		return SecondChAckOffsetKhz400
	case 3:
		return SecondChAckOffsetKhz800
	case 4:
		return SecondChAckOffsetKhz1600
	case 5:
		return SecondChAckOffsetKhz3200
	default:
		return SecondChAckOffsetUnknown
	}
}
