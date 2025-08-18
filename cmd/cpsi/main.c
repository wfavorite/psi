
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <signal.h>
#include <poll.h>
#include <errno.h>


#include "clargs.h"
#include "usock.h"
#include "debugstr.h"

/*
    ToDo:
    [ ] Make stdout/debug mode stdout 'consistent'.
    [ ] Design standardized return values that can be used for observability.
    [ ] Move to a Linux host and add the actual poll implementation.
        (Currently written/running on Darwin.)
    [ ] Should not write to stdout. Perhaps to syslog, or a local log?
    [ ] The Unix socket path should probably be an environmental variable
        (passed by the parent).
    [ ] Find a more appropriate heartbeat time (or drop it).
    

    Done:
    [X] Write a proper signal handler.
    [X] Accept the three arguments.
*/



/* ======================================================================== */

void TERM_handler(int sig)
{
    shutdown_socket();
    fprintf(debug_fp, "Monitor exiting.\n"); fflush(debug_fp);
    fclose(debug_fp);
    exit(0);
}




/* ======================================================================== */

int main(int argc, char *argv[])
{
    CLArgs *cmdline;
    struct pollfd pfd;       /* One member array of pollfd for poll()       */
    int poll_rv;             /* poll() return value                         */

    // Register the signal handler.
    signal(SIGTERM, TERM_handler);
    signal(SIGQUIT, TERM_handler);
    signal(SIGINT, TERM_handler);
    signal(SIGHUP, SIG_IGN);

    // Read the command line.
    if ( NULL == (cmdline = parse_cmdline(argc, argv)) )
    {
        // A true assert-level issue. (I just don't use assert() for the check.)
        fprintf(stderr, "ASSERT: Failed to allocate memory during app start.\n");
        exit(1);
    }

    // Handle the command line (parsing) error - if one exists.
    if (cmdline->error[0] != 0)
    {
        fprintf(stderr, "ERROR: %s\n", cmdline->error);
        exit(1);
    }

    // (Conditionally) Setup the debug stream.
    establish_debug_stream();

    // Write invocation info to the debug stream.
    fprintf(debug_fp, "Write events to      : %s\n", cmdline->sendto);
    fprintf(debug_fp, "Poll target          : %s\n", cmdline->target);
    fprintf(debug_fp, "Alarm args-threshold : %s\n", cmdline->args);
    fflush(debug_fp);
    
    // Setup the Unix socket.
    init_unix_socket(cmdline->sendto);


    if (0 > (pfd.fd = open(cmdline->target, O_RDWR | O_NONBLOCK)))
    {
        fprintf(stderr, "ERROR: open failure - %s\n", strerror(errno));
        exit(1);
    }
    pfd.events = POLLPRI;

    if (0 > write(pfd.fd, cmdline->args, strlen(cmdline->args) + 1))
    {
        fprintf(stderr, "ERROR: proc write failure - %s\n", strerror(errno));
        exit(1);
    }

    while ( 1 )
    {
        poll_rv = poll(&pfd, 1, -1);
        
        if ( poll_rv < 0 )
        {
            fprintf(stderr, "ERROR: poll failed - %s\n", strerror(errno));
            exit(1);
        }

        if (pfd.revents & POLLERR)
        {
            fprintf(stderr, "ERROR: got POLLERR, event source is gone\n");
            exit(1);
        }
        
        if (pfd.revents & POLLPRI)
        {
            send_threshold_event();
        }
        else
        {
            fprintf(stderr, "ERROR: Unknown event received - 0x%x\n", pfd.revents);
            exit(1);
        }

      
    }

    // STUB: Unreachable.




    // STUB: Just exit for now.
    //close(s);
    exit(0);

}