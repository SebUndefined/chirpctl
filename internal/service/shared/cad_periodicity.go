package shared

import "strings"

type CadPeriodicity string

const (
	CadPeriodicityUnknown CadPeriodicity = "UNKNOWN"
	CadPeriodicitySEC1    CadPeriodicity = "SEC_1"
	CadPeriodicityMS500   CadPeriodicity = "MS_500"
	CadPeriodicityMS250   CadPeriodicity = "MS_250"
	CadPeriodicityMS100   CadPeriodicity = "MS_100"
	CadPeriodicityMS20    CadPeriodicity = "MS_20"
	CadPeriodicityMS50    CadPeriodicity = "MS_50"
)

func (c CadPeriodicity) String() string {
	if len(strings.TrimSpace(string(c))) == 0 {
		return CadPeriodicityUnknown.String()
	}
	return string(c)
}

func ParseCadPeriodicitySourceFromInt32(r int32) CadPeriodicity {
	switch r {
	case 0:
		return CadPeriodicitySEC1
	case 1:
		return CadPeriodicityMS500
	case 2:
		return CadPeriodicityMS250
	case 3:
		return CadPeriodicityMS100
	case 4:
		return CadPeriodicityMS20
	case 5:
		return CadPeriodicityMS50
	default:
		return CadPeriodicityUnknown
	}
}
