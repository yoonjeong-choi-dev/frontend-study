package tftp_server

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
)

const (
	// DatagramSize 최대 지원하는 데이터그램 크기
	// => UDP 통신의 파편화 방지
	DatagramSize = 516

	// BlockSize 데이터그램에서 헤더를 뺀 본문 크기
	BlockSize = DatagramSize - 4
)

// OpCode 헤더에 포함되는 수행할 명령어 코드
// TCP 연결에서의 수신 관련 패킷(ACK, SYN)을 어플리케이션 레벨에서 구현
type OpCode uint16

// 헤더에 포함시킬 OpCode 정의
const (
	OpRRQ OpCode = iota + 1
	_
	OpData
	OpAck
	OpErr
)

// ErrCode 요청 처리 중 에러 발생 시 응답으로 전송할 에러 타입
type ErrCode uint16

const (
	ErrUnknown ErrCode = iota
	ErrNotFound
	ErrAccessViolation
	ErrDiskFull
	ErrIllegalOp
	ErrUnknownID
	ErrFileExists
	ErrNoUser
)

/*
다음 각 타입들에 대해서 마샬링 및 언마샬 메서드 구현 필요
- ReadRequest
  - 읽기 요청
  - OpCode(2) + filename + null(1) + mode + null(0)
- Data
  - 데이터
  - OpCode(2) + Block Number(2) + payload
- Ack
  - 확인 패킷
  - OpCode(2) + Block Number(2)
- Err
  - 에러 패킷
  - OpCode(2) + ErrCode(2) + Error Message + null(0)

각 타입은 encoding 패키지의 BinaryMarshaler 및 BinaryUnmarshaler 인터페이스 구현
- 각 타입을 네트워크 연결에서 사용 가능
- MarshalBinary: BinaryMarshaler 인터페이스 메서드
- UnmarshalBinary: BinaryUnmarshaler 인터페이스 메서드
*/

type ReadRequest struct {
	Filename string
	Mode     string
}

func (r *ReadRequest) MarshalBinary() ([]byte, error) {
	// default mode: octet
	mode := "octet"
	if r.Mode != "" {
		mode = r.Mode
	}

	// 읽기 요청 패킷 구조
	// OpCode(2) + filename + null(1) + mode + null(0)
	capacity := 2 + 2 + len(r.Filename) + 1 + len(r.Mode) + 1

	b := new(bytes.Buffer)
	b.Grow(capacity)

	// 읽기 요청 패킷 구조에 맞게 바이너리 데이터 쓰기
	err := binary.Write(b, binary.BigEndian, OpRRQ)
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(r.Filename)
	if err != nil {
		return nil, err
	}

	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(mode)
	if err != nil {
		return nil, err
	}

	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (r *ReadRequest) UnmarshalBinary(payload []byte) error {
	req := bytes.NewBuffer(payload)

	var code OpCode

	// 읽기 요청 패킷 구조에 맞게 바이너리 데이터 언마샬
	err := binary.Read(req, binary.BigEndian, &code)
	if err != nil {
		return err
	}

	if code != OpRRQ {
		return errors.New("invalid RRQ")
	}

	r.Filename, err = req.ReadString(0)
	if err != nil {
		return errors.New("invalid RRQ")
	}

	// null 제거
	r.Filename = strings.TrimRight(r.Filename, "\x00")
	if len(r.Filename) == 0 {
		return errors.New("invalid RRQ")
	}

	r.Mode, err = req.ReadString(0)
	if err != nil {
		return errors.New("invalid RRQ")
	}

	// null 제거
	r.Mode = strings.TrimRight(r.Mode, "\x00")
	if len(r.Mode) == 0 {
		return errors.New("invalid RRQ")
	}

	// only "octet" mode supported
	mode := strings.ToLower(r.Mode)
	if mode != "octet" {
		return errors.New("only binary transfers supported")
	}

	return nil
}

type Data struct {
	Block   uint16
	Payload io.Reader
}

func (d *Data) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	b.Grow(DatagramSize)

	// 블럭 번호는 1씩 증가
	d.Block++

	// 패킷 구조: OpCode(2) + Block Number(2) + payload
	err := binary.Write(b, binary.BigEndian, OpData)
	if err != nil {
		return nil, err
	}

	err = binary.Write(b, binary.BigEndian, d.Block)
	if err != nil {
		return nil, err
	}

	_, err = io.CopyN(b, d.Payload, BlockSize)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return b.Bytes(), nil
}

func (d *Data) UnmarshalBinary(payload []byte) error {
	if l := len(payload); l < 4 || l > DatagramSize {
		return errors.New("invalid DATA")
	}

	var opCode OpCode

	// 패킷 구조: OpCode(2) + Block Number(2) + payload
	err := binary.Read(bytes.NewReader(payload[:2]), binary.BigEndian, &opCode)
	if err != nil || opCode != OpData {
		return errors.New("invalid DATA")
	}

	err = binary.Read(bytes.NewReader(payload[2:4]), binary.BigEndian, &d.Block)
	if err != nil {
		return errors.New("invalid DATA")
	}

	d.Payload = bytes.NewBuffer(payload[4:])
	return nil
}

type Ack uint16

func (a *Ack) MarshalBinary() ([]byte, error) {
	// 패킷 구조: OpCode(2) + Block Number(2)
	capacity := 2 + 2

	b := new(bytes.Buffer)
	b.Grow(capacity)

	err := binary.Write(b, binary.BigEndian, OpAck)
	if err != nil {
		return nil, err
	}

	err = binary.Write(b, binary.BigEndian, a)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), err
}

func (a *Ack) UnmarshalBinary(payload []byte) error {
	// 패킷 구조: OpCode(2) + Block Number(2)
	var opCode OpCode
	r := bytes.NewReader(payload)

	err := binary.Read(r, binary.BigEndian, &opCode)
	if err != nil {
		return err
	}
	if opCode != OpAck {
		return errors.New("invalid ACK")
	}

	return binary.Read(r, binary.BigEndian, a)
}

type Err struct {
	Error   ErrCode
	Message string
}

func (e *Err) MarshalBinary() ([]byte, error) {
	// 패킷 구조: OpCode(2) + ErrCode(2) + Error Message + null(0)
	capacity := 2 + 2 + len(e.Message) + 1

	b := new(bytes.Buffer)
	b.Grow(capacity)

	err := binary.Write(b, binary.BigEndian, OpErr)
	if err != nil {
		return nil, err
	}

	err = binary.Write(b, binary.BigEndian, e.Error)
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(e.Message)
	if err != nil {
		return nil, err
	}

	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), err
}

func (e *Err) UnmarshalBinary(payload []byte) error {
	// 패킷 구조: OpCode(2) + ErrCode(2) + Error Message + null(0)
	r := bytes.NewBuffer(payload)
	var opCode OpCode

	err := binary.Read(r, binary.BigEndian, &opCode)
	if err != nil {
		return err
	}
	if opCode != OpErr {
		return errors.New("invalid ERROR")
	}

	err = binary.Read(r, binary.BigEndian, &e.Error)
	if err != nil {
		return err
	}

	e.Message, err = r.ReadString(0)
	if err != nil {
		return err
	}

	e.Message = strings.TrimRight(e.Message, "\x00")
	return nil
}
