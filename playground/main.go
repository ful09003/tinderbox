package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/ful09003/tinderbox/pkg/scrape"
	"github.com/ful09003/tinderbox/pkg/types"
	dto "github.com/prometheus/client_model/go"
)

var data = `# HELP joy_felt_total A counter of joy experienced.
	# TYPE joy_felt_total counter
	joy_felt_total{developer="me"} 9000
	# HELP despair_felt_total A counter of despair experienced.
	# TYPE despair_felt_total counter
	despair_felt_total{developer="me"} 9001
`

var nodeExporterData = `# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 7.732e-06
go_gc_duration_seconds{quantile="0.25"} 3.0485e-05
go_gc_duration_seconds{quantile="0.5"} 3.3451e-05
go_gc_duration_seconds{quantile="0.75"} 3.902e-05
go_gc_duration_seconds{quantile="1"} 0.027804582
go_gc_duration_seconds_sum 0.040211961
go_gc_duration_seconds_count 333
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 7
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.14.4"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 793536
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 9.449888e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.446319e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 61488
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 2.5884041344404004e-06
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 3.574024e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 793536
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 6.3709184e+07
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 2.318336e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 3577
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 6.3586304e+07
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.602752e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.6383327349963167e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 65065
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 13888
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 117368
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 131072
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.730121e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 1.081344e+06
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 1.081344e+06
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.4006784e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 19
# HELP node_arp_entries ARP entries by device
# TYPE node_arp_entries gauge
node_arp_entries{device="wlp2s0"} 6
# HELP node_boot_time_seconds Node boot time, in unixtime.
# TYPE node_boot_time_seconds gauge
node_boot_time_seconds 1.633374242e+09
# HELP node_context_switches_total Total number of context switches.
# TYPE node_context_switches_total counter
node_context_switches_total 5.84482941e+08
# HELP node_cooling_device_cur_state Current throttle state of the cooling device
# TYPE node_cooling_device_cur_state gauge
node_cooling_device_cur_state{name="0",type="Processor"} 0
node_cooling_device_cur_state{name="1",type="Processor"} 0
node_cooling_device_cur_state{name="10",type="ath10k_thermal"} 0
node_cooling_device_cur_state{name="2",type="Processor"} 0
node_cooling_device_cur_state{name="3",type="Processor"} 0
node_cooling_device_cur_state{name="4",type="Processor"} 0
node_cooling_device_cur_state{name="5",type="Processor"} 0
node_cooling_device_cur_state{name="6",type="Processor"} 0
node_cooling_device_cur_state{name="7",type="Processor"} 0
node_cooling_device_cur_state{name="8",type="intel_powerclamp"} -1
node_cooling_device_cur_state{name="9",type="TCC Offset"} 0
# HELP node_cooling_device_max_state Maximum throttle state of the cooling device
# TYPE node_cooling_device_max_state gauge
node_cooling_device_max_state{name="0",type="Processor"} 3
node_cooling_device_max_state{name="1",type="Processor"} 3
node_cooling_device_max_state{name="10",type="ath10k_thermal"} 100
node_cooling_device_max_state{name="2",type="Processor"} 3
node_cooling_device_max_state{name="3",type="Processor"} 3
node_cooling_device_max_state{name="4",type="Processor"} 3
node_cooling_device_max_state{name="5",type="Processor"} 3
node_cooling_device_max_state{name="6",type="Processor"} 3
node_cooling_device_max_state{name="7",type="Processor"} 3
node_cooling_device_max_state{name="8",type="intel_powerclamp"} 50
node_cooling_device_max_state{name="9",type="TCC Offset"} 63
# HELP node_cpu_core_throttles_total Number of times this cpu core has been throttled.
# TYPE node_cpu_core_throttles_total counter
node_cpu_core_throttles_total{core="0",package="0"} 3734
node_cpu_core_throttles_total{core="1",package="0"} 1220
node_cpu_core_throttles_total{core="2",package="0"} 12320
node_cpu_core_throttles_total{core="3",package="0"} 3792
# HELP node_cpu_frequency_max_hertz Maximum cpu thread frequency in hertz.
# TYPE node_cpu_frequency_max_hertz gauge
node_cpu_frequency_max_hertz{cpu="0"} 4.6e+09
node_cpu_frequency_max_hertz{cpu="1"} 4.6e+09
node_cpu_frequency_max_hertz{cpu="2"} 4.6e+09
node_cpu_frequency_max_hertz{cpu="3"} 4.6e+09
node_cpu_frequency_max_hertz{cpu="4"} 4.6e+09
node_cpu_frequency_max_hertz{cpu="5"} 4.6e+09
node_cpu_frequency_max_hertz{cpu="6"} 4.6e+09
node_cpu_frequency_max_hertz{cpu="7"} 4.6e+09
# HELP node_cpu_frequency_min_hertz Minimum cpu thread frequency in hertz.
# TYPE node_cpu_frequency_min_hertz gauge
node_cpu_frequency_min_hertz{cpu="0"} 4e+08
node_cpu_frequency_min_hertz{cpu="1"} 4e+08
node_cpu_frequency_min_hertz{cpu="2"} 4e+08
node_cpu_frequency_min_hertz{cpu="3"} 4e+08
node_cpu_frequency_min_hertz{cpu="4"} 4e+08
node_cpu_frequency_min_hertz{cpu="5"} 4e+08
node_cpu_frequency_min_hertz{cpu="6"} 4e+08
node_cpu_frequency_min_hertz{cpu="7"} 4e+08
# HELP node_cpu_guest_seconds_total Seconds the cpus spent in guests (VMs) for each mode.
# TYPE node_cpu_guest_seconds_total counter
node_cpu_guest_seconds_total{cpu="0",mode="nice"} 0
node_cpu_guest_seconds_total{cpu="0",mode="user"} 0
node_cpu_guest_seconds_total{cpu="1",mode="nice"} 0
node_cpu_guest_seconds_total{cpu="1",mode="user"} 0
node_cpu_guest_seconds_total{cpu="2",mode="nice"} 0
node_cpu_guest_seconds_total{cpu="2",mode="user"} 0
node_cpu_guest_seconds_total{cpu="3",mode="nice"} 0
node_cpu_guest_seconds_total{cpu="3",mode="user"} 0
node_cpu_guest_seconds_total{cpu="4",mode="nice"} 0
node_cpu_guest_seconds_total{cpu="4",mode="user"} 0
node_cpu_guest_seconds_total{cpu="5",mode="nice"} 0
node_cpu_guest_seconds_total{cpu="5",mode="user"} 0
node_cpu_guest_seconds_total{cpu="6",mode="nice"} 0
node_cpu_guest_seconds_total{cpu="6",mode="user"} 0
node_cpu_guest_seconds_total{cpu="7",mode="nice"} 0
node_cpu_guest_seconds_total{cpu="7",mode="user"} 0
# HELP node_cpu_package_throttles_total Number of times this cpu package has been throttled.
# TYPE node_cpu_package_throttles_total counter
node_cpu_package_throttles_total{package="0"} 18598
# HELP node_cpu_scaling_frequency_hertz Current scaled cpu thread frequency in hertz.
# TYPE node_cpu_scaling_frequency_hertz gauge
node_cpu_scaling_frequency_hertz{cpu="0"} 1.854604e+09
node_cpu_scaling_frequency_hertz{cpu="1"} 1.917634e+09
node_cpu_scaling_frequency_hertz{cpu="2"} 1.886364e+09
node_cpu_scaling_frequency_hertz{cpu="3"} 1.822633e+09
node_cpu_scaling_frequency_hertz{cpu="4"} 1.845184e+09
node_cpu_scaling_frequency_hertz{cpu="5"} 1.841617e+09
node_cpu_scaling_frequency_hertz{cpu="6"} 1.825986e+09
node_cpu_scaling_frequency_hertz{cpu="7"} 1.849999e+09
# HELP node_cpu_scaling_frequency_max_hertz Maximum scaled cpu thread frequency in hertz.
# TYPE node_cpu_scaling_frequency_max_hertz gauge
node_cpu_scaling_frequency_max_hertz{cpu="0"} 4.6e+09
node_cpu_scaling_frequency_max_hertz{cpu="1"} 4.6e+09
node_cpu_scaling_frequency_max_hertz{cpu="2"} 4.6e+09
node_cpu_scaling_frequency_max_hertz{cpu="3"} 4.6e+09
node_cpu_scaling_frequency_max_hertz{cpu="4"} 4.6e+09
node_cpu_scaling_frequency_max_hertz{cpu="5"} 4.6e+09
node_cpu_scaling_frequency_max_hertz{cpu="6"} 4.6e+09
node_cpu_scaling_frequency_max_hertz{cpu="7"} 4.6e+09
# HELP node_cpu_scaling_frequency_min_hertz Minimum scaled cpu thread frequency in hertz.
# TYPE node_cpu_scaling_frequency_min_hertz gauge
node_cpu_scaling_frequency_min_hertz{cpu="0"} 4e+08
node_cpu_scaling_frequency_min_hertz{cpu="1"} 4e+08
node_cpu_scaling_frequency_min_hertz{cpu="2"} 4e+08
node_cpu_scaling_frequency_min_hertz{cpu="3"} 4e+08
node_cpu_scaling_frequency_min_hertz{cpu="4"} 4e+08
node_cpu_scaling_frequency_min_hertz{cpu="5"} 4e+08
node_cpu_scaling_frequency_min_hertz{cpu="6"} 4e+08
node_cpu_scaling_frequency_min_hertz{cpu="7"} 4e+08
# HELP node_cpu_seconds_total Seconds the cpus spent in each mode.
# TYPE node_cpu_seconds_total counter
node_cpu_seconds_total{cpu="0",mode="idle"} 85869.94
node_cpu_seconds_total{cpu="0",mode="iowait"} 160.16
node_cpu_seconds_total{cpu="0",mode="irq"} 0
node_cpu_seconds_total{cpu="0",mode="nice"} 4.92
node_cpu_seconds_total{cpu="0",mode="softirq"} 21.14
node_cpu_seconds_total{cpu="0",mode="steal"} 0
node_cpu_seconds_total{cpu="0",mode="system"} 1987.75
node_cpu_seconds_total{cpu="0",mode="user"} 5698.3
node_cpu_seconds_total{cpu="1",mode="idle"} 86245.07
node_cpu_seconds_total{cpu="1",mode="iowait"} 152.46
node_cpu_seconds_total{cpu="1",mode="irq"} 0
node_cpu_seconds_total{cpu="1",mode="nice"} 6.13
node_cpu_seconds_total{cpu="1",mode="softirq"} 18.08
node_cpu_seconds_total{cpu="1",mode="steal"} 0
node_cpu_seconds_total{cpu="1",mode="system"} 2174.86
node_cpu_seconds_total{cpu="1",mode="user"} 5730.36
node_cpu_seconds_total{cpu="2",mode="idle"} 86837.57
node_cpu_seconds_total{cpu="2",mode="iowait"} 155.94
node_cpu_seconds_total{cpu="2",mode="irq"} 0
node_cpu_seconds_total{cpu="2",mode="nice"} 4.85
node_cpu_seconds_total{cpu="2",mode="softirq"} 7.47
node_cpu_seconds_total{cpu="2",mode="steal"} 0
node_cpu_seconds_total{cpu="2",mode="system"} 1986.4
node_cpu_seconds_total{cpu="2",mode="user"} 5734.02
node_cpu_seconds_total{cpu="3",mode="idle"} 86701.02
node_cpu_seconds_total{cpu="3",mode="iowait"} 160.13
node_cpu_seconds_total{cpu="3",mode="irq"} 0
node_cpu_seconds_total{cpu="3",mode="nice"} 5.89
node_cpu_seconds_total{cpu="3",mode="softirq"} 6.69
node_cpu_seconds_total{cpu="3",mode="steal"} 0
node_cpu_seconds_total{cpu="3",mode="system"} 2060.08
node_cpu_seconds_total{cpu="3",mode="user"} 5779.27
node_cpu_seconds_total{cpu="4",mode="idle"} 86991.45
node_cpu_seconds_total{cpu="4",mode="iowait"} 154.76
node_cpu_seconds_total{cpu="4",mode="irq"} 0
node_cpu_seconds_total{cpu="4",mode="nice"} 5.23
node_cpu_seconds_total{cpu="4",mode="softirq"} 17.65
node_cpu_seconds_total{cpu="4",mode="steal"} 0
node_cpu_seconds_total{cpu="4",mode="system"} 1862.98
node_cpu_seconds_total{cpu="4",mode="user"} 5646.78
node_cpu_seconds_total{cpu="5",mode="idle"} 86807.78
node_cpu_seconds_total{cpu="5",mode="iowait"} 154.53
node_cpu_seconds_total{cpu="5",mode="irq"} 0
node_cpu_seconds_total{cpu="5",mode="nice"} 4.88
node_cpu_seconds_total{cpu="5",mode="softirq"} 137.72
node_cpu_seconds_total{cpu="5",mode="steal"} 0
node_cpu_seconds_total{cpu="5",mode="system"} 1962.61
node_cpu_seconds_total{cpu="5",mode="user"} 5134.66
node_cpu_seconds_total{cpu="6",mode="idle"} 86438.93
node_cpu_seconds_total{cpu="6",mode="iowait"} 144.32
node_cpu_seconds_total{cpu="6",mode="irq"} 0
node_cpu_seconds_total{cpu="6",mode="nice"} 4.54
node_cpu_seconds_total{cpu="6",mode="softirq"} 5.68
node_cpu_seconds_total{cpu="6",mode="steal"} 0
node_cpu_seconds_total{cpu="6",mode="system"} 2329.95
node_cpu_seconds_total{cpu="6",mode="user"} 5966.72
node_cpu_seconds_total{cpu="7",mode="idle"} 86790.92
node_cpu_seconds_total{cpu="7",mode="iowait"} 173.74
node_cpu_seconds_total{cpu="7",mode="irq"} 0
node_cpu_seconds_total{cpu="7",mode="nice"} 4.12
node_cpu_seconds_total{cpu="7",mode="softirq"} 127.47
node_cpu_seconds_total{cpu="7",mode="steal"} 0
node_cpu_seconds_total{cpu="7",mode="system"} 1891.72
node_cpu_seconds_total{cpu="7",mode="user"} 5622.01
# HELP node_disk_discard_time_seconds_total This is the total number of seconds spent by all discards.
# TYPE node_disk_discard_time_seconds_total counter
node_disk_discard_time_seconds_total{device="mmcblk0"} 0
node_disk_discard_time_seconds_total{device="mmcblk0p1"} 0
node_disk_discard_time_seconds_total{device="mmcblk0p2"} 0
node_disk_discard_time_seconds_total{device="nvme0n1"} 11.605
# HELP node_disk_discarded_sectors_total The total number of sectors discarded successfully.
# TYPE node_disk_discarded_sectors_total counter
node_disk_discarded_sectors_total{device="mmcblk0"} 0
node_disk_discarded_sectors_total{device="mmcblk0p1"} 0
node_disk_discarded_sectors_total{device="mmcblk0p2"} 0
node_disk_discarded_sectors_total{device="nvme0n1"} 8.76179368e+08
# HELP node_disk_discards_completed_total The total number of discards completed successfully.
# TYPE node_disk_discards_completed_total counter
node_disk_discards_completed_total{device="mmcblk0"} 0
node_disk_discards_completed_total{device="mmcblk0p1"} 0
node_disk_discards_completed_total{device="mmcblk0p2"} 0
node_disk_discards_completed_total{device="nvme0n1"} 111091
# HELP node_disk_discards_merged_total The total number of discards merged.
# TYPE node_disk_discards_merged_total counter
node_disk_discards_merged_total{device="mmcblk0"} 0
node_disk_discards_merged_total{device="mmcblk0p1"} 0
node_disk_discards_merged_total{device="mmcblk0p2"} 0
node_disk_discards_merged_total{device="nvme0n1"} 0
# HELP node_disk_flush_requests_time_seconds_total This is the total number of seconds spent by all flush requests.
# TYPE node_disk_flush_requests_time_seconds_total counter
node_disk_flush_requests_time_seconds_total{device="mmcblk0"} 0
node_disk_flush_requests_time_seconds_total{device="mmcblk0p1"} 0
node_disk_flush_requests_time_seconds_total{device="mmcblk0p2"} 0
node_disk_flush_requests_time_seconds_total{device="nvme0n1"} 188.594
# HELP node_disk_flush_requests_total The total number of flush requests completed successfully
# TYPE node_disk_flush_requests_total counter
node_disk_flush_requests_total{device="mmcblk0"} 0
node_disk_flush_requests_total{device="mmcblk0p1"} 0
node_disk_flush_requests_total{device="mmcblk0p2"} 0
node_disk_flush_requests_total{device="nvme0n1"} 194093
# HELP node_disk_io_now The number of I/Os currently in progress.
# TYPE node_disk_io_now gauge
node_disk_io_now{device="mmcblk0"} 0
node_disk_io_now{device="mmcblk0p1"} 0
node_disk_io_now{device="mmcblk0p2"} 0
node_disk_io_now{device="nvme0n1"} 0
# HELP node_disk_io_time_seconds_total Total seconds spent doing I/Os.
# TYPE node_disk_io_time_seconds_total counter
node_disk_io_time_seconds_total{device="mmcblk0"} 8.624
node_disk_io_time_seconds_total{device="mmcblk0p1"} 3.18
node_disk_io_time_seconds_total{device="mmcblk0p2"} 8.388
node_disk_io_time_seconds_total{device="nvme0n1"} 1993.988
# HELP node_disk_io_time_weighted_seconds_total The weighted # of seconds spent doing I/Os.
# TYPE node_disk_io_time_weighted_seconds_total counter
node_disk_io_time_weighted_seconds_total{device="mmcblk0"} 13.297
node_disk_io_time_weighted_seconds_total{device="mmcblk0p1"} 3.182
node_disk_io_time_weighted_seconds_total{device="mmcblk0p2"} 8.78
node_disk_io_time_weighted_seconds_total{device="nvme0n1"} 8429.999
# HELP node_disk_read_bytes_total The total number of bytes read successfully.
# TYPE node_disk_read_bytes_total counter
node_disk_read_bytes_total{device="mmcblk0"} 1.9167232e+07
node_disk_read_bytes_total{device="mmcblk0p1"} 4.330496e+06
node_disk_read_bytes_total{device="mmcblk0p2"} 1.1998208e+07
node_disk_read_bytes_total{device="nvme0n1"} 5.179343872e+09
# HELP node_disk_read_time_seconds_total The total number of seconds spent by all reads.
# TYPE node_disk_read_time_seconds_total counter
node_disk_read_time_seconds_total{device="mmcblk0"} 10.41
node_disk_read_time_seconds_total{device="mmcblk0p1"} 1.992
node_disk_read_time_seconds_total{device="mmcblk0p2"} 7.082
node_disk_read_time_seconds_total{device="nvme0n1"} 214.466
# HELP node_disk_reads_completed_total The total number of reads completed successfully.
# TYPE node_disk_reads_completed_total counter
node_disk_reads_completed_total{device="mmcblk0"} 1401
node_disk_reads_completed_total{device="mmcblk0p1"} 188
node_disk_reads_completed_total{device="mmcblk0p2"} 1028
node_disk_reads_completed_total{device="nvme0n1"} 177039
# HELP node_disk_reads_merged_total The total number of reads merged.
# TYPE node_disk_reads_merged_total counter
node_disk_reads_merged_total{device="mmcblk0"} 93
node_disk_reads_merged_total{device="mmcblk0p1"} 11
node_disk_reads_merged_total{device="mmcblk0p2"} 82
node_disk_reads_merged_total{device="nvme0n1"} 79246
# HELP node_disk_write_time_seconds_total This is the total number of seconds spent by all writes.
# TYPE node_disk_write_time_seconds_total counter
node_disk_write_time_seconds_total{device="mmcblk0"} 2.887
node_disk_write_time_seconds_total{device="mmcblk0p1"} 1.189
node_disk_write_time_seconds_total{device="mmcblk0p2"} 1.697
node_disk_write_time_seconds_total{device="nvme0n1"} 8015.332
# HELP node_disk_writes_completed_total The total number of writes completed successfully.
# TYPE node_disk_writes_completed_total counter
node_disk_writes_completed_total{device="mmcblk0"} 35
node_disk_writes_completed_total{device="mmcblk0p1"} 12
node_disk_writes_completed_total{device="mmcblk0p2"} 23
node_disk_writes_completed_total{device="nvme0n1"} 1.182857e+06
# HELP node_disk_writes_merged_total The number of writes merged.
# TYPE node_disk_writes_merged_total counter
node_disk_writes_merged_total{device="mmcblk0"} 125
node_disk_writes_merged_total{device="mmcblk0p1"} 6
node_disk_writes_merged_total{device="mmcblk0p2"} 119
node_disk_writes_merged_total{device="nvme0n1"} 1.156501e+06
# HELP node_disk_written_bytes_total The total number of bytes written successfully.
# TYPE node_disk_written_bytes_total counter
node_disk_written_bytes_total{device="mmcblk0"} 655360
node_disk_written_bytes_total{device="mmcblk0p1"} 73728
node_disk_written_bytes_total{device="mmcblk0p2"} 581632
node_disk_written_bytes_total{device="nvme0n1"} 3.2066113536e+10
# HELP node_entropy_available_bits Bits of available entropy.
# TYPE node_entropy_available_bits gauge
node_entropy_available_bits 3792
# HELP node_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which node_exporter was built.
# TYPE node_exporter_build_info gauge
node_exporter_build_info{branch="HEAD",goversion="go1.14.4",revision="3715be6ae899f2a9b9dbfd9c39f3e09a7bd4559f",version="1.0.1"} 1
# HELP node_filefd_allocated File descriptor statistics: allocated.
# TYPE node_filefd_allocated gauge
node_filefd_allocated 16896
# HELP node_filefd_maximum File descriptor statistics: maximum.
# TYPE node_filefd_maximum gauge
node_filefd_maximum 9.223372036854776e+18
# HELP node_filesystem_avail_bytes Filesystem space available to non-root users in bytes.
# TYPE node_filesystem_avail_bytes gauge
node_filesystem_avail_bytes{device="/dev/nvme0n1p1",fstype="vfat",mountpoint="/boot/efi"} 6.025216e+07
node_filesystem_avail_bytes{device="/dev/nvme0n1p2",fstype="vfat",mountpoint="/recovery"} 1.247993856e+09
node_filesystem_avail_bytes{device="/dev/nvme0n1p3",fstype="ext4",mountpoint="/"} 3.63241611264e+11
node_filesystem_avail_bytes{device="tmpfs",fstype="tmpfs",mountpoint="/run"} 1.63831808e+09
node_filesystem_avail_bytes{device="tmpfs",fstype="tmpfs",mountpoint="/run/lock"} 5.24288e+06
node_filesystem_avail_bytes{device="tmpfs",fstype="tmpfs",mountpoint="/run/user/1000"} 1.637691392e+09
# HELP node_filesystem_device_error Whether an error occurred while getting statistics for the given device.
# TYPE node_filesystem_device_error gauge
node_filesystem_device_error{device="/dev/mmcblk0p1",fstype="ext4",mountpoint="/media/michael/8bbb0642-cf6e-42cf-9478-0066a8fbdb12"} 1
node_filesystem_device_error{device="/dev/mmcblk0p2",fstype="ext4",mountpoint="/media/michael/945137cc-4ae5-4f5a-84f5-8a40247a4e53"} 1
node_filesystem_device_error{device="/dev/nvme0n1p1",fstype="vfat",mountpoint="/boot/efi"} 0
node_filesystem_device_error{device="/dev/nvme0n1p2",fstype="vfat",mountpoint="/recovery"} 0
node_filesystem_device_error{device="/dev/nvme0n1p3",fstype="ext4",mountpoint="/"} 0
node_filesystem_device_error{device="gvfsd-fuse",fstype="fuse.gvfsd-fuse",mountpoint="/run/user/1000/gvfs"} 1
node_filesystem_device_error{device="portal",fstype="fuse.portal",mountpoint="/run/user/1000/doc"} 1
node_filesystem_device_error{device="tmpfs",fstype="tmpfs",mountpoint="/run"} 0
node_filesystem_device_error{device="tmpfs",fstype="tmpfs",mountpoint="/run/lock"} 0
node_filesystem_device_error{device="tmpfs",fstype="tmpfs",mountpoint="/run/user/1000"} 0
# HELP node_filesystem_files Filesystem total file nodes.
# TYPE node_filesystem_files gauge
node_filesystem_files{device="/dev/nvme0n1p1",fstype="vfat",mountpoint="/boot/efi"} 0
node_filesystem_files{device="/dev/nvme0n1p2",fstype="vfat",mountpoint="/recovery"} 0
node_filesystem_files{device="/dev/nvme0n1p3",fstype="ext4",mountpoint="/"} 2.9138944e+07
node_filesystem_files{device="tmpfs",fstype="tmpfs",mountpoint="/run"} 2.002615e+06
node_filesystem_files{device="tmpfs",fstype="tmpfs",mountpoint="/run/lock"} 2.002615e+06
node_filesystem_files{device="tmpfs",fstype="tmpfs",mountpoint="/run/user/1000"} 400523
# HELP node_filesystem_files_free Filesystem total free file nodes.
# TYPE node_filesystem_files_free gauge
node_filesystem_files_free{device="/dev/nvme0n1p1",fstype="vfat",mountpoint="/boot/efi"} 0
node_filesystem_files_free{device="/dev/nvme0n1p2",fstype="vfat",mountpoint="/recovery"} 0
node_filesystem_files_free{device="/dev/nvme0n1p3",fstype="ext4",mountpoint="/"} 2.8474805e+07
node_filesystem_files_free{device="tmpfs",fstype="tmpfs",mountpoint="/run"} 2.001463e+06
node_filesystem_files_free{device="tmpfs",fstype="tmpfs",mountpoint="/run/lock"} 2.002612e+06
node_filesystem_files_free{device="tmpfs",fstype="tmpfs",mountpoint="/run/user/1000"} 400270
# HELP node_filesystem_free_bytes Filesystem free space in bytes.
# TYPE node_filesystem_free_bytes gauge
node_filesystem_free_bytes{device="/dev/nvme0n1p1",fstype="vfat",mountpoint="/boot/efi"} 6.025216e+07
node_filesystem_free_bytes{device="/dev/nvme0n1p2",fstype="vfat",mountpoint="/recovery"} 1.247993856e+09
node_filesystem_free_bytes{device="/dev/nvme0n1p3",fstype="ext4",mountpoint="/"} 3.87122925568e+11
node_filesystem_free_bytes{device="tmpfs",fstype="tmpfs",mountpoint="/run"} 1.63831808e+09
node_filesystem_free_bytes{device="tmpfs",fstype="tmpfs",mountpoint="/run/lock"} 5.24288e+06
node_filesystem_free_bytes{device="tmpfs",fstype="tmpfs",mountpoint="/run/user/1000"} 1.637691392e+09
# HELP node_filesystem_readonly Filesystem read-only status.
# TYPE node_filesystem_readonly gauge
node_filesystem_readonly{device="/dev/nvme0n1p1",fstype="vfat",mountpoint="/boot/efi"} 0
node_filesystem_readonly{device="/dev/nvme0n1p2",fstype="vfat",mountpoint="/recovery"} 0
node_filesystem_readonly{device="/dev/nvme0n1p3",fstype="ext4",mountpoint="/"} 0
node_filesystem_readonly{device="tmpfs",fstype="tmpfs",mountpoint="/run"} 0
node_filesystem_readonly{device="tmpfs",fstype="tmpfs",mountpoint="/run/lock"} 0
node_filesystem_readonly{device="tmpfs",fstype="tmpfs",mountpoint="/run/user/1000"} 0
# HELP node_filesystem_size_bytes Filesystem size in bytes.
# TYPE node_filesystem_size_bytes gauge
node_filesystem_size_bytes{device="/dev/nvme0n1p1",fstype="vfat",mountpoint="/boot/efi"} 5.21146368e+08
node_filesystem_size_bytes{device="/dev/nvme0n1p2",fstype="vfat",mountpoint="/recovery"} 4.2865664e+09
node_filesystem_size_bytes{device="/dev/nvme0n1p3",fstype="ext4",mountpoint="/"} 4.68724166656e+11
node_filesystem_size_bytes{device="tmpfs",fstype="tmpfs",mountpoint="/run"} 1.640546304e+09
node_filesystem_size_bytes{device="tmpfs",fstype="tmpfs",mountpoint="/run/lock"} 5.24288e+06
node_filesystem_size_bytes{device="tmpfs",fstype="tmpfs",mountpoint="/run/user/1000"} 1.640542208e+09
# HELP node_forks_total Total number of forks.
# TYPE node_forks_total counter
node_forks_total 505018
# HELP node_hwmon_chip_names Annotation metric for human-readable chip names
# TYPE node_hwmon_chip_names gauge
node_hwmon_chip_names{chip="0000:00:1c_6_0000:02:00_0",chip_name="ath10k_hwmon"} 1
node_hwmon_chip_names{chip="dell_smm",chip_name="dell_smm"} 1
node_hwmon_chip_names{chip="nvme_nvme0",chip_name="nvme"} 1
node_hwmon_chip_names{chip="platform_coretemp_0",chip_name="coretemp"} 1
node_hwmon_chip_names{chip="power_supply_ac",chip_name="ac"} 1
node_hwmon_chip_names{chip="power_supply_bat0",chip_name="bat0"} 1
node_hwmon_chip_names{chip="power_supply_ucsi_source_psy_usbc000:001",chip_name="ucsi_source_psy_usbc000:001"} 1
node_hwmon_chip_names{chip="power_supply_ucsi_source_psy_usbc000:002",chip_name="ucsi_source_psy_usbc000:002"} 1
node_hwmon_chip_names{chip="power_supply_ucsi_source_psy_usbc000:003",chip_name="ucsi_source_psy_usbc000:003"} 1
node_hwmon_chip_names{chip="thermal_thermal_zone0",chip_name="acpitz"} 1
node_hwmon_chip_names{chip="thermal_thermal_zone8",chip_name="pch_cannonlake"} 1
# HELP node_hwmon_curr_amps Hardware monitor for current (input)
# TYPE node_hwmon_curr_amps gauge
node_hwmon_curr_amps{chip="power_supply_bat0",sensor="curr1"} 0.001
node_hwmon_curr_amps{chip="power_supply_ucsi_source_psy_usbc000:001",sensor="curr1"} 0
node_hwmon_curr_amps{chip="power_supply_ucsi_source_psy_usbc000:002",sensor="curr1"} 0
node_hwmon_curr_amps{chip="power_supply_ucsi_source_psy_usbc000:003",sensor="curr1"} 2.25
# HELP node_hwmon_curr_max_amps Hardware monitor for current (max)
# TYPE node_hwmon_curr_max_amps gauge
node_hwmon_curr_max_amps{chip="power_supply_ucsi_source_psy_usbc000:001",sensor="curr1"} 0
node_hwmon_curr_max_amps{chip="power_supply_ucsi_source_psy_usbc000:002",sensor="curr1"} 0
node_hwmon_curr_max_amps{chip="power_supply_ucsi_source_psy_usbc000:003",sensor="curr1"} 2.25
# HELP node_hwmon_fan_rpm Hardware monitor for fan revolutions per minute (input)
# TYPE node_hwmon_fan_rpm gauge
node_hwmon_fan_rpm{chip="dell_smm",sensor="fan1"} 0
node_hwmon_fan_rpm{chip="dell_smm",sensor="fan2"} 0
# HELP node_hwmon_in_max_volts Hardware monitor for voltage (max)
# TYPE node_hwmon_in_max_volts gauge
node_hwmon_in_max_volts{chip="power_supply_ucsi_source_psy_usbc000:001",sensor="in0"} 5
node_hwmon_in_max_volts{chip="power_supply_ucsi_source_psy_usbc000:002",sensor="in0"} 5
node_hwmon_in_max_volts{chip="power_supply_ucsi_source_psy_usbc000:003",sensor="in0"} 20
# HELP node_hwmon_in_min_volts Hardware monitor for voltage (min)
# TYPE node_hwmon_in_min_volts gauge
node_hwmon_in_min_volts{chip="power_supply_ucsi_source_psy_usbc000:001",sensor="in0"} 5
node_hwmon_in_min_volts{chip="power_supply_ucsi_source_psy_usbc000:002",sensor="in0"} 5
node_hwmon_in_min_volts{chip="power_supply_ucsi_source_psy_usbc000:003",sensor="in0"} 5
# HELP node_hwmon_in_volts Hardware monitor for voltage (input)
# TYPE node_hwmon_in_volts gauge
node_hwmon_in_volts{chip="power_supply_bat0",sensor="in0"} 8.647
node_hwmon_in_volts{chip="power_supply_ucsi_source_psy_usbc000:001",sensor="in0"} 5
node_hwmon_in_volts{chip="power_supply_ucsi_source_psy_usbc000:002",sensor="in0"} 5
node_hwmon_in_volts{chip="power_supply_ucsi_source_psy_usbc000:003",sensor="in0"} 20
# HELP node_hwmon_pwm Hardware monitor pwm element 
# TYPE node_hwmon_pwm gauge
node_hwmon_pwm{chip="dell_smm",sensor="pwm1"} 0
node_hwmon_pwm{chip="dell_smm",sensor="pwm2"} 0
# HELP node_hwmon_sensor_label Label for given chip and sensor
# TYPE node_hwmon_sensor_label gauge
node_hwmon_sensor_label{chip="nvme_nvme0",label="composite",sensor="temp1"} 1
node_hwmon_sensor_label{chip="nvme_nvme0",label="sensor_1",sensor="temp2"} 1
node_hwmon_sensor_label{chip="nvme_nvme0",label="sensor_2",sensor="temp3"} 1
node_hwmon_sensor_label{chip="platform_coretemp_0",label="core_0",sensor="temp2"} 1
node_hwmon_sensor_label{chip="platform_coretemp_0",label="core_1",sensor="temp3"} 1
node_hwmon_sensor_label{chip="platform_coretemp_0",label="core_2",sensor="temp4"} 1
node_hwmon_sensor_label{chip="platform_coretemp_0",label="core_3",sensor="temp5"} 1
node_hwmon_sensor_label{chip="platform_coretemp_0",label="package_id_0",sensor="temp1"} 1
# HELP node_hwmon_temp_alarm Hardware sensor alarm status (temp)
# TYPE node_hwmon_temp_alarm gauge
node_hwmon_temp_alarm{chip="nvme_nvme0",sensor="temp1"} 0
# HELP node_hwmon_temp_celsius Hardware monitor for temperature (input)
# TYPE node_hwmon_temp_celsius gauge
node_hwmon_temp_celsius{chip="0000:00:1c_6_0000:02:00_0",sensor="temp1"} 39
node_hwmon_temp_celsius{chip="nvme_nvme0",sensor="temp1"} 42.85
node_hwmon_temp_celsius{chip="nvme_nvme0",sensor="temp2"} 42.85
node_hwmon_temp_celsius{chip="nvme_nvme0",sensor="temp3"} 43.85
node_hwmon_temp_celsius{chip="platform_coretemp_0",sensor="temp1"} 52
node_hwmon_temp_celsius{chip="platform_coretemp_0",sensor="temp2"} 52
node_hwmon_temp_celsius{chip="platform_coretemp_0",sensor="temp3"} 50
node_hwmon_temp_celsius{chip="platform_coretemp_0",sensor="temp4"} 49
node_hwmon_temp_celsius{chip="platform_coretemp_0",sensor="temp5"} 50
node_hwmon_temp_celsius{chip="thermal_thermal_zone0",sensor="temp0"} 25
node_hwmon_temp_celsius{chip="thermal_thermal_zone0",sensor="temp1"} 25
node_hwmon_temp_celsius{chip="thermal_thermal_zone8",sensor="temp0"} 49
node_hwmon_temp_celsius{chip="thermal_thermal_zone8",sensor="temp1"} 48
# HELP node_hwmon_temp_crit_alarm_celsius Hardware monitor for temperature (crit_alarm)
# TYPE node_hwmon_temp_crit_alarm_celsius gauge
node_hwmon_temp_crit_alarm_celsius{chip="platform_coretemp_0",sensor="temp1"} 0
node_hwmon_temp_crit_alarm_celsius{chip="platform_coretemp_0",sensor="temp2"} 0
node_hwmon_temp_crit_alarm_celsius{chip="platform_coretemp_0",sensor="temp3"} 0
node_hwmon_temp_crit_alarm_celsius{chip="platform_coretemp_0",sensor="temp4"} 0
node_hwmon_temp_crit_alarm_celsius{chip="platform_coretemp_0",sensor="temp5"} 0
# HELP node_hwmon_temp_crit_celsius Hardware monitor for temperature (crit)
# TYPE node_hwmon_temp_crit_celsius gauge
node_hwmon_temp_crit_celsius{chip="nvme_nvme0",sensor="temp1"} 80.85000000000001
node_hwmon_temp_crit_celsius{chip="platform_coretemp_0",sensor="temp1"} 100
node_hwmon_temp_crit_celsius{chip="platform_coretemp_0",sensor="temp2"} 100
node_hwmon_temp_crit_celsius{chip="platform_coretemp_0",sensor="temp3"} 100
node_hwmon_temp_crit_celsius{chip="platform_coretemp_0",sensor="temp4"} 100
node_hwmon_temp_crit_celsius{chip="platform_coretemp_0",sensor="temp5"} 100
node_hwmon_temp_crit_celsius{chip="thermal_thermal_zone0",sensor="temp1"} 107
# HELP node_hwmon_temp_max_celsius Hardware monitor for temperature (max)
# TYPE node_hwmon_temp_max_celsius gauge
node_hwmon_temp_max_celsius{chip="nvme_nvme0",sensor="temp1"} 78.85000000000001
node_hwmon_temp_max_celsius{chip="nvme_nvme0",sensor="temp2"} 65261.85
node_hwmon_temp_max_celsius{chip="nvme_nvme0",sensor="temp3"} 65261.85
node_hwmon_temp_max_celsius{chip="platform_coretemp_0",sensor="temp1"} 100
node_hwmon_temp_max_celsius{chip="platform_coretemp_0",sensor="temp2"} 100
node_hwmon_temp_max_celsius{chip="platform_coretemp_0",sensor="temp3"} 100
node_hwmon_temp_max_celsius{chip="platform_coretemp_0",sensor="temp4"} 100
node_hwmon_temp_max_celsius{chip="platform_coretemp_0",sensor="temp5"} 100
# HELP node_hwmon_temp_min_celsius Hardware monitor for temperature (min)
# TYPE node_hwmon_temp_min_celsius gauge
node_hwmon_temp_min_celsius{chip="nvme_nvme0",sensor="temp1"} -0.15
node_hwmon_temp_min_celsius{chip="nvme_nvme0",sensor="temp2"} -273.15000000000003
node_hwmon_temp_min_celsius{chip="nvme_nvme0",sensor="temp3"} -273.15000000000003
# HELP node_intr_total Total number of interrupts serviced.
# TYPE node_intr_total counter
node_intr_total 2.78642008e+08
# HELP node_load1 1m load average.
# TYPE node_load1 gauge
node_load1 1.35
# HELP node_load15 15m load average.
# TYPE node_load15 gauge
node_load15 1.48
# HELP node_load5 5m load average.
# TYPE node_load5 gauge
node_load5 1.61
# HELP node_memory_Active_anon_bytes Memory information field Active_anon_bytes.
# TYPE node_memory_Active_anon_bytes gauge
node_memory_Active_anon_bytes 1.9144704e+07
# HELP node_memory_Active_bytes Memory information field Active_bytes.
# TYPE node_memory_Active_bytes gauge
node_memory_Active_bytes 4.096299008e+09
# HELP node_memory_Active_file_bytes Memory information field Active_file_bytes.
# TYPE node_memory_Active_file_bytes gauge
node_memory_Active_file_bytes 4.077154304e+09
# HELP node_memory_AnonHugePages_bytes Memory information field AnonHugePages_bytes.
# TYPE node_memory_AnonHugePages_bytes gauge
node_memory_AnonHugePages_bytes 2.097152e+06
# HELP node_memory_AnonPages_bytes Memory information field AnonPages_bytes.
# TYPE node_memory_AnonPages_bytes gauge
node_memory_AnonPages_bytes 4.347121664e+09
# HELP node_memory_Bounce_bytes Memory information field Bounce_bytes.
# TYPE node_memory_Bounce_bytes gauge
node_memory_Bounce_bytes 0
# HELP node_memory_Buffers_bytes Memory information field Buffers_bytes.
# TYPE node_memory_Buffers_bytes gauge
node_memory_Buffers_bytes 3.93560064e+08
# HELP node_memory_Cached_bytes Memory information field Cached_bytes.
# TYPE node_memory_Cached_bytes gauge
node_memory_Cached_bytes 1.0119647232e+10
# HELP node_memory_CommitLimit_bytes Memory information field CommitLimit_bytes.
# TYPE node_memory_CommitLimit_bytes gauge
node_memory_CommitLimit_bytes 8.20271104e+09
# HELP node_memory_Committed_AS_bytes Memory information field Committed_AS_bytes.
# TYPE node_memory_Committed_AS_bytes gauge
node_memory_Committed_AS_bytes 1.458753536e+10
# HELP node_memory_DirectMap1G_bytes Memory information field DirectMap1G_bytes.
# TYPE node_memory_DirectMap1G_bytes gauge
node_memory_DirectMap1G_bytes 3.221225472e+09
# HELP node_memory_DirectMap2M_bytes Memory information field DirectMap2M_bytes.
# TYPE node_memory_DirectMap2M_bytes gauge
node_memory_DirectMap2M_bytes 1.3092519936e+10
# HELP node_memory_DirectMap4k_bytes Memory information field DirectMap4k_bytes.
# TYPE node_memory_DirectMap4k_bytes gauge
node_memory_DirectMap4k_bytes 4.89394176e+08
# HELP node_memory_Dirty_bytes Memory information field Dirty_bytes.
# TYPE node_memory_Dirty_bytes gauge
node_memory_Dirty_bytes 905216
# HELP node_memory_FileHugePages_bytes Memory information field FileHugePages_bytes.
# TYPE node_memory_FileHugePages_bytes gauge
node_memory_FileHugePages_bytes 0
# HELP node_memory_FilePmdMapped_bytes Memory information field FilePmdMapped_bytes.
# TYPE node_memory_FilePmdMapped_bytes gauge
node_memory_FilePmdMapped_bytes 0
# HELP node_memory_HardwareCorrupted_bytes Memory information field HardwareCorrupted_bytes.
# TYPE node_memory_HardwareCorrupted_bytes gauge
node_memory_HardwareCorrupted_bytes 0
# HELP node_memory_HugePages_Free Memory information field HugePages_Free.
# TYPE node_memory_HugePages_Free gauge
node_memory_HugePages_Free 0
# HELP node_memory_HugePages_Rsvd Memory information field HugePages_Rsvd.
# TYPE node_memory_HugePages_Rsvd gauge
node_memory_HugePages_Rsvd 0
# HELP node_memory_HugePages_Surp Memory information field HugePages_Surp.
# TYPE node_memory_HugePages_Surp gauge
node_memory_HugePages_Surp 0
# HELP node_memory_HugePages_Total Memory information field HugePages_Total.
# TYPE node_memory_HugePages_Total gauge
node_memory_HugePages_Total 0
# HELP node_memory_Hugepagesize_bytes Memory information field Hugepagesize_bytes.
# TYPE node_memory_Hugepagesize_bytes gauge
node_memory_Hugepagesize_bytes 2.097152e+06
# HELP node_memory_Hugetlb_bytes Memory information field Hugetlb_bytes.
# TYPE node_memory_Hugetlb_bytes gauge
node_memory_Hugetlb_bytes 0
# HELP node_memory_Inactive_anon_bytes Memory information field Inactive_anon_bytes.
# TYPE node_memory_Inactive_anon_bytes gauge
node_memory_Inactive_anon_bytes 4.557848576e+09
# HELP node_memory_Inactive_bytes Memory information field Inactive_bytes.
# TYPE node_memory_Inactive_bytes gauge
node_memory_Inactive_bytes 9.843113984e+09
# HELP node_memory_Inactive_file_bytes Memory information field Inactive_file_bytes.
# TYPE node_memory_Inactive_file_bytes gauge
node_memory_Inactive_file_bytes 5.285265408e+09
# HELP node_memory_KReclaimable_bytes Memory information field KReclaimable_bytes.
# TYPE node_memory_KReclaimable_bytes gauge
node_memory_KReclaimable_bytes 8.33896448e+08
# HELP node_memory_KernelStack_bytes Memory information field KernelStack_bytes.
# TYPE node_memory_KernelStack_bytes gauge
node_memory_KernelStack_bytes 1.9808256e+07
# HELP node_memory_Mapped_bytes Memory information field Mapped_bytes.
# TYPE node_memory_Mapped_bytes gauge
node_memory_Mapped_bytes 1.300570112e+09
# HELP node_memory_MemAvailable_bytes Memory information field MemAvailable_bytes.
# TYPE node_memory_MemAvailable_bytes gauge
node_memory_MemAvailable_bytes 1.018525696e+10
# HELP node_memory_MemFree_bytes Memory information field MemFree_bytes.
# TYPE node_memory_MemFree_bytes gauge
node_memory_MemFree_bytes 3.41798912e+08
# HELP node_memory_MemTotal_bytes Memory information field MemTotal_bytes.
# TYPE node_memory_MemTotal_bytes gauge
node_memory_MemTotal_bytes 1.6405426176e+10
# HELP node_memory_Mlocked_bytes Memory information field Mlocked_bytes.
# TYPE node_memory_Mlocked_bytes gauge
node_memory_Mlocked_bytes 81920
# HELP node_memory_NFS_Unstable_bytes Memory information field NFS_Unstable_bytes.
# TYPE node_memory_NFS_Unstable_bytes gauge
node_memory_NFS_Unstable_bytes 0
# HELP node_memory_PageTables_bytes Memory information field PageTables_bytes.
# TYPE node_memory_PageTables_bytes gauge
node_memory_PageTables_bytes 5.0929664e+07
# HELP node_memory_Percpu_bytes Memory information field Percpu_bytes.
# TYPE node_memory_Percpu_bytes gauge
node_memory_Percpu_bytes 9.306112e+06
# HELP node_memory_SReclaimable_bytes Memory information field SReclaimable_bytes.
# TYPE node_memory_SReclaimable_bytes gauge
node_memory_SReclaimable_bytes 8.33896448e+08
# HELP node_memory_SUnreclaim_bytes Memory information field SUnreclaim_bytes.
# TYPE node_memory_SUnreclaim_bytes gauge
node_memory_SUnreclaim_bytes 1.73948928e+08
# HELP node_memory_ShmemHugePages_bytes Memory information field ShmemHugePages_bytes.
# TYPE node_memory_ShmemHugePages_bytes gauge
node_memory_ShmemHugePages_bytes 0
# HELP node_memory_ShmemPmdMapped_bytes Memory information field ShmemPmdMapped_bytes.
# TYPE node_memory_ShmemPmdMapped_bytes gauge
node_memory_ShmemPmdMapped_bytes 0
# HELP node_memory_Shmem_bytes Memory information field Shmem_bytes.
# TYPE node_memory_Shmem_bytes gauge
node_memory_Shmem_bytes 1.176133632e+09
# HELP node_memory_Slab_bytes Memory information field Slab_bytes.
# TYPE node_memory_Slab_bytes gauge
node_memory_Slab_bytes 1.007845376e+09
# HELP node_memory_SwapCached_bytes Memory information field SwapCached_bytes.
# TYPE node_memory_SwapCached_bytes gauge
node_memory_SwapCached_bytes 0
# HELP node_memory_SwapFree_bytes Memory information field SwapFree_bytes.
# TYPE node_memory_SwapFree_bytes gauge
node_memory_SwapFree_bytes 0
# HELP node_memory_SwapTotal_bytes Memory information field SwapTotal_bytes.
# TYPE node_memory_SwapTotal_bytes gauge
node_memory_SwapTotal_bytes 0
# HELP node_memory_Unevictable_bytes Memory information field Unevictable_bytes.
# TYPE node_memory_Unevictable_bytes gauge
node_memory_Unevictable_bytes 9.30336768e+08
# HELP node_memory_VmallocChunk_bytes Memory information field VmallocChunk_bytes.
# TYPE node_memory_VmallocChunk_bytes gauge
node_memory_VmallocChunk_bytes 0
# HELP node_memory_VmallocTotal_bytes Memory information field VmallocTotal_bytes.
# TYPE node_memory_VmallocTotal_bytes gauge
node_memory_VmallocTotal_bytes 3.5184372087808e+13
# HELP node_memory_VmallocUsed_bytes Memory information field VmallocUsed_bytes.
# TYPE node_memory_VmallocUsed_bytes gauge
node_memory_VmallocUsed_bytes 5.490688e+07
# HELP node_memory_WritebackTmp_bytes Memory information field WritebackTmp_bytes.
# TYPE node_memory_WritebackTmp_bytes gauge
node_memory_WritebackTmp_bytes 0
# HELP node_memory_Writeback_bytes Memory information field Writeback_bytes.
# TYPE node_memory_Writeback_bytes gauge
node_memory_Writeback_bytes 0
# HELP node_netstat_Icmp6_InErrors Statistic Icmp6InErrors.
# TYPE node_netstat_Icmp6_InErrors untyped
node_netstat_Icmp6_InErrors 5
# HELP node_netstat_Icmp6_InMsgs Statistic Icmp6InMsgs.
# TYPE node_netstat_Icmp6_InMsgs untyped
node_netstat_Icmp6_InMsgs 31134
# HELP node_netstat_Icmp6_OutMsgs Statistic Icmp6OutMsgs.
# TYPE node_netstat_Icmp6_OutMsgs untyped
node_netstat_Icmp6_OutMsgs 60585
# HELP node_netstat_Icmp_InErrors Statistic IcmpInErrors.
# TYPE node_netstat_Icmp_InErrors untyped
node_netstat_Icmp_InErrors 200
# HELP node_netstat_Icmp_InMsgs Statistic IcmpInMsgs.
# TYPE node_netstat_Icmp_InMsgs untyped
node_netstat_Icmp_InMsgs 311
# HELP node_netstat_Icmp_OutMsgs Statistic IcmpOutMsgs.
# TYPE node_netstat_Icmp_OutMsgs untyped
node_netstat_Icmp_OutMsgs 113
# HELP node_netstat_Ip6_InOctets Statistic Ip6InOctets.
# TYPE node_netstat_Ip6_InOctets untyped
node_netstat_Ip6_InOctets 1.796939309e+09
# HELP node_netstat_Ip6_OutOctets Statistic Ip6OutOctets.
# TYPE node_netstat_Ip6_OutOctets untyped
node_netstat_Ip6_OutOctets 5.3975087e+07
# HELP node_netstat_IpExt_InOctets Statistic IpExtInOctets.
# TYPE node_netstat_IpExt_InOctets untyped
node_netstat_IpExt_InOctets 1.052233883e+09
# HELP node_netstat_IpExt_OutOctets Statistic IpExtOutOctets.
# TYPE node_netstat_IpExt_OutOctets untyped
node_netstat_IpExt_OutOctets 6.7498903e+07
# HELP node_netstat_Ip_Forwarding Statistic IpForwarding.
# TYPE node_netstat_Ip_Forwarding untyped
node_netstat_Ip_Forwarding 1
# HELP node_netstat_TcpExt_ListenDrops Statistic TcpExtListenDrops.
# TYPE node_netstat_TcpExt_ListenDrops untyped
node_netstat_TcpExt_ListenDrops 0
# HELP node_netstat_TcpExt_ListenOverflows Statistic TcpExtListenOverflows.
# TYPE node_netstat_TcpExt_ListenOverflows untyped
node_netstat_TcpExt_ListenOverflows 0
# HELP node_netstat_TcpExt_SyncookiesFailed Statistic TcpExtSyncookiesFailed.
# TYPE node_netstat_TcpExt_SyncookiesFailed untyped
node_netstat_TcpExt_SyncookiesFailed 0
# HELP node_netstat_TcpExt_SyncookiesRecv Statistic TcpExtSyncookiesRecv.
# TYPE node_netstat_TcpExt_SyncookiesRecv untyped
node_netstat_TcpExt_SyncookiesRecv 0
# HELP node_netstat_TcpExt_SyncookiesSent Statistic TcpExtSyncookiesSent.
# TYPE node_netstat_TcpExt_SyncookiesSent untyped
node_netstat_TcpExt_SyncookiesSent 0
# HELP node_netstat_TcpExt_TCPSynRetrans Statistic TcpExtTCPSynRetrans.
# TYPE node_netstat_TcpExt_TCPSynRetrans untyped
node_netstat_TcpExt_TCPSynRetrans 995
# HELP node_netstat_Tcp_ActiveOpens Statistic TcpActiveOpens.
# TYPE node_netstat_Tcp_ActiveOpens untyped
node_netstat_Tcp_ActiveOpens 12514
# HELP node_netstat_Tcp_CurrEstab Statistic TcpCurrEstab.
# TYPE node_netstat_Tcp_CurrEstab untyped
node_netstat_Tcp_CurrEstab 16
# HELP node_netstat_Tcp_InErrs Statistic TcpInErrs.
# TYPE node_netstat_Tcp_InErrs untyped
node_netstat_Tcp_InErrs 117
# HELP node_netstat_Tcp_InSegs Statistic TcpInSegs.
# TYPE node_netstat_Tcp_InSegs untyped
node_netstat_Tcp_InSegs 568606
# HELP node_netstat_Tcp_OutSegs Statistic TcpOutSegs.
# TYPE node_netstat_Tcp_OutSegs untyped
node_netstat_Tcp_OutSegs 540924
# HELP node_netstat_Tcp_PassiveOpens Statistic TcpPassiveOpens.
# TYPE node_netstat_Tcp_PassiveOpens untyped
node_netstat_Tcp_PassiveOpens 230
# HELP node_netstat_Tcp_RetransSegs Statistic TcpRetransSegs.
# TYPE node_netstat_Tcp_RetransSegs untyped
node_netstat_Tcp_RetransSegs 2896
# HELP node_netstat_Udp6_InDatagrams Statistic Udp6InDatagrams.
# TYPE node_netstat_Udp6_InDatagrams untyped
node_netstat_Udp6_InDatagrams 575997
# HELP node_netstat_Udp6_InErrors Statistic Udp6InErrors.
# TYPE node_netstat_Udp6_InErrors untyped
node_netstat_Udp6_InErrors 15277
# HELP node_netstat_Udp6_NoPorts Statistic Udp6NoPorts.
# TYPE node_netstat_Udp6_NoPorts untyped
node_netstat_Udp6_NoPorts 0
# HELP node_netstat_Udp6_OutDatagrams Statistic Udp6OutDatagrams.
# TYPE node_netstat_Udp6_OutDatagrams untyped
node_netstat_Udp6_OutDatagrams 27775
# HELP node_netstat_Udp6_RcvbufErrors Statistic Udp6RcvbufErrors.
# TYPE node_netstat_Udp6_RcvbufErrors untyped
node_netstat_Udp6_RcvbufErrors 15277
# HELP node_netstat_Udp6_SndbufErrors Statistic Udp6SndbufErrors.
# TYPE node_netstat_Udp6_SndbufErrors untyped
node_netstat_Udp6_SndbufErrors 0
# HELP node_netstat_UdpLite6_InErrors Statistic UdpLite6InErrors.
# TYPE node_netstat_UdpLite6_InErrors untyped
node_netstat_UdpLite6_InErrors 0
# HELP node_netstat_UdpLite_InErrors Statistic UdpLiteInErrors.
# TYPE node_netstat_UdpLite_InErrors untyped
node_netstat_UdpLite_InErrors 0
# HELP node_netstat_Udp_InDatagrams Statistic UdpInDatagrams.
# TYPE node_netstat_Udp_InDatagrams untyped
node_netstat_Udp_InDatagrams 793628
# HELP node_netstat_Udp_InErrors Statistic UdpInErrors.
# TYPE node_netstat_Udp_InErrors untyped
node_netstat_Udp_InErrors 126
# HELP node_netstat_Udp_NoPorts Statistic UdpNoPorts.
# TYPE node_netstat_Udp_NoPorts untyped
node_netstat_Udp_NoPorts 62
# HELP node_netstat_Udp_OutDatagrams Statistic UdpOutDatagrams.
# TYPE node_netstat_Udp_OutDatagrams untyped
node_netstat_Udp_OutDatagrams 127582
# HELP node_netstat_Udp_RcvbufErrors Statistic UdpRcvbufErrors.
# TYPE node_netstat_Udp_RcvbufErrors untyped
node_netstat_Udp_RcvbufErrors 126
# HELP node_netstat_Udp_SndbufErrors Statistic UdpSndbufErrors.
# TYPE node_netstat_Udp_SndbufErrors untyped
node_netstat_Udp_SndbufErrors 1
# HELP node_network_address_assign_type address_assign_type value of /sys/class/net/<iface>.
# TYPE node_network_address_assign_type gauge
node_network_address_assign_type{device="br-c1aacbf10933"} 3
node_network_address_assign_type{device="docker0"} 3
node_network_address_assign_type{device="lo"} 0
node_network_address_assign_type{device="wlp2s0"} 0
# HELP node_network_carrier carrier value of /sys/class/net/<iface>.
# TYPE node_network_carrier gauge
node_network_carrier{device="br-c1aacbf10933"} 0
node_network_carrier{device="docker0"} 0
node_network_carrier{device="lo"} 1
node_network_carrier{device="wlp2s0"} 1
# HELP node_network_carrier_changes_total carrier_changes_total value of /sys/class/net/<iface>.
# TYPE node_network_carrier_changes_total counter
node_network_carrier_changes_total{device="br-c1aacbf10933"} 1
node_network_carrier_changes_total{device="docker0"} 1
node_network_carrier_changes_total{device="lo"} 0
node_network_carrier_changes_total{device="wlp2s0"} 30
# HELP node_network_carrier_down_changes_total carrier_down_changes_total value of /sys/class/net/<iface>.
# TYPE node_network_carrier_down_changes_total counter
node_network_carrier_down_changes_total{device="br-c1aacbf10933"} 1
node_network_carrier_down_changes_total{device="docker0"} 1
node_network_carrier_down_changes_total{device="lo"} 0
node_network_carrier_down_changes_total{device="wlp2s0"} 15
# HELP node_network_carrier_up_changes_total carrier_up_changes_total value of /sys/class/net/<iface>.
# TYPE node_network_carrier_up_changes_total counter
node_network_carrier_up_changes_total{device="br-c1aacbf10933"} 0
node_network_carrier_up_changes_total{device="docker0"} 0
node_network_carrier_up_changes_total{device="lo"} 0
node_network_carrier_up_changes_total{device="wlp2s0"} 15
# HELP node_network_device_id device_id value of /sys/class/net/<iface>.
# TYPE node_network_device_id gauge
node_network_device_id{device="br-c1aacbf10933"} 0
node_network_device_id{device="docker0"} 0
node_network_device_id{device="lo"} 0
node_network_device_id{device="wlp2s0"} 0
# HELP node_network_dormant dormant value of /sys/class/net/<iface>.
# TYPE node_network_dormant gauge
node_network_dormant{device="br-c1aacbf10933"} 0
node_network_dormant{device="docker0"} 0
node_network_dormant{device="lo"} 0
node_network_dormant{device="wlp2s0"} 0
# HELP node_network_flags flags value of /sys/class/net/<iface>.
# TYPE node_network_flags gauge
node_network_flags{device="br-c1aacbf10933"} 4099
node_network_flags{device="docker0"} 4099
node_network_flags{device="lo"} 9
node_network_flags{device="wlp2s0"} 4099
# HELP node_network_iface_id iface_id value of /sys/class/net/<iface>.
# TYPE node_network_iface_id gauge
node_network_iface_id{device="br-c1aacbf10933"} 4
node_network_iface_id{device="docker0"} 3
node_network_iface_id{device="lo"} 1
node_network_iface_id{device="wlp2s0"} 2
# HELP node_network_iface_link iface_link value of /sys/class/net/<iface>.
# TYPE node_network_iface_link gauge
node_network_iface_link{device="br-c1aacbf10933"} 4
node_network_iface_link{device="docker0"} 3
node_network_iface_link{device="lo"} 1
node_network_iface_link{device="wlp2s0"} 2
# HELP node_network_iface_link_mode iface_link_mode value of /sys/class/net/<iface>.
# TYPE node_network_iface_link_mode gauge
node_network_iface_link_mode{device="br-c1aacbf10933"} 0
node_network_iface_link_mode{device="docker0"} 0
node_network_iface_link_mode{device="lo"} 0
node_network_iface_link_mode{device="wlp2s0"} 1
# HELP node_network_info Non-numeric data from /sys/class/net/<iface>, value is always 1.
# TYPE node_network_info gauge
node_network_info{address="00:00:00:00:00:00",broadcast="00:00:00:00:00:00",device="lo",duplex="",ifalias="",operstate="unknown"} 1
node_network_info{address="02:42:86:a2:66:ce",broadcast="ff:ff:ff:ff:ff:ff",device="br-c1aacbf10933",duplex="unknown",ifalias="",operstate="down"} 1
node_network_info{address="02:42:f9:dc:7a:9e",broadcast="ff:ff:ff:ff:ff:ff",device="docker0",duplex="unknown",ifalias="",operstate="down"} 1
node_network_info{address="9c:b6:d0:98:41:b7",broadcast="ff:ff:ff:ff:ff:ff",device="wlp2s0",duplex="",ifalias="",operstate="up"} 1
# HELP node_network_mtu_bytes mtu_bytes value of /sys/class/net/<iface>.
# TYPE node_network_mtu_bytes gauge
node_network_mtu_bytes{device="br-c1aacbf10933"} 1500
node_network_mtu_bytes{device="docker0"} 1500
node_network_mtu_bytes{device="lo"} 65536
node_network_mtu_bytes{device="wlp2s0"} 1500
# HELP node_network_name_assign_type name_assign_type value of /sys/class/net/<iface>.
# TYPE node_network_name_assign_type gauge
node_network_name_assign_type{device="br-c1aacbf10933"} 3
node_network_name_assign_type{device="docker0"} 3
node_network_name_assign_type{device="wlp2s0"} 4
# HELP node_network_net_dev_group net_dev_group value of /sys/class/net/<iface>.
# TYPE node_network_net_dev_group gauge
node_network_net_dev_group{device="br-c1aacbf10933"} 0
node_network_net_dev_group{device="docker0"} 0
node_network_net_dev_group{device="lo"} 0
node_network_net_dev_group{device="wlp2s0"} 0
# HELP node_network_protocol_type protocol_type value of /sys/class/net/<iface>.
# TYPE node_network_protocol_type gauge
node_network_protocol_type{device="br-c1aacbf10933"} 1
node_network_protocol_type{device="docker0"} 1
node_network_protocol_type{device="lo"} 772
node_network_protocol_type{device="wlp2s0"} 1
# HELP node_network_receive_bytes_total Network device statistic receive_bytes.
# TYPE node_network_receive_bytes_total counter
node_network_receive_bytes_total{device="br-c1aacbf10933"} 0
node_network_receive_bytes_total{device="docker0"} 0
node_network_receive_bytes_total{device="lo"} 1.0347663e+07
node_network_receive_bytes_total{device="wlp2s0"} 2.966265253e+09
# HELP node_network_receive_compressed_total Network device statistic receive_compressed.
# TYPE node_network_receive_compressed_total counter
node_network_receive_compressed_total{device="br-c1aacbf10933"} 0
node_network_receive_compressed_total{device="docker0"} 0
node_network_receive_compressed_total{device="lo"} 0
node_network_receive_compressed_total{device="wlp2s0"} 0
# HELP node_network_receive_drop_total Network device statistic receive_drop.
# TYPE node_network_receive_drop_total counter
node_network_receive_drop_total{device="br-c1aacbf10933"} 0
node_network_receive_drop_total{device="docker0"} 0
node_network_receive_drop_total{device="lo"} 0
node_network_receive_drop_total{device="wlp2s0"} 0
# HELP node_network_receive_errs_total Network device statistic receive_errs.
# TYPE node_network_receive_errs_total counter
node_network_receive_errs_total{device="br-c1aacbf10933"} 0
node_network_receive_errs_total{device="docker0"} 0
node_network_receive_errs_total{device="lo"} 0
node_network_receive_errs_total{device="wlp2s0"} 0
# HELP node_network_receive_fifo_total Network device statistic receive_fifo.
# TYPE node_network_receive_fifo_total counter
node_network_receive_fifo_total{device="br-c1aacbf10933"} 0
node_network_receive_fifo_total{device="docker0"} 0
node_network_receive_fifo_total{device="lo"} 0
node_network_receive_fifo_total{device="wlp2s0"} 0
# HELP node_network_receive_frame_total Network device statistic receive_frame.
# TYPE node_network_receive_frame_total counter
node_network_receive_frame_total{device="br-c1aacbf10933"} 0
node_network_receive_frame_total{device="docker0"} 0
node_network_receive_frame_total{device="lo"} 0
node_network_receive_frame_total{device="wlp2s0"} 0
# HELP node_network_receive_multicast_total Network device statistic receive_multicast.
# TYPE node_network_receive_multicast_total counter
node_network_receive_multicast_total{device="br-c1aacbf10933"} 0
node_network_receive_multicast_total{device="docker0"} 0
node_network_receive_multicast_total{device="lo"} 0
node_network_receive_multicast_total{device="wlp2s0"} 0
# HELP node_network_receive_packets_total Network device statistic receive_packets.
# TYPE node_network_receive_packets_total counter
node_network_receive_packets_total{device="br-c1aacbf10933"} 0
node_network_receive_packets_total{device="docker0"} 0
node_network_receive_packets_total{device="lo"} 84160
node_network_receive_packets_total{device="wlp2s0"} 2.480574e+06
# HELP node_network_speed_bytes speed_bytes value of /sys/class/net/<iface>.
# TYPE node_network_speed_bytes gauge
node_network_speed_bytes{device="br-c1aacbf10933"} -125000
node_network_speed_bytes{device="docker0"} -125000
# HELP node_network_transmit_bytes_total Network device statistic transmit_bytes.
# TYPE node_network_transmit_bytes_total counter
node_network_transmit_bytes_total{device="br-c1aacbf10933"} 0
node_network_transmit_bytes_total{device="docker0"} 0
node_network_transmit_bytes_total{device="lo"} 1.0347663e+07
node_network_transmit_bytes_total{device="wlp2s0"} 1.33755756e+08
# HELP node_network_transmit_carrier_total Network device statistic transmit_carrier.
# TYPE node_network_transmit_carrier_total counter
node_network_transmit_carrier_total{device="br-c1aacbf10933"} 0
node_network_transmit_carrier_total{device="docker0"} 0
node_network_transmit_carrier_total{device="lo"} 0
node_network_transmit_carrier_total{device="wlp2s0"} 0
# HELP node_network_transmit_colls_total Network device statistic transmit_colls.
# TYPE node_network_transmit_colls_total counter
node_network_transmit_colls_total{device="br-c1aacbf10933"} 0
node_network_transmit_colls_total{device="docker0"} 0
node_network_transmit_colls_total{device="lo"} 0
node_network_transmit_colls_total{device="wlp2s0"} 0
# HELP node_network_transmit_compressed_total Network device statistic transmit_compressed.
# TYPE node_network_transmit_compressed_total counter
node_network_transmit_compressed_total{device="br-c1aacbf10933"} 0
node_network_transmit_compressed_total{device="docker0"} 0
node_network_transmit_compressed_total{device="lo"} 0
node_network_transmit_compressed_total{device="wlp2s0"} 0
# HELP node_network_transmit_drop_total Network device statistic transmit_drop.
# TYPE node_network_transmit_drop_total counter
node_network_transmit_drop_total{device="br-c1aacbf10933"} 0
node_network_transmit_drop_total{device="docker0"} 0
node_network_transmit_drop_total{device="lo"} 0
node_network_transmit_drop_total{device="wlp2s0"} 0
# HELP node_network_transmit_errs_total Network device statistic transmit_errs.
# TYPE node_network_transmit_errs_total counter
node_network_transmit_errs_total{device="br-c1aacbf10933"} 0
node_network_transmit_errs_total{device="docker0"} 0
node_network_transmit_errs_total{device="lo"} 0
node_network_transmit_errs_total{device="wlp2s0"} 0
# HELP node_network_transmit_fifo_total Network device statistic transmit_fifo.
# TYPE node_network_transmit_fifo_total counter
node_network_transmit_fifo_total{device="br-c1aacbf10933"} 0
node_network_transmit_fifo_total{device="docker0"} 0
node_network_transmit_fifo_total{device="lo"} 0
node_network_transmit_fifo_total{device="wlp2s0"} 0
# HELP node_network_transmit_packets_total Network device statistic transmit_packets.
# TYPE node_network_transmit_packets_total counter
node_network_transmit_packets_total{device="br-c1aacbf10933"} 0
node_network_transmit_packets_total{device="docker0"} 0
node_network_transmit_packets_total{device="lo"} 84160
node_network_transmit_packets_total{device="wlp2s0"} 666512
# HELP node_network_transmit_queue_length transmit_queue_length value of /sys/class/net/<iface>.
# TYPE node_network_transmit_queue_length gauge
node_network_transmit_queue_length{device="br-c1aacbf10933"} 0
node_network_transmit_queue_length{device="docker0"} 0
node_network_transmit_queue_length{device="lo"} 1000
node_network_transmit_queue_length{device="wlp2s0"} 1000
# HELP node_network_up Value is 1 if operstate is 'up', 0 otherwise.
# TYPE node_network_up gauge
node_network_up{device="br-c1aacbf10933"} 0
node_network_up{device="docker0"} 0
node_network_up{device="lo"} 0
node_network_up{device="wlp2s0"} 1
# HELP node_nf_conntrack_entries Number of currently allocated flow entries for connection tracking.
# TYPE node_nf_conntrack_entries gauge
node_nf_conntrack_entries 86
# HELP node_nf_conntrack_entries_limit Maximum size of connection tracking table.
# TYPE node_nf_conntrack_entries_limit gauge
node_nf_conntrack_entries_limit 262144
# HELP node_power_supply_capacity capacity value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_capacity gauge
node_power_supply_capacity{power_supply="BAT0"} 100
# HELP node_power_supply_charge_ampere charge_ampere value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_charge_ampere gauge
node_power_supply_charge_ampere{power_supply="BAT0"} 5.978
# HELP node_power_supply_charge_full charge_full value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_charge_full gauge
node_power_supply_charge_full{power_supply="BAT0"} 5.978
# HELP node_power_supply_charge_full_design charge_full_design value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_charge_full_design gauge
node_power_supply_charge_full_design{power_supply="BAT0"} 6.842
# HELP node_power_supply_current_ampere current_ampere value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_current_ampere gauge
node_power_supply_current_ampere{power_supply="BAT0"} 0.001
node_power_supply_current_ampere{power_supply="ucsi-source-psy-USBC000:001"} 0
node_power_supply_current_ampere{power_supply="ucsi-source-psy-USBC000:002"} 0
node_power_supply_current_ampere{power_supply="ucsi-source-psy-USBC000:003"} 2.25
# HELP node_power_supply_current_max current_max value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_current_max gauge
node_power_supply_current_max{power_supply="ucsi-source-psy-USBC000:001"} 0
node_power_supply_current_max{power_supply="ucsi-source-psy-USBC000:002"} 0
node_power_supply_current_max{power_supply="ucsi-source-psy-USBC000:003"} 2.25
# HELP node_power_supply_cyclecount cyclecount value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_cyclecount gauge
node_power_supply_cyclecount{power_supply="BAT0"} 0
# HELP node_power_supply_info info of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_info gauge
node_power_supply_info{power_supply="AC",type="Mains"} 1
node_power_supply_info{power_supply="ucsi-source-psy-USBC000:001",type="USB",usb_type="[C] PD PD_PPS"} 1
node_power_supply_info{power_supply="ucsi-source-psy-USBC000:002",type="USB",usb_type="[C] PD PD_PPS"} 1
node_power_supply_info{power_supply="ucsi-source-psy-USBC000:003",type="USB",usb_type="C [PD] PD_PPS"} 1
node_power_supply_info{capacity_level="Full",manufacturer="LGC-LGC6.73",model_name="DELL H754V91",power_supply="BAT0",serial_number="10239",status="Full",technology="Li-ion",type="Battery"} 1
# HELP node_power_supply_online online value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_online gauge
node_power_supply_online{power_supply="AC"} 1
node_power_supply_online{power_supply="ucsi-source-psy-USBC000:001"} 0
node_power_supply_online{power_supply="ucsi-source-psy-USBC000:002"} 0
node_power_supply_online{power_supply="ucsi-source-psy-USBC000:003"} 1
# HELP node_power_supply_present present value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_present gauge
node_power_supply_present{power_supply="BAT0"} 1
# HELP node_power_supply_voltage_max voltage_max value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_voltage_max gauge
node_power_supply_voltage_max{power_supply="ucsi-source-psy-USBC000:001"} 5
node_power_supply_voltage_max{power_supply="ucsi-source-psy-USBC000:002"} 5
node_power_supply_voltage_max{power_supply="ucsi-source-psy-USBC000:003"} 20
# HELP node_power_supply_voltage_min voltage_min value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_voltage_min gauge
node_power_supply_voltage_min{power_supply="ucsi-source-psy-USBC000:001"} 5
node_power_supply_voltage_min{power_supply="ucsi-source-psy-USBC000:002"} 5
node_power_supply_voltage_min{power_supply="ucsi-source-psy-USBC000:003"} 5
# HELP node_power_supply_voltage_min_design voltage_min_design value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_voltage_min_design gauge
node_power_supply_voltage_min_design{power_supply="BAT0"} 7.6
# HELP node_power_supply_voltage_volt voltage_volt value of /sys/class/power_supply/<power_supply>.
# TYPE node_power_supply_voltage_volt gauge
node_power_supply_voltage_volt{power_supply="BAT0"} 8.647
node_power_supply_voltage_volt{power_supply="ucsi-source-psy-USBC000:001"} 5
node_power_supply_voltage_volt{power_supply="ucsi-source-psy-USBC000:002"} 5
node_power_supply_voltage_volt{power_supply="ucsi-source-psy-USBC000:003"} 20
# HELP node_pressure_cpu_waiting_seconds_total Total time in seconds that processes have waited for CPU time
# TYPE node_pressure_cpu_waiting_seconds_total counter
node_pressure_cpu_waiting_seconds_total 1184.506429
# HELP node_pressure_io_stalled_seconds_total Total time in seconds no process could make progress due to IO congestion
# TYPE node_pressure_io_stalled_seconds_total counter
node_pressure_io_stalled_seconds_total 186.537609
# HELP node_pressure_io_waiting_seconds_total Total time in seconds that processes have waited due to IO congestion
# TYPE node_pressure_io_waiting_seconds_total counter
node_pressure_io_waiting_seconds_total 205.363775
# HELP node_pressure_memory_stalled_seconds_total Total time in seconds no process could make progress due to memory congestion
# TYPE node_pressure_memory_stalled_seconds_total counter
node_pressure_memory_stalled_seconds_total 0.18931
# HELP node_pressure_memory_waiting_seconds_total Total time in seconds that processes have waited for memory
# TYPE node_pressure_memory_waiting_seconds_total counter
node_pressure_memory_waiting_seconds_total 0.21434
# HELP node_procs_blocked Number of processes blocked waiting for I/O to complete.
# TYPE node_procs_blocked gauge
node_procs_blocked 0
# HELP node_procs_running Number of processes in runnable state.
# TYPE node_procs_running gauge
node_procs_running 15
# HELP node_schedstat_running_seconds_total Number of seconds CPU spent running a process.
# TYPE node_schedstat_running_seconds_total counter
node_schedstat_running_seconds_total{cpu="0"} 8367.546013203
node_schedstat_running_seconds_total{cpu="1"} 8779.625482945
node_schedstat_running_seconds_total{cpu="2"} 8483.501914414
node_schedstat_running_seconds_total{cpu="3"} 8579.478600437
node_schedstat_running_seconds_total{cpu="4"} 8208.399226906
node_schedstat_running_seconds_total{cpu="5"} 8150.364406191
node_schedstat_running_seconds_total{cpu="6"} 8819.178879814
node_schedstat_running_seconds_total{cpu="7"} 8304.991088246
# HELP node_schedstat_timeslices_total Number of timeslices executed by CPU.
# TYPE node_schedstat_timeslices_total counter
node_schedstat_timeslices_total{cpu="0"} 3.6018372e+07
node_schedstat_timeslices_total{cpu="1"} 3.9830578e+07
node_schedstat_timeslices_total{cpu="2"} 3.5520383e+07
node_schedstat_timeslices_total{cpu="3"} 3.6529128e+07
node_schedstat_timeslices_total{cpu="4"} 3.6492172e+07
node_schedstat_timeslices_total{cpu="5"} 4.7452046e+07
node_schedstat_timeslices_total{cpu="6"} 3.84226e+07
node_schedstat_timeslices_total{cpu="7"} 3.5749898e+07
# HELP node_schedstat_waiting_seconds_total Number of seconds spent by processing waiting for this CPU.
# TYPE node_schedstat_waiting_seconds_total counter
node_schedstat_waiting_seconds_total{cpu="0"} 238.162047506
node_schedstat_waiting_seconds_total{cpu="1"} 281.071365466
node_schedstat_waiting_seconds_total{cpu="2"} 231.701570019
node_schedstat_waiting_seconds_total{cpu="3"} 256.912884535
node_schedstat_waiting_seconds_total{cpu="4"} 240.79591714
node_schedstat_waiting_seconds_total{cpu="5"} 360.025331095
node_schedstat_waiting_seconds_total{cpu="6"} 289.094068091
node_schedstat_waiting_seconds_total{cpu="7"} 257.390036257
# HELP node_scrape_collector_duration_seconds node_exporter: Duration of a collector scrape.
# TYPE node_scrape_collector_duration_seconds gauge
node_scrape_collector_duration_seconds{collector="arp"} 0.000807492
node_scrape_collector_duration_seconds{collector="bcache"} 5.8869e-05
node_scrape_collector_duration_seconds{collector="bonding"} 6.413e-05
node_scrape_collector_duration_seconds{collector="btrfs"} 0.000491615
node_scrape_collector_duration_seconds{collector="conntrack"} 0.000125947
node_scrape_collector_duration_seconds{collector="cpu"} 0.002046468
node_scrape_collector_duration_seconds{collector="cpufreq"} 0.025466964
node_scrape_collector_duration_seconds{collector="diskstats"} 0.000869965
node_scrape_collector_duration_seconds{collector="edac"} 7.9918e-05
node_scrape_collector_duration_seconds{collector="entropy"} 0.000875119
node_scrape_collector_duration_seconds{collector="filefd"} 6.6079e-05
node_scrape_collector_duration_seconds{collector="filesystem"} 0.000879219
node_scrape_collector_duration_seconds{collector="hwmon"} 0.066628277
node_scrape_collector_duration_seconds{collector="infiniband"} 2.3705e-05
node_scrape_collector_duration_seconds{collector="ipvs"} 6.7042e-05
node_scrape_collector_duration_seconds{collector="loadavg"} 0.000125964
node_scrape_collector_duration_seconds{collector="mdadm"} 6.3023e-05
node_scrape_collector_duration_seconds{collector="meminfo"} 0.000505476
node_scrape_collector_duration_seconds{collector="netclass"} 0.009034155
node_scrape_collector_duration_seconds{collector="netdev"} 0.000380078
node_scrape_collector_duration_seconds{collector="netstat"} 0.003080062
node_scrape_collector_duration_seconds{collector="nfs"} 3.5682e-05
node_scrape_collector_duration_seconds{collector="nfsd"} 0.000984803
node_scrape_collector_duration_seconds{collector="powersupplyclass"} 0.033197049
node_scrape_collector_duration_seconds{collector="pressure"} 0.000456751
node_scrape_collector_duration_seconds{collector="rapl"} 0.002369126
node_scrape_collector_duration_seconds{collector="schedstat"} 0.001009569
node_scrape_collector_duration_seconds{collector="sockstat"} 0.001080297
node_scrape_collector_duration_seconds{collector="softnet"} 0.000175264
node_scrape_collector_duration_seconds{collector="stat"} 0.001533125
node_scrape_collector_duration_seconds{collector="textfile"} 2.8733e-05
node_scrape_collector_duration_seconds{collector="thermal_zone"} 0.041344967
node_scrape_collector_duration_seconds{collector="time"} 2.2438e-05
node_scrape_collector_duration_seconds{collector="timex"} 2.4433e-05
node_scrape_collector_duration_seconds{collector="udp_queues"} 0.001056615
node_scrape_collector_duration_seconds{collector="uname"} 1.7594e-05
node_scrape_collector_duration_seconds{collector="vmstat"} 0.003503415
node_scrape_collector_duration_seconds{collector="xfs"} 0.000214515
node_scrape_collector_duration_seconds{collector="zfs"} 0.001107325
# HELP node_scrape_collector_success node_exporter: Whether a collector succeeded.
# TYPE node_scrape_collector_success gauge
node_scrape_collector_success{collector="arp"} 1
node_scrape_collector_success{collector="bcache"} 1
node_scrape_collector_success{collector="bonding"} 0
node_scrape_collector_success{collector="btrfs"} 1
node_scrape_collector_success{collector="conntrack"} 1
node_scrape_collector_success{collector="cpu"} 1
node_scrape_collector_success{collector="cpufreq"} 1
node_scrape_collector_success{collector="diskstats"} 1
node_scrape_collector_success{collector="edac"} 1
node_scrape_collector_success{collector="entropy"} 1
node_scrape_collector_success{collector="filefd"} 1
node_scrape_collector_success{collector="filesystem"} 1
node_scrape_collector_success{collector="hwmon"} 1
node_scrape_collector_success{collector="infiniband"} 0
node_scrape_collector_success{collector="ipvs"} 0
node_scrape_collector_success{collector="loadavg"} 1
node_scrape_collector_success{collector="mdadm"} 1
node_scrape_collector_success{collector="meminfo"} 1
node_scrape_collector_success{collector="netclass"} 1
node_scrape_collector_success{collector="netdev"} 1
node_scrape_collector_success{collector="netstat"} 1
node_scrape_collector_success{collector="nfs"} 0
node_scrape_collector_success{collector="nfsd"} 0
node_scrape_collector_success{collector="powersupplyclass"} 1
node_scrape_collector_success{collector="pressure"} 1
node_scrape_collector_success{collector="rapl"} 0
node_scrape_collector_success{collector="schedstat"} 1
node_scrape_collector_success{collector="sockstat"} 1
node_scrape_collector_success{collector="softnet"} 1
node_scrape_collector_success{collector="stat"} 1
node_scrape_collector_success{collector="textfile"} 1
node_scrape_collector_success{collector="thermal_zone"} 1
node_scrape_collector_success{collector="time"} 1
node_scrape_collector_success{collector="timex"} 1
node_scrape_collector_success{collector="udp_queues"} 1
node_scrape_collector_success{collector="uname"} 1
node_scrape_collector_success{collector="vmstat"} 1
node_scrape_collector_success{collector="xfs"} 1
node_scrape_collector_success{collector="zfs"} 1
# HELP node_sockstat_FRAG6_inuse Number of FRAG6 sockets in state inuse.
# TYPE node_sockstat_FRAG6_inuse gauge
node_sockstat_FRAG6_inuse 0
# HELP node_sockstat_FRAG6_memory Number of FRAG6 sockets in state memory.
# TYPE node_sockstat_FRAG6_memory gauge
node_sockstat_FRAG6_memory 0
# HELP node_sockstat_FRAG_inuse Number of FRAG sockets in state inuse.
# TYPE node_sockstat_FRAG_inuse gauge
node_sockstat_FRAG_inuse 0
# HELP node_sockstat_FRAG_memory Number of FRAG sockets in state memory.
# TYPE node_sockstat_FRAG_memory gauge
node_sockstat_FRAG_memory 0
# HELP node_sockstat_RAW6_inuse Number of RAW6 sockets in state inuse.
# TYPE node_sockstat_RAW6_inuse gauge
node_sockstat_RAW6_inuse 1
# HELP node_sockstat_RAW_inuse Number of RAW sockets in state inuse.
# TYPE node_sockstat_RAW_inuse gauge
node_sockstat_RAW_inuse 0
# HELP node_sockstat_TCP6_inuse Number of TCP6 sockets in state inuse.
# TYPE node_sockstat_TCP6_inuse gauge
node_sockstat_TCP6_inuse 8
# HELP node_sockstat_TCP_alloc Number of TCP sockets in state alloc.
# TYPE node_sockstat_TCP_alloc gauge
node_sockstat_TCP_alloc 26
# HELP node_sockstat_TCP_inuse Number of TCP sockets in state inuse.
# TYPE node_sockstat_TCP_inuse gauge
node_sockstat_TCP_inuse 16
# HELP node_sockstat_TCP_mem Number of TCP sockets in state mem.
# TYPE node_sockstat_TCP_mem gauge
node_sockstat_TCP_mem 7
# HELP node_sockstat_TCP_mem_bytes Number of TCP sockets in state mem_bytes.
# TYPE node_sockstat_TCP_mem_bytes gauge
node_sockstat_TCP_mem_bytes 28672
# HELP node_sockstat_TCP_orphan Number of TCP sockets in state orphan.
# TYPE node_sockstat_TCP_orphan gauge
node_sockstat_TCP_orphan 0
# HELP node_sockstat_TCP_tw Number of TCP sockets in state tw.
# TYPE node_sockstat_TCP_tw gauge
node_sockstat_TCP_tw 2
# HELP node_sockstat_UDP6_inuse Number of UDP6 sockets in state inuse.
# TYPE node_sockstat_UDP6_inuse gauge
node_sockstat_UDP6_inuse 11
# HELP node_sockstat_UDPLITE6_inuse Number of UDPLITE6 sockets in state inuse.
# TYPE node_sockstat_UDPLITE6_inuse gauge
node_sockstat_UDPLITE6_inuse 0
# HELP node_sockstat_UDPLITE_inuse Number of UDPLITE sockets in state inuse.
# TYPE node_sockstat_UDPLITE_inuse gauge
node_sockstat_UDPLITE_inuse 0
# HELP node_sockstat_UDP_inuse Number of UDP sockets in state inuse.
# TYPE node_sockstat_UDP_inuse gauge
node_sockstat_UDP_inuse 16
# HELP node_sockstat_UDP_mem Number of UDP sockets in state mem.
# TYPE node_sockstat_UDP_mem gauge
node_sockstat_UDP_mem 21
# HELP node_sockstat_UDP_mem_bytes Number of UDP sockets in state mem_bytes.
# TYPE node_sockstat_UDP_mem_bytes gauge
node_sockstat_UDP_mem_bytes 86016
# HELP node_sockstat_sockets_used Number of IPv4 sockets in use.
# TYPE node_sockstat_sockets_used gauge
node_sockstat_sockets_used 1288
# HELP node_softnet_dropped_total Number of dropped packets
# TYPE node_softnet_dropped_total counter
node_softnet_dropped_total{cpu="0"} 0
node_softnet_dropped_total{cpu="1"} 0
node_softnet_dropped_total{cpu="2"} 0
node_softnet_dropped_total{cpu="3"} 0
node_softnet_dropped_total{cpu="4"} 0
node_softnet_dropped_total{cpu="5"} 0
node_softnet_dropped_total{cpu="6"} 0
node_softnet_dropped_total{cpu="7"} 0
# HELP node_softnet_processed_total Number of processed packets
# TYPE node_softnet_processed_total counter
node_softnet_processed_total{cpu="0"} 15453
node_softnet_processed_total{cpu="1"} 13762
node_softnet_processed_total{cpu="2"} 13462
node_softnet_processed_total{cpu="3"} 14108
node_softnet_processed_total{cpu="4"} 12756
node_softnet_processed_total{cpu="5"} 12692
node_softnet_processed_total{cpu="6"} 13475
node_softnet_processed_total{cpu="7"} 1.011975e+06
# HELP node_softnet_times_squeezed_total Number of times processing packets ran out of quota
# TYPE node_softnet_times_squeezed_total counter
node_softnet_times_squeezed_total{cpu="0"} 0
node_softnet_times_squeezed_total{cpu="1"} 0
node_softnet_times_squeezed_total{cpu="2"} 0
node_softnet_times_squeezed_total{cpu="3"} 0
node_softnet_times_squeezed_total{cpu="4"} 0
node_softnet_times_squeezed_total{cpu="5"} 0
node_softnet_times_squeezed_total{cpu="6"} 0
node_softnet_times_squeezed_total{cpu="7"} 0
# HELP node_textfile_scrape_error 1 if there was an error opening or reading a file, 0 otherwise
# TYPE node_textfile_scrape_error gauge
node_textfile_scrape_error 0
# HELP node_thermal_zone_temp Zone temperature in Celsius
# TYPE node_thermal_zone_temp gauge
node_thermal_zone_temp{type="B0D4",zone="9"} 50.05
node_thermal_zone_temp{type="INT3400 Thermal",zone="4"} 20
node_thermal_zone_temp{type="SEN1",zone="1"} 43.05
node_thermal_zone_temp{type="SEN2",zone="2"} 36.05
node_thermal_zone_temp{type="SEN3",zone="3"} 44.05
node_thermal_zone_temp{type="SEN4",zone="5"} 42.05
node_thermal_zone_temp{type="TMEM",zone="7"} 41.05
node_thermal_zone_temp{type="TSKN",zone="6"} 40.05
node_thermal_zone_temp{type="acpitz",zone="0"} 25
node_thermal_zone_temp{type="pch_cannonlake",zone="8"} 48
node_thermal_zone_temp{type="x86_pkg_temp",zone="10"} 51
# HELP node_time_seconds System time in seconds since epoch (1970).
# TYPE node_time_seconds gauge
node_time_seconds 1.6383328421125576e+09
# HELP node_timex_estimated_error_seconds Estimated error in seconds.
# TYPE node_timex_estimated_error_seconds gauge
node_timex_estimated_error_seconds 0
# HELP node_timex_frequency_adjustment_ratio Local clock frequency adjustment.
# TYPE node_timex_frequency_adjustment_ratio gauge
node_timex_frequency_adjustment_ratio 1.0003692013397216
# HELP node_timex_loop_time_constant Phase-locked loop time constant.
# TYPE node_timex_loop_time_constant gauge
node_timex_loop_time_constant 3
# HELP node_timex_maxerror_seconds Maximum error in seconds.
# TYPE node_timex_maxerror_seconds gauge
node_timex_maxerror_seconds 0.0435
# HELP node_timex_offset_seconds Time offset in between local system and reference clock.
# TYPE node_timex_offset_seconds gauge
node_timex_offset_seconds -0.001531009
# HELP node_timex_pps_calibration_total Pulse per second count of calibration intervals.
# TYPE node_timex_pps_calibration_total counter
node_timex_pps_calibration_total 0
# HELP node_timex_pps_error_total Pulse per second count of calibration errors.
# TYPE node_timex_pps_error_total counter
node_timex_pps_error_total 0
# HELP node_timex_pps_frequency_hertz Pulse per second frequency.
# TYPE node_timex_pps_frequency_hertz gauge
node_timex_pps_frequency_hertz 0
# HELP node_timex_pps_jitter_seconds Pulse per second jitter.
# TYPE node_timex_pps_jitter_seconds gauge
node_timex_pps_jitter_seconds 0
# HELP node_timex_pps_jitter_total Pulse per second count of jitter limit exceeded events.
# TYPE node_timex_pps_jitter_total counter
node_timex_pps_jitter_total 0
# HELP node_timex_pps_shift_seconds Pulse per second interval duration.
# TYPE node_timex_pps_shift_seconds gauge
node_timex_pps_shift_seconds 0
# HELP node_timex_pps_stability_exceeded_total Pulse per second count of stability limit exceeded events.
# TYPE node_timex_pps_stability_exceeded_total counter
node_timex_pps_stability_exceeded_total 0
# HELP node_timex_pps_stability_hertz Pulse per second stability, average of recent frequency changes.
# TYPE node_timex_pps_stability_hertz gauge
node_timex_pps_stability_hertz 0
# HELP node_timex_status Value of the status array bits.
# TYPE node_timex_status gauge
node_timex_status 8193
# HELP node_timex_sync_status Is clock synchronized to a reliable server (1 = yes, 0 = no).
# TYPE node_timex_sync_status gauge
node_timex_sync_status 1
# HELP node_timex_tai_offset_seconds International Atomic Time (TAI) offset.
# TYPE node_timex_tai_offset_seconds gauge
node_timex_tai_offset_seconds 0
# HELP node_timex_tick_seconds Seconds between clock ticks.
# TYPE node_timex_tick_seconds gauge
node_timex_tick_seconds 0.01
# HELP node_udp_queues Number of allocated memory in the kernel for UDP datagrams in bytes.
# TYPE node_udp_queues gauge
node_udp_queues{ip="v4",queue="rx"} 0
node_udp_queues{ip="v4",queue="tx"} 0
node_udp_queues{ip="v6",queue="rx"} 0
node_udp_queues{ip="v6",queue="tx"} 0
# HELP node_uname_info Labeled system information as provided by the uname system call.
# TYPE node_uname_info gauge
node_uname_info{domainname="(none)",machine="x86_64",nodename="plank",release="5.13.0-7614-generic",sysname="Linux",version="#14~1631647151~21.04~930e87c-Ubuntu SMP Fri Sep 17 00:24:58 UTC "} 1
# HELP node_vmstat_oom_kill /proc/vmstat information field oom_kill.
# TYPE node_vmstat_oom_kill untyped
node_vmstat_oom_kill 0
# HELP node_vmstat_pgfault /proc/vmstat information field pgfault.
# TYPE node_vmstat_pgfault untyped
node_vmstat_pgfault 3.03454016e+08
# HELP node_vmstat_pgmajfault /proc/vmstat information field pgmajfault.
# TYPE node_vmstat_pgmajfault untyped
node_vmstat_pgmajfault 22041
# HELP node_vmstat_pgpgin /proc/vmstat information field pgpgin.
# TYPE node_vmstat_pgpgin untyped
node_vmstat_pgpgin 5.076415e+06
# HELP node_vmstat_pgpgout /proc/vmstat information field pgpgout.
# TYPE node_vmstat_pgpgout untyped
node_vmstat_pgpgout 3.1315204e+07
# HELP node_vmstat_pswpin /proc/vmstat information field pswpin.
# TYPE node_vmstat_pswpin untyped
node_vmstat_pswpin 0
# HELP node_vmstat_pswpout /proc/vmstat information field pswpout.
# TYPE node_vmstat_pswpout untyped
node_vmstat_pswpout 0
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 2.67
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1.048576e+06
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 9
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 9.89184e+06
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.63529809911e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 7.35645696e+08
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes -1
# HELP promhttp_metric_handler_errors_total Total number of internal errors encountered by the promhttp metric handler.
# TYPE promhttp_metric_handler_errors_total counter
promhttp_metric_handler_errors_total{cause="encoding"} 0
promhttp_metric_handler_errors_total{cause="gathering"} 0
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 2
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
`

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	done := make(chan struct{})
	defer close(done)

	fams, err := scrape.Parse([]byte(nodeExporterData))
	if err != nil {
		panic(err)
	}

	allFams := make([]*dto.MetricFamily, 0)

	for _, v := range fams {
		allFams = append(allFams, v)
	}

	gen := types.NewGenerator(os.Stdout, types.Chaotic, allFams).WithStepDuration(1 * time.Minute).WithGaugeVariance(1.0)

	gen.WriteOpenMetrics(done, time.Now().Add(- 12 * time.Hour), time.Now().Add(- 2 * time.Hour))
		
}