package proto

import (
	_ "unsafe"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalItemStackRequest(r protocol.IO, x *protocol.ItemStackRequest) {
	r.Varint32(&x.RequestID)
	protocol.FuncSlice(r, &x.Actions, func(p *protocol.StackRequestAction) {
		IOStackRequestAction(r, p)
	})
	protocol.FuncSlice(r, &x.FilterStrings, r.String)
	r.Int32(&x.FilterCause)
}

func MarshalStackRequestAction(r protocol.IO, x protocol.StackRequestAction) {
	switch act := x.(type) {
	case *protocol.TakeStackRequestAction:
		MarshalTakeStackRequestAction(r, act)
	case *protocol.PlaceStackRequestAction:
		MarshalPlaceStackRequestAction(r, act)
	case *protocol.SwapStackRequestAction:
		MarshalSwapStackRequestAction(r, act)
	case *protocol.DropStackRequestAction:
		MarshalDropStackRequestAction(r, act)
	case *protocol.DestroyStackRequestAction:
		MarshalDestroyStackRequestAction(r, act)
	case *protocol.ConsumeStackRequestAction:
		MarshalConsumeStackRequestAction(r, act)
	case *protocol.CraftRecipeStackRequestAction:
		MarshalCraftRecipeStackRequestAction(r, act)
	case *protocol.AutoCraftRecipeStackRequestAction:
		MarshalAutoCraftRecipeStackRequestAction(r, act)
	case *protocol.CraftCreativeStackRequestAction:
		MarshalCraftCreativeStackRequestAction(r, act)
	case *protocol.CraftGrindstoneRecipeStackRequestAction:
		MarshalCraftGrindstoneRecipeStackRequestAction(r, act)
	case *protocol.CraftLoomRecipeStackRequestAction:
		MarshalCraftLoomRecipeStackRequestAction(r, act)
	case *protocol.CraftResultsDeprecatedStackRequestAction:
		MarshalCraftResultsDeprecatedStackRequestAction(r, act)
	default:
		x.Marshal(r)
	}
}

func MarshalItemStackResponse(r protocol.IO, x *protocol.ItemStackResponse) {
	r.Uint8(&x.Status)
	r.Varint32(&x.RequestID)
	if x.Status == protocol.ItemStackResponseStatusOK {
		protocol.FuncIOSlice(r, &x.ContainerInfo, MarshalStackResponseContainerInfo)
	}
}

func MarshalStackResponseContainerInfo(r protocol.IO, x *protocol.StackResponseContainerInfo) {
	MarshalFullContainerName(r, &x.Container)
	protocol.FuncIOSlice(r, &x.SlotInfo, MarshalStackResponseSlotInfo)
}

func MarshalStackResponseSlotInfo(r protocol.IO, x *protocol.StackResponseSlotInfo) {
	r.Uint8(&x.Slot)
	r.Uint8(&x.HotbarSlot)
	r.Uint8(&x.Count)
	r.Varint32(&x.StackNetworkID)
	if x.Slot != x.HotbarSlot {
		r.InvalidValue(x.HotbarSlot, "hotbar slot", "hot bar slot must be equal to normal slot")
	}
	r.String(&x.CustomName)
	if IsProtoGTE(r, ID766) {
		r.String(&x.FilteredCustomName)
	}
	r.Varint32(&x.DurabilityCorrection)
}

func MarshalTakeStackRequestAction(r protocol.IO, a *protocol.TakeStackRequestAction) {
	r.Uint8(&a.Count)
	StackReqSlotInfo(r, &a.Source)
	StackReqSlotInfo(r, &a.Destination)
}

func MarshalPlaceStackRequestAction(r protocol.IO, a *protocol.PlaceStackRequestAction) {
	r.Uint8(&a.Count)
	StackReqSlotInfo(r, &a.Source)
	StackReqSlotInfo(r, &a.Destination)
}

func MarshalSwapStackRequestAction(r protocol.IO, a *protocol.SwapStackRequestAction) {
	StackReqSlotInfo(r, &a.Source)
	StackReqSlotInfo(r, &a.Destination)
}

func MarshalDropStackRequestAction(r protocol.IO, a *protocol.DropStackRequestAction) {
	r.Uint8(&a.Count)
	StackReqSlotInfo(r, &a.Source)
	r.Bool(&a.Randomly)
}

func MarshalDestroyStackRequestAction(r protocol.IO, a *protocol.DestroyStackRequestAction) {
	r.Uint8(&a.Count)
	StackReqSlotInfo(r, &a.Source)
}

func MarshalConsumeStackRequestAction(r protocol.IO, a *protocol.ConsumeStackRequestAction) {
	MarshalDestroyStackRequestAction(r, &a.DestroyStackRequestAction)
}

func MarshalPlaceInContainerStackRequestAction(r protocol.IO, a *protocol.PlaceInContainerStackRequestAction) {
	r.Uint8(&a.Count)
	StackReqSlotInfo(r, &a.Source)
}

func MarshalTakeOutContainerStackRequestAction(r protocol.IO, a *protocol.TakeOutContainerStackRequestAction) {
	r.Uint8(&a.Count)
	StackReqSlotInfo(r, &a.Source)
}

func MarshalCraftRecipeStackRequestAction(r protocol.IO, a *protocol.CraftRecipeStackRequestAction) {
	r.Varuint32(&a.RecipeNetworkID)
	if IsProtoGTE(r, ID712) {
		r.Uint8(&a.NumberOfCrafts)
	}
}

func MarshalAutoCraftRecipeStackRequestAction(r protocol.IO, a *protocol.AutoCraftRecipeStackRequestAction) {
	r.Varuint32(&a.RecipeNetworkID)
	r.Uint8(&a.NumberOfCrafts)
	if IsProtoGTE(r, ID712) {
		r.Uint8(&a.TimesCrafted)
	}
	protocol.FuncSlice(r, &a.Ingredients, r.ItemDescriptorCount)
}

func MarshalCraftCreativeStackRequestAction(r protocol.IO, a *protocol.CraftCreativeStackRequestAction) {
	r.Varuint32(&a.CreativeItemNetworkID)
	if IsProtoGTE(r, ID712) {
		r.Uint8(&a.NumberOfCrafts)
	}
}

func MarshalCraftGrindstoneRecipeStackRequestAction(r protocol.IO, c *protocol.CraftGrindstoneRecipeStackRequestAction) {
	r.Varuint32(&c.RecipeNetworkID)
	if IsProtoGTE(r, ID712) {
		r.Uint8(&c.NumberOfCrafts)
	}
	r.Varint32(&c.Cost)
}

func MarshalCraftLoomRecipeStackRequestAction(r protocol.IO, c *protocol.CraftLoomRecipeStackRequestAction) {
	r.String(&c.Pattern)
	if IsProtoGTE(r, ID712) {
		r.Uint8(&c.TimesCrafted)
	}
}

func MarshalCraftResultsDeprecatedStackRequestAction(r protocol.IO, a *protocol.CraftResultsDeprecatedStackRequestAction) {
	protocol.FuncSlice(r, &a.ResultItems, r.Item)
	r.Uint8(&a.TimesCrafted)
}

// StackReqSlotInfo reads/writes a StackRequestSlotInfo x using IO r.
func StackReqSlotInfo(r protocol.IO, x *protocol.StackRequestSlotInfo) {
	MarshalFullContainerName(r, &x.Container)
	r.Uint8(&x.Slot)
	r.Varint32(&x.StackNetworkID)
}

//go:linkname lookupStackRequestAction github.com/sandertv/gophertunnel/minecraft/protocol.lookupStackRequestAction
func lookupStackRequestAction(id uint8, x *protocol.StackRequestAction) bool

//go:linkname lookupStackRequestActionType github.com/sandertv/gophertunnel/minecraft/protocol.lookupStackRequestActionType
func lookupStackRequestActionType(x protocol.StackRequestAction, id *uint8) bool
