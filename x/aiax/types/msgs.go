package types

import (
	"fmt"

	// cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ sdk.Msg = &MsgSendToEthereum{}
)

// NewMsgSendToEthereum returns a new MsgSendToEthereum
func NewMsgSendToEthereum(sender sdk.AccAddress, destAddress string, send sdk.Coin, bridgeFee sdk.Coin) *MsgSendToEthereum {
	return &MsgSendToEthereum{
		Sender:            sender.String(),
		EthereumRecipient: destAddress,
		Amount:            send,
		BridgeFee:         bridgeFee,
	}
}

// Route should return the name of the module
func (msg MsgSendToEthereum) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSendToEthereum) Type() string { return "send_to_eth" }

// ValidateBasic runs stateless checks on the message
// Checks if the Eth address is valid
func (msg MsgSendToEthereum) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	// fee and send must be of the same denom
	// this check is VERY IMPORTANT
	if msg.Amount.Denom != msg.BridgeFee.Denom {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins,
			fmt.Sprintf("fee and amount must be the same type %s != %s", msg.Amount.Denom, msg.BridgeFee.Denom))
	}

	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "amount")
	}
	if !msg.BridgeFee.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "fee")
	}
	if !common.IsHexAddress(msg.EthereumRecipient) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "ethereum address")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSendToEthereum) GetSignBytes() []byte {
	panic(fmt.Errorf("deprecated"))
}

// GetSigners defines whose signature is required
func (msg MsgSendToEthereum) GetSigners() []sdk.AccAddress {
	acc, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{acc}
}
