#include <stdio.h>

int main() {
    char name[6];
    printf("Hello. Please enter your name\n");
    scanf("%5s", name);
    printf("Hello %s.", name);
    return 0;
}