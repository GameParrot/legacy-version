package proto

func EntityDataFlagsLength(protoID int32) int {
	if protoID >= ID898 {
		return 127
	}
	if protoID >= ID844 {
		return 126
	}
	if protoID >= ID818 {
		return 125
	}
	if protoID >= ID800 {
		return 124
	}
	if protoID >= ID786 {
		return 123
	}
	return 120
}
