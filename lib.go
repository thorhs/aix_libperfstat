package aix_libperfstat

//go:generate go build -o tools/generator/generator tools/generator/main.go
//go:generate tools/generator/generator data_sources/disk.multiple
//go:generate tools/generator/generator data_sources/diskpath.multiple
//go:generate tools/generator/generator data_sources/cpu.multiple
//go:generate tools/generator/generator data_sources/diskadapter.multiple
//go:generate tools/generator/generator data_sources/memory_page.multiple
//go:generate tools/generator/generator data_sources/netadapter.multiple
//go:generate tools/generator/generator data_sources/netbuffer.multiple
//go:generate tools/generator/generator data_sources/netinterface.multiple
// disabled go:generate tools/generator/generator data_sources/bio_stat.multiple

//go:generate tools/generator/generator data_sources/memory.total
//go:generate tools/generator/generator data_sources/partition.total
