// +build windows

package diskinfo

import (
	"log"
	"syscall"
	"unsafe"
)

func GetDiskInfo(partition string) DiskInfo {
	var di DiskInfo

	kernel32, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		log.Panic(err)
	}

	defer syscall.FreeLibrary(kernel32)
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")

	if err != nil {
		log.Panic(err)
	}

	var lpFreeBytesAvailable, lpTotalNumberOfBytes, lpTotalNumberOfFreeBytes int64
	_, _, _ = syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(partition))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)), 0, 0)

	di.Total = uint64(lpTotalNumberOfBytes)
	di.Free = uint64(lpTotalNumberOfFreeBytes)
	di.Used = di.Total - di.Free

	return di
}

func GetTotalBytes(partition string) uint64 {
	di := GetDiskInfo(partition)
	return di.Total
}

func GetUsedBytes(partition string) uint64 {
	di := GetDiskInfo(partition)
	return di.Used
}

func GetFreeBytes(partition string) uint64 {
	di := GetDiskInfo(partition)
	return di.Free
}
