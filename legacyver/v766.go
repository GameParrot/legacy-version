package legacyver

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New766() *Protocol {
	return &Protocol{
		ver: "1.21.50",
		id:  proto.ID766}
}
