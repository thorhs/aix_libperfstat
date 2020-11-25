package generated

/*
#cgo LDFLAGS: -lperfstat
#include <libperfstat.h>

*/
import "C"

import (

)

type DiskpathStat struct {

	Xrate float64
	Xfers float64
	Rblks float64
	Wblks float64
	Time float64
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

type DiskpathStats map[string]DiskpathStat

func CollectDiskpaths() DiskpathStats {
	output := make(DiskpathStats)

num := C.perfstat_diskpath(nil, nil, C.sizeof_perfstat_diskpath_t, 0)
if num == 0 {
		return output
	}

	initStruct := C.perfstat_id_t{}
	initStruct.name[0]=0

	diskpathstats := make([]C.perfstat_diskpath_t, num)

	_ = C.perfstat_diskpath(&initStruct, &diskpathstats[0], C.sizeof_perfstat_diskpath_t, num)

	for _, data := range diskpathstats {

		name := C.GoString(&data.name[0])


		output[name] = DiskpathStat{
			Xrate: float64(data.xrate),
			Xfers: float64(data.xfers),
			Rblks: float64(data.rblks),
			Wblks: float64(data.wblks),
			Time: float64(data.time),
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
