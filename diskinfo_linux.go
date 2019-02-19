// +build !windows

package diskinfo

import (
	"syscall"
	"log"
)

func GetDiskInfo() DiskInfo {
	var di DiskInfo

	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		log.Panic(err)
	}

	di.Total = fs.Blocks * uint64(fs.Bsize)
	di.Free = fs.Bfree * uint64(fs.Bsize)
	di.Used = di.Total - di.Free

	return di
}

func GetTotalBytes() uint64 {
	di := GetDiskInfo()
	return di.Free
}

func GetUsedBytes() uint64 {
	di := GetDiskInfo()
	return di.Used
}

func GetFreeBytes() uint64 {
	di := GetDiskInfo()
	return di.Free
}