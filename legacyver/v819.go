package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/internal/chunk"
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/akmalfairuz/legacy-version/mapping"
)

const (
	// ItemVersion819 ...
	ItemVersion819 = 261
	// BlockVersion819 ...
	BlockVersion819 int32 = (1 << 24) | (21 << 16) | (90 << 8)
)

var (
	//go:embed data/required_item_list_819.json
	requiredItemList819 []byte
	//go:embed data/block_states_819.nbt
	blockStateData819 []byte
)

func New819(dragonflyMapping bool) *Protocol {
	itemMapping := mapping.NewItemMapping(requiredItemList819, ItemVersion819)
	blockMapping := mapping.NewBlockMapping(blockStateData819)
	return &Protocol{
		ver:             "1.21.90",
		id:              proto.ID819,
		blockTranslator: NewBlockTranslator(blockMapping, blockMappingLatest, chunk.NewNetworkPersistentEncoding(blockMapping, BlockVersion819), chunk.NewBlockPaletteEncoding(blockMapping, BlockVersion819), false),
		itemTranslator:  NewItemTranslator(itemMapping, itemMappingLatest(dragonflyMapping), blockMapping, blockMappingLatest),
	}
}
