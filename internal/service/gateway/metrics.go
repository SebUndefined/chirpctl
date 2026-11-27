package gateway

import (
	"time"

	"github.com/chirpstack/chirpstack/api/go/v4/common"
	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type GetGatewayMetricsParams struct {
	GatewayID   string
	Start       *time.Time
	End         *time.Time
	Aggregation shared.Aggregation
}

type GetGatewayMetricsResult struct {
	RxPackets *shared.Metric `protobuf:"bytes,1,opt,name=rx_packets,json=rxPackets,proto3" json:"rx_packets,omitempty"`
	// TX packets.
	TxPackets *common.Metric `protobuf:"bytes,2,opt,name=tx_packets,json=txPackets,proto3" json:"tx_packets,omitempty"`
	// TX packets / frequency.
	TxPacketsPerFreq *common.Metric `protobuf:"bytes,3,opt,name=tx_packets_per_freq,json=txPacketsPerFreq,proto3" json:"tx_packets_per_freq,omitempty"`
	// RX packets / frequency.
	RxPacketsPerFreq *common.Metric `protobuf:"bytes,4,opt,name=rx_packets_per_freq,json=rxPacketsPerFreq,proto3" json:"rx_packets_per_freq,omitempty"`
	// TX packets / DR.
	TxPacketsPerDr *common.Metric `protobuf:"bytes,5,opt,name=tx_packets_per_dr,json=txPacketsPerDr,proto3" json:"tx_packets_per_dr,omitempty"`
	// RX packets / DR.
	RxPacketsPerDr *common.Metric `protobuf:"bytes,6,opt,name=rx_packets_per_dr,json=rxPacketsPerDr,proto3" json:"rx_packets_per_dr,omitempty"`
	// TX packets per status.
	TxPacketsPerStatus *common.Metric
}
