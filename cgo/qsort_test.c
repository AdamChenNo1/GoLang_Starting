#include <stdio.h>
#include <stdlib.h>

#define DIM(X) (sizeof(x)/sizeof((x)[0]))

static int cmp(const void* a,const void* b) {
    const void* pa = (int*)a;
    const void* pb = (int*)b;
    return *pa - *pb;
}

int main() {
    int values[] = {42,18,36,57,79,6}
    int i;

    qsort(values,DIM(values),sizeof(values[0]),cmp);

    for(i = 0;i < DIM(values);i++){
        printf("%d",values[i]);
    }
    return 0;
}