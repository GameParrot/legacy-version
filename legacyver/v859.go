package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/mapping"
)

const (
	// ItemVersion859 ...
	ItemVersion859 = 241
	// BlockVersion859 ...
	BlockVersion859 int32 = (1 << 24) | (21 << 16) | (120 << 8)
)

var (
	//go:embed data/dragonfly_items.json
	dragonflyLatestItemList []byte
	//go:embed data/required_item_list_859.json
	requiredItemList859 []byte
	//go:embed data/block_states_859.nbt
	blockStateData859 []byte

	itemMappingLatestPocketMine = mapping.NewItemMapping(requiredItemList859, ItemVersion859)
	itemMappingLatestDragonfly  = mapping.NewItemMapping(dragonflyLatestItemList, ItemVersion859)
	blockMappingLatest          = mapping.NewBlockMapping(blockStateData859)
)

func itemMappingLatest(dragonflyMapping bool) mapping.Item {
	if dragonflyMapping {
		return itemMappingLatestDragonfly
	}
	return itemMappingLatestPocketMine
}
