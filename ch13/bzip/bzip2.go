package bzip

/*
//#cgo CFLAGS: -I/usr/include
//#cgo LDFLAGS: -L/usr/lib -lbz2
#include<bzlib.h>
int bz2compress(bz_stream *s,int action,char *in,unsigned *intlen,char *out,unsigned *outlen);
*/
import (
	"unsafe"
	"C"
	"io"
)

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
	C.Bz2_bzCompressInit(w.stream,blockSize,verbosity,workFactor)
	return w
}

func (w *writer) Write(data []byte) (int,error) {
	if w.stream == nil{
		panic("closed")
	}
	var total int
	for len(data)>0 {
		inlen,outlen:=C.uint(len(data),C.uint(cap(w.outbuf)))
		c.bz2compress(w.stream,c.BZ_RUN,
			(*C.char)(unsafe.Pointer(&data[0])),&inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)),&outlen)
		total+=int(inlen)
		data=data[inlen:]
		if _,err:=w.w.Write(w.outbuf[:outlen]);err!=nil{
			return total,err
		}
	}
	return total,nil
}

func (w *writer) Close() error{
	if w.stream==nil{
		panic("closed")
	}
	defer func(){
		C.Bz2_bzCompressInit(w.stream)
		w.stream=nil
	}()
	for {
		inlen,outlen:=C.uint(0),C.uint(cap(w.outbuf))
		r:=C.bz2compress(w.stream,C.BZ_FINISH,nil,&inlen,(*C.char)(unsafe.Pointer(&w.outbuf)),&outlen)
		if _,err:=w.w.write(w.outbuf[:outlen]);err!=nil{
			return err
		}
		if r==C.BZ_STREAM_END{
			return nil
		}
	}
}