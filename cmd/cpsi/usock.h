#ifndef USOCK_H
#define USOCK_H

void init_unix_socket(char *filename);
void send_threshold_event(void);
void send_heartbeat(void);
void shutdown_socket(void);

#endif
