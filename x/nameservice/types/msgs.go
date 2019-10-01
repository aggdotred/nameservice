package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey is defined in key.go file
const RouterKey = ModuleName

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
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
func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value Cannot be empty")
	}
	return nil
}
