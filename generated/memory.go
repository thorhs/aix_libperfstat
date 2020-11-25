package generated

/*
#cgo LDFLAGS: -lperfstat
#include <libperfstat.h>

*/
import "C"

type MemoryStat struct {

	VirtTotal float64
	RealTotal float64
	RealFree float64
	RealPinned float64
	RealInuse float64
	Pgbad float64
	Pgexct float64
	Pgins float64
	Pgouts float64
	Pgspins float64
	Pgspouts float64
	Scans float64
	Cycles float64
	Pgsteals float64
	Numperm float64
	PgspTotal float64
	PgspFree float64
	PgspRsvd float64
	RealSystem float64
	RealUser float64
	RealProcess float64
	VirtActive float64
	Iome float64
	Iomu float64
	Iohwm float64
	Pmem float64
	ComprsdTotal float64
	ComprsdWsegPgs float64
	Cpgins float64
	Cpgouts float64
	TrueSize float64
	ExpandedMemory float64
	ComprsdWsegSize float64
	TargetCpoolSize float64
	MaxCpoolSize float64
	MinUcpoolSize float64
	CpoolSize float64
	UcpoolSize float64
	CpoolInuse float64
	UcpoolInuse float64
	RealAvail float64
	BytesCoalesced float64
	BytesCoalescedMempool float64
}

func CollectMemory() MemoryStat {
	memorystat := C.perfstat_memory_total_t{}

	_ = C.perfstat_memory_total(nil, &memorystat, C.sizeof_perfstat_memory_total_t, 1)

	return MemoryStat{
		VirtTotal: float64(memorystat.virt_total),
		RealTotal: float64(memorystat.real_total),
		RealFree: float64(memorystat.real_free),
		RealPinned: float64(memorystat.real_pinned),
		RealInuse: float64(memorystat.real_inuse),
		Pgbad: float64(memorystat.pgbad),
		Pgexct: float64(memorystat.pgexct),
		Pgins: float64(memorystat.pgins),
		Pgouts: float64(memorystat.pgouts),
		Pgspins: float64(memorystat.pgspins),
		Pgspouts: float64(memorystat.pgspouts),
		Scans: float64(memorystat.scans),
		Cycles: float64(memorystat.cycles),
		Pgsteals: float64(memorystat.pgsteals),
		Numperm: float64(memorystat.numperm),
		PgspTotal: float64(memorystat.pgsp_total),
		PgspFree: float64(memorystat.pgsp_free),
		PgspRsvd: float64(memorystat.pgsp_rsvd),
		RealSystem: float64(memorystat.real_system),
		RealUser: float64(memorystat.real_user),
		RealProcess: float64(memorystat.real_process),
		VirtActive: float64(memorystat.virt_active),
		Iome: float64(memorystat.iome),
		Iomu: float64(memorystat.iomu),
		Iohwm: float64(memorystat.iohwm),
		Pmem: float64(memorystat.pmem),
		ComprsdTotal: float64(memorystat.comprsd_total),
		ComprsdWsegPgs: float64(memorystat.comprsd_wseg_pgs),
		Cpgins: float64(memorystat.cpgins),
		Cpgouts: float64(memorystat.cpgouts),
		TrueSize: float64(memorystat.true_size),
		ExpandedMemory: float64(memorystat.expanded_memory),
		ComprsdWsegSize: float64(memorystat.comprsd_wseg_size),
		TargetCpoolSize: float64(memorystat.target_cpool_size),
		MaxCpoolSize: float64(memorystat.max_cpool_size),
		MinUcpoolSize: float64(memorystat.min_ucpool_size),
		CpoolSize: float64(memorystat.cpool_size),
		UcpoolSize: float64(memorystat.ucpool_size),
		CpoolInuse: float64(memorystat.cpool_inuse),
		UcpoolInuse: float64(memorystat.ucpool_inuse),
		RealAvail: float64(memorystat.real_avail),
		BytesCoalesced: float64(memorystat.bytes_coalesced),
		BytesCoalescedMempool: float64(memorystat.bytes_coalesced_mempool),
	}
}
