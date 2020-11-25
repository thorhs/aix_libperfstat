package generated

/*
#cgo LDFLAGS: -lperfstat
#include <libperfstat.h>

*/
import "C"

import (

)

type DiskStat struct {
	Vgname string

	Size float64
	Free float64
	Bsize float64
	Xfers float64
	Wblks float64
	Rblks float64
	Qdepth float64
	Time float64
	PathsCount float64
	QFull float64
	Rserv float64
	Rtimeout float64
	Rfailed float64
	MinRserv float64
	MaxRserv float64
	Wserv float64
	Wtimeout float64
	Wfailed float64
	MinWserv float64
	MaxWserv float64
	WqDepth float64
	WqSampled float64
	WqTime float64
	WqMinTime float64
	WqMaxTime float64
	QSampled float64
}

type DiskStats map[string]DiskStat

func CollectDisks() DiskStats {
	output := make(DiskStats)

num := C.perfstat_disk(nil, nil, C.sizeof_perfstat_disk_t, 0)
if num == 0 {
		return output
	}

	initStruct := C.perfstat_id_t{}
	initStruct.name[0]=0

	diskstats := make([]C.perfstat_disk_t, num)

	_ = C.perfstat_disk(&initStruct, &diskstats[0], C.sizeof_perfstat_disk_t, num)

	for _, data := range diskstats {

		name := C.GoString(&data.name[0])


		output[name] = DiskStat{
			Vgname: C.GoString(&data.vgname[0]),
			Size: float64(data.size),
			Free: float64(data.free),
			Bsize: float64(data.bsize),
			Xfers: float64(data.xfers),
			Wblks: float64(data.wblks),
			Rblks: float64(data.rblks),
			Qdepth: float64(data.qdepth),
			Time: float64(data.time),
			PathsCount: float64(data.paths_count),
			QFull: float64(data.q_full),
			Rserv: float64(data.rserv),
			Rtimeout: float64(data.rtimeout),
			Rfailed: float64(data.rfailed),
			MinRserv: float64(data.min_rserv),
			MaxRserv: float64(data.max_rserv),
			Wserv: float64(data.wserv),
			Wtimeout: float64(data.wtimeout),
			Wfailed: float64(data.wfailed),
			MinWserv: float64(data.min_wserv),
			MaxWserv: float64(data.max_wserv),
			WqDepth: float64(data.wq_depth),
			WqSampled: float64(data.wq_sampled),
			WqTime: float64(data.wq_time),
			WqMinTime: float64(data.wq_min_time),
			WqMaxTime: float64(data.wq_max_time),
			QSampled: float64(data.q_sampled),
		}
	}

	return output
}
