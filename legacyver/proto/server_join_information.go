package proto

import (
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalGatheringJoinInfo(r protocol.IO, x *protocol.GatheringJoinInfo) {
	if IsProtoLT(r, ID944) {
		uuidStr := x.ExperienceID.String()
		r.String(&uuidStr)
		if uuidStr != "" {
			x.ExperienceID = uuid.MustParse(uuidStr)
		}
	} else {
		r.UUID(&x.ExperienceID)
	}
	r.String(&x.ExperienceName)
	if IsProtoLT(r, ID944) {
		uuidStr := x.ExperienceWorldID.String()
		r.String(&uuidStr)
		if uuidStr != "" {
			x.ExperienceWorldID = uuid.MustParse(uuidStr)
		}
	} else {
		r.UUID(&x.ExperienceWorldID)
	}
	r.String(&x.ExperienceWorldName)
	r.String(&x.CreatorID)
	if IsProtoGTE(r, ID944) {
		r.UUID(&x.TargetID)
		if IsProtoLT(r, ID975) {
			sid, _ := uuid.Parse(x.ScenarioID)
			r.UUID(&sid)
			x.ScenarioID = sid.String()
		} else {
			r.String(&x.ScenarioID)
		}
		r.String(&x.ServerID)
	} else {
		storeId := ""
		r.String(&storeId)
	}
}

func MarshalServerJoinInformation(r protocol.IO, x *protocol.ServerJoinInformation) {
	protocol.OptionalFuncIO(r, &x.GatheringJoinInfo, MarshalGatheringJoinInfo)
	if IsProtoGTE(r, ID944) {
		protocol.OptionalMarshaler(r, &x.StoreEntryPointInfo)
		protocol.OptionalFuncIO(r, &x.PresenceInfo, MarshalPresenceInfo)
	}
}

func MarshalPresenceInfo(r protocol.IO, x *protocol.PresenceInfo) {
	if IsProtoGTE(r, ID1001) {
		protocol.OptionalFunc(r, &x.ExperienceName, r.String)
		protocol.OptionalFunc(r, &x.WorldName, r.String)
		r.String(&x.RichPresenceID)
		return
	}
	experienceName, _ := x.ExperienceName.Value()
	r.String(&experienceName)
	x.ExperienceName = protocol.Option(experienceName)
	worldName, _ := x.WorldName.Value()
	r.String(&worldName)
	x.WorldName = protocol.Option(worldName)
}
