#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <stdbool.h>

#define MAX_STR 32

bool str_compare(char* str1, char* str2) {
    int min_len;
    if (strlen(str1) > strlen(str2)) {min_len = strlen(str1);}
    else {min_len = strlen(str2);}

    for (int i = 0; i < min_len; i++) {
        if (str1[i] != str2[i]) { return false; }
    }
    return true;
}

int main() {

    char name[MAX_STR];
    printf("Please enter your name: ");
    scanf("%s", name);
    for (int i = 0; name[i]; i++) {
        name[i] = tolower(name[i]);
    }
    printf("Your name is: %s", name);

    if (str_compare(name, "lucas")) {
        printf("\nHey Lucas, you're a champ!\n");
    } else if (str_compare(name, "emilie")) {
        printf("\nHeyyy beautiful!\n");
    } else {
        printf("\nWho are you really?\n");
    }

    return 0;
}