// +build !windows

package diskinfo

import (
	"syscall"
	"log"
)

func GetDiskInfo(partition string) DiskInfo {
	var di DiskInfo

	fs := syscall.Statfs_t{}
	err := syscall.Statfs(partition, &fs)
	if err != nil {
		log.Panic(err)
	}

	di.Total = fs.Blocks * uint64(fs.Bsize)
	di.Free = fs.Bfree * uint64(fs.Bsize)
	di.Used = di.Total - di.Free

	return di
}

func GetTotalBytes(partition string) uint64 {
	di := GetDiskInfo(partition)
	return di.Free
}

func GetUsedBytes(partition string) uint64 {
	di := GetDiskInfo(partition)
	return di.Used
}

func GetFreeBytes(partition string) uint64 {
	di := GetDiskInfo(partition)
	return di.Free
}