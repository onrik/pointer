package pointer

import (
	"time"
)

// GetInt returns value from pointer or default value
func GetInt(i *int, d int) int {
	if i != nil {
		return *i
	}

	return d
}

// GetInt8 returns value from pointer or default value
func GetInt8(i *int8, d int8) int8 {
	if i != nil {
		return *i
	}

	return d
}

// GetInt16 returns value from pointer or default value
func GetInt16(i *int16, d int16) int16 {
	if i != nil {
		return *i
	}

	return d
}

// GetInt32 returns value from pointer or default value
func GetInt32(i *int32, d int32) int32 {
	if i != nil {
		return *i
	}

	return d
}

// GetInt64 returns value from pointer or default value
func GetInt64(i *int64, d int64) int64 {
	if i != nil {
		return *i
	}

	return d
}

// GetUint returns value from pointer or default value
func GetUint(i *uint, d uint) uint {
	if i != nil {
		return *i
	}

	return d
}

// GetUint8 returns value from pointer or default value
func GetUint8(i *uint8, d uint8) uint8 {
	if i != nil {
		return *i
	}

	return d
}

// GetUint16 returns value from pointer or default value
func GetUint16(i *uint16, d uint16) uint16 {
	if i != nil {
		return *i
	}

	return d
}

// GetUint32 returns value from pointer or default value
func GetUint32(i *uint32, d uint32) uint32 {
	if i != nil {
		return *i
	}

	return d
}

// GetUint64 returns value from pointer or default value
func GetUint64(i *uint64, d uint64) uint64 {
	if i != nil {
		return *i
	}

	return d
}

// GetFloat32 returns value from pointer or default value
func GetFloat32(f *float32, d float32) float32 {
	if f != nil {
		return *f
	}

	return d
}

// GetFloat64 returns value from pointer or default value
func GetFloat64(f *float64, d float64) float64 {
	if f != nil {
		return *f
	}

	return d
}

// GetComplex64 returns value from pointer or default value
func GetComplex64(c *complex64, d complex64) complex64 {
	if c != nil {
		return *c
	}

	return d
}

// GetComplex128 returns value from pointer or default value
func GetComplex128(c *complex128, d complex128) complex128 {
	if c != nil {
		return *c
	}

	return d
}

// GetBool returns value from pointer or default value
func GetBool(b *bool, d bool) bool {
	if b != nil {
		return *b
	}

	return d
}

// GetByte returns value from pointer or default value
func GetByte(b *byte, d byte) byte {
	if b != nil {
		return *b
	}

	return d
}

// GetRune returns value from pointer or default value
func GetRune(r *rune, d rune) rune {
	if r != nil {
		return *r
	}

	return d
}

// GetString returns value from pointer or default value
func GetString(s *string, d string) string {
	if s != nil {
		return *s
	}

	return d
}

// GetTime returns value from pointer or default value
func GetTime(t *time.Time, d time.Time) time.Time {
	if t != nil {
		return *t
	}

	return d
}
