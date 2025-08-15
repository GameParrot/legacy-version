package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New819() *Protocol {
	return &Protocol{
		ver: "1.21.90",
		id:  proto.ID819,
	}
}
