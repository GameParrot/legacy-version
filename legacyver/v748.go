package legacyver

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New748() *Protocol {
	return &Protocol{
		ver: "1.21.40",
		id:  proto.ID748}
}
