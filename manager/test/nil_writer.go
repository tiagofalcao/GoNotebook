package test

// Buf simulate a io.Writer
type NilWriter bool

// Write simulate write and check bytes
func (w NilWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

//Close
func (w NilWriter) Close() error {
	return nil
}
