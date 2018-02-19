package pointer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type GetSuite struct {
	suite.Suite
}

func (s *GetSuite) TestGetInt() {
	a := 238
	s.Require().Equal(10, GetInt(nil, 10))
	s.Require().Equal(a, GetInt(&a, 10))
}

func (s *GetSuite) TestGetInt8() {
	a := int8(52)
	b := int8(18)
	s.Require().Equal(b, GetInt8(nil, b))
	s.Require().Equal(a, GetInt8(&a, b))
}

func (s *GetSuite) TestGetInt16() {
	a := int16(3487)
	b := int16(38)
	s.Require().Equal(b, GetInt16(nil, b))
	s.Require().Equal(a, GetInt16(&a, b))
}

func (s *GetSuite) TestGetInt32() {
	a := int32(3902)
	b := int32(990)
	s.Require().Equal(b, GetInt32(nil, b))
	s.Require().Equal(a, GetInt32(&a, b))
}

func (s *GetSuite) TestGetInt64() {
	a := int64(902338974)
	b := int64(833984)
	s.Require().Equal(b, GetInt64(nil, b))
	s.Require().Equal(a, GetInt64(&a, b))
}

func (s *GetSuite) TestGetUint() {
	a := uint(38463)
	b := uint(734)
	s.Require().Equal(b, GetUint(nil, b))
	s.Require().Equal(a, GetUint(&a, b))
}

func (s *GetSuite) TestGetUint8() {
	a := uint8(34)
	b := uint8(22)
	s.Require().Equal(b, GetUint8(nil, b))
	s.Require().Equal(a, GetUint8(&a, b))
}

func (s *GetSuite) TestGetUint16() {
	a := uint16(188)
	b := uint16(13)
	s.Require().Equal(b, GetUint16(nil, b))
	s.Require().Equal(a, GetUint16(&a, b))
}

func (s *GetSuite) TestGetUint32() {
	a := uint32(83743)
	b := uint32(238)
	s.Require().Equal(b, GetUint32(nil, b))
	s.Require().Equal(a, GetUint32(&a, b))
}

func (s *GetSuite) TestGetUint64() {
	a := uint64(28319)
	b := uint64(98343)
	s.Require().Equal(b, GetUint64(nil, b))
	s.Require().Equal(a, GetUint64(&a, b))
}

func (s *GetSuite) TestGetFloat32() {
	a := float32(34.2343)
	b := float32(22.3284)
	s.Require().Equal(b, GetFloat32(nil, b))
	s.Require().Equal(a, GetFloat32(&a, b))
}

func (s *GetSuite) TestGetFloat64() {
	a := float64(9324.32434)
	b := float64(9823.2323)
	s.Require().Equal(b, GetFloat64(nil, b))
	s.Require().Equal(a, GetFloat64(&a, b))
}

func (s *GetSuite) TestGetComplex64() {
	a := complex64(348973)
	b := complex64(324863)
	s.Require().Equal(b, GetComplex64(nil, b))
	s.Require().Equal(a, GetComplex64(&a, b))
}

func (s *GetSuite) TestGetComplex128() {
	a := complex128(3463)
	b := complex128(9823)
	s.Require().Equal(b, GetComplex128(nil, b))
	s.Require().Equal(a, GetComplex128(&a, b))
}

func (s *GetSuite) TestGetBool() {
	a := true
	b := false
	s.Require().Equal(b, GetBool(nil, b))
	s.Require().Equal(a, GetBool(&a, b))
}

func (s *GetSuite) TestGetByte() {
	a := byte(32)
	b := byte(33)
	s.Require().Equal(b, GetByte(nil, b))
	s.Require().Equal(a, GetByte(&a, b))
}

func (s *GetSuite) TestGetRune() {
	a := 'a'
	b := 'b'
	s.Require().Equal(b, GetRune(nil, b))
	s.Require().Equal(a, GetRune(&a, b))
}

func (s *GetSuite) TestGetString() {
	a := "111"
	b := "222"
	s.Require().Equal(b, GetString(nil, b))
	s.Require().Equal(a, GetString(&a, b))
}

func (s *GetSuite) TestGetTime() {
	a := time.Now()
	b := a.Add(10 * time.Second)
	s.Require().Equal(b, GetTime(nil, b))
	s.Require().Equal(a, GetTime(&a, b))
}

func TestGetSuite(t *testing.T) {
	suite.Run(t, new(GetSuite))
}
