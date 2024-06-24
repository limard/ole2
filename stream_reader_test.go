package ole2

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	bts := make([]byte, 1<<10)
	for i := 0; i < 1<<10; i++ {
		bts[i] = byte(i)
	}
	ole := &Ole{Lsector: 8, Lssector: 1, SecID: []uint32{2, 1, ENDOFCHAIN}, reader: bytes.NewReader(bts)}
	r := ole.stream_read(0, 30)
	res := make([]byte, 14)
	fmt.Println(r.Read(res))
	fmt.Println(res)
}

func TestSeek(t *testing.T) {
	bts := make([]byte, 1<<10)
	for i := 0; i < 1<<10; i++ {
		bts[i] = byte(i)
	}
	ole := &Ole{Lsector: 8, Lssector: 1, SecID: []uint32{2, 1, ENDOFCHAIN}, reader: bytes.NewReader(bts)}
	r := ole.stream_read(0, 30)
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
}

func TestSeek1(t *testing.T) {
	bts := make([]byte, 1<<10)
	for i := 0; i < 1<<10; i++ {
		bts[i] = byte(i)
	}
	ole := &Ole{Lsector: 8, Lssector: 1, SecID: []uint32{2, 1, ENDOFCHAIN}, reader: bytes.NewReader(bts)}
	r := ole.stream_read(0, 30)
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
}
