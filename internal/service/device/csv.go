package device

import (
	"fmt"

	"github.com/sebundefined/chirpctl/internal/csvx"
	"github.com/sebundefined/chirpctl/internal/validation"
	"github.com/sebundefined/chirpctl/utils"
)

type Specs struct {
	DevEUI, Name, Description, ApplicationID, DeviceProfileID, SkipFcntCheck int
	IsDisabled, Variables, Tags, JoinEUI                                     int
	maxUsedIdx                                                               int
}

func (s *Specs) finalize() {
	s.maxUsedIdx = max(
		s.DevEUI,
		s.Name,
		s.Description,
		s.ApplicationID,
		s.DeviceProfileID,
		s.SkipFcntCheck,
		s.Variables,
		s.Tags,
		s.JoinEUI)
}

func (s *Specs) set(col string, value int) {
	switch col {
	case "devEUI":
		s.DevEUI = value
	case "name":
		s.Name = value
	case "description":
		s.Description = value
	case "applicationID":
		s.ApplicationID = value
	case "deviceProfileID":
		s.DeviceProfileID = value
	case "skipFcntCheck":
		s.SkipFcntCheck = value
	case "isDisabled":
		s.IsDisabled = value
	case "variables":
		s.Variables = value
	case "tags":
		s.Tags = value
	case "joinEUI":
		s.JoinEUI = value
	}
}

var requiredCols = []string{
	"name",
	"description",
	"devEUI",
	"applicationID",
	"deviceProfileID",
}
var optionalCols = []string{
	"skipFcntCheck",
	"isDisabled",
	"variables",
	"tags",
	"joinEUI",
}

func CompileDeviceCsvSpecs(idx csvx.ColumnIndex) (Specs, error) {
	specs := Specs{}
	for _, col := range requiredCols {
		if pos, ok := idx[col]; !ok {
			return specs, fmt.Errorf("column '%s' does not exist", col)
		} else {
			specs.set(col, pos)
		}
	}
	for _, col := range optionalCols {
		if pos, ok := idx[col]; ok {
			specs.set(col, pos)
		} else {
			specs.set(col, -1)
		}
	}
	specs.finalize()
	return specs, nil
}

func DecodeDeviceCsvSpecs(row []string, s Specs) (Device, error) {
	if len(row) <= s.maxUsedIdx {
		return Device{}, fmt.Errorf("invalid row length %d, expected at least %d", len(row), s.maxUsedIdx+1)
	}
	d := Device{
		Description: row[s.Description],
	}
	if err := validation.ValidateHexLen(row[s.DevEUI], 8); err != nil {
		return Device{}, fmt.Errorf("invalid devEUI: %w", err)
	} else {
		d.DevEUI = row[s.DevEUI]
	}
	if err := validation.NotEmptyString(row[s.Name]); err != nil {
		return Device{}, fmt.Errorf("invalid name %s: %w", row[s.Name], err)
	} else {
		d.Name = row[s.Name]
	}
	if err := validation.ValidateUUID(row[s.ApplicationID]); err != nil {
		return Device{}, fmt.Errorf("invalid applicationID: %w", err)
	} else {
		d.ApplicationID = row[s.ApplicationID]
	}
	if err := validation.ValidateUUID(row[s.DeviceProfileID]); err != nil {
		return Device{}, fmt.Errorf("invalid DeviceProfileID: %w", err)
	} else {
		d.DeviceProfileID = row[s.DeviceProfileID]
	}
	if v, exist, err := utils.ReadOptionalBool(s.SkipFcntCheck, row); err != nil {
		return Device{}, fmt.Errorf("invalid SkipFcntCheck: %w", err)
	} else if exist {
		d.SkipFcntCheck = v
	}
	if v, exist, err := utils.ReadOptionalBool(s.IsDisabled, row); err != nil {
		return Device{}, fmt.Errorf("invalid IsDisabled: %w", err)
	} else if exist {
		d.IsDisabled = v
	}
	if v, exist, err := utils.ReadOptionalMap(s.Variables, row); err != nil {
		return Device{}, fmt.Errorf("invalid Variables: %w", err)
	} else if exist {
		d.Variables = v
	}
	if v, exist, err := utils.ReadOptionalMap(s.Tags, row); err != nil {
		return Device{}, fmt.Errorf("invalid Tags: %w", err)
	} else if exist {
		d.Tags = v
	}
	if s.JoinEUI != -1 {
		if err := validation.ValidateHexLen(row[s.JoinEUI], 8); err != nil {
			return Device{}, fmt.Errorf("invalid JoinEUI: %w", err)
		} else {
			d.JoinEUI = row[s.JoinEUI]
		}
	}

	return d, nil
}
