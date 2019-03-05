package coco

/*
#include "maskApi.h"
#include "stdlib.h"

*/
import "C"
import (
	"runtime"
	"unsafe"
)

//RLE contains a pointer to an array of C.RLE
type RLE struct {
	r          *C.RLE
	h, w, size C.siz
}

//BB bounding box
type BB []float64

func (b BB) siz() C.siz {
	return C.siz(len(b))
}
func (b BB) c() C.BB {
	return (C.BB)(unsafe.Pointer(&b[0]))
}

//InitRLEs creates an array of *RLE which holds a pointer to an array of C.RLEs
func InitRLEs(size uint32) *RLE {
	r := new(RLE)
	r.size = C.siz(size)
	C.rlesInit(&r.r, r.size)
	runtime.SetFinalizer(r, rlesfree)
	return r
}

func rlesfree(r *RLE) {
	C.rlesFree(&r.r, r.size)
	r = nil
}

//EncodeRLE binary masks using RLE.
//void rleEncode( RLE *R, const byte *mask, siz h, siz w, siz n );
func EncodeRLE(mask []byte, h, w, n uint32) *RLE {
	r := InitRLEs(n)
	r.h = C.siz(h)
	r.w = C.siz(w)
	C.rleEncode(r.r, (*C.byte)(&mask[0]), (C.siz)(h), (C.siz)(w), (C.siz)(n))
	return r
}

// Decode binary masks encoded via RLE
//void rleDecode( const RLE *R, byte *mask, siz n );
func (r *RLE) Decode() (mask []byte) {

	mask = make([]byte, r.h*r.w*r.size)
	C.rleDecode(r.r, (*C.byte)(&mask[0]), r.size)
	return mask
}

//MergeFrom - Compute union or intersection of encoded masks.
//merges m into r
func (r *RLE) MergeFrom(m *RLE, intersect bool) {

	var inter C.int
	if intersect {
		inter = 255
	}
	C.rleMerge(m.r, r.r, r.size, (inter))
}

//AreaRLE -  Compute area of encoded masks.
//void rleArea( const RLE *R, siz n, uint *a );
func (r *RLE) AreaRLE() []uint32 {

	x := make([]uint32, r.size)
	C.rleArea(r.r, r.size, (*C.uint)(&x[0]))

	return x
}

//IoURLE Compute intersection over union between masks.
//void rleIou( RLE *dt, RLE *gt, siz m, siz n, byte *iscrowd, double *o );
func IoURLE(dt, gt *RLE, iscrowd []byte) (out []float64) {
	out = make([]float64, gt.size*dt.size)
	C.rleIou(dt.r, gt.r, dt.size, gt.size, (*C.byte)(&iscrowd[0]), (*C.double)(&out[0]))
	return out
}

//NonMaxSup - Compute non-maximum suppression between bounding masks
//void rleNms( RLE *dt, siz n, uint *keep, double thr );
func (r *RLE) NonMaxSup(thresh float64) (keep []bool) {

	keep = make([]bool, r.size)
	kp := make([]C.uint, r.size)
	C.rleNms(r.r, r.size, &kp[0], (C.double)(thresh))
	for i := range keep {
		if kp[i] > 0 {
			keep[i] = true
		}

	}
	return keep
}

//IoUBB -Compute intersection over union between bounding boxes.
//void bbIou( BB dt, BB gt, siz m, siz n, byte *iscrowd, double *o );
func IoUBB(dt, gt BB, iscrowd []byte) (out []float64) {
	out = make([]float64, dt.siz()*gt.siz())
	C.bbIou(dt.c(), gt.c(), dt.siz(), gt.siz(), (*C.byte)(&iscrowd[0]), (*C.double)(&out[0]))
	return out
}

//NonMaxSupBB non-maximum suppression between bounding boxes
//void bbNms( BB dt, siz n, uint *keep, double thr );
func NonMaxSupBB(dt BB, thresh float64) (keep []bool) {
	keep = make([]bool, dt.siz())
	kp := make([]C.uint, dt.siz())
	var x C.BB
	C.bbNms(x, dt.siz(), &kp[0], (C.double)(thresh))
	for i := range keep {
		if kp[i] > 0 {
			keep[i] = true
		}
	}
	return keep
}

//ToBB bounding boxes surrounding encoded masks.
//void rleToBbox( const RLE *R, BB bb, siz n );
func (r *RLE) ToBB() (bb BB) {

	vol := r.size * r.w * r.h

	bb = make(BB, vol)

	C.rleToBbox(r.r, bb.c(), r.size)
	return bb
}

//ToRLE Convert bounding boxes to encoded masks.
//void rleFrBbox( RLE *R, const BB bb, siz h, siz w, siz n );
func (b BB) ToRLE(h, w, n uint32) *RLE {
	r := InitRLEs(n)
	r.h = (C.siz)(h)
	r.w = (C.siz)(w)
	C.rleFrBbox(r.r, b.c(), r.h, r.w, r.size)
	return r
}

// RLEFromPoly Convert polygon to encoded mask.
//void rleFrPoly( RLE *R, const double *xy, siz k, siz h, siz w );
func RLEFromPoly(poly *float64, k, h, w uint32) *RLE {
	r := InitRLEs(1)
	r.h = (C.siz)(h)
	r.w = (C.siz)(w)
	C.rleFrPoly(r.r, (*C.double)(poly), (C.siz)(k), (C.siz)(h), (C.siz)(w))
	return r

}

//Char contains a pointer to a c char string
type Char struct {
	c unsafe.Pointer
}

//ToChar Get compressed string representation of encoded mask.
//char* rleToString( const RLE *R );
func (r *RLE) ToChar() *Char {
	x := new(Char)
	x.c = unsafe.Pointer(C.rleToString(r.r))
	runtime.SetFinalizer(x, freechar)
	return x
}
func freechar(c *Char) {
	C.free(c.c)
	c = nil
}

//ToRLE Convert from compressed string representation of encoded mask.
//void rleFrString( RLE *R, char *s, siz h, siz w );
func (c *Char) ToRLE(h, w uint32) *RLE {
	r := InitRLEs(1)
	C.rleFrString(r.r, (*C.char)(c.c), (C.siz)(h), (C.siz)(w))
	return r
}
