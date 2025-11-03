package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New844() *Protocol {
	return &Protocol{
		ver: "1.21.111",
		id:  proto.ID844,
	}
}
