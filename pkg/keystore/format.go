package keystore

// Keystore is the contents of a Java Keystore; a collection of key pairs or public certificate chains
type Keystore struct {
	Entries []interface{}
}

const (
	magic    uint32 = 0xfeedfeed
	version1 int32  = 0x01
	version2 int32  = 0x02
)
