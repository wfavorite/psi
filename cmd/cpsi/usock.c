#include <sys/socket.h>
#include <sys/un.h>
#include <fcntl.h>
#include <unistd.h>
#include <stdlib.h>
#include <errno.h>

#include "debugstr.h"
#include "usock.h"

#define BUFFER_SIZE 1024

static int use_unix_socket;
static char mbuf[BUFFER_SIZE];

/* ======================================================================== */

void init_unix_socket(char *filename)
{
    //int s;
    int rv;
    struct sockaddr_un addr;
    

    // If we are not writing to stdout, then bail.
    if ( 0 == strcmp(filename, "stdout") )
    {
        use_unix_socket = -1;
        return;
    }

    fprintf(debug_fp, "Creating the socket..."); fflush(debug_fp);
    use_unix_socket = socket(AF_UNIX, SOCK_STREAM, 0);
    if (use_unix_socket < 0)
    {
        fprintf(debug_fp, "Failed.\n"); fflush(debug_fp);
        fprintf(stderr, "ERROR: socket failed - %s\n", strerror(errno));
        exit(1);
    }
    fprintf(debug_fp, "Done.\n"); fflush(debug_fp);

    fprintf(debug_fp, "Connecting to the unix socket."); fflush(debug_fp);
    memset(&addr, 0, sizeof(struct sockaddr_un));
    fprintf(debug_fp, "."); fflush(debug_fp);
    addr.sun_family = AF_UNIX;
    strncpy(addr.sun_path, filename, sizeof(addr.sun_path) - 1);
    fprintf(debug_fp, "."); fflush(debug_fp);

    rv = connect(use_unix_socket,
        (const struct sockaddr *) &addr,
        sizeof(struct sockaddr_un));

    if (rv < 0) {
        fprintf(debug_fp, "Failed.\n"); fflush(debug_fp);
        fprintf(stderr, "ERROR: connect failed - %s\n", strerror(errno));
        exit(1);
    }
    fprintf(debug_fp, "Done.\n"); fflush(debug_fp);

    printf("Reporting up..."); fflush(debug_fp);
    snprintf(mbuf, BUFFER_SIZE, "ClientUp(%d)", getpid());
    printf("."); fflush(debug_fp);
    rv = write(use_unix_socket, mbuf, strlen(mbuf));
    if (rv < 0)
    {
        printf("Failed.\n"); fflush(debug_fp);
        fprintf(stderr, "ERROR: Failed to send data on the Unix socket - %s\n", strerror(errno));
        exit(1);
    }
    fprintf(debug_fp, "Done.\n"); fflush(debug_fp);
   
}

/* ======================================================================== */

void send_threshold_event(void)
{
    int rv;

    if ( use_unix_socket < 0 )
    {
        printf("ThresholdEvent\n");
        fflush(stdout);
        return;
    }

    snprintf(mbuf, BUFFER_SIZE, "ThresholdEvent");
    rv = write(use_unix_socket, mbuf, strlen(mbuf));
    if (rv < 0)
    {
        fprintf(stderr, "ERROR: Failed Unix socket write - %s\n", strerror(errno));
        exit(1);
    }
}



/* ======================================================================== */

void shutdown_socket(void)
{
    if ( use_unix_socket < 0 )
        return;

    close(use_unix_socket);
}



  /*
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
        */
        