package tenant

import (
	"errors"
	"maps"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/shared"
	"github.com/sebundefined/chirpctl/internal/service/tenant"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toTenantListItem(src *apiv4.TenantListItem) (tenant.ListItem, error) {
	if src == nil {
		return tenant.ListItem{}, errors.New("tenant list item is nil")
	}
	return tenant.ListItem{
		ID:                  src.Id,
		CreatedAt:           shared.TsProtobufToPtrTime(src.GetCreatedAt()),
		UpdatedAt:           shared.TsProtobufToPtrTime(src.GetUpdatedAt()),
		Name:                src.GetName(),
		CanHaveGateways:     src.CanHaveGateways,
		MaxGatewayCount:     src.MaxGatewayCount,
		MaxDeviceCount:      src.MaxDeviceCount,
		PrivateGatewaysUp:   src.PrivateGatewaysUp,
		PrivateGatewaysDown: src.PrivateGatewaysDown,
	}, nil
}

func toTenant(src *apiv4.Tenant, createdAt *timestamppb.Timestamp, updatedAt *timestamppb.Timestamp) (*tenant.Tenant, error) {
	if src == nil {
		return nil, errors.New("tenant response is nil")
	}
	result := &tenant.Tenant{
		ID:                  src.Id,
		CreatedAt:           shared.TsProtobufToPtrTime(createdAt),
		UpdatedAt:           shared.TsProtobufToPtrTime(updatedAt),
		Name:                src.Name,
		Description:         src.Description,
		CanHaveGateways:     src.CanHaveGateways,
		MaxGatewayCount:     src.MaxGatewayCount,
		MaxDeviceCount:      src.MaxDeviceCount,
		PrivateGatewaysUp:   src.PrivateGatewaysUp,
		PrivateGatewaysDown: src.PrivateGatewaysDown,
		Tags:                maps.Clone(src.Tags),
		DevAddrPrefixes:     src.DevAddrPrefixes,
	}
	return result, nil
}

func toTenantUsersListItem(src *apiv4.TenantUserListItem) (tenant.ListUserItem, error) {
	if src == nil {
		return tenant.ListUserItem{}, errors.New("tenant list user item is nil")
	}
	return tenant.ListUserItem{
		ID:             src.UserId,
		CreatedAt:      shared.TsProtobufToPtrTime(src.GetCreatedAt()),
		UpdatedAt:      shared.TsProtobufToPtrTime(src.GetUpdatedAt()),
		Email:          src.Email,
		IsAdmin:        src.IsAdmin,
		IsGatewayAdmin: src.IsGatewayAdmin,
		IsDeviceAdmin:  src.IsDeviceAdmin,
	}, nil
}

func toTenantUser(src *apiv4.TenantUser, createdAt *timestamppb.Timestamp, updatedAt *timestamppb.Timestamp) (*tenant.User, error) {
	if src == nil {
		return nil, errors.New("tenant response is nil")
	}
	result := &tenant.User{
		ID:             src.UserId,
		CreatedAt:      shared.TsProtobufToPtrTime(createdAt),
		UpdatedAt:      shared.TsProtobufToPtrTime(updatedAt),
		TenantID:       src.TenantId,
		IsAdmin:        src.IsAdmin,
		IsDeviceAdmin:  src.IsDeviceAdmin,
		IsGatewayAdmin: src.IsGatewayAdmin,
		Email:          src.Email,
		DeviceProfiles: nil,
		Applications:   nil,
	}
	return result, nil
}
