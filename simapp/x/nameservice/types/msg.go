package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

var _ sdk.Msg = &MsgSetName{}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) *MsgSetName {
	return &MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgSetName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetName) Type() string { return "set_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

func (msg MsgSetName) Reset() {
	msg.Name = ""
	msg.Owner = sdk.AccAddress{}
	msg.Value = ""
}

func (msg MsgSetName) String() string {
	return fmt.Sprintf("Name: (%s), Owner: (%s), Value: (%s)", msg.Name, msg.Owner.String, msg.Value)
}

func (_ MsgSetName) ProtoMessage() {}

// MsgBuyName defines the BuyName message
type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

var _ sdk.Msg = &MsgBuyName{}

// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) *MsgBuyName {
	return &MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

// Route should return the name of the module
func (msg MsgBuyName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBuyName) Type() string { return "buy_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyName) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

func (msg MsgBuyName) Reset() {
	msg.Name = ""
	msg.Bid = sdk.Coins{}
	msg.Buyer = sdk.AccAddress{}
}

func (msg MsgBuyName) String() string {
	return fmt.Sprintf("Name: (%s), Bid: (%s), Buyer: (%s)", msg.Name, msg.Bid.String(), msg.Buyer.String())
}

func (_ MsgBuyName) ProtoMessage() {}

// MsgDeleteName defines a DeleteName message
type MsgDeleteName struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

var _ sdk.Msg = &MsgDeleteName{}

// NewMsgDeleteName is a constructor function for MsgDeleteName
func NewMsgDeleteName(name string, owner sdk.AccAddress) *MsgDeleteName {
	return &MsgDeleteName{
		Name:  name,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgDeleteName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteName) Type() string { return "delete_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

func (msg MsgDeleteName) Reset() {
	msg.Name = ""
	msg.Owner = sdk.AccAddress{}
}

func (msg MsgDeleteName) String() string {
	return fmt.Sprintf("Name: (%s), Owner: (%s)", msg.Name, msg.Owner.String())
}

func (_ MsgDeleteName) ProtoMessage() {}
