package gateway

import (
	"errors"
	"maps"

	"github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/chirpstack/chirpstack/api/go/v4/common"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"github.com/sebundefined/chirpctl/internal/service/shared"
)

// toGatewayListItem converts a ChirpStack v4 list item into the internal model.
func toGatewayListItem(src *api.GatewayListItem) (gateway.ListItem, error) {
	if src == nil {
		return gateway.ListItem{}, errors.New("gateway list item is nil")
	}

	return gateway.ListItem{
		GatewayID:        src.GetGatewayId(),
		CreatedAt:        shared.TsProtobufToPtrTime(src.GetCreatedAt()),
		UpdatedAt:        shared.TsProtobufToPtrTime(src.GetUpdatedAt()),
		LastSeenAt:       shared.TsProtobufToPtrTime(src.GetLastSeenAt()),
		Name:             src.GetName(),
		Description:      src.GetDescription(),
		Location:         toLocation(src.GetLocation()),
		State:            gateway.ParseStateFromInt32(int32(src.GetState())),
		Metadata:         maps.Clone(src.GetProperties()),
		TenantID:         src.GetTenantId(),
		DownlinkPriority: src.GetDownlinkPriority(),
	}, nil
}

// toGateway converts a ChirpStack v4 response into the internal Gateway model.
func toGateway(resp *api.GetGatewayResponse) (*gateway.Gateway, error) {
	if resp == nil || resp.GetGateway() == nil {
		return nil, errors.New("gateway response is empty")
	}

	src := resp.GetGateway()
	return &gateway.Gateway{
		GatewayID:        src.GatewayId,
		CreatedAt:        shared.TsProtobufToPtrTime(resp.CreatedAt),
		UpdatedAt:        shared.TsProtobufToPtrTime(resp.UpdatedAt),
		LastSeenAt:       shared.TsProtobufToPtrTime(resp.LastSeenAt),
		Name:             src.Name,
		Description:      src.Description,
		Location:         toLocation(src.GetLocation()),
		TenantID:         src.TenantId,
		Tags:             maps.Clone(src.Tags),
		Metadata:         maps.Clone(src.Metadata),
		StatsInterval:    src.StatsInterval,
		DownlinkPriority: src.DownlinkPriority,
	}, nil
}

// toLocation converts an optional ChirpStack v4 location.
func toLocation(src *common.Location) *shared.Location {
	if src == nil {
		return nil
	}

	return &shared.Location{
		Latitude:  src.Latitude,
		Longitude: src.Longitude,
		Altitude:  src.Altitude,
		Source:    shared.ParseLocationSourceFromInt32(int32(src.Source)),
		Accuracy:  src.Accuracy,
	}
}
