
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <signal.h>
#include <poll.h>
#include <errno.h>
#include <time.h>

#include "clargs.h"
#include "usock.h"
#include "debugstr.h"

// How often to send heartbeats into the Unix socket / framework.
#define HEARTBEAT_INTERVAL_SEC 60 * 5

/*
    ToDo:
    [ ] Should have a version - that is reported in the debug log.
    [ ] Design standardized return values that can be used for observability.

    Done:
    [X] Write a readme - with usage info (as markdown).
    [X] Find a more appropriate heartbeat time (or drop it).
    [X] Consider wrapping all debug mesages. (A bit safer if the fp is NULL.)
    [X] Make stdout/debug mode stdout 'consistent'.
    [X] The Unix socket path should probably be an environmental variable
        (passed by the parent).
    [X] Should not write to stdout. Perhaps to syslog, or a local log?
    [X] Move to a Linux host and add the actual poll implementation.
        (Currently written/running on Darwin.)
    [X] Write a proper signal handler.
    [X] Accept the three arguments.
*/

/* ======================================================================== */

void TERM_handler(int sig)
{
    shutdown_socket();
    log_printf("Monitor exiting.\n"); 
    close_debug_stream();
    exit(0);
}

/* ======================================================================== */

int main(int argc, char *argv[])
{
    CLArgs *cmdline;
    struct pollfd pfd;       /* One member array of pollfd for poll()       */
    int poll_rv;             /* poll() return value                         */
    time_t now;              /* The latest time sampling                    */
    time_t next_hb;          /* Next heartbeat time                         */
    int ensave;

    // Register the signal handler.
    signal(SIGTERM, TERM_handler);
    signal(SIGQUIT, TERM_handler);
    signal(SIGINT, TERM_handler);
    signal(SIGHUP, SIG_IGN);

    // Read the command line.
    if ( NULL == (cmdline = parse_cmdline(argc, argv)) )
    {
        // A true assert-level issue. (I just don't use assert() for the check.)
        err_printf("ASSERT: Failed to allocate memory during app start.\n");
        exit(1);
    }

    // Handle the command line (parsing) error - if one exists.
    if (cmdline->error[0] != 0)
    {
        err_printf("ERROR: %s\n", cmdline->error);
        exit(1);
    }

    // (Conditionally) Setup the debug stream.
    establish_debug_stream();

    // Write invocation info to the debug stream.
    log_printf("Write events to      : %s\n", cmdline->sendto);
    log_printf("Poll target          : %s\n", cmdline->target);
    log_printf("Alarm args-threshold : %s\n", cmdline->args);

    // Setup the Unix socket.
    init_unix_socket(cmdline->sendto);

    log_printf("Opening the target file...");
    if (0 > (pfd.fd = open(cmdline->target, O_RDWR | O_NONBLOCK)))
    {
        log_printf("Failed.\n");
        err_printf("ERROR: open failure - %s\n", strerror(errno));
        exit(1);
    }
    log_printf("Done.\n");
    pfd.events = POLLPRI;

    log_printf("Writing event threshold to the target file...");
    if (0 > write(pfd.fd, cmdline->args, strlen(cmdline->args) + 1))
    {
        ensave = errno;
        log_printf("Failed.\n");
        err_printf("ERROR: proc write failure - %s\n", strerror(ensave));
        exit(1);
    }
    log_printf("Done.\n");

    // Grab the time and set next heartbeat time.
    time(&now);
    next_hb = now + HEARTBEAT_INTERVAL_SEC;

    while ( 1 )
    {
        // 1000 is in milli (1 second).
        poll_rv = poll(&pfd, 1, 1000);
        
        if ( poll_rv < 0 )
        {
            err_printf("ERROR: poll failed - %s\n", strerror(errno));
            exit(1);
        }

        if ( poll_rv == 0 )
        {
            // The poll() call timed out.
            time(&now);
            if ( now >= next_hb )
            {
                send_heartbeat();
                // Set next heartbeat.
                next_hb = now + HEARTBEAT_INTERVAL_SEC;
            }
            continue;
        }

        if (pfd.revents & POLLERR)
        {
            err_printf("ERROR: got POLLERR, event source is gone\n");
            exit(1);
        }
        
        if (pfd.revents & POLLPRI)
        {
            send_threshold_event();
        }
        else
        {
            err_printf("ERROR: Unknown event received - 0x%x\n", pfd.revents);
            exit(1);
        }
    }

    // Unreachable.
}