#include <stdio.h>

#define MAX_LST 32
#define MAX_STR 10

int main() {

    int lst[MAX_LST] = {0,1,2};
    printf("%d\n", lst[0]);

    char str[MAX_STR] = {'H', 'E', 'L', 'L', 'O', 0};
    char *another_str = "hello";
    printf("%s\n", str);
    printf("%s\n", another_str);

    return 0;
}