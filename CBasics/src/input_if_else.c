#include <stdio.h>
#include <string.h>

int main() {
    char name[6];
    printf("Hello. Please enter your name\n");
    scanf("%5s", name);
    if (strcmp(name, "Lucas") == 0) {
        printf("Hello %s.", name);
    } else {
        printf("Wrong Name, %s", name);
    };
    return 0;
}
// counterfactuals: what if