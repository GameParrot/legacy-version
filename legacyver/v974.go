package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New974() *Protocol {
	return &Protocol{
		ver: "1.26.20",
		id:  proto.ID974,
	}
}
