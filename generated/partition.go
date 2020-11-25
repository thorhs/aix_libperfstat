package generated

/*
#cgo LDFLAGS: -lperfstat
#include <libperfstat.h>

*/
import "C"

type PartitionStat struct {

	OnlineCpus float64
	MaxCpus float64
	MinCpus float64
	OnlineMemory float64
	MaxMemory float64
	MinMemory float64
	EntitledProcCapacity float64
	MaxProcCapacity float64
	MinProcCapacity float64
	ProcCapacityIncrement float64
	UnallocProcCapacity float64
	VarProcCapacityWeight float64
	UnallocVarProcCapacityWeight float64
	OnlinePhysCpusSys float64
	MaxPhysCpusSys float64
	PhysCpusPool float64
	Puser float64
	Psys float64
	Pidle float64
	Pwait float64
	PoolIdleTime float64
	Phantintrs float64
	InvolVirtCswitch float64
	VolVirtCswitch float64
	TimebaseLast float64
	ReservedPages float64
	ReservedPagesize float64
	IdleDonatedPurr float64
	IdleDonatedSpurr float64
	BusyDonatedPurr float64
	BusyDonatedSpurr float64
	IdleStolenPurr float64
	IdleStolenSpurr float64
	BusyStolenPurr float64
	BusyStolenSpurr float64
	ShcpusInSys float64
	MaxPoolCapacity float64
	EntitledPoolCapacity float64
	PoolMaxTime float64
	PoolBusyTime float64
	PoolScaledBusyTime float64
	ShcpuTotTime float64
	ShcpuBusyTime float64
	ShcpuScaledBusyTime float64
	VarMemWeight float64
	Iome float64
	Pmem float64
	Hpi float64
	Hpit float64
	HypvPagesize float64
	OnlineLcpus float64
	SmtThrds float64
	PuserSpurr float64
	PsysSpurr float64
	PidleSpurr float64
	PwaitSpurr float64
	Spurrflag float64
	PowerSaveMode float64
	TrueMemory float64
	ExpandedMemory float64
	TargetMemexpFactr float64
	CurrentMemexpFactr float64
	TargetCpoolSize float64
	MaxCpoolSize float64
	MinUcpoolSize float64
	AmeDeficitSize float64
	CmcsTotalTime float64
	PurrCoalescing float64
	SpurrCoalescing float64
	 MemPoolSize  float64
	  IOMemEntInUse float64
	  IOMemEntFree float64
	  IOHighWaterMark float64
	PurrCounter float64
	SpurrCounter float64
	RealFree float64
	RealAvail float64
}

func CollectPartition() PartitionStat {
	partitionstat := C.perfstat_partition_total_t{}

	_ = C.perfstat_partition_total(nil, &partitionstat, C.sizeof_perfstat_partition_total_t, 1)

	return PartitionStat{
		OnlineCpus: float64(partitionstat.online_cpus),
		MaxCpus: float64(partitionstat.max_cpus),
		MinCpus: float64(partitionstat.min_cpus),
		OnlineMemory: float64(partitionstat.online_memory),
		MaxMemory: float64(partitionstat.max_memory),
		MinMemory: float64(partitionstat.min_memory),
		EntitledProcCapacity: float64(partitionstat.entitled_proc_capacity),
		MaxProcCapacity: float64(partitionstat.max_proc_capacity),
		MinProcCapacity: float64(partitionstat.min_proc_capacity),
		ProcCapacityIncrement: float64(partitionstat.proc_capacity_increment),
		UnallocProcCapacity: float64(partitionstat.unalloc_proc_capacity),
		VarProcCapacityWeight: float64(partitionstat.var_proc_capacity_weight),
		UnallocVarProcCapacityWeight: float64(partitionstat.unalloc_var_proc_capacity_weight),
		OnlinePhysCpusSys: float64(partitionstat.online_phys_cpus_sys),
		MaxPhysCpusSys: float64(partitionstat.max_phys_cpus_sys),
		PhysCpusPool: float64(partitionstat.phys_cpus_pool),
		Puser: float64(partitionstat.puser),
		Psys: float64(partitionstat.psys),
		Pidle: float64(partitionstat.pidle),
		Pwait: float64(partitionstat.pwait),
		PoolIdleTime: float64(partitionstat.pool_idle_time),
		Phantintrs: float64(partitionstat.phantintrs),
		InvolVirtCswitch: float64(partitionstat.invol_virt_cswitch),
		VolVirtCswitch: float64(partitionstat.vol_virt_cswitch),
		TimebaseLast: float64(partitionstat.timebase_last),
		ReservedPages: float64(partitionstat.reserved_pages),
		ReservedPagesize: float64(partitionstat.reserved_pagesize),
		IdleDonatedPurr: float64(partitionstat.idle_donated_purr),
		IdleDonatedSpurr: float64(partitionstat.idle_donated_spurr),
		BusyDonatedPurr: float64(partitionstat.busy_donated_purr),
		BusyDonatedSpurr: float64(partitionstat.busy_donated_spurr),
		IdleStolenPurr: float64(partitionstat.idle_stolen_purr),
		IdleStolenSpurr: float64(partitionstat.idle_stolen_spurr),
		BusyStolenPurr: float64(partitionstat.busy_stolen_purr),
		BusyStolenSpurr: float64(partitionstat.busy_stolen_spurr),
		ShcpusInSys: float64(partitionstat.shcpus_in_sys),
		MaxPoolCapacity: float64(partitionstat.max_pool_capacity),
		EntitledPoolCapacity: float64(partitionstat.entitled_pool_capacity),
		PoolMaxTime: float64(partitionstat.pool_max_time),
		PoolBusyTime: float64(partitionstat.pool_busy_time),
		PoolScaledBusyTime: float64(partitionstat.pool_scaled_busy_time),
		ShcpuTotTime: float64(partitionstat.shcpu_tot_time),
		ShcpuBusyTime: float64(partitionstat.shcpu_busy_time),
		ShcpuScaledBusyTime: float64(partitionstat.shcpu_scaled_busy_time),
		VarMemWeight: float64(partitionstat.var_mem_weight),
		Iome: float64(partitionstat.iome),
		Pmem: float64(partitionstat.pmem),
		Hpi: float64(partitionstat.hpi),
		Hpit: float64(partitionstat.hpit),
		HypvPagesize: float64(partitionstat.hypv_pagesize),
		OnlineLcpus: float64(partitionstat.online_lcpus),
		SmtThrds: float64(partitionstat.smt_thrds),
		PuserSpurr: float64(partitionstat.puser_spurr),
		PsysSpurr: float64(partitionstat.psys_spurr),
		PidleSpurr: float64(partitionstat.pidle_spurr),
		PwaitSpurr: float64(partitionstat.pwait_spurr),
		Spurrflag: float64(partitionstat.spurrflag),
		PowerSaveMode: float64(partitionstat.power_save_mode),
		TrueMemory: float64(partitionstat.true_memory),
		ExpandedMemory: float64(partitionstat.expanded_memory),
		TargetMemexpFactr: float64(partitionstat.target_memexp_factr),
		CurrentMemexpFactr: float64(partitionstat.current_memexp_factr),
		TargetCpoolSize: float64(partitionstat.target_cpool_size),
		MaxCpoolSize: float64(partitionstat.max_cpool_size),
		MinUcpoolSize: float64(partitionstat.min_ucpool_size),
		AmeDeficitSize: float64(partitionstat.ame_deficit_size),
		CmcsTotalTime: float64(partitionstat.cmcs_total_time),
		PurrCoalescing: float64(partitionstat.purr_coalescing),
		SpurrCoalescing: float64(partitionstat.spurr_coalescing),
		 MemPoolSize : float64(partitionstat. MemPoolSize ),
		  IOMemEntInUse: float64(partitionstat.  IOMemEntInUse),
		  IOMemEntFree: float64(partitionstat.  IOMemEntFree),
		  IOHighWaterMark: float64(partitionstat.  IOHighWaterMark),
		PurrCounter: float64(partitionstat.purr_counter),
		SpurrCounter: float64(partitionstat.spurr_counter),
		RealFree: float64(partitionstat.real_free),
		RealAvail: float64(partitionstat.real_avail),
	}
}
