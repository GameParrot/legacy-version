package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New998() *Protocol {
	return &Protocol{
		ver: "1.26.30",
		id:  proto.ID1001,
	}
}
