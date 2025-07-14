package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New827_Pre() *Protocol {
	return &Protocol{
		ver: "1.21.100",
		id:  proto.ID827,
	}
}
