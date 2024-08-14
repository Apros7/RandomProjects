#include <stdio.h>
// #define DEBUG // Or use -DDEBUG at compile time

int main() {
    int age = 21;
    char name[6] = "Lucas";
    printf("Age: %d, Name: %s", age, name);

    #ifdef DEBUG
    printf("\nWE ARE IN DEBUG MODE %s:%d", __FILE__, __LINE__);
    #endif
}