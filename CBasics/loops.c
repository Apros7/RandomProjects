#include <stdio.h>
#include <string.h>

#define MAX_IDS 32

int add(int x, int y) {
   int z = x+y;
   return z;
}

int main() {
   int ids[MAX_IDS] = {0};
   int other_int = add(1,2);
   printf("%d\n", other_int);

   for (int i = 0; i < MAX_IDS; i++) {
    ids[i] = i;
    printf("%d: %d\n", i, ids[i]);
   }
   
   int i = 0;
   memset(ids, 0, sizeof(ids));
   while (i < MAX_IDS) {
    ids[i] = i;
    printf("%d: %d\n", i, ids[i]);
    i++;
   }
   return 0;
}

