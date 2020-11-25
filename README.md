# AIX libperfstat wrapper for GO

libperfstat is a library for detailed statistics gathering from the AIX
operating system from IBM.  It provides detailed measurements for things
like CPU utilization, disk utilization, disk adapters, disk paths, network
interface, memory and more.  A complete list if what is available is found
in the documentation from IBM:
(Perfstat API programming)[https://www.ibm.com/support/knowledgecenter/en/ssw_aix_72/performancetools/idprftools_perfstat.html)

## Data sources implemented

Currently the following data sources are available:

* cpu.multiple
* disk.multiple
* diskadapter.multiple
* diskpath.multiple
* memory.total
* memory_page.multiple
* netadapter.multiple
* netbuffer.multiple
* netinterface.multiple
* partition.total

Data soruces which have a .total suffix are returning 'global' data, such
as memory usage and partition information.

Data sources wich have a .multiple suffix are returning detail data
on an individual resoruce.

The data sources above do not expose the full data set available from the
libperfstat library, but mainly the ones I found useful for monitoring.
Adding additional data points or data soruces is fairly simple and involvs
only chaning the data source definitions and running `go generate`.

## Using this library
Prerequisites:

* AIX 7.1 (older may work, not tested)
* GCC compiler (AIX Toolbox)
* GNU Make (AIX Toolbox)
* Go 1.4
* bos.perf.libperfstat LPP installed

Example:
```go
package main

use "fmt"

use perfstat "github.com/thorhs/aix_libperfstat/generated"


func main() {
	fmt.Printf("%+v\n", perfstat.CollectCpus())
}
```

## Generating library

Prerequisites:

* AIX 7.1 (older may work, not tested)
* GCC compiler (AIX Toolbox)
* GNU Make (AIX Toolbox)
* Go 1.4

Building:

    export PATH=/opt/freeware/bin:$PATH # If required, to use GNU make
    go generate

## Format of data sources
### General
The data source defenitions have been copied from the previous C++
(node_exporter_aix)[https://github.com/thorhs/node_exporter_aix].
Some fields are not being used, at least not yet.

The name of the data source file is used for the data source name in libperstat.
memory.total is translated to the struct perfstat_memory_total_t an function perfstat_memory_total.
disk.multiple is translated to the struct perfstat_disk_t and function perfstat_disk.

### .total

The .total files contain raw elements of the repsective C struct which we want to capture.
Each line has three fields, seperated by Tab:

1. Metric type, gauge or counter. Not used in this library
2. Field name, as it appears in the C struct.  This will be converted to CamelCase in the Go struct.
3. Description.  Not used in this library.

### .multiple

The first line is the header and contains the following fields, seperated by space:
1. Resource identity (field in struct used for the identity).
2. ID field, not used in.
3. Optional list of fields to be extracted, strings, used for additional labels.  Multiples are separated by space.

The second line is the initializer line, it consists of two fields separated by space:
1. Name of structure for 'first' item.  Used in the libperfstat library if you only want a subset of the resources.
2. Struct initializer, should be compatible with C++ and Go, I've only ever used an empty string or 0, depending on resource type.

The rest of the lines are identical to the .total format.

# Contributing
Please see [CONTRIBUTING.md](CONTRIBUTING.md).  We welcome issues, pull requests and general suggestions.

# License
This program is distributed under the terms of the MIT license.

See [LICENSE](LICENSE) for details.
