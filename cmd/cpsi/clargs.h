#ifndef CLARGS_H
#define CLARGS_H


/*
    I just use the same arg length for all arguments.
    80 chars is sufficient to hold anything here.
*/
#define MAX_CLARG_LENGTH 80

/* ------------------------------------------------------------------------ */

typedef struct clargs
{
    char sendto[MAX_CLARG_LENGTH];
    char target[MAX_CLARG_LENGTH];
    char args[MAX_CLARG_LENGTH];
    char error[MAX_CLARG_LENGTH];

} CLArgs;



CLArgs *parse_cmdline(int argc, char *argv[]);



#endif
