package legacyver

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New786() *Protocol {
	return &Protocol{
		ver: "1.21.70",
		id:  proto.ID786}
}
