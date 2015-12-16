package main

import (
	"encoding/binary"
	"io"
	"math"
)

type binaryWriter struct {
	w io.Writer
}

func newBinaryWriter(w io.Writer) *binaryWriter {
	return &binaryWriter{w}
}

func (bw *binaryWriter) WriteBool(b bool) error {
	bs := []byte{0}
	if b {
		bs = []byte{1}
	}
	_, err := bw.w.Write(bs)
	return err
}

func (bw *binaryWriter) WriteInt8(i int8) error {
	return binary.Write(bw.w, binary.BigEndian, i)
}

func (bw *binaryWriter) WriteInt16(i int16) error {
	return binary.Write(bw.w, binary.BigEndian, i)
}

func (bw *binaryWriter) WriteInt32(i int32) error {
	return binary.Write(bw.w, binary.BigEndian, i)
}

func (bw *binaryWriter) WriteInt64(i int64) error {
	return binary.Write(bw.w, binary.BigEndian, i)
}

func (bw *binaryWriter) WriteUint8(n uint8) error {
	return binary.Write(bw.w, binary.BigEndian, n)
}

func (bw *binaryWriter) WriteUint16(n uint16) error {
	return binary.Write(bw.w, binary.BigEndian, n)
}

func (bw *binaryWriter) WriteUint32(n uint32) error {
	return binary.Write(bw.w, binary.BigEndian, n)
}

func (bw *binaryWriter) WriteUint64(n uint64) error {
	return binary.Write(bw.w, binary.BigEndian, n)
}

func (bw *binaryWriter) WriteFloat(f float32) error {
	return binary.Write(bw.w, binary.BigEndian, math.Float32bits(f))
}

func (bw *binaryWriter) WriteDouble(f float64) error {
	return binary.Write(bw.w, binary.BigEndian, math.Float64bits(f))
}

func (bw *binaryWriter) WriteString(s string) error {
	if err := bw.WriteUint32(uint32(len(s))); err != nil {
		return err
	}
	if _, err := io.WriteString(bw.w, s); err != nil {
		return err
	}
	return nil
}
