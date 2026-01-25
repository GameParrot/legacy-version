package legacyver

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
)

func New860() *Protocol {
	p := New859()
	p.ver = "1.21.124"
	p.id = proto.ID860
	return p
}
