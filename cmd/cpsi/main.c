
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <sys/un.h>

#define UNIX_SOCKET_NAME "/tmp/cpsi.sock"

#define BUFFER_SIZE 1024

/*
    ToDo:
    [ ] Should not write to stdout. Perhaps to syslog, or a local log?
    [ ] The Unix socket path should probably be an environmental variable
        (passed by the parent).

    Done:

*/


int main(int argc, char *argv[])
{
    int s;                   /* Socket */
    struct sockaddr_un addr;
    int rv;                  /* Return value - multiple uses. */
    char mbuf[BUFFER_SIZE];
    int loop_exit;

    // STUB: You could check to see if this exists first.
    // STUB: ...and conditionally remove it.
    //printf("Removing the unix socket...");
    //unlink(UNIX_SOCKET_NAME);
    //printf("Done.\n");



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


    loop_exit = 1;
    while ( loop_exit )
    {
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