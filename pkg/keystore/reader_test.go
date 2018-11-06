package keystore

import (
	"fmt"
	"testing"
)

func TestPassingNilReaderIsError(t *testing.T) {
	if _, err := Decode(nil); err == nil {
		t.Error("expected error, got none")
	}
}

func TestShouldRecogniseMagicConstant(t *testing.T) {
	header := []byte{0xfe, 0xed, 0xfe, 0xed}
	fmt.Printf("my header: %+v\n", header)
	fmt.Printf("magicBytes: %+v", magicBytes)
	if _, err := readHeader(header); err != nil {
		t.Errorf("Error when reading header: %+v", err)
	}
	header = []byte{}
	if _, err := readHeader(header); err == nil {
		t.Error("Expected error reading header, got none")
	}
	header = []byte{0xde, 0xad, 0xbe, 0xef}
	if _, err := readHeader(header); err == nil {
		t.Error("Expected error reading magic, got none")
	}
}

func TestShouldAcceptVersionOneOrVersionTwo(t *testing.T) {
	header := []byte{0xfe, 0xed, 0xfe, 0xed, 0x0, 0x0, 0x0, 0x0}
	if _, err := readHeader(header); err == nil {
		t.Error("Expected version error (0), got none")
	}
	header[7] = 0x1
	if _, err := readHeader(header); err != nil {
		t.Errorf("Unexpected error reading version (1): %+v", err)
	}
	header[7] = 0x2
	if _, err := readHeader(header); err != nil {
		t.Errorf("Unexpected error reading version (2): %+v", err)
	}
	header[6] = 0xf
	header[7] = 0xf
	if _, err := readHeader(header); err == nil {
		t.Error("Expected version error (255), got none")
	}
}

func TestShouldReadCount(t *testing.T) {
	header := make([]byte, 12)
	header = append(header, 0xfe, 0xed, 0xfe, 0xed)
	header = append(header, 0x0, 0x0, 0x0, 0x2)
	header = append(header, 0x0, 0x0, 0x0, 0x5)

	count, err := readHeader(header)
	if err == nil {

	}
	if count != 5 {

	}
}
