package main

import (
	"os/exec"
	"syscall"
	"time"

	"github.com/wfavorite/initq"
)

/*
	STUB: This is heavily stubbed.
	STUB:
	STUB: What should happen
	STUB: - Launch clients from a local (internal) copy.
	STUB: - Launch client(s) based on config.
	STUB: - Pass things like args and socket name to client.
	STUB: - Keeps PID and process info.
*/

/* ------------------------------------------------------------------------ */

// ClientCmd is the stats and info wrapper for the launched client.
type ClientCmd struct {
	cmd           *exec.Cmd `json:"-"`
	PID           int       `json:"pid"`
	MonitorTarget string    `json:"monitor_target"`
	MonitorArgs   string    `json:"monitor_arg"`
	Started       string    `json:"start_timestamp"`
}

/* ------------------------------------------------------------------------ */

type ClientLaunch struct {
	List []*ClientCmd `json:"list"`
}

/* ======================================================================== */

func (cd *CoreData) ClientLaunch() initq.ReqResult {

	if cd.List == nil {
		return initq.TryAgain
	}

	clil := new(ClientLaunch)
	clil.List = make([]*ClientCmd, 0)
	cd.Clil = clil

	// STUB: Here is where a config would be applied to what to monitor.

	// This has a tendency to trigger occasionally even on a fairly quiesced
	// system.
	cd.LaunchClient("io", "some 15000 1000000")
	// From the kernel.org example.
	cd.LaunchClient("memory", "some 150000 1000000")

	// This is an invalid value and will cause errno 22 (invalid argument) error
	// before polling will begin. Such a client will exit prematurely.
	// cd.LaunchClient("cpu", "some 150000000 1000000000")

	return initq.Satisfied

}

/* ======================================================================== */

// ShutdownClients walks the process list and sends TERM signals to each.
func (cd *CoreData) ShutdownClients() {
	for _, cmdWrap := range cd.Clil.List {
		cmdWrap.cmd.Process.Signal(syscall.SIGTERM)
	}
}

/* ======================================================================== */

// LaunchClient starts a client using the provided arguments.
func (cd *CoreData) LaunchClient(target string, arg string) {

	cmdWrap := new(ClientCmd)
	cmdWrap.MonitorTarget = target
	cmdWrap.MonitorArgs = arg

	cmdWrap.cmd = exec.Command("../cpsi/cpsi", cd.List.Filename, target, arg)

	// Debugggery. This should be optional in a production instance.
	cmdWrap.cmd.Env = append(cmdWrap.cmd.Env, "CPSI_VERBOSE_BN=cpsi.out")

	cmdWrap.cmd.Start()
	cmdWrap.PID = cmdWrap.cmd.Process.Pid
	cmdWrap.Started = time.Now().UTC().Format(time.RFC3339)
	// Don't wait.\

	// Add to list
	cd.Clil.List = append(cd.Clil.List, cmdWrap)

}
