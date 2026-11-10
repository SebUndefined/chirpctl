package device

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/device"
)

// buildDeviceListTable builds a table from device list items.
func buildDeviceListTable(items []device.ListItem, styled bool, useColor bool) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{
			"DevEUI", "Name", "Profile", "Last Seen", "Battery", "Margin",
		},
		Rows: make([]printing.Row, 0, len(items)),
	}

	for _, item := range items {
		battery := "-"
		margin := "-"

		if item.DeviceStatus != nil {
			status := item.DeviceStatus

			margin = fmt.Sprintf("%d dB", status.Margin)

			if status.ExternalPowerSource {
				battery = "EXT"
			} else {
				battery = fmt.Sprintf("%.0f%%", status.BatteryLevel)
			}

			if styled {
				healthy := status.ExternalPowerSource ||
					status.BatteryLevel >= 20

				battery = statusLabel(battery, healthy, styled, useColor)
			}
		}

		td.Rows = append(td.Rows, printing.Row{
			item.DevEUI,
			item.Name,
			item.DeviceProfileName,
			printing.FormatTime(item.LastSeenAt, ""),
			battery,
			margin,
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

// buildDeviceDetailTable builds the human-readable device details.
func buildDeviceDetailTable(d *device.Device, showVariables bool, styled bool, useColor bool) printing.TableData {
	if showVariables {
		td := printing.TableData{
			Headers: printing.Row{
				"Name", "Value",
			},
		}
		appendStringMapRows(&td, "Variables", d.Variables)
		return td
	}
	disabled := "No"
	if d.IsDisabled {
		disabled = "Yes"
	}

	fcntCheck := "Enabled"
	if d.SkipFcntCheck {
		fcntCheck = "Disabled (unsafe)"
	}

	disabled = statusLabel(disabled, !d.IsDisabled, styled, useColor)
	fcntCheck = statusLabel(fcntCheck, !d.SkipFcntCheck, styled, useColor)

	td := printing.TableData{
		Headers: printing.Row{"Field", "Value"},
		Rows: []printing.Row{
			{"Dev EUI", d.DevEUI},
			{"Name", d.Name},
			{"Description", strings.Join(strings.Fields(d.Description), " ")},
			{"Application ID", d.ApplicationID},
			{"Device Profile ID", d.DeviceProfileID},
			{"Join EUI", d.JoinEUI},
			{"Disabled", disabled},
			{"Frame Counter Check", fcntCheck},
			{"Created At", printing.FormatTime(d.CreatedAt, "")},
			{"Updated At", printing.FormatTime(d.UpdatedAt, "")},
			{"Last Seen", printing.FormatTime(d.LastSeenAt, "")},
		},
	}

	appendStringMapRows(&td, "Tags", d.Tags)

	td.Rows = append(td.Rows, printing.Row{
		"Variables", fmt.Sprintf("%d entries", len(d.Variables)),
	})

	return td
}

// statusLabel adds a visual indicator in styled table mode.
func statusLabel(label string, healthy, styled, useColor bool) string {
	if !styled {
		return label
	}

	if !useColor {
		return "● " + label
	}

	if healthy {
		return printing.StatusDot(printing.DotGreen) + " " + label
	}

	return printing.StatusDot(printing.DotRed) + " " + label
}
