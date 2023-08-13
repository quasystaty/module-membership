package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "membership"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_membership"
)

// Keys for membership store
// Items are stored with the following key: values
//
// - 0x00<memberAddrLen (1 Byte)><memberAddr_Bytes>: Member
//
// - 0x01: memberCount
//
// - 0x02<memberStatus (1 Byte)><memberAddrLen (1 Byte)><memberAddr_Bytes>: Status-Filtered Member
//
// - 0x03<memberStatus (1 Byte)>: Status-Filtered Member Count
var (
	MembersKeyPrefix           = []byte{0x00} // prefix for each key to a member
	MemberCountKey             = []byte{0x01} // key for the member count
	MemberStatusKeyPrefix      = []byte{0x02} // prefix for each key to a member filtered by status
	MemberStatusCountKeyPrefix = []byte{0x03} // prefix for the count of members filtered by status
	MemberMetadataKeyPrefix    = []byte{0x04} // prefix for each key to a member's metadata
	VotesToDeleteKeyPrefix     = []byte{0x05} // prefix for each key to a vote
	DirectDemocracyKey         = []byte{0x06} // key for the Direct Democracy settings

	// Add keys here so that we can check for duplicate values in a unit test
	AllKeys = [][]byte{
		MembersKeyPrefix,
		MemberCountKey,
		MemberStatusKeyPrefix,
		MemberStatusCountKeyPrefix,
		MemberMetadataKeyPrefix,
		VotesToDeleteKeyPrefix,
		DirectDemocracyKey,
	}
)

// MembersKey returns the key for the member with the given address
func MemberKey(addr sdk.AccAddress) []byte {
	return append(MembersKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

// MemberStatusKey returns the key for the member with the given address and status
func MemberStatusKey(status MembershipStatus, addr sdk.AccAddress) []byte {
	// Convert MembershipStatus to byte
	byteValue := byte(status)
	return append(MemberStatusKeyPrefix, append([]byte{byteValue}, address.MustLengthPrefix(addr.Bytes())...)...)
}

// MemberStatusCountKey returns the key for the count of members with the given status
func MemberStatusCountKey(status MembershipStatus) []byte {
	// Convert MembershipStatus to byte
	byteValue := byte(status)
	return append(MemberStatusCountKeyPrefix, []byte{byteValue}...)
}

// MemberMetadataKey returns the key for the metadata of the member with the given address and name
func MemberMetadataKey(addr sdk.AccAddress, name string) []byte {
	return append(MemberMetadataKeyPrefix, append(address.MustLengthPrefix(addr.Bytes()), []byte(name)...)...)
}

// VotesToDeleteKey returns the key for the votes of the proposal with the given ID
func VotesToDeleteKey(proposalID uint64) []byte {
	return append(VotesToDeleteKeyPrefix, govtypes.GetProposalIDBytes(proposalID)...)
}

// VoteToDeleteKey returns the key for the vote of the proposal with the given ID and voter
func VoteToDeleteKey(proposalID uint64, voterAddr sdk.AccAddress) []byte {
	return append(VotesToDeleteKey(proposalID), address.MustLengthPrefix(voterAddr.Bytes())...)
}
