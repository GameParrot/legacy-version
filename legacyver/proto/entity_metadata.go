package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func EntityDataFlagsLength(protoID int32) int {
	if protoID >= ID2168 {
		return protocol.EntityDataFlagCount
	}
	if protoID >= ID975 {
		return 130
	}
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

// ClientMovementPredictionFlagsLength returns the actor-flag count used by
// ClientMovementPredictionSync. Protocol 800 retained the 123-bit movement
// prediction layout even though entity metadata itself had already grown to
// 124 flags.
func ClientMovementPredictionFlagsLength(protoID int32) int {
	if protoID < ID818 {
		return 123
	}
	return EntityDataFlagsLength(protoID)
}
