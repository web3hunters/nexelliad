package subnetworks

import (
	"encoding/hex"

	"github.com/Nexellia-Network/nexelliad/domain/consensus/model/externalapi"
)

// FromString creates a DomainSubnetworkID from the given byte slice
func FromString(str string) (*externalapi.DomainSubnetworkID, error) {
	subnetworkIDBytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return FromBytes(subnetworkIDBytes)
}