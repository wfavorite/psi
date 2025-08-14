package main

/* ------------------------------------------------------------------------ */
/*
	Version history

	0.1.0    13-8-25 - Original creation.
	0.2.0    13-8-25 - Code cleanup and commentary. Prep for another day of
	                   work.
			 14-8-25 - New command line options and output functionality.
			         - Lotta changes all over the place.
	0.3.0    14-8-25 - Cleanup, documentation, and test coverage.
	                 - I added Paul Reynaud as the primary author. I found a
					   'diary' of his notes during his imprisonment by the
					   Vichy regime on the discount shelf at the Strand in
					   NYC. While interesting, most of what i read (without
					   actually buying it) was his delayed knowledge of what
					   was happening in the war. There was some interesting
					   reflections - and promise of more - but not enough
					   for me to commit to reading the whole book.
*/

/* ------------------------------------------------------------------------ */

// VersionString must be commented.
const VersionString = "0.3.0"

/* ------------------------------------------------------------------------ */
/*
	ToDos:

	Done:
	[X] Author.
	[X] Test coverage for bad command line options.
	[X] All (public) command line options should be in the usage output.
	[X] Determine command line calling convention.
	[X] Fix the import (linter) error/problem.
	[X] Interval should first assume time in seconds, and if not then attempt
	    to parse as a time.Duration.
	[X] Create the tabular output.
	[X] Allow for -t(imestamp) in tabular output.
	[X] Consider renaming. The "psi" name is sure to exist. Check for name
	    collisions.
	[X] Establish thresholds for ANSI colours.
	[X] Wrap the read-file in a read-all function/class.
	[X] The read-all/some class needs an Update() method for the iteration
	    option.
	[X] Consider moving the actual file read to a sub-class.
	[X] Handle command line basics.

*/
