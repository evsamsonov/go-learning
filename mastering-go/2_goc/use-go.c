#include <stdio.h>
#include "c-shared.h"

// gcc -o use-go use-go.c ./c-shared.o

int main(int argc, char **argv) {
    GoInt a = 12;
    GoInt b = 23;

    printf("Call a Go function!\n");
    PrintMessage();

    GoInt p = Multiply(a, b);
    printf("Go func result: %d\n", (int)p);

    return 0;
}
