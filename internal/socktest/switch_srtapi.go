// Copyright (c) 2018 CyberAgent, Inc. All rights reserved.
// https://github.com/openfresh/gosrt

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// https://github.com/golang/go

// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris windows

package socktest

// Sockets maps a socket descriptor to the status of socket.
type Sockets map[int]Status

func (sw *Switch) sockso(s int) *Status {
	sw.smu.RLock()
	defer sw.smu.RUnlock()
	so, ok := sw.sotab[s]
	if !ok {
		return nil
	}
	return &so
}

// addLocked returns a new Status without locking.
// sw.smu must be held before call.
func (sw *Switch) addLocked(s, family, sotype, proto int) *Status {
	sw.once.Do(sw.init)
	so := Status{Cookie: cookie(family, sotype, proto)}
	sw.sotab[s] = so
	return &so
}
