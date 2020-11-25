package generated

/*
#cgo LDFLAGS: -lperfstat
#include <libperfstat.h>

*/
import "C"

import (

)

type NetinterfaceStat struct {

	Type float64
	Mtu float64
	Ipackets float64
	Ibytes float64
	Ierrors float64
	Opackets float64
	Obytes float64
	Oerrors float64
	Collisions float64
	Bitrate float64
	Xmitdrops float64
	IfIqdrops float64
	IfArpdrops float64
}

type NetinterfaceStats map[string]NetinterfaceStat

func CollectNetinterfaces() NetinterfaceStats {
	num := C.perfstat_netinterface(nil, nil, C.sizeof_perfstat_netinterface_t, 0)

	initStruct := C.perfstat_id_t{}
	initStruct.name[0]=0

	netinterfacestats := make([]C.perfstat_netinterface_t, num)

	_ = C.perfstat_netinterface(&initStruct, &netinterfacestats[0], C.sizeof_perfstat_netinterface_t, num)

	output := make(NetinterfaceStats)

	for _, data := range netinterfacestats {

		name := C.GoString(&data.name[0])


		output[name] = NetinterfaceStat{
			Type: float64(data._type),
			Mtu: float64(data.mtu),
			Ipackets: float64(data.ipackets),
			Ibytes: float64(data.ibytes),
			Ierrors: float64(data.ierrors),
			Opackets: float64(data.opackets),
			Obytes: float64(data.obytes),
			Oerrors: float64(data.oerrors),
			Collisions: float64(data.collisions),
			Bitrate: float64(data.bitrate),
			Xmitdrops: float64(data.xmitdrops),
			IfIqdrops: float64(data.if_iqdrops),
			IfArpdrops: float64(data.if_arpdrops),
		}
	}

	return output
}
