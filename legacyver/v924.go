package legacyver

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New924() *Protocol {
	return &Protocol{
		ver: "1.26.0",
		id:  proto.ID924,
	}
}
