package diskinfo

type DiskInfo struct {
	Total	uint64	`json:"all"`
	Used	uint64	`json:"used"`
	Free 	uint64	`json:"free"`
}