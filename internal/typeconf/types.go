package typeconf

import (
	"fmt"
	"github.com/samber/lo"
)

func StringToInt[T ~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64](s string, fallback int) T {
	var result T
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		return T(fallback)
	}
	return result
}

func IntToString[T ~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64](i T, fallback string) string {
	return fmt.Sprintf("%d", i)
}

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func ByteToBool(b byte) bool {
	return b != 0
}

func IntToInt[T1 ~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64, T2 ~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64](i T1) T2 {
	return T2(i)
}

func SliceIntToSliceInt[T1 ~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64, T2 ~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64](in []T1) []T2 {
	return lo.Map(in, func(i T1, _ int) T2 {
		return T2(i)
	})
}
