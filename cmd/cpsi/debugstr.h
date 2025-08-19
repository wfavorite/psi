#ifndef DEBUGSTR_H
#define DEBUGSTR_H

#include <stdio.h>
#include <stdarg.h>

// establish_debug_stream conditionally opens a debug file.
void establish_debug_stream(void);

// log_printf conditionally writes to the open debug file.
void log_printf(const char *fmt, ...);

// err_printf writes to stderr and conditionally writes to the debug file.
void err_printf(const char *fmt, ...);

// close_debug_stream conditionally closes the debug file.
void close_debug_stream(void);

#endif
