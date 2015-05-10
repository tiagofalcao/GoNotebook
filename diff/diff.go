package diff

import (
	"bytes"
	"fmt"
	"os"
)

// Callback is called when bytes differ
type Callback func(base, received []byte, value interface{})

// Buf simulate a io.Writer to check differences
type Buf struct {
	base     *os.File
	callback Callback
	value    interface{}
}

// NewBuf create a new Buf
func NewBuf(b *os.File, c Callback, v interface{}) Buf {
	return Buf{b, c, v}
}

// Write simulate write and check bytes
func (d Buf) Write(p []byte) (n int, err error) {
	b := make([]byte, len(p))
	d.base.Read(b)
	if bytes.Equal(p, b) {
		return len(p), nil
	}
	d.callback(b, p, d.value)
	return 0, fmt.Errorf("Output differ")

}

//Close check end of buffer
func (d Buf) Close() error {
	return nil
}
