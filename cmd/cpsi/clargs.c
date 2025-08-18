#include <string.h>
#include <stdlib.h>
#include <stdio.h>

#include "clargs.h"


/* ======================================================================== */

CLArgs *parse_cmdline(int argc, char *argv[])
{
    CLArgs *cla;
    int i;
    char c;
    unsigned long long nano_one, nano_two;

    // Allocate it.
    if (NULL == (cla = malloc(sizeof(CLArgs))))
    {
        // Something went horribly wrong. Return NULL.
        return NULL;
    }

    // Clear it.
    memset(cla, 0, sizeof(CLArgs));

    // Start processing arguments.

    // There should be three arguments (to the base command).
    if ( argc != 4 )
    {
        sprintf(cla->error, "Expected three arguments: <dest> <monitor_target> <monitor_args>");
        return cla;
    }

    // argv[1] = sendto

    // Some basic rules about the first arg...
    // - Should be absolute path, or
    // - Should be (exactly) 'stdout'.

    if ( 0 == strcmp(argv[1], "stdout") )
    {
        // Will send to stdout. (Effectively a debug option.)
        strcpy(cla->sendto, "stdout");
    }
    else
    {
        if ( argv[1][0] != '/' )
        {
            sprintf(cla->error, "The Unix socket must be an absolute path.");
            return cla;
        }
        else
        {
            if (strlen(argv[1]) > 79)
            {
                sprintf(cla->error, "The Unix socket name length is out of bounds.");
                return cla;
            }

            // This error should not be a thing - i just checked for it above.
            if ( NULL == strncpy(cla->sendto, argv[1], MAX_CLARG_LENGTH))
            {
                sprintf(cla->error, "Failed to copy the Unix socket filename.");
                return cla;
            }
        }
    }

    // argv[2] = target (the filename in /proc/pressure)

    // Solitary rule for second arg...
    // - Must match one of the known items.

    if ( 0 == strcmp(argv[2], "cpu"))
    {
        strcpy(cla->target, "/proc/pressure/cpu");
    }

    if ( 0 == strcmp(argv[2], "io"))
    {
        sprintf(cla->target, "/proc/pressure/io");
    }

    if ( 0 == strcmp(argv[2], "irq"))
    {
        sprintf(cla->target, "/proc/pressure/irq");
    }

    if ( 0 == strcmp(argv[2], "memory"))
    {
        sprintf(cla->target, "/proc/pressure/memory");
    }

    if ( cla->target[0] == 0  )
    {
        printf(cla->error, "Target (second argument) was not recognized/valid target.");
        return cla;
    }

    // argv[3] = (threshold) args to write to the proc file

    // Rules for the third argument
    // - Must be three 'words' (two spaces).
    // - First word must be "some" or "full".
    // - Second and third words must be numeric.
    // - Second < third.

    if (( argv[3] != strstr(argv[3], "some ") ) && ( argv[3] != strstr(argv[3], "full ") ))
    {
        strcpy(cla->error, "The first word in the args option must be \"some\" or \"full\".");
        return cla;
    }

    // A space was tested for above so the first checked item must be 'something'.
    nano_one = 0;
    i = 5;
    c = argv[3][i];
    while (( c != ' ' ) && ( c != 0 ))
    {
        if ((c >= '0') && ( c <= '9'))
        {
            nano_one *= 10;
            nano_one += c - '0';
        }
        else
        {
            strcpy(cla->error, "Encountered a non-numeric in the first numeric argument.");
            return cla;           
        }
        i++;
        c = argv[3][i];
    }

    nano_two = 0;
    i++;
    c = argv[3][i];
    while ( c != 0 )
    {
        if ((c >= '0') && ( c <= '9'))
        {
            nano_two *= 10;
            nano_two += c - '0';
        }
        else
        {
            strcpy(cla->error, "Encountered a non-numeric in the second numeric argument.");
            return cla;           
        }
        i++;
        c = argv[3][i];
    }

    if ( nano_one > nano_two )
    {
        strcpy(cla->error, "The argument sampling threshold is greater than the size.");
        return cla;           
    }

    strcpy(cla->args, argv[3]);

    return cla;
}