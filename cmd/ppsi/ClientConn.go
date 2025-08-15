package main

import (
	"slices"

	"github.com/wfavorite/initq"
)

/* ------------------------------------------------------------------------ */

// ClientConn is a client connection from the C-based monitor.
type ClientConn struct {
	PID int
}

/* ------------------------------------------------------------------------ */

// ClientCache is the list of all clients currently active.
type ClientCache struct {
	Cache []*ClientConn
}

/* ======================================================================== */

// ClientCache initializes the ClientCache.
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

// New creates a new ClientConn and adds it to the cache. The data is
// initialized-empty and must be filled in after.
func (cc *ClientCache) New() (cli *ClientConn) {

	cli = new(ClientConn)
	cli.PID = -1

	cc.Cache = append(cc.Cache, cli)

	return
}

/* ======================================================================== */

// Remove deletes the client instance from the cache.
func (cc *ClientCache) Remove(pid int) {
	for i, c := range cc.Cache {
		if c.PID == pid {
			cc.Cache = slices.Delete(cc.Cache, i, i)
		}
	}
}
