package legacyver

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New729() *Protocol {
	return &Protocol{
		ver: "1.21.30",
		id:  proto.ID729}
}
