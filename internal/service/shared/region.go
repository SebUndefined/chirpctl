package shared

import "strings"

type Region string

const (
	RegionUnknown Region = "UNKNOWN"
	RegionEU868   Region = "EU_868"
	RegionUS915   Region = "US_915"
	RegionCN779   Region = "CN_779"
	RegionEU433   Region = "EU_433"
	RegionAU915   Region = "AU_915"
	RegionCN470   Region = "CN_470"
	RegionAS923   Region = "AS_923"
	RegionAS9232  Region = "AS_923_2"
	RegionAS9233  Region = "AS_923_3"
	RegionAS9234  Region = "AS_923_4"
	RegionKR920   Region = "KR_920"
	RegionIN865   Region = "IN_865"
	RegionRU864   Region = "RU_864"
	RegionISM2400 Region = "ISM_2400"
)

func (r Region) String() string {
	if len(strings.TrimSpace(string(r))) == 0 {
		return RegionUnknown.String()
	}
	return string(r)
}

func ParseRegionSourceFromInt32(r int32) Region {
	switch r {
	case 0:
		return RegionEU868
	case 2:
		return RegionUS915
	case 3:
		return RegionCN779
	case 4:
		return RegionEU433
	case 5:
		return RegionAU915
	case 6:
		return RegionCN470
	case 7:
		return RegionAS923
	case 8:
		return RegionKR920
	case 9:
		return RegionIN865
	case 10:
		return RegionRU864
	case 11:
		return RegionISM2400
	case 12:
		return RegionAS9232
	case 13:
		return RegionAS9233
	case 14:
		return RegionAS9234
	default:
		return RegionUnknown
	}
}
