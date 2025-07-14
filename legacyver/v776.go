package legacyver

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New776() *Protocol {
	return &Protocol{
		ver: "1.21.60",
		id:  proto.ID776}
}
