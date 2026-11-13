package tenant

import (
	"strings"

	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/tenant"
)

// buildTenantListTable builds a table from gateway list items.
func buildTenantListTable(items []tenant.ListItem, styled bool, useColor bool) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{
			"ID", "NAME", "CAN GW", "PV GW UP", "PV GW DOWN",
			"MAX GW", "MAX DEV",
			"CREATED AT", "UPDATED AT",
		},
		Rows: make([]printing.Row, 0, len(items)),
	}

	for _, item := range items {
		td.Rows = append(td.Rows, printing.Row{
			item.ID,
			item.Name,
			printing.FormatBool(printing.FormatBoolInput{
				Value:       item.CanHaveGateways,
				Styled:      styled,
				Colored:     useColor,
				TextTrue:    "YES",
				TextFalse:   "NO",
				ColorTrue:   printing.ColorGreen,
				ColorFalse:  printing.ColorGray,
				SymbolTrue:  "✓",
				SymbolFalse: "—",
			}),
			printing.FormatBool(printing.FormatBoolInput{
				Value:       item.PrivateGatewaysUp,
				Styled:      styled,
				Colored:     useColor,
				TextTrue:    "YES",
				TextFalse:   "NO",
				ColorTrue:   printing.ColorGreen,
				ColorFalse:  printing.ColorGray,
				SymbolTrue:  "✓",
				SymbolFalse: "—",
			}),
			printing.FormatBool(printing.FormatBoolInput{
				Value:       item.PrivateGatewaysDown,
				Styled:      styled,
				Colored:     useColor,
				TextTrue:    "YES",
				TextFalse:   "NO",
				ColorTrue:   printing.ColorGreen,
				ColorFalse:  printing.ColorGray,
				SymbolTrue:  "✓",
				SymbolFalse: "—",
			}),
			item.MaxGatewayCount,
			item.MaxDeviceCount,
			printing.FormatTime(item.CreatedAt, ""),
			printing.FormatTime(item.UpdatedAt, ""),
		})
	}
	return td
}

// buildTenantDetailTable builds the human-readable device details.
func buildTenantDetailTable(t *tenant.Tenant, styled bool, useColor bool) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{"Field", "Value"},
		Rows: []printing.Row{
			{"ID", t.ID},
			{"Name", t.Name},
			{"Description", strings.Join(strings.Fields(t.Description), " ")},
			{"Can have gateways", printing.FormatBool(printing.FormatBoolInput{
				Value:       t.CanHaveGateways,
				Styled:      styled,
				Colored:     useColor,
				TextTrue:    "YES",
				TextFalse:   "NO",
				ColorTrue:   printing.ColorGreen,
				ColorFalse:  printing.ColorGray,
				SymbolTrue:  "✓",
				SymbolFalse: "—",
			})},
			{"Private gateways up", printing.FormatBool(printing.FormatBoolInput{
				Value:       t.PrivateGatewaysUp,
				Styled:      styled,
				Colored:     useColor,
				TextTrue:    "YES",
				TextFalse:   "NO",
				ColorTrue:   printing.ColorGreen,
				ColorFalse:  printing.ColorGray,
				SymbolTrue:  "✓",
				SymbolFalse: "—",
			})},
			{"Private gateways down", printing.FormatBool(printing.FormatBoolInput{
				Value:       t.PrivateGatewaysDown,
				Styled:      styled,
				Colored:     useColor,
				TextTrue:    "YES",
				TextFalse:   "NO",
				ColorTrue:   printing.ColorGreen,
				ColorFalse:  printing.ColorGray,
				SymbolTrue:  "✓",
				SymbolFalse: "—",
			})},
			{"Max gateways", t.MaxGatewayCount},
			{"Max devices", t.MaxDeviceCount},
			{"Address prefix", strings.Join(t.DevAddrPrefixes, ", ")},
		},
	}

	printing.AppendStringMapRows(&td, "Tags", t.Tags)

	return td
}
