#include<bzlib.h>

int bx2compress(bz_stream *s,int action,char *in,unsigned *inlen,char *out,unsigned *outlen){
    s->next_in=in;
    s->avail_in=*inlen;
    s->next_out=out;
    s->avail_out=*outlen;
    int r=BZ@_bzCompress(s,action);
    *inlen-=s->avail_in;
    *outlen-=s->avail_out;
    s->next_in=s->next_out=NULL:
    return r
}