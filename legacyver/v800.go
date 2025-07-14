package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New800() *Protocol {
	return &Protocol{
		ver: "1.21.80",
		id:  proto.ID800,
	}
}
