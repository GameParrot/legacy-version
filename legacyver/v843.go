package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New843() *Protocol {
	return &Protocol{
		ver: "1.21.110",
		id:  proto.ID843,
	}
}
