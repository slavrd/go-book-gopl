package limit

import (
	"io"
)

func LimitReader(r io.Reader, n int64) io.Reader {

	return &limitReader{r: r, limit: n}

}

type limitReader struct {
	r           io.Reader
	read, limit int64
}

func (lr *limitReader) Read(b []byte) (n int, err error) {
	n, err = lr.r.Read(b[:lr.limit-lr.read])
	lr.read += int64(n)
	if lr.read >= lr.limit {
		err = io.EOF
	}
	return n, err
}
