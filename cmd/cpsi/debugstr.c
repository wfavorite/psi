#include <errno.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "debugstr.h"


#define FN_BUFFER_SIZE 80


FILE *debug_fp;


void establish_debug_stream(void)
{
    char debug_fn[FN_BUFFER_SIZE];
    char *env_basename;

    if ( NULL != (env_basename = getenv("CPSI_VERBOSE_BN")) )
    {
        if (FN_BUFFER_SIZE < snprintf(debug_fn, FN_BUFFER_SIZE, "%s.%d", env_basename, getpid()))
        {
            fprintf(stderr, "ERROR: Debug stream ENV variable too large.\n");
            exit(1);
        }
    }
    else
    {
        sprintf(debug_fn, "/dev/null");
    }


    // Open the debug stream - whatever it might be.
    debug_fp = fopen(debug_fn, "w");
    if ( debug_fp == NULL )
    {
        fprintf(stderr, "ERROR: Failed to debug stream - %s\n", strerror(errno));
        exit(1);
    }

}

