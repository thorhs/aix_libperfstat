package generated

/*
#cgo LDFLAGS: -lperfstat
#include <libperfstat.h>

*/
import "C"

import (

	"fmt"

)

type MemoryPageStat struct {

	RealTotal float64
	RealFree float64
	RealPinned float64
	RealInuse float64
	Pgexct float64
	Pgins float64
	Pgouts float64
	Pgspins float64
	Pgspouts float64
	Scans float64
	Cycles float64
	Pgsteals float64
	Numperm float64
	Numpgsp float64
	RealSystem float64
	RealUser float64
	RealProcess float64
	VirtActive float64
	ComprsdTotal float64
	ComprsdWsegPgs float64
	Cpgins float64
	Cpgouts float64
	CpoolInuse float64
	UcpoolSize float64
	ComprsdWsegSize float64
	RealAvail float64
}

type MemoryPageStats map[string]MemoryPageStat

func CollectMemoryPages() MemoryPageStats {
	num := C.perfstat_memory_page(nil, nil, C.sizeof_perfstat_memory_page_t, 0)

	initStruct := C.perfstat_psize_t{}
	initStruct.psize=0

	memory_pagestats := make([]C.perfstat_memory_page_t, num)

	_ = C.perfstat_memory_page(&initStruct, &memory_pagestats[0], C.sizeof_perfstat_memory_page_t, num)

	output := make(MemoryPageStats)

	for _, data := range memory_pagestats {

		psize := fmt.Sprintf("%d", C.long(data.psize))


		output[psize] = MemoryPageStat{
			RealTotal: float64(data.real_total),
			RealFree: float64(data.real_free),
			RealPinned: float64(data.real_pinned),
			RealInuse: float64(data.real_inuse),
			Pgexct: float64(data.pgexct),
			Pgins: float64(data.pgins),
			Pgouts: float64(data.pgouts),
			Pgspins: float64(data.pgspins),
			Pgspouts: float64(data.pgspouts),
			Scans: float64(data.scans),
			Cycles: float64(data.cycles),
			Pgsteals: float64(data.pgsteals),
			Numperm: float64(data.numperm),
			Numpgsp: float64(data.numpgsp),
			RealSystem: float64(data.real_system),
			RealUser: float64(data.real_user),
			RealProcess: float64(data.real_process),
			VirtActive: float64(data.virt_active),
			ComprsdTotal: float64(data.comprsd_total),
			ComprsdWsegPgs: float64(data.comprsd_wseg_pgs),
			Cpgins: float64(data.cpgins),
			Cpgouts: float64(data.cpgouts),
			CpoolInuse: float64(data.cpool_inuse),
			UcpoolSize: float64(data.ucpool_size),
			ComprsdWsegSize: float64(data.comprsd_wseg_size),
			RealAvail: float64(data.real_avail),
		}
	}

	return output
}
