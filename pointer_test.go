package pointer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func (s *Suite) TestInt() {
	a := 143
	p := Int(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestUint() {
	a := uint(19)
	p := Uint(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestInt8() {
	a := int8(67)
	p := Int8(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestUint8() {
	a := uint8(90)
	p := Uint8(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestInt16() {
	a := int16(84)
	p := Int16(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestUint16() {
	a := uint16(145)
	p := Uint16(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestInt32() {
	a := int32(334)
	p := Int32(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestUint32() {
	a := uint32(2837)
	p := Uint32(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestInt64() {
	a := int64(982304)
	p := Int64(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestUint64() {
	a := uint64(39240983)
	p := Uint64(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestFloat32() {
	a := float32(334.234)
	p := Float32(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestFloat64() {
	a := float64(90.213)
	p := Float64(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestComplex64() {
	a := complex64(522)
	p := Complex64(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestComplex128() {
	a := complex128(522)
	p := Complex128(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestBool() {
	a := true
	p := Bool(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestByte() {
	a := byte(48)
	p := Byte(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestRune() {
	a := 'a'
	p := Rune(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestString() {
	a := "123"
	p := String(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func (s *Suite) TestTime() {
	a := time.Now()
	p := Time(a)
	s.Require().NotNil(p)
	s.Require().Equal(a, *p)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
