package deviceprofile

import (
	"fmt"
	"strings"

	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/deviceprofile"
)

// buildGatewayListTable builds a table from gateway list items.
func buildDeviceProfileListTable(items []deviceprofile.ListItem, styled bool, useColor bool) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{
			"ID", "NAME", "REGION", "MAC VERSION",
			"REG PARAMS", "OTAA", "CLASS B", "CLASS C",
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
			printing.FormatTime(item.CreatedAt, ""),
			printing.FormatTime(item.UpdatedAt, ""),
		})
	}
	return td
}

// buildDeviceProfileDetailTable builds the human-readable gateway details.
func buildDeviceProfileDetailTable(p *deviceprofile.DeviceProfile, styled, useColor bool) printing.TableData {
	td := printing.TableData{
		Headers: printing.Row{"Field", "Value"},
		Rows:    []printing.Row{},
	}
	appendGeneralRows(&td, p, styled, useColor)
	appendLoRaWANRows(&td, p, styled, useColor)
	appendCapabilitiesRows(&td, p, styled, useColor)
	appendActivationRows(&td, p, styled, useColor)
	appendClassBRows(&td, p, styled, useColor)
	appendClassCRows(&td, p, styled, useColor)
	appendRelayRows(&td, p, styled, useColor)
	appendMeasurementRows(&td, p, styled, useColor)
	return td
}

// appendGeneralRows appends general device profile information.
func appendGeneralRows(td *printing.TableData, p *deviceprofile.DeviceProfile, styled, useColor bool) {
	td.Rows = append(td.Rows,
		printing.Row{"ID", p.ID},
		printing.Row{"Name", p.Name},
		printing.Row{"Description", strings.Join(strings.Fields(p.Description), " ")},
		printing.Row{"Tenant ID", p.TenantID},
		printing.Row{"Created At", printing.FormatTime(p.CreatedAt, "")},
		printing.Row{"Updated At", printing.FormatTime(p.UpdatedAt, "")},
	)
}

func appendLoRaWANRows(td *printing.TableData, p *deviceprofile.DeviceProfile, styled, useColor bool) {
	td.Rows = append(td.Rows,
		printing.Row{"Region", p.Region},
		printing.Row{"Region Config", p.RegionConfigID},
		printing.Row{"MAC Version", p.MacVersion},
		printing.Row{"Revision", p.RegParamsRevision},
		printing.Row{"ADR Algorithm", p.AdrAlgorithmID},
	)
}

func appendCapabilitiesRows(td *printing.TableData, p *deviceprofile.DeviceProfile, styled, useColor bool) {
	td.Rows = append(td.Rows,
		printing.Row{"OTAA", formatCapability(p.SupportsOtaa, styled, useColor)},
		printing.Row{"Class B", formatCapability(p.SupportsClassB, styled, useColor)},
		printing.Row{"Class C", formatCapability(p.SupportsClassC, styled, useColor)},
		printing.Row{"Roaming", formatCapability(p.AllowRoaming, styled, useColor)},
	)
}
func appendActivationRows(td *printing.TableData, p *deviceprofile.DeviceProfile, styled, useColor bool) {
	td.Rows = append(td.Rows,
		printing.Row{"Flush Queue On Activate", formatCapability(p.FlushQueueOnActivate, styled, useColor)},
		printing.Row{"Uplink Interval (s)", p.UplinkInterval},
		printing.Row{"Device Status Interval (s)", p.DeviceStatusReqInterval},
		printing.Row{"RX1 Delay", p.Rx1Delay},
		printing.Row{"ABP RX1 Delay", p.AbpRx1Delay},
		printing.Row{"ABP RX1 DR Offset", p.AbpRx1DrOffset},
		printing.Row{"ABP RX2 DR", p.AbpRx2Dr},
		printing.Row{"ABP RX2 Frequency (Hz)", p.AbpRx2Freq},
	)
}
func appendClassBRows(td *printing.TableData, p *deviceprofile.DeviceProfile, styled, useColor bool) {
	td.Rows = append(td.Rows,
		printing.Row{"Class B Timeout", p.ClassBTimeout},
		printing.Row{"Ping Slot Periodicity", p.ClassBPingSlotPeriodicity},
		printing.Row{"Ping Slot DR", p.ClassBPingSlotDr},
		printing.Row{"Ping Slot Frequency (Hz)", p.ClassBPingSlotFreq},
		printing.Row{"Downlink Only", formatCapability(p.ClassBDownlinkOnly, styled, useColor)},
	)
}
func appendClassCRows(td *printing.TableData, p *deviceprofile.DeviceProfile, styled, useColor bool) {
	td.Rows = append(td.Rows,
		printing.Row{"Class C Timeout", p.ClassCTimeout},
	)
}
func appendRelayRows(td *printing.TableData, p *deviceprofile.DeviceProfile, styled, useColor bool) {
	td.Rows = append(td.Rows,
		printing.Row{"Is Relay", formatCapability(p.IsRelay, styled, useColor)},
		printing.Row{"Relay Enabled", formatCapability(p.RelayEnabled, styled, useColor)},
		printing.Row{"Is Relay End-Device", formatCapability(p.IsRelayEd, styled, useColor)},
		printing.Row{"Relay ED Only", formatCapability(p.RelayEdRelayOnly, styled, useColor)},
		printing.Row{"CAD Periodicity", p.RelayCadPeriodicity},
		printing.Row{"Default Channel Index", p.RelayDefaultChannelIndex},
		printing.Row{"Second Channel Frequency (Hz)", p.RelaySecondChannelFreq},
		printing.Row{"Second Channel DR", p.RelaySecondChannelDr},
		printing.Row{"Second Channel ACK Offset", p.RelaySecondChannelAckOffset},
		printing.Row{"ED Activation Mode", p.RelayEdActivationMode},
		printing.Row{"ED Smart Enable Level", p.RelayEdSmartEnableLevel},
		printing.Row{"ED Backoff", p.RelayEdBackOff},
		printing.Row{"ED Uplink Bucket Size", p.RelayEdUplinkLimitBucketSize},
		printing.Row{"ED Uplink Reload Rate", p.RelayEdUplinkLimitReloadRate},
		printing.Row{"Join Request Reload Rate", p.RelayJoinReqLimitReloadRate},
		printing.Row{"Join Request Bucket Size", p.RelayJoinReqLimitBucketSize},
		printing.Row{"Notify Reload Rate", p.RelayNotifyLimitReloadRate},
		printing.Row{"Notify Bucket Size", p.RelayNotifyLimitBucketSize},
		printing.Row{"Global Uplink Reload Rate", p.RelayGlobalUplinkLimitReloadRate},
		printing.Row{"Global Uplink Bucket Size", p.RelayGlobalUplinkLimitBucketSize},
		printing.Row{"Overall Reload Rate", p.RelayOverallLimitReloadRate},
		printing.Row{"Overall Bucket Size", p.RelayOverallLimitBucketSize},
	)
}
func appendMeasurementRows(td *printing.TableData, p *deviceprofile.DeviceProfile, styled, useColor bool) {
	td.Rows = append(td.Rows,
		printing.Row{"Payload Codec Runtime", p.PayloadCodecRuntime},
		printing.Row{"Payload Codec Script",
			formatCapability(strings.TrimSpace(p.PayloadCodecScript) != "", styled, useColor)},
		printing.Row{"Auto Detect Measurements",
			formatCapability(p.AutoDetectMeasurements, styled, useColor)},
		printing.Row{"Measurements Count", len(p.Measurements)},
		printing.Row{"Device ID", p.DeviceID},
		printing.Row{"Firmware Version", p.FirmwareVersion},
		printing.Row{"Supported Uplink Data Rates", fmt.Sprint(p.SupportedUplinkDataRates)},
		printing.Row{"Application Layer Parameters",
			formatCapability(p.AppLayerParams != nil, styled, useColor)},
	)
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
