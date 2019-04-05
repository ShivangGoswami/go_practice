package bzip

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include<bzlib.h>
int bz2compress(bz_stream *s,int action,char *in,unsigned *intlen,char *out,unsigned *outlen);
*/
import "C"
import "io"

type Writer struct {
	w      io.Writer
	stream *C.bz_stream
	outbuf [64 * 1024]byte
}

func NewWriter(out io.Writer) io.WriteCloser{
	const(
		blockSize=9
		verbosity=0
		workFactor=30
	)
	w:=&Writer{w:out,stream:new(c.bz_stream)}
	c.Bz2_bzCompressInit(w.stream,blockSize,verbosity,workFactor)
	return w
}
