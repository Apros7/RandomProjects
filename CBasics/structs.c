#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

#define MAX_IDS 32
#define MAX_EMPLOYEES 1000

struct __attribute__((__packed__)) employee {
    int id;
    char firstName[64];
    char lastName[64];
    float income;
    bool isManager;
};

// Without __attribute__ size is 140
// With __attribute__ size is 137

union test_u {
    int x;
    char c;
};

int main() {
    int n = 100;
    struct employee *employees = malloc(sizeof(struct employee) * n);

    printf("size of employee: %lu\n", sizeof(struct employee));
        
    int i = 0;

    for (i = 0; i < n; i++) {
        employees[i].income = 100.0;
        employees[i].isManager = false;
    }

    printf("%f\n%d\n%s\n", employees[10].income, employees[10].isManager, employees[10].firstName);

    free(employees);
    employees = NULL;

    return 0;

}