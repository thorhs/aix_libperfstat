package aix_libperfstat

import "testing"
import "github.com/thorhs/aix_libperfstat/generated"

func TestCpus(t *testing.T) {
	cpus := generated.CollectCpus()

	if len(cpus) < 1 {
		t.Error("Empty set of CPUs returned")
	}

	for cpu, data := range cpus {
		if data.User == 0 {
			t.Errorf("CPU %s is not reporting any user CPU time", cpu)
		}
	}
}

func TestDisks(t *testing.T) {
	disks := generated.CollectDisks()

	if len(disks) < 1 {
		t.Error("Empty set of disks returned")
	}

	for disk, data := range disks {
		if data.Bsize == 0 {
			t.Errorf("Disk %s is reporting 0 in block size", disk)
		}
	}
}

func TestDiskadapters(t *testing.T) {
	diskadapters := generated.CollectDiskadapters()

	if len(diskadapters) < 1 {
		t.Error("Empty set of diskadapters returned")
	}

	for diskadapter, data := range diskadapters {
		if data.Xfers == 0 && diskadapter != "vscsi0" {
			t.Errorf("Diskadapter '%s' is reporting 0 in xfers", diskadapter)
		}
	}
}

func TestDiskpaths(t *testing.T) {
	diskpaths := generated.CollectDiskpaths()

	if len(diskpaths) < 1 {
		t.Error("Empty set of diskpaths returned")
	}

	/*
	for diskpath, data := range diskpaths {
		if data.Xfers == 0 {
			t.Errorf("Diskpath '%s' is reporting 0 in xfers", diskpath)
		}
	}
	*/
}

func TestMemoryPages(t *testing.T) {
	memorypages := generated.CollectMemoryPages()

	if len(memorypages) < 1 {
		t.Error("Empty set of memorypages returned")
	}

	for memorypage, data := range memorypages {
		if data.RealTotal == 0 && memorypage != "0" {
			t.Errorf("MemoryPage '%s' is reporting 0 in RealTotal", memorypage)
		}
	}
}

func TestNetadapters(t *testing.T) {
	netadapters := generated.CollectNetadapters()

	if len(netadapters) < 1 {
		t.Error("Empty set of netadapters returned")
	}

	for netadapter, data := range netadapters {
		if data.TxPackets == 0  {
			t.Errorf("Netadapter '%s' is reporting 0 in TxPackets", netadapter)
		}
	}
}

func TestNetbuffers(t *testing.T) {
	netbuffers := generated.CollectNetbuffers()

	if len(netbuffers) < 1 {
		t.Error("Empty set of netbuffers returned")
	}

	for netbuffer, data := range netbuffers {
		if data.Calls == 0  {
			t.Errorf("Netbuffer '%s' is reporting 0 in calls", netbuffer)
		}
	}
}

func TestNetinterfaces(t *testing.T) {
	netinterfaces := generated.CollectNetinterfaces()

	if len(netinterfaces) < 1 {
		t.Error("Empty set of netinterfaces returned")
	}

	for netinterface, data := range netinterfaces {
		if data.Opackets == 0  {
			t.Errorf("Netinterface '%s' is reporting 0 in Opackets", netinterface)
		}
	}
}

/*
func TestBio(t *testing.T) {
	bio_stats := generated.CollectBioStats()

	if len(bio_stats) < 1 {
		t.Error("Empty set of bio_stat returned")
	}

	for bio_stat, data := range bio_stats {
		if data.Nread == 0  {
			t.Errorf("Bio_stat '%s' is reporting 0 in nread", bio_stat)
		}
	}
}
*/

func TestMemory(t *testing.T) {
	memory := generated.CollectMemory()

	if memory.VirtTotal == 0  {
		t.Errorf("Memory is reporting 0 in VirtTotal")
	}
}

func TestPartition(t *testing.T) {
	partition := generated.CollectPartition()

	if partition.Pmem == 0  {
		t.Errorf("Partition is reporting 0 in Pmem")
	}
}
