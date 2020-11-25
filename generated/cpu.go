package generated

/*
#cgo LDFLAGS: -lperfstat
#include <libperfstat.h>

*/
import "C"

import (

)

type CpuStat struct {

	User float64
	Sys float64
	Idle float64
	Wait float64
	Pswitch float64
	Syscall float64
	Sysread float64
	Syswrite float64
	Sysfork float64
	Sysexec float64
	Readch float64
	Writech float64
	Bread float64
	Bwrite float64
	Lread float64
	Lwrite float64
	Phread float64
	Phwrite float64
	Iget float64
	Namei float64
	Dirblk float64
	Msg float64
	Sema float64
	Minfaults float64
	Majfaults float64
	Puser float64
	Psys float64
	Pidle float64
	Pwait float64
	RedispSd0 float64
	RedispSd1 float64
	RedispSd2 float64
	RedispSd3 float64
	RedispSd4 float64
	RedispSd5 float64
	MigrationPush float64
	MigrationS3grq float64
	MigrationS3pul float64
	InvolCswitch float64
	VolCswitch float64
	Runque float64
	Bound float64
	Decrintrs float64
	Mpcrintrs float64
	Mpcsintrs float64
	Devintrs float64
	Softintrs float64
	Phantintrs float64
	IdleDonatedPurr float64
	IdleDonatedSpurr float64
	BusyDonatedPurr float64
	BusyDonatedSpurr float64
	IdleStolenPurr float64
	IdleStolenSpurr float64
	BusyStolenPurr float64
	BusyStolenSpurr float64
	Hpi float64
	Hpit float64
	PuserSpurr float64
	PsysSpurr float64
	PidleSpurr float64
	PwaitSpurr float64
	Spurrflag float64
	Localdispatch float64
	Neardispatch float64
	Fardispatch float64
	Cswitches float64
	TbLast float64
	State float64
	VtbLast float64
	IcountLast float64
}

type CpuStats map[string]CpuStat

func CollectCpus() CpuStats {
	output := make(CpuStats)

num := C.perfstat_cpu(nil, nil, C.sizeof_perfstat_cpu_t, 0)
if num == 0 {
		return output
	}

	initStruct := C.perfstat_id_t{}
	initStruct.name[0]=0

	cpustats := make([]C.perfstat_cpu_t, num)

	_ = C.perfstat_cpu(&initStruct, &cpustats[0], C.sizeof_perfstat_cpu_t, num)

	for _, data := range cpustats {

		name := C.GoString(&data.name[0])


		output[name] = CpuStat{
			User: float64(data.user),
			Sys: float64(data.sys),
			Idle: float64(data.idle),
			Wait: float64(data.wait),
			Pswitch: float64(data.pswitch),
			Syscall: float64(data.syscall),
			Sysread: float64(data.sysread),
			Syswrite: float64(data.syswrite),
			Sysfork: float64(data.sysfork),
			Sysexec: float64(data.sysexec),
			Readch: float64(data.readch),
			Writech: float64(data.writech),
			Bread: float64(data.bread),
			Bwrite: float64(data.bwrite),
			Lread: float64(data.lread),
			Lwrite: float64(data.lwrite),
			Phread: float64(data.phread),
			Phwrite: float64(data.phwrite),
			Iget: float64(data.iget),
			Namei: float64(data.namei),
			Dirblk: float64(data.dirblk),
			Msg: float64(data.msg),
			Sema: float64(data.sema),
			Minfaults: float64(data.minfaults),
			Majfaults: float64(data.majfaults),
			Puser: float64(data.puser),
			Psys: float64(data.psys),
			Pidle: float64(data.pidle),
			Pwait: float64(data.pwait),
			RedispSd0: float64(data.redisp_sd0),
			RedispSd1: float64(data.redisp_sd1),
			RedispSd2: float64(data.redisp_sd2),
			RedispSd3: float64(data.redisp_sd3),
			RedispSd4: float64(data.redisp_sd4),
			RedispSd5: float64(data.redisp_sd5),
			MigrationPush: float64(data.migration_push),
			MigrationS3grq: float64(data.migration_S3grq),
			MigrationS3pul: float64(data.migration_S3pul),
			InvolCswitch: float64(data.invol_cswitch),
			VolCswitch: float64(data.vol_cswitch),
			Runque: float64(data.runque),
			Bound: float64(data.bound),
			Decrintrs: float64(data.decrintrs),
			Mpcrintrs: float64(data.mpcrintrs),
			Mpcsintrs: float64(data.mpcsintrs),
			Devintrs: float64(data.devintrs),
			Softintrs: float64(data.softintrs),
			Phantintrs: float64(data.phantintrs),
			IdleDonatedPurr: float64(data.idle_donated_purr),
			IdleDonatedSpurr: float64(data.idle_donated_spurr),
			BusyDonatedPurr: float64(data.busy_donated_purr),
			BusyDonatedSpurr: float64(data.busy_donated_spurr),
			IdleStolenPurr: float64(data.idle_stolen_purr),
			IdleStolenSpurr: float64(data.idle_stolen_spurr),
			BusyStolenPurr: float64(data.busy_stolen_purr),
			BusyStolenSpurr: float64(data.busy_stolen_spurr),
			Hpi: float64(data.hpi),
			Hpit: float64(data.hpit),
			PuserSpurr: float64(data.puser_spurr),
			PsysSpurr: float64(data.psys_spurr),
			PidleSpurr: float64(data.pidle_spurr),
			PwaitSpurr: float64(data.pwait_spurr),
			Spurrflag: float64(data.spurrflag),
			Localdispatch: float64(data.localdispatch),
			Neardispatch: float64(data.neardispatch),
			Fardispatch: float64(data.fardispatch),
			Cswitches: float64(data.cswitches),
			TbLast: float64(data.tb_last),
			State: float64(data.state),
			VtbLast: float64(data.vtb_last),
			IcountLast: float64(data.icount_last),
		}
	}

	return output
}
