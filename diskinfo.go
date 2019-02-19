package diskinfo

type DiskInfo struct {
	Total	uint64	`json:"total"`
	Used	uint64	`json:"used"`
	Free 	uint64	`json:"free"`
}