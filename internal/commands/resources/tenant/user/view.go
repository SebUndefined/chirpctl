package user

import (
	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/tenant"
)

// buildTenantUsersListTable builds a table from gateway list items.
func buildTenantUsersListTable(items []tenant.ListUserItem, styled bool, useColor bool) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{
			"ID", "EMAIL", "Is Admin", "Is GW admin", "Is DEV admin",
			"CREATED AT", "UPDATED AT",
		},
		Rows: make([]printing.Row, 0, len(items)),
	}

	for _, item := range items {
		td.Rows = append(td.Rows, printing.Row{
			item.ID,
			item.Email,
			printing.FormatBool(printing.FormatBoolInput{
				Value:       item.IsAdmin,
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
				Value:       item.IsGatewayAdmin,
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
				Value:       item.IsDeviceAdmin,
				Styled:      styled,
				Colored:     useColor,
				TextTrue:    "YES",
				TextFalse:   "NO",
				ColorTrue:   printing.ColorGreen,
				ColorFalse:  printing.ColorGray,
				SymbolTrue:  "✓",
				SymbolFalse: "—",
			}),
			printing.FormatTime(item.CreatedAt, ""),
			printing.FormatTime(item.UpdatedAt, ""),
		})
	}
	return td
}

// buildTenantDetailTable builds the human-readable device details.
func buildTenantUserDetailTable(t *tenant.User, styled bool, useColor bool) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{"Field", "Value"},
		Rows: []printing.Row{
			{"ID", t.ID},
			{"Email", t.Email},
			{"Admin", printing.FormatBool(printing.FormatBoolInput{
				Value:       t.IsAdmin,
				Styled:      styled,
				Colored:     useColor,
				TextTrue:    "YES",
				TextFalse:   "NO",
				ColorTrue:   printing.ColorGreen,
				ColorFalse:  printing.ColorGray,
				SymbolTrue:  "✓",
				SymbolFalse: "—",
			})},
			{"Device admin", printing.FormatBool(printing.FormatBoolInput{
				Value:       t.IsDeviceAdmin,
				Styled:      styled,
				Colored:     useColor,
				TextTrue:    "YES",
				TextFalse:   "NO",
				ColorTrue:   printing.ColorGreen,
				ColorFalse:  printing.ColorGray,
				SymbolTrue:  "✓",
				SymbolFalse: "—",
			})},
			{"Gateway admin", printing.FormatBool(printing.FormatBoolInput{
				Value:       t.IsGatewayAdmin,
				Styled:      styled,
				Colored:     useColor,
				TextTrue:    "YES",
				TextFalse:   "NO",
				ColorTrue:   printing.ColorGreen,
				ColorFalse:  printing.ColorGray,
				SymbolTrue:  "✓",
				SymbolFalse: "—",
			})},
		},
	}

	return td
}
