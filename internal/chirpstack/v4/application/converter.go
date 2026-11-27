package application

import (
	"errors"
	"time"

	"github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/application"
	"github.com/sebundefined/chirpctl/internal/service/shared"
)

func toApplicationListItem(src *api.ApplicationListItem) (application.ListItem, error) {
	if src == nil {
		return application.ListItem{}, errors.New("nil application listItem")
	}
	return application.ListItem{
		ID:          src.Id,
		CreatedAt:   shared.TsProtobufToPtrTime(src.GetCreatedAt()),
		UpdatedAt:   shared.TsProtobufToPtrTime(src.GetUpdatedAt()),
		Name:        src.GetName(),
		Description: src.GetDescription(),
	}, nil
}

func toApplication(src *api.GetApplicationResponse) (application.Application, error) {
	if src == nil {
		return application.Application{}, errors.New("nil application")
	}
	return application.Application{
		ID:          src.Application.Id,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		Name:        "",
		Description: "",
		TenantID:    "",
		Tags:        nil,
	}, nil
}
