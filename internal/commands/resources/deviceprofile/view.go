package deviceprofile

import (
	"fmt"

	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/deviceprofile"
)

// buildGatewayListTable builds a table from gateway list items.
func buildDeviceProfileListTable(items []deviceprofile.ListItem, styled bool, useColor bool) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{
			"ID", "NAME", "REGION", "MAC VERSION",
			"REG PARAMS", "OTAA", "CLASS B", "CLASS C",
			"VENDOR ID", "VENDOR NAME",
			"DEVICE ID", "DEVICE NAME", "FIRMWARE",
			"CREATED AT", "UPDATED AT",
		},
		Rows: make([]printing.Row, 0, len(items)),
	}

	for _, item := range items {
		td.Rows = append(td.Rows, printing.Row{
			item.ID,
			item.Name,
			fmt.Sprint(item.Region),
			fmt.Sprint(item.MacVersion),
			fmt.Sprint(item.RegParamsRevision),
			formatCapability(item.SupportsOtaa, styled, useColor),
			formatCapability(item.SupportsClassB, styled, useColor),
			formatCapability(item.SupportsClassC, styled, useColor),
			item.VendorID,
			item.VendorName,
			item.DeviceID,
			item.DeviceName,
			item.FirmwareVersion,
			printing.FormatTime(item.CreatedAt, ""),
			printing.FormatTime(item.UpdatedAt, ""),
		})
	}
	return td
}

func formatCapability(enabled, styled, useColor bool) string {
	if !styled {
		if enabled {
			return "YES"
		}
		return "NO"
	}

	if enabled {
		symbol := "✓"
		if useColor {
			symbol = printing.StatusCheck(printing.ColorGreen)
		}
		return symbol + " YES"
	}

	symbol := "—"
	if useColor {
		symbol = printing.StatusSymbol(symbol, printing.ColorGray)
	}
	return symbol + " NO"
}
