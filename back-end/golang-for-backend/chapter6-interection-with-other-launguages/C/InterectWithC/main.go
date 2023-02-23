package main

//#include <stdio.h>
//#include <stdlib.h>
//void PrintString(char* str){
//  printf("====================\n");
//  printf("This is C code\n");
//	printf("%s\n", str);
//  printf("====================\n");
//}
//char* GetName(int idx) {
//	if(idx == 1) return "Yoonjeong";
//	else if(idx == 2) return "YJ";
//	else return "Anonymous";
//}
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Pass a pointer to C")
	strPtrToPass := C.CString("This is from Go!")
	C.PrintString(strPtrToPass)
	C.free(unsafe.Pointer(strPtrToPass))

	fmt.Println("\nGet a value from C")
	for i := 1; i < 4; i++ {
		cStr := C.GetName(C.int(i))
		fmt.Println(C.GoString(cStr))
	}
}
