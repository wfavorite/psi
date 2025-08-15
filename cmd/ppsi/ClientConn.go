package main

import (
	"slices"

	"github.com/wfavorite/initq"
)

/* ------------------------------------------------------------------------ */

type ClientConn struct {
	PID int
}

/* ------------------------------------------------------------------------ */

type ClientCache struct {
	Cache []*ClientConn
}

/* ======================================================================== */

func (cd *CoreData) ClientCache() initq.ReqResult {

	if cd.Cash != nil {
		return initq.Satisfied
	}

	if cd.Logr == nil {
		return initq.TryAgain
	}

	cc := new(ClientCache)
	cc.Cache = make([]*ClientConn, 0)

	cd.Cash = cc

	return initq.Satisfied
}

/* ======================================================================== */

func (cc *ClientCache) New() (cli *ClientConn) {

	cli = new(ClientConn)
	cli.PID = -1

	cc.Cache = append(cc.Cache, cli)

	return
}

/* ======================================================================== */

func (cc *ClientCache) Remove(pid int) {
	for i, c := range cc.Cache {
		if c.PID == pid {
			cc.Cache = slices.Delete(cc.Cache, i, i)
		}
	}
}
