package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New944() *Protocol {
	return &Protocol{
		ver: "1.26.10",
		id:  proto.ID944,
	}
}
