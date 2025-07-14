package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

// New671 ...
func New671() *Protocol {
	return &Protocol{
		ver: "1.20.80",
		id:  proto.ID671,
	}
}
