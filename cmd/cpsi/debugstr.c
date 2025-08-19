#include <errno.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "debugstr.h"

#define DEBUG_ENV_VAR_NAME "CPSI_VERBOSE_BN"

#define FN_BUFFER_SIZE 80


static FILE *debug_fp;

/* ======================================================================== */

void establish_debug_stream(void)
{
    char debug_fn[FN_BUFFER_SIZE];
    char *env_basename;

    if ( NULL != (env_basename = getenv(DEBUG_ENV_VAR_NAME) ) )
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

/* ======================================================================== */

void log_printf(const char *fmt, ...)
{
    va_list args;

    // Don't try to write to no-file.
    if (debug_fp == NULL)
        return;

    va_start(args, fmt);
    vfprintf(debug_fp, fmt, args);
    va_end(args);

    // Always flush.
    fflush(debug_fp);
}

/* ======================================================================== */

void err_printf(const char *fmt, ...)
{
    va_list args;

    va_start(args, fmt);
    if (debug_fp != NULL)
    {
        vfprintf(debug_fp, fmt, args);
        fflush(debug_fp);
    }
    vfprintf(stderr, fmt, args);
    va_end(args);
}

/* ======================================================================== */

void close_debug_stream(void)
{
   // Don't try to close something not open.
    if (debug_fp == NULL)
        return;

    fclose(debug_fp);
    debug_fp = NULL;
}