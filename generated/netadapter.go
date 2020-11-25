package generated

/*
#cgo LDFLAGS: -lperfstat
#include <libperfstat.h>

*/
import "C"

import (

)

type NetadapterStat struct {

	TxPackets float64
	TxBytes float64
	TxInterrupts float64
	TxErrors float64
	TxPacketsDropped float64
	TxQueueSize float64
	TxQueueLen float64
	TxQueueOverflow float64
	TxBroadcastPackets float64
	TxMulticastPackets float64
	TxCarrierSense float64
	TxDMAUnderrun float64
	TxLostCTSErrors float64
	TxMaxCollisionErrors float64
	TxLateCollisionErrors float64
	TxDeferred float64
	TxTimeoutErrors float64
	TxSingleCollisionCount float64
	TxMultipleCollisionCount float64
	RxPackets float64
	RxBytes float64
	RxInterrupts float64
	RxErrors float64
	RxPacketsDropped float64
	RxBadPackets float64
	RxMulticastPackets float64
	RxBroadcastPackets float64
	RxCRCErrors float64
	RxDMAOverrun float64
	RxAlignmentErrors float64
	RxNoresourceErrors float64
	RxCollisionErrors float64
	RxPacketTooshortErrors float64
	RxPacketToolongErrors float64
	RxPacketsDiscardedbyadapter float64
}

type NetadapterStats map[string]NetadapterStat

func CollectNetadapters() NetadapterStats {
	output := make(NetadapterStats)

num := C.perfstat_netadapter(nil, nil, C.sizeof_perfstat_netadapter_t, 0)
if num == 0 {
		return output
	}

	initStruct := C.perfstat_id_t{}
	initStruct.name[0]=0

	netadapterstats := make([]C.perfstat_netadapter_t, num)

	_ = C.perfstat_netadapter(&initStruct, &netadapterstats[0], C.sizeof_perfstat_netadapter_t, num)

	for _, data := range netadapterstats {

		name := C.GoString(&data.name[0])


		output[name] = NetadapterStat{
			TxPackets: float64(data.tx_packets),
			TxBytes: float64(data.tx_bytes),
			TxInterrupts: float64(data.tx_interrupts),
			TxErrors: float64(data.tx_errors),
			TxPacketsDropped: float64(data.tx_packets_dropped),
			TxQueueSize: float64(data.tx_queue_size),
			TxQueueLen: float64(data.tx_queue_len),
			TxQueueOverflow: float64(data.tx_queue_overflow),
			TxBroadcastPackets: float64(data.tx_broadcast_packets),
			TxMulticastPackets: float64(data.tx_multicast_packets),
			TxCarrierSense: float64(data.tx_carrier_sense),
			TxDMAUnderrun: float64(data.tx_DMA_underrun),
			TxLostCTSErrors: float64(data.tx_lost_CTS_errors),
			TxMaxCollisionErrors: float64(data.tx_max_collision_errors),
			TxLateCollisionErrors: float64(data.tx_late_collision_errors),
			TxDeferred: float64(data.tx_deferred),
			TxTimeoutErrors: float64(data.tx_timeout_errors),
			TxSingleCollisionCount: float64(data.tx_single_collision_count),
			TxMultipleCollisionCount: float64(data.tx_multiple_collision_count),
			RxPackets: float64(data.rx_packets),
			RxBytes: float64(data.rx_bytes),
			RxInterrupts: float64(data.rx_interrupts),
			RxErrors: float64(data.rx_errors),
			RxPacketsDropped: float64(data.rx_packets_dropped),
			RxBadPackets: float64(data.rx_bad_packets),
			RxMulticastPackets: float64(data.rx_multicast_packets),
			RxBroadcastPackets: float64(data.rx_broadcast_packets),
			RxCRCErrors: float64(data.rx_CRC_errors),
			RxDMAOverrun: float64(data.rx_DMA_overrun),
			RxAlignmentErrors: float64(data.rx_alignment_errors),
			RxNoresourceErrors: float64(data.rx_noresource_errors),
			RxCollisionErrors: float64(data.rx_collision_errors),
			RxPacketTooshortErrors: float64(data.rx_packet_tooshort_errors),
			RxPacketToolongErrors: float64(data.rx_packet_toolong_errors),
			RxPacketsDiscardedbyadapter: float64(data.rx_packets_discardedbyadapter),
		}
	}

	return output
}
