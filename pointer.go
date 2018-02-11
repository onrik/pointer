package pointer

import (
	"time"
)

// Int returns pointer to int
func Int(i int) *int {
	return &i
}

// Uint returns pointer to uint
func Uint(i uint) *uint {
	return &i
}

// Int8 returns pointer to int8
func Int8(i int8) *int8 {
	return &i
}

// Uint8 returns pointer to int8
func Uint8(i uint8) *uint8 {
	return &i
}

// Int16 returns pointer to int16
func Int16(i int16) *int16 {
	return &i
}

// Uint16 returns pointer to uint16
func Uint16(i uint16) *uint16 {
	return &i
}

// Int32 returns pointer to int32
func Int32(i int32) *int32 {
	return &i
}

// Uint32 returns pointer to uint32
func Uint32(i uint32) *uint32 {
	return &i
}

// Int64 returns pointer to int64
func Int64(i int64) *int64 {
	return &i
}

// Uint64 returns pointer to uint64
func Uint64(i uint64) *uint64 {
	return &i
}

// Float32 returns pointer to float32
func Float32(f float32) *float32 {
	return &f
}

// Float64 returns pointer to float64
func Float64(f float64) *float64 {
	return &f
}

// Complex64 returns pointer to complex64
func Complex64(c complex64) *complex64 {
	return &c
}

// Complex128 returns pointer to complex128
func Complex128(c complex128) *complex128 {
	return &c
}

// Bool returns pointer to bool
func Bool(b bool) *bool {
	return &b
}

// Byte returns pointer to byte
func Byte(b byte) *byte {
	return &b
}

// Rune returns pointer to rune
func Rune(r rune) *rune {
	return &r
}

// String returns pointer to string
func String(s string) *string {
	return &s
}

// Time returns pointer to Time
func Time(t time.Time) *time.Time {
	return &t
}
