package legacyver

import "github.com/akmalfairuz/legacy-version/legacyver/proto"

// New1001 returns the protocol used by Minecraft 1.26.30.
func New1001() *Protocol {
	return &Protocol{
		ver: "1.26.30",
		id:  proto.ID1001,
	}
}
