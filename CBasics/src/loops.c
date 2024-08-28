#include <stdio.h>
#include <string.h>

#define MAX_IDS 32

int main() {
   int ids[MAX_IDS] = {0};

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

