package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_address "github.com/cosmos/cosmos-sdk/types/address"
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
	MemberCountKey             = []byte{0x01} // prefix for the count of members
	MemberStatusKeyPrefix      = []byte{0x02} // prefix for each key to a member filtered by status
	MemberStatusCountKeyPrefix = []byte{0x03} // prefix for the count of members filtered by status
)

// MembersKey returns the key for the member with the given address
func MemberKey(address sdk.AccAddress) []byte {
	return append(MembersKeyPrefix, sdk_address.MustLengthPrefix(address.Bytes())...)
}

// MemberStatusKey returns the key for the member with the given address and status
func MemberStatusKey(status MembershipStatus, a sdk.AccAddress) []byte {
	// Convert MembershipStatus to byte
	byteValue := byte(status)
	return append(MemberStatusKeyPrefix, append([]byte{byteValue}, sdk_address.MustLengthPrefix(a.Bytes())...)...)
}

// MemberStatusCountKey returns the key for the count of members with the given status
func MemberStatusCountKey(status MembershipStatus) []byte {
	// Convert MembershipStatus to byte
	byteValue := byte(status)
	return append(MemberStatusCountKeyPrefix, []byte{byteValue}...)
}
