package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

// New685 uses same data as 686
func New685() *Protocol {
	return &Protocol{
		ver: "1.21.0",
		id:  proto.ID685,
	}
}
