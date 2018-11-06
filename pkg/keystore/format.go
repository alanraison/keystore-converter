package keystore

type Keystore struct {
	Magic   uint32
	Version int32
	Count   int32
	Entries []interface{}
}

const (
	magic    uint32 = 0xfeedfeed
	version1 int32  = 0x01
	version2 int32  = 0x02
)
