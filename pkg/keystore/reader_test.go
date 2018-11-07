package keystore

import (
	"bytes"
	"testing"
)

func TestPassingNilReaderIsError(t *testing.T) {
	if _, err := Decode(nil); err == nil {
		t.Error("expected error, got none")
	}
}

func TestShouldRecogniseMagicConstant(t *testing.T) {
	header := []byte{0xfe, 0xed, 0xfe, 0xed}
	if err := readMagic(bytes.NewReader(header)); err != nil {
		t.Errorf("Unexpected error when reading header: %+v", err)
	}
	header = []byte{}
	if err := readMagic(bytes.NewReader(header)); err == nil {
		t.Error("Expected error reading header, got none")
	}
	header = []byte{0xde, 0xad, 0xbe, 0xef}
	if err := readMagic(bytes.NewReader(header)); err == nil {
		t.Error("Expected error reading magic, got none")
	}
}

func TestShouldAcceptVersionOneOrVersionTwo(t *testing.T) {
	header := []byte{0x0, 0x0, 0x0, 0x0}
	if err := readVersion(bytes.NewReader(header)); err == nil {
		t.Error("Expected version error (0), got none")
	}
	header[3] = 0x1
	if err := readVersion(bytes.NewReader(header)); err != nil {
		t.Errorf("Unexpected error reading version (1): %+v", err)
	}
	header[3] = 0x2
	if err := readVersion(bytes.NewReader(header)); err != nil {
		t.Errorf("Unexpected error reading version (2): %+v", err)
	}
	header[2] = 0xf
	header[3] = 0xf
	if err := readVersion(bytes.NewReader(header)); err == nil {
		t.Error("Expected version error (255), got none")
	}
}
