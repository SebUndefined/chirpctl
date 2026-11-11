package shared

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TsProtobufToPtrTime(ts *timestamppb.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	t := ts.AsTime()
	return &t
}

type MapPatch struct {
	Set    map[string]string
	Remove []string
}

func UpdateMap(patch *MapPatch, m map[string]string) map[string]string {
	if patch == nil {
		return m
	}

	if m == nil && len(patch.Set) > 0 {
		m = make(map[string]string, len(patch.Set))
	}
	if patch.Set != nil && len(patch.Set) > 0 {
		for k, v := range patch.Set {
			m[k] = v
		}
	}

	for _, v := range patch.Remove {
		delete(m, v)
	}

	return m
}
