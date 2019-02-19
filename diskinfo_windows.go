// +build windows

package diskinfo

import (
	"syscall"
	"log"
	"unsafe"
)

func GetDiskInfo() DiskInfo {
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
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("D:"))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)), 0, 0)

	di.Total = uint64(lpTotalNumberOfBytes)
	di.Free = uint64(lpTotalNumberOfFreeBytes)
	di.Used = di.Total - di.Free

	return di
}

func GetTotalBytes() uint64 {
	di := GetDiskInfo()
	return di.Total
}

func GetUsedBytes() uint64 {
	di := GetDiskInfo()
	return di.Used
}

func GetFreeBytes() uint64 {
	di := GetDiskInfo()
	return di.Free
}