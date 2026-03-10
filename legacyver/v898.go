package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New898() *Protocol {
	return &Protocol{
		ver: "1.21.130",
		id:  proto.ID898,
	}
}
