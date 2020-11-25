package generated

/*
#cgo LDFLAGS: -lperfstat
#include <libperfstat.h>

*/
import "C"

import (

)

type DiskadapterStat struct {

	Number float64
	Size float64
	Free float64
	Xrate float64
	Xfers float64
	Rblks float64
	Wblks float64
	Time float64
	AdapterType float64
	 DkBsize float64
	 DkRxfers float64
	 DkRserv float64
	 DkWserv float64
	 MinRserv float64
	 MaxRserv float64
	 MinWserv float64
	 MaxWserv float64
	 WqDepth float64
	 WqSampled float64
	 WqTime float64
	 WqMinTime float64
	 WqMaxTime float64
	 QFull float64
	 QSampled float64
}

type DiskadapterStats map[string]DiskadapterStat

func CollectDiskadapters() DiskadapterStats {
	output := make(DiskadapterStats)

num := C.perfstat_diskadapter(nil, nil, C.sizeof_perfstat_diskadapter_t, 0)
if num == 0 {
		return output
	}

	initStruct := C.perfstat_id_t{}
	initStruct.name[0]=0

	diskadapterstats := make([]C.perfstat_diskadapter_t, num)

	_ = C.perfstat_diskadapter(&initStruct, &diskadapterstats[0], C.sizeof_perfstat_diskadapter_t, num)

	for _, data := range diskadapterstats {

		name := C.GoString(&data.name[0])


		output[name] = DiskadapterStat{
			Number: float64(data.number),
			Size: float64(data.size),
			Free: float64(data.free),
			Xrate: float64(data.xrate),
			Xfers: float64(data.xfers),
			Rblks: float64(data.rblks),
			Wblks: float64(data.wblks),
			Time: float64(data.time),
			AdapterType: float64(data.adapter_type),
			 DkBsize: float64(data. dk_bsize),
			 DkRxfers: float64(data. dk_rxfers),
			 DkRserv: float64(data. dk_rserv),
			 DkWserv: float64(data. dk_wserv),
			 MinRserv: float64(data. min_rserv),
			 MaxRserv: float64(data. max_rserv),
			 MinWserv: float64(data. min_wserv),
			 MaxWserv: float64(data. max_wserv),
			 WqDepth: float64(data. wq_depth),
			 WqSampled: float64(data. wq_sampled),
			 WqTime: float64(data. wq_time),
			 WqMinTime: float64(data. wq_min_time),
			 WqMaxTime: float64(data. wq_max_time),
			 QFull: float64(data. q_full),
			 QSampled: float64(data. q_sampled),
		}
	}

	return output
}
