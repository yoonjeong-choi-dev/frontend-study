#include <stdio.h>
void PrintHelloWorld() {
    printf("====================\n");
    printf("This is from C file\n");
    printf("Hello World~!\n");
    printf("====================\n");
}

int Factorial(int x) {
    if(x<=1) return x;
    return Factorial(x-1) * x;
}