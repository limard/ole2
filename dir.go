package ole2

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"unicode/utf16"
)

type DIR_TYPE byte

const (
	EMPTY       DIR_TYPE = 0
	USERSTORAGE DIR_TYPE = 1
	USERSTREAM  DIR_TYPE = 2
	LOCKBYTES   DIR_TYPE = 3
	PROPERTY    DIR_TYPE = 4
	ROOT        DIR_TYPE = 5
)

func (t DIR_TYPE) String() string {
	switch t {
	case EMPTY:
		return "EMPTY"
	case USERSTORAGE:
		return "STORAGE"
	case USERSTREAM:
		return "STREAM"
	case LOCKBYTES:
		return "LOCKBYTES"
	case PROPERTY:
		return "PROPERTY"
	case ROOT:
		return "ROOT"
	}
	return ""
}

type File struct {
	NameBts   [32]uint16
	Bsize     uint16
	Type      DIR_TYPE
	Flag      byte
	Left      uint32
	Right     uint32
	Child     uint32
	Guid      [8]uint16
	Userflags uint32
	Time      [2]uint64
	Sstart    uint32
	Size      uint32
	Proptype  uint32
}

func (d *File) Name() string {
	runes := utf16.Decode(d.NameBts[:d.Bsize/2-1])
	return string(runes)
}

func ParseOle10Native(r io.Reader) (obj *Ole10Native, e error) {
	reader := ole10NativeReader{
		r: bufio.NewReader(r),
	}

	obj = &Ole10Native{}
	obj.NativeSize, e = reader.readUint32()
	if e != nil {
		return nil, fmt.Errorf("ole2: invalid ole10native. %s", e.Error())
	}

	reader.r.Discard(2) // 02 00
	obj.Name, e = reader.read0EndString()
	if e != nil {
		return nil, fmt.Errorf("ole2: invalid ole10native. %s", e.Error())
	}
	obj.CacheName, e = reader.read0EndString()
	if e != nil {
		return nil, fmt.Errorf("ole2: invalid ole10native. %s", e.Error())
	}
	reader.read0EndString()
	reader.read0EndString()

	reader.r.Discard(2) // 03 00
	reader.readAnis()
	obj.NativeData, e = reader.readAnis()
	if e != nil {
		return nil, fmt.Errorf("ole2: invalid ole10native. %s", e.Error())
	}
	return obj, nil
}

type Ole10Native struct {
	NativeSize uint32 // file size
	Name       string
	CacheName  string
	NativeData []byte
}

type ole10NativeReader struct {
	r *bufio.Reader
}

func (t *ole10NativeReader) readUint32() (n uint32, e error) {
	if e = binary.Read(t.r, binary.LittleEndian, &n); e != nil {
		return
	}
	return
}
func (t *ole10NativeReader) read0EndString() (s string, e error) {
	s, e = t.r.ReadString(0x00)
	if e != nil {
		return
	}
	if len(s) > 0 && s[len(s)-1] == 0x00 {
		s = s[:len(s)-1]
	}
	return s, nil
}
func (t *ole10NativeReader) readAnis() (b []byte, e error) {
	var length uint32
	if e = binary.Read(t.r, binary.LittleEndian, &length); e != nil {
		return nil, e
	}
	b = make([]byte, length)
	_, e = io.ReadFull(t.r, b)
	if e != nil {
		return nil, e
	}
	return b, nil
}
