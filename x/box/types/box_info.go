package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//Box interface
type Box interface {
	GetBoxId() string
	SetBoxId(string)

	GetBoxType() string
	SetBoxType(string)

	GetOwner() sdk.AccAddress
	SetOwner(sdk.AccAddress)

	GetCreatedTime() time.Time
	SetCreatedTime(time.Time)

	GetName() string
	SetName(string)

	GetTotalAmount() sdk.Coin
	SetTotalAmount(sdk.Coin)

	GetDescription() string
	SetDescription(string)

	IsTradeDisabled() bool
	SetTradeDisabled(bool)

	GetLock() LockBox
	SetLock(LockBox)

	GetDeposit() DepositBox
	SetDeposit(DepositBox)

	GetFuture() FutureBox
	SetFuture(FutureBox)

	String() string
}

// BoxInfos is an array of BoxInfo
type BoxInfos []BoxInfo

//type BaseBoxInfo struct {
//}
type BoxInfo struct {
	BoxId         string         `json:"box_id"`
	Owner         sdk.AccAddress `json:"owner"`
	Name          string         `json:"name"`
	BoxType       string         `json:"type"`
	CreatedTime   time.Time      `json:"created_time"`
	TotalAmount   sdk.Coin       `json:"total_amount"`
	Description   string         `json:"description"`
	TradeDisabled bool           `json:"trade_disabled"`
	Lock          LockBox        `json:"lock"`
	Deposit       DepositBox     `json:"deposit"`
	Future        FutureBox      `json:"future"`
}

// Implements Box Interface
var _ Box = (*BoxInfo)(nil)

func (bi BoxInfo) GetBoxId() string {
	return bi.BoxId
}
func (bi *BoxInfo) SetBoxId(boxId string) {
	bi.BoxId = boxId
}
func (bi BoxInfo) GetBoxType() string {
	return bi.BoxType
}
func (bi *BoxInfo) SetBoxType(boxType string) {
	bi.BoxType = boxType
}
func (bi BoxInfo) GetOwner() sdk.AccAddress {
	return bi.Owner
}
func (bi *BoxInfo) SetOwner(owner sdk.AccAddress) {
	bi.Owner = owner
}
func (bi BoxInfo) GetCreatedTime() time.Time {
	return bi.CreatedTime
}
func (bi *BoxInfo) SetCreatedTime(createdTime time.Time) {
	bi.CreatedTime = createdTime
}
func (bi BoxInfo) GetName() string {
	return bi.Name
}
func (bi *BoxInfo) SetName(name string) {
	bi.Name = name
}
func (bi BoxInfo) GetTotalAmount() sdk.Coin {
	return bi.TotalAmount
}
func (bi *BoxInfo) SetTotalAmount(totalAmount sdk.Coin) {
	bi.TotalAmount = totalAmount
}
func (bi BoxInfo) GetDescription() string {
	return bi.Description
}
func (bi *BoxInfo) SetDescription(description string) {
	bi.Description = description
}

func (bi BoxInfo) IsTradeDisabled() bool {
	return bi.TradeDisabled
}

func (bi *BoxInfo) SetTradeDisabled(tradeDisabled bool) {
	bi.TradeDisabled = tradeDisabled
}

func (bi BoxInfo) GetLock() LockBox {
	return bi.Lock
}
func (bi *BoxInfo) SetLock(lock LockBox) {
	bi.Lock = lock
}

func (bi BoxInfo) GetDeposit() DepositBox {
	return bi.Deposit
}
func (bi *BoxInfo) SetDeposit(deposit DepositBox) {
	bi.Deposit = deposit
}

func (bi BoxInfo) GetFuture() FutureBox {
	return bi.Future
}
func (bi *BoxInfo) SetFuture(future FutureBox) {
	bi.Future = future
}

type AddressDeposit struct {
	Address sdk.AccAddress `json:"address"`
	Amount  sdk.Int        `json:"amount"`
}

func NewAddressDeposit(address sdk.AccAddress, amount sdk.Int) AddressDeposit {
	return AddressDeposit{address, amount}
}
func (bi AddressDeposit) String() string {
	return fmt.Sprintf(`
  Address:			%s
  Amount:			%s`,
		bi.Address.String(), bi.Amount.String())
}

//nolint
func (bi BoxInfo) String() string {
	return fmt.Sprintf(`Box:
  BoxId: 	         			%s
  Owner:           				%s
  Name:             			%s
  TotalAmount:      			%s
  CreatedTime:					%s
  Description:	    			%s
  TradeDisabled:  				%t`,
		bi.BoxId, bi.Owner.String(), bi.Name, bi.TotalAmount.String(),
		bi.CreatedTime.String(), bi.Description, bi.TradeDisabled)
}

//nolint
func (bi BoxInfos) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-16s|%s\n",
		"BoxID", "Owner", "Name", "BoxTime")
	for _, box := range bi {
		out += fmt.Sprintf("%-17s|%-44s|%-16s|%s\n",
			box.GetBoxId(), box.GetOwner().String(), box.GetName(), box.GetCreatedTime().String())
	}
	return strings.TrimSpace(out)
}
