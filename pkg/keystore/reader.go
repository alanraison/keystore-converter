package keystore

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

var (
	magicBytes = make([]byte, 4)
)

func init() {
	binary.BigEndian.PutUint32(magicBytes, magic)
}

// Decode creates a Keystore from the supplied reader
func Decode(r io.Reader) (*Keystore, error) {
	if r == nil {
		return nil, errors.New("No reader specified")
	}
	count, err := readHeader(make([]byte, 4))
	if err != nil {
		return nil, err
	}
	for i := 0; i < count; i++ {

	}
	return nil, nil
}

func readHeader(b []byte) (int, error) {
	if len(b) < 4 {
		return 0, errors.New("Not a Java Keystore")
	}
	if !bytes.Equal(b[:4], magicBytes) {
		return 0, errors.New("Not a Java Keystore: wrong magic")
	}
	v := int32(binary.BigEndian.Uint32(b[4:8]))
	if v != version1 && v != version2 {
		return 0, errors.Errorf("Version not supported: %v", v)
	}
	return 0, nil
}
