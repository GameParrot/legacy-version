package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New859() *Protocol {
	return &Protocol{
		ver: "1.21.120",
		id:  proto.ID859,
	}
}
