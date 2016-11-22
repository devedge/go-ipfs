package network

import (
	"context"

	bsmsg "github.com/ipfs/go-ipfs/exchange/bitswap/message"
	cid "gx/ipfs/QmX4hxL9LDFVpYtNfBEBgVSynRGsooVf4F8nrvJiCZuxqq/go-cid"
	protocol "gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
	peer "gx/ipfs/QmfMmLGoKzCHDN7cGgk64PJr4iipzidDRME8HABSJqvmhC/go-libp2p-peer"
)

var (
	// These two are equivalent, legacy
	ProtocolBitswapOne    protocol.ID = "/ipfs/bitswap/1.0.0"
	ProtocolBitswapNoVers protocol.ID = "/ipfs/bitswap"

	ProtocolBitswap protocol.ID = "/ipfs/bitswap/1.1.0"
)

// BitSwapNetwork provides network connectivity for BitSwap sessions
type BitSwapNetwork interface {

	// SendMessage sends a BitSwap message to a peer.
	SendMessage(
		context.Context,
		peer.ID,
		bsmsg.BitSwapMessage) error

	// SetDelegate registers the Reciver to handle messages received from the
	// network.
	SetDelegate(Receiver)

	ConnectTo(context.Context, peer.ID) error

	NewMessageSender(context.Context, peer.ID) (MessageSender, error)

	Routing
}

type MessageSender interface {
	SendMsg(bsmsg.BitSwapMessage) error
	Close() error
}

// Implement Receiver to receive messages from the BitSwapNetwork
type Receiver interface {
	ReceiveMessage(
		ctx context.Context,
		sender peer.ID,
		incoming bsmsg.BitSwapMessage)

	ReceiveError(error)

	// Connected/Disconnected warns bitswap about peer connections
	PeerConnected(peer.ID)
	PeerDisconnected(peer.ID)
}

type Routing interface {
	// FindProvidersAsync returns a channel of providers for the given key
	FindProvidersAsync(context.Context, *cid.Cid, int) <-chan peer.ID

	// Provide provides the key to the network
	Provide(context.Context, *cid.Cid) error
}
