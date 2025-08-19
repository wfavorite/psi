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

    log_printf("Creating the socket...");
    use_unix_socket = socket(AF_UNIX, SOCK_STREAM, 0);
    if (use_unix_socket < 0)
    {
        log_printf("Failed.\n");
        err_printf("ERROR: socket failed - %s\n", strerror(errno));
        exit(1);
    }
    log_printf("Done.\n");

    log_printf("Connecting to the unix socket.");
    memset(&addr, 0, sizeof(struct sockaddr_un));
    log_printf(".");
    addr.sun_family = AF_UNIX;
    strncpy(addr.sun_path, filename, sizeof(addr.sun_path) - 1);
    log_printf(".");

    rv = connect(use_unix_socket,
        (const struct sockaddr *) &addr,
        sizeof(struct sockaddr_un));

    if (rv < 0) {
        log_printf("Failed.\n");
        err_printf("ERROR: connect failed - %s\n", strerror(errno));
        exit(1);
    }
    log_printf("Done.\n");

    log_printf("Reporting up...");
    snprintf(mbuf, BUFFER_SIZE, "ClientUp(%d)", getpid());
    log_printf(".");
    rv = write(use_unix_socket, mbuf, strlen(mbuf));
    if (rv < 0)
    {
        log_printf("Failed.\n");
        err_printf("ERROR: Failed to send data on the Unix socket - %s\n", strerror(errno));
        exit(1);
    }
    log_printf("Done.\n");
   
}

/* ======================================================================== */

void send_threshold_event(void)
{
    int rv;

    // 'Always' (conditionally) log it.
    log_printf("ThresholdEvent\n");

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
        err_printf("ERROR: Failed Unix socket write - %s\n", strerror(errno));
        exit(1);
    }
}

/* ======================================================================== */

void send_heartbeat(void)
{
    int rv;

    // 'Always' (conditionally) log it.
    log_printf("HeartBeat\n");

    if ( use_unix_socket < 0 )
    {
        printf("HeartBeat\n");
        fflush(stdout);
        return;
    }

    snprintf(mbuf, BUFFER_SIZE, "HeartBeat");
    rv = write(use_unix_socket, mbuf, strlen(mbuf));
    if (rv < 0)
    {
        err_printf("ERROR: Failed Unix socket write - %s\n", strerror(errno));
        exit(1);
    } 
}

/* ======================================================================== */

void shutdown_socket(void)
{
    log_printf("Closing Unix socket...");
    if ( use_unix_socket < 0 )
    {
        log_printf("Skipped.\n");
        return;
    }

    close(use_unix_socket);
    log_printf("Done.\n");
}
