
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <signal.h>

#define UNIX_SOCKET_NAME "/tmp/cpsi.sock"

#define BUFFER_SIZE 1024

/*
    ToDo:
    [ ] Design standardized return values that can be used for observability.
    [ ] Accept the three arguments.
    [ ] Move to a Linux host and add the actual poll implementation.
        (Currently written/running on Darwin.)
    [ ] Should not write to stdout. Perhaps to syslog, or a local log?
    [ ] The Unix socket path should probably be an environmental variable
        (passed by the parent).
    [ ] Find a more appropriate heartbeat time (or drop it).
    [ ] Write a proper signal handler.

    Done:

*/

int continue_loop;

/* ======================================================================== */

void TERM_handler(int sig)
{
    continue_loop = 0;

    // STUB: Shutdown should happen here - rather than setting the loop-stop
    // STUB: variable.
}




/* ======================================================================== */

int main(int argc, char *argv[])
{
    int s;                   /* Socket */
    struct sockaddr_un addr;
    int rv;                  /* Return value - multiple uses. */
    char mbuf[BUFFER_SIZE];

    
    // Initialize our global.
    continue_loop = 1;

    // Register the signal handler.
    signal(SIGTERM, TERM_handler);
    signal(SIGHUP, SIG_IGN);

   // STUB: Check for and grab the arguments here.


    printf("Creating the socket...");
    s = socket(AF_UNIX, SOCK_STREAM, 0);
    if (s < 0)
    {
        printf("Failed.\n");
        // STUB: Use strerror()
        fprintf(stderr, "ERROR: Failed to create a Unix socket.\n");
        exit(1);
    }
    printf("Done.\n");


    printf("Connecting to the unix socket.");
    memset(&addr, 0, sizeof(struct sockaddr_un));
    printf(".");
    addr.sun_family = AF_UNIX;
    strncpy(addr.sun_path, UNIX_SOCKET_NAME, sizeof(addr.sun_path) - 1);
    printf(".");

    rv = connect(s,
            (const struct sockaddr *) &addr,
            sizeof(struct sockaddr_un));

    if (rv < 0) {
        printf("Failed.\n");
        // STUB: Use strerror()
        fprintf(stderr, "ERROR: Failed to connect to the Unix socket.\n");
        exit(1);
    }
    printf("Done.\n");

    printf("Reporting up...");
    snprintf(mbuf, BUFFER_SIZE, "ClientUp(%d)", getpid());
    printf(".");
    rv = write(s, mbuf, strlen(mbuf));
    if (rv < 0) {
        printf("Failed.\n");
        // STUB: Use strerror()
        fprintf(stderr, "ERROR: Failed to send data on the Unix socket.\n");
        exit(1);
    }
    printf("Done.\n");


    
    while ( continue_loop )
    {
        // STUB: This is incompatible with the shutdown.
        // STUB: The shutdown should probably just be handled in the signal handler.
        sleep(3);
        



        printf("Reporting heartbeat...");
        snprintf(mbuf, BUFFER_SIZE, "HeartBeat");
        printf(".");
        rv = write(s, mbuf, strlen(mbuf));
        if (rv < 0) {
            printf("Failed.\n");
            // STUB: Use strerror()
            fprintf(stderr, "ERROR: Failed to send data on the Unix socket.\n");
            exit(1);
        }
        printf("Done.\n");
        
    }






    // STUB: Just exit for now.
    close(s);
    exit(0);

}