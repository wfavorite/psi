package main

/*
	Version history:
	0.0.0     15-8-25 - Skeletal beginnings.
	0.1.0     15-8-25 - Now launches clients on start.
	0.2.0     19-8-25 - Works. PoC mostly feature complete.
*/

// VersionString should be commented.
const VersionString = "0.2.0"

/*
	ToDos:
	[ ] Heartbeats are stats, not logged items.
	[ ] Define some sort of config file.
	[ ] Drop a pid file.
	[ ] Logging should go to a file.
	[ ] Probably need to check on clients - to see if/when they exited (in the
	    launch list). There is a 'bad' argument example in ClientLaunch.go that
		will cause the cpsi instance to exit. Currently these become zombies.

	Done:
	[X] Resolve the logging problem.
*/
