package output

import "strings"

type Action string

const (
	ActionUnknown Action = "UNKNOWN"
	ActionCreate  Action = "CREATE"
	ActionDelete  Action = "DELETE"
	ActionUpdate  Action = "UPDATE"
)

func (a Action) String() string {
	if len(strings.TrimSpace(string(a))) == 0 {
		return ActionUnknown.String()
	}
	return string(a)
}

func (a Action) PastTense() string {
	switch a {
	case ActionCreate:
		return "created"
	case ActionDelete:
		return "deleted"
	case ActionUpdate:
		return "updated"
	default:
		return strings.ToLower(ActionUnknown.String())
	}
}

type MutationResult struct {
	Kind   string `json:"kind" yaml:"kind"`
	ID     string `json:"id" yaml:"id"`
	Action Action `json:"action" yaml:"action"`
}
