package packet

import (
	"bytes"
	"encoding/binary"
	"log"
	"sync"
)

// Writer will be write bytes
type Writer struct {
	buffer *bytes.Buffer
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// NewPlainWriter will create Writer without data
func NewPlainWriter() *Writer {
	writer := &Writer{
		buffer: bufPool.Get().(*bytes.Buffer),
	}

	writer.buffer.Reset()

	return writer
}

// NewWriter will create Writer will new buffer
func NewWriter(packetID uint16) *Writer {
	writer := &Writer{
		buffer: bufPool.Get().(*bytes.Buffer),
	}

	writer.buffer.Reset()

	writer.WriteUInt16(packetID)
	return writer
}

// GetData will return bytes array in buffer prefix with header
func (pw *Writer) GetData() []byte {
	defer bufPool.Put(pw.buffer)

	header := make([]byte, 2)
	binary.LittleEndian.PutUint16(header, uint16(pw.buffer.Len()-2))

	return append(header, pw.buffer.Bytes()...)
}

// GetDataWithoutPrefixHeader will return bytes array in buffer
func (pw *Writer) GetDataWithoutPrefixHeader() []byte {
	return pw.buffer.Bytes()
}

// GetDataWithoutRemoveHeader will return bytes array in buffer without remove header
func (pw *Writer) GetDataWithoutRemoveHeader() []byte {
	return pw.buffer.Bytes()
}

// WriteUInt8 perform writing uint8 data to byte buffer.
func (pw *Writer) WriteUInt8(data uint8) {
	pw.write(data)
}

// WriteUInt16 perform writing uint16 data to byte buffer.
func (pw *Writer) WriteUInt16(data uint16) {
	pw.write(data)
}

// WriteUInt32 perform writing uint32 data to byte buffer.
func (pw *Writer) WriteUInt32(data uint32) {
	pw.write(data)
}

// WriteUInt64 perform writing uint64 data to byte buffer.
func (pw *Writer) WriteUInt64(data uint64) {
	pw.write(data)
}

// WriteInt8 perform writing int8 data to byte buffer.
func (pw *Writer) WriteInt8(data int8) {
	pw.write(data)
}

// WriteInt16 perform writing int16 data to byte buffer.
func (pw *Writer) WriteInt16(data int16) {
	pw.write(data)
}

// WriteInt32 perform writing int32 data to byte buffer.
func (pw *Writer) WriteInt32(data int32) {
	pw.write(data)
}

// WriteInt64 perform writing int64 data to byte buffer.
func (pw *Writer) WriteInt64(data int64) {
	pw.write(data)
}

// WriteFloat32 perform writing float32 data to byte buffer.
func (pw *Writer) WriteFloat32(data float32) {
	pw.write(data)
}

// WriteFloat64 perform writing float64 data to byte buffer.
func (pw *Writer) WriteFloat64(data float64) {
	pw.write(data)
}

// WriteString perform writing string data to byte buffer.
func (pw *Writer) WriteString(data string) {
	pw.write(data)
}

// WriteBoolean perform writing boolean data to byte buffer.
func (pw *Writer) WriteBoolean(data bool) {
	pw.write(data)
}

func (pw *Writer) write(data interface{}) {
	switch v := data.(type) {
	case string:
		pw.buffer.Write([]byte(v))
		pw.buffer.WriteByte(uint8(0))
	default:
		if err := binary.Write(pw.buffer, binary.LittleEndian, data); err != nil {
			log.Fatal("binary.Write failed: ", data, err)
		}
	}
}
