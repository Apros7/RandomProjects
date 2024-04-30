#include <stdio.h>
#include <string.h>

int main() {
    char name[6];
    int age;
    printf("Hello. Please enter your name\n");
    scanf("%5s", name);
    printf("And your age\n");
    scanf("%d", &age);
    if (strcmp(name, "Lucas") == 0 && age == 21) {
        printf("Hello %s.", name);
    } else {
        printf("Wrong Name, %s", name);
    };
    return 0;
}
// counterfactuals: what if