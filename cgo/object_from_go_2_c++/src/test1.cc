#include "person/person.h"

#include <stdio.h>
// using namespace goobj2c;

int main(){
    auto p = new Person("GOPPHER",15);

    char buf[64];
    char* name = p->GetName(buf,sizeof(buf)-1);
    int age = p->GetAge();

    printf("%s,%d years old.\n",name,age);
    delete p;

    return 0;
}