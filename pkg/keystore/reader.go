package keystore

import (
	"bufio"
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
	if err := readMagic(r); err != nil {
		return nil, err
	}
	if err := readVersion(r); err != nil {
		return nil, err
	}
	// for i := 0; i < count; i++ {

	// }
	return nil, nil
}

func readMagic(r io.Reader) error {
	br := bufio.NewReader(r)
	magic, err := br.Peek(4)
	if err != nil {
		return errors.Wrap(err, "Reading from input")
	}
	if !bytes.Equal(magic, magicBytes) {
		return errors.New("Not a Java Keystore: wrong magic")
	}
	return nil
}

func readVersion(r io.Reader) error {
	vb := make([]byte, 4)
	if c, err := r.Read(vb); c != 4 {
		return errors.New("Could not read Keystore version")
	} else if err != nil {
		return err
	}
	v := int32(binary.BigEndian.Uint32(vb))
	if v != version1 && v != version2 {
		return errors.Errorf("Version not supported: %v", v)
	}
	return nil
}
