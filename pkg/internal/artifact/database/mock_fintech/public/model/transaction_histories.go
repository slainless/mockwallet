//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type TransactionHistories struct {
	ID              int32
	UUID            string
	AccountUUID     string
	DestUUID        *string
	Mutation        int64
	Currency        string
	Status          int16
	Address         *string
	TransactionNote *string
	TransactionDate time.Time
	TransactionType int16
}
