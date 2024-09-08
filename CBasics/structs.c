#include <stdio.h>
#include <stdbool.h>

#define MAX_IDS 32
#define MAX_EMPLOYEES 1000

struct employee {
    int id;
    char firstName[64];
    char lastName[64];
    float income;
    bool isManager;
};

int main() {
    struct employee employees[MAX_EMPLOYEES];

    int i = 0;

    for (i = 0; i < MAX_EMPLOYEES; i++) {
        employees[i].income = 100.0;
        employees[i].isManager = false;
    }

    printf("%f\n%d\n%s\n", employees[10].income, employees[10].isManager, employees[10].firstName);

    return 0;

}