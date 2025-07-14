package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

// New662 ...
func New662() *Protocol {
	return &Protocol{
		ver: "1.20.70",
		id:  proto.ID662,
	}
}
