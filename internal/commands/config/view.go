package config

import (
	"maps"
	"slices"

	"github.com/sebundefined/chirpctl/internal/config"
	"github.com/sebundefined/chirpctl/internal/printing"
)

type ContextView struct {
	Name       string `json:"name" yaml:"name"`
	Current    bool   `json:"current" yaml:"current"`
	APIVersion string `json:"apiVersion" yaml:"apiVersion"`
	Address    string `json:"address" yaml:"address"`
	TenantID   string `json:"tenantID" yaml:"tenantID"`
	TLS        bool   `json:"tls" yaml:"tls"`
	HasToken   bool   `json:"hasToken" yaml:"hasToken"`
}

func buildCurrentContextTable(view ContextView) printing.TableData {
	tls := "Secure"
	if !view.TLS {
		tls = "Plaintext"
	}

	token := "Missing"
	if view.HasToken {
		token = "Configured"
	}
	td := printing.TableData{
		Headers: printing.Row{"Field", "Value"},
		Rows: []printing.Row{
			{"Name", view.Name},
			{"API Version", view.APIVersion},
			{"Address", view.Address},
			{"Tenant ID", view.TenantID},
			{"TLS", tls},
			{"Token", token},
		},
	}
	return td
}

func buildContextViews(cfgFile *config.File) []ContextView {
	if cfgFile == nil {
		return []ContextView{}
	}
	names := slices.Sorted(maps.Keys(cfgFile.Contexts))
	views := make([]ContextView, 0, len(names))
	for _, name := range names {
		ctx := cfgFile.Contexts[name]

		view := ContextView{
			Name:       name,
			Current:    name == cfgFile.DefaultContext,
			APIVersion: ctx.ApiVersion,
			Address:    ctx.Addr,
			TenantID:   ctx.TenantID,
			TLS:        !ctx.Insecure,
			HasToken:   ctx.Token != "",
		}

		views = append(views, view)
	}
	return views
}

func buildContextsTable(views []ContextView, styled bool, useColor bool) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{
			"Current", "Name", "API version", "Address",
			"TenantID", "TLS", "HasToken",
		},
		Rows: make([]printing.Row, 0, len(views)),
	}
	for _, item := range views {
		current := ""
		if item.Current {
			current = "*"
			if styled {
				current = "●"
				if useColor {
					current = printing.StatusDot(printing.DotGreen)
				}
			}
		}
		tls := "Secure"
		if !item.TLS {
			tls = "Plaintext"
		}

		token := "Missing"
		if item.HasToken {
			token = "Configured"
		}

		tenantID := item.TenantID
		if tenantID == "" {
			tenantID = "-"
		}
		td.Rows = append(td.Rows, printing.Row{
			current,
			item.Name,
			item.APIVersion,
			item.Address,
			tenantID,
			tls,
			token,
		})
	}

	return td
}
