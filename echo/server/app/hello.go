/*****************************************************
# DESC    : Hello Service
# AUTHOR  : Alex Stocks
# VERSION : 1.0
# LICENCE : LGPL V3
# EMAIL   : alexstocks@foxmail.com
# MOD     : 2016-07-29 15:49
# FILE    : hello.go
******************************************************/

package main

import (
	"golang.org/x/net/context"
)

import (
// "github.com/AlexStocks/gocolor"
)

type Hello struct{}

func (this *Hello) Echo(ctx context.Context, req []string, rsp *string) error {
	// gocolor.Info("req:%#v", req)
	*rsp = req[0]
	// gocolor.Info("rsp:%#v", *rsp)

	return nil
}

func (this *Hello) Service() string {
	return "com.youni.HelloService"
}

func (this *Hello) Version() string {
	return ""
}
