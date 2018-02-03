package paasio

import "io"

type typeReadCounter struct {
	//nReadBytes int
	//nReadOps   int
	reader io.Reader
}

func NewReadCounter(reader io.Reader) ReadCounter {
	//nr = reader
	return ReadCounter{reader}
}

type NewWriteCounter struct {
}

// type NewReadWriteCounter struct {
// 	NewReadCounter
// 	NewWriteCounter
// }

func (nrc *typeReadCounter) ReadCount() (int64, int) {
	//x := nrc.(io.Reader)
	//nrc.ReadCount()
	return 0, 0
}

func (*NewWriteCounter) WriteCount() (int64, int) {
	return 0, 0
}
