package main

//#include <stdio.h>
//void printHelloWorld() {
//	printf("Hello World~!\n");
//  printf("This is written in go file\n");
//}
import "C"

func main() {
	C.printHelloWorld()
}
