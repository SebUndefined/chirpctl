package gateway

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
)

// buildGatewayListTable builds a table from gateway list items.
func buildGatewayListTable(items []gateway.ListItem) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{
			"Gateway ID", "Name", "State",
			"Created At", "Updated At", "Last Seen",
			"Tenant ID",
		},
		Rows: make([]printing.Row, 0, len(items)),
	}

	for _, item := range items {
		// Reuse the StateOnline / StateOffline / StateNeverSeen
		// switch from your existing listGateways implementation.
		state := "UNKNOWN"
		dot := printing.StatusDot(printing.DotGray)

		switch item.State {
		case gateway.StateOnline:
			state = "ONLINE"
			dot = printing.StatusDot(printing.DotGreen)

		case gateway.StateOffline:
			state = "OFFLINE"
			dot = printing.StatusDot(printing.DotRed)

		case gateway.StateNeverSeen:
			state = "NEVER_SEEN"
		}

		td.Rows = append(td.Rows, printing.Row{
			item.GatewayID,
			item.Name,
			fmt.Sprintf("%s %s", dot, state),
			printing.FormatTime(item.CreatedAt, ""),
			printing.FormatTime(item.UpdatedAt, ""),
			printing.FormatTime(item.LastSeenAt, ""),
			item.TenantID,
		})
	}

	return td
}

// appendStringMapRows appends map entries in alphabetical order.
func appendStringMapRows(td *printing.TableData, prefix string, values map[string]string) {
	if len(values) == 0 {
		td.Rows = append(td.Rows, printing.Row{prefix, "-"})
		return
	}

	for _, key := range slices.Sorted(maps.Keys(values)) {
		td.Rows = append(td.Rows, printing.Row{
			prefix + "." + key,
			values[key],
		})
	}
}

// buildGatewayDetailTable builds the human-readable gateway details.
func buildGatewayDetailTable(g *gateway.Gateway) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{"Field", "Value"},
		Rows: []printing.Row{
			{"Gateway ID", g.GatewayID},
			{"Name", g.Name},
			{"Description", strings.Join(strings.Fields(g.Description), " ")},
			{"Tenant ID", g.TenantID},
			{"Location", g.Location},
			{"Created At", printing.FormatTime(g.CreatedAt, "")},
			{"Updated At", printing.FormatTime(g.UpdatedAt, "")},
			{"Last Seen", printing.FormatTime(g.LastSeenAt, "")},
			{"Stats Interval (s)", g.StatsInterval},
			{"Downlink Priority", g.DownlinkPriority},
		},
	}

	appendStringMapRows(&td, "Tags", g.Tags)
	appendStringMapRows(&td, "Metadata", g.Metadata)

	return td
}
