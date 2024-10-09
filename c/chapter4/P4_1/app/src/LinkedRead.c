#include <stdio.h>
#include<stdlib.h>
#include "../include/LinkedRead.h"

void LinkedRead() {
    Node *head = NULL;
    Node *tail = NULL;
    Node *cur = NULL;

    Node *newNode = NULL;
    int readData;

    while(1) {
        printf("자연수 입력: ");
        scanf("%d", &readData);
        if(readData < 1) {
            break;
        }

        newNode = (Node*)malloc(sizeof(Node));
        newNode->data = readData;
        newNode->next = NULL;

        if(tail == NULL) {
            tail = newNode;
        } else {
            head->next = newNode;
        }
        head = newNode;
    }
    printf("\n");

    // 입력 받은 데이터의 출력 과정 //////
    printf("입력 받은 데이터의 전체출력! \n");
    if(head == NULL) {
        printf("저장된 자연수가 존재하지 않습니다. \n");
    } else {
        cur = tail;
        printf("%d ", cur->data);

        while(cur->next != NULL) {
            cur = cur->next;
            printf("%d ", cur->data);
        }
    }
    printf("\n\n");

    // 메모리 해제과정 //
    if(tail == NULL) {
        return;
    } else {
        Node * delNode = tail;
        Node * delNextNode = tail->next;

        printf("%d을(를) 삭제합닏다. \n", tail->data);
        free(delNode);

        while(delNextNode != NULL) {
            delNode = delNextNode;
            delNextNode = delNextNode->next;

            printf("%d을(를) 삭제합니다. \n", delNode->data);
            free(delNode);
        }
    }
}