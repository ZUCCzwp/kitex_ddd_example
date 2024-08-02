// Package entity -----------------------------
// @file      : user.go
// @author    : Jimmy
// @contact   : 2114272829@qq.com
// @time      : 2024/7/29 下午1:24
// -------------------------------------------
package entity

import "fmt"

// User represent entity of the user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// NewUser initialize user
func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name")
	}

	return &User{
		Name: name,
	}, nil
}
