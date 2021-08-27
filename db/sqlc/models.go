// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"accountID"`
	// can be positive or negative
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"fromAccountID"`
	ToAccountID   int64 `json:"toAccountID"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	Username          string    `json:"username"`
	FullName          string    `json:"fullName"`
	CreatedAt         time.Time `json:"createdAt"`
	HashedPassword    string    `json:"hashedPassword"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	Email             string    `json:"email"`
}
