package main

/* ------------------------------------------------------------------------ */
/*
	Version history

	0.1.0    13-8-25 - Original creation.
	0.2.0    13-8-25 - Code cleanup and commentary. Prep for another day of
	                   work.
*/

/* ------------------------------------------------------------------------ */

// VersionString must be commented.
const VersionString = "0.1.0"

/* ------------------------------------------------------------------------ */
/*
	ToDos:
	[_] Determine command line calling convention.
	[ ] Consider moving the actual file read to a sub-class.
	[ ] Author.
	[ ] Establish thresholds for ANSI colours.
	[ ] Consider renaming. The "psi" name is sure to exist. Check for name
	    collisions.
	[ ] Create the tabular output.
	[ ] Allow for -t(imestamp) in tabular output.
	[ ] Interval should first assume time in seconds, and if not then attempt
	    to parse as a time.Duration.
	[ ] Wrap the read-file in a read-all function/class.
	[ ] The read-all/some class needs an Update() method for the iteration
	    option.

	Done:
	[X] Handle command line basics.

*/
