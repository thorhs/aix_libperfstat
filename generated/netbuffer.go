package generated

/*
#cgo LDFLAGS: -lperfstat
#include <libperfstat.h>

*/
import "C"

import (

)

type NetbufferStat struct {

	Inuse float64
	Calls float64
	Delayed float64
	Free float64
	Failed float64
	Highwatermark float64
	Freed float64
}

type NetbufferStats map[string]NetbufferStat

func CollectNetbuffers() NetbufferStats {
	output := make(NetbufferStats)

num := C.perfstat_netbuffer(nil, nil, C.sizeof_perfstat_netbuffer_t, 0)
if num == 0 {
		return output
	}

	initStruct := C.perfstat_id_t{}
	initStruct.name[0]=0

	netbufferstats := make([]C.perfstat_netbuffer_t, num)

	_ = C.perfstat_netbuffer(&initStruct, &netbufferstats[0], C.sizeof_perfstat_netbuffer_t, num)

	for _, data := range netbufferstats {

		name := C.GoString(&data.name[0])


		output[name] = NetbufferStat{
			Inuse: float64(data.inuse),
			Calls: float64(data.calls),
			Delayed: float64(data.delayed),
			Free: float64(data.free),
			Failed: float64(data.failed),
			Highwatermark: float64(data.highwatermark),
			Freed: float64(data.freed),
		}
	}

	return output
}
