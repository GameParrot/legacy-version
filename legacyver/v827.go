package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/mapping"
)

const (
	// ItemVersion827 ...
	ItemVersion827 = 271
	// BlockVersion827 ...
	BlockVersion827 int32 = (1 << 24) | (21 << 16) | (100 << 8)
)

var (
	//go:embed data/dragonfly_items.json
	dragonflyLatestItemList []byte
	//go:embed data/required_item_list_827.json
	requiredItemList827 []byte
	//go:embed data/block_states_827.nbt
	blockStateData827 []byte

	itemMappingLatestPocketMine = mapping.NewItemMapping(requiredItemList827, ItemVersion827)
	itemMappingLatestDragonfly  = mapping.NewItemMapping(dragonflyLatestItemList, ItemVersion827)
	blockMappingLatest          = mapping.NewBlockMapping(blockStateData827)
)

func itemMappingLatest(dragonflyMapping bool) mapping.Item {
	if dragonflyMapping {
		return itemMappingLatestDragonfly
	}
	return itemMappingLatestPocketMine
}
