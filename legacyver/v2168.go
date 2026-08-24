package legacyver

import "github.com/akmalfairuz/legacy-version/legacyver/proto"

// New2168 returns the protocol used by Minecraft 1.26.40.
func New2168() *Protocol {
	return &Protocol{
		ver: "1.26.40",
		id:  proto.ID2168,
	}
}
