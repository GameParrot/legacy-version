package legacyver

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New712() *Protocol {
	return &Protocol{
		ver: "1.21.20",
		id:  proto.ID712}
}
