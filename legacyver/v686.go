package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New686() *Protocol {
	return &Protocol{
		ver: "1.21.2",
		id:  proto.ID686}
}
