package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New827() *Protocol {
	return &Protocol{
		ver: "1.21.100",
		id:  proto.ID827,
	}
}
