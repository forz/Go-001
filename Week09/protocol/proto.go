package protocol

import (
	"bufio"
	"encoding/binary"
)

// Proto p.
type Proto struct {
	PackLen   int32  // package length
	HeaderLen int16  // header length
	Ver       int16  // protocol version
	Operation int32  // operation for request
	Seq       int32  // sequence number chosen by client
	Body      []byte // body
}

const (
	rawHeaderLen = uint16(16)
)

func (p *Proto) Read(rd *bufio.Reader) (err error) {
	var (
		packLen   int32
		headerLen int16
	)
	// read
	if err = binary.Read(rd, binary.BigEndian, &packLen); err != nil {
		return
	}
	if err = binary.Read(rd, binary.BigEndian, &headerLen); err != nil {
		return
	}
	if err = binary.Read(rd, binary.BigEndian, &p.Ver); err != nil {
		return
	}
	if err = binary.Read(rd, binary.BigEndian, &p.Operation); err != nil {
		return
	}
	if err = binary.Read(rd, binary.BigEndian, &p.Seq); err != nil {
		return
	}
	var (
		n, t    int
		bodyLen = int(packLen - int32(headerLen))
	)
	if bodyLen > 0 {
		p.Body = make([]byte, bodyLen)
		for {
			if t, err = rd.Read(p.Body[n:]); err != nil {
				return
			}
			if n += t; n == bodyLen {
				break
			}
		}
	} else {
		p.Body = nil
	}
	return
}

func (p *Proto) Write(wr *bufio.Writer) (err error) {
	// write
	if err = binary.Write(wr, binary.BigEndian, uint32(rawHeaderLen)+uint32(len(p.Body))); err != nil {
		return
	}
	if err = binary.Write(wr, binary.BigEndian, rawHeaderLen); err != nil {
		return
	}
	if err = binary.Write(wr, binary.BigEndian, p.Ver); err != nil {
		return
	}
	if err = binary.Write(wr, binary.BigEndian, p.Operation); err != nil {
		return
	}
	if err = binary.Write(wr, binary.BigEndian, p.Seq); err != nil {
		return
	}
	if p.Body != nil {
		if err = binary.Write(wr, binary.BigEndian, p.Body); err != nil {
			return
		}
	}
	err = wr.Flush()
	return
}
