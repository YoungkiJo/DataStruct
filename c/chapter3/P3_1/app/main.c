#include <stdio.h>
#include "./include/ArrayList.h"

int main() {
    List list;
    int data;
    ListInit(&list);

    LInsert(&list, 1); LInsert(&list, 2);
    LInsert(&list, 3); LInsert(&list, 4);
    LInsert(&list, 5); LInsert(&list, 6);
    LInsert(&list, 7); LInsert(&list, 8);
    LInsert(&list, 9);


    printf("현재 데이터의 수: %d \n", LCount(&list));

    if(LFirst(&list, &data)) {
        printf("%d ", data);

        while(LNext(&list, &data)) {
            printf("%d ", data);
        }
    }
    printf("\n\n");

    int sum = 0;
    if(LFirst(&list, &data)) {
        sum += data;

        while(LNext(&list, &data)) {
            sum += data;
        }
    }
    printf("현재 데이터의 총합: %d \n", sum);

    if(LFirst(&list, &data)) {
        if(data%2==0 || data%3==0) {
            LRemove(&list);
        }

        while(LNext(&list, &data)) {
            if(data%2==0 || data%3==0) {
                LRemove(&list); 
            }
            
        }
    }
    printf("현재 데이터의 수: %d \n", LCount(&list));

    if(LFirst(&list, &data)) {
        printf("%d ", data);

        while(LNext(&list, &data)) {
            printf("%d ", data);
        }
    }
    printf("\n\n");
    return 0;

}
