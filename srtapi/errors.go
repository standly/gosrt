// Copyright (c) 2018 CyberAgent, Inc. All rights reserved.
// https://github.com/openfresh/gosrt

package srtapi

// #cgo LDFLAGS: -lsrt
// #include <srt/srt.h>
import "C"

// Errors
const (
	EUNKNOWN        = Errno(C.SRT_EUNKNOWN)
	SUCCESS         = Errno(C.SRT_SUCCESS)
	ECONNSETUP      = Errno(C.SRT_ECONNSETUP)
	ENOSERVER       = Errno(C.SRT_ENOSERVER)
	ECONNREJ        = Errno(C.SRT_ECONNREJ)
	ESOCKFAIL       = Errno(C.SRT_ESOCKFAIL)
	ESECFAIL        = Errno(C.SRT_ESECFAIL)
	ECONNFAIL       = Errno(C.SRT_ECONNFAIL)
	ECONNLOST       = Errno(C.SRT_ECONNLOST)
	ENOCONN         = Errno(C.SRT_ENOCONN)
	ERESOURCE       = Errno(C.SRT_ERESOURCE)
	ETHREAD         = Errno(C.SRT_ETHREAD)
	ENOBUF          = Errno(C.SRT_ENOBUF)
	EFILE           = Errno(C.SRT_EFILE)
	EINVRDOFF       = Errno(C.SRT_EINVRDOFF)
	ERDPERM         = Errno(C.SRT_ERDPERM)
	EINVWROFF       = Errno(C.SRT_EINVWROFF)
	EWRPERM         = Errno(C.SRT_EWRPERM)
	EINVOP          = Errno(C.SRT_EINVOP)
	EBOUNDSOCK      = Errno(C.SRT_EBOUNDSOCK)
	ECONNSOCK       = Errno(C.SRT_ECONNSOCK)
	EINVPARAM       = Errno(C.SRT_EINVPARAM)
	EINVSOCK        = Errno(C.SRT_EINVSOCK)
	EUNBOUNDSOCK    = Errno(C.SRT_EUNBOUNDSOCK)
	ENOLISTEN       = Errno(C.SRT_ENOLISTEN)
	ERDVNOSERV      = Errno(C.SRT_ERDVNOSERV)
	ERDVUNBOUND     = Errno(C.SRT_ERDVUNBOUND)
	EINVALMSGAPI    = Errno(C.SRT_EINVALMSGAPI)
	EINVALBUFFERAPI = Errno(C.SRT_EINVALBUFFERAPI)
	EDUPLISTEN      = Errno(C.SRT_EDUPLISTEN)
	ELARGEMSG       = Errno(C.SRT_ELARGEMSG)
	EINVPOLLID      = Errno(C.SRT_EINVPOLLID)
	EASYNCFAIL      = Errno(C.SRT_EASYNCFAIL)
	EASYNCSND       = Errno(C.SRT_EASYNCSND)
	EASYNCRCV       = Errno(C.SRT_EASYNCRCV)
	ETIMEOUT        = Errno(C.SRT_ETIMEOUT)
	ECONGEST        = Errno(C.SRT_ECONGEST)
	EPEERERR        = Errno(C.SRT_EPEERERR)
)
