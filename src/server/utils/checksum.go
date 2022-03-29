package utils

import (
	"hash/crc32"
)

func GetHash(id string) uint32{
	return crc32.ChecksumIEEE([]byte(id))
}





