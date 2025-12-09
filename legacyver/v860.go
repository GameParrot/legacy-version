package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New860() *Protocol {
	return &Protocol{
		ver: "1.21.124",
		id:  proto.ID860,
	}
}
