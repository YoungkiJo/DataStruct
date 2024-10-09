#ifndef __LINKED_LIST_H__
#define __LINKED_LIST_H__

typedef struct _node {
    int data;
    struct _node *next;
} Node;

void LinkedRead();

#endif