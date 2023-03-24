package bytes_and_strings

import (
	"fmt"
	"io"
	"strings"
)

func SearchString(str string) {
	fmt.Printf("Test String: %s\n", str)
	fmt.Printf("Contain 'this' :%t\n", strings.Contains(str, "this"))
	fmt.Printf("Contain one of [a,b,c]: %t\n", strings.ContainsAny(str, "abc"))
	fmt.Printf("Start with 'this': %t\n", strings.HasPrefix(str, "this"))
	fmt.Printf("End with 'test': %t\n", strings.HasSuffix(str, "test"))
}

func ModifyString(str string) {
	fmt.Printf("Test String: %s\n", str)
	fmt.Printf("Split with space: %#v\n", strings.Split(str, " "))
	fmt.Printf("Upper Case for each word first char: %s\n", strings.Title(str))
	fmt.Printf("Trim start and end spaces: %s\n", strings.TrimSpace(str))
}

func StringReader(str string) io.Reader {
	return strings.NewReader(str)
}

func StringReaderToWriter(str string, w io.Writer) {
	r := StringReader(str)
	io.Copy(w, r)
}
