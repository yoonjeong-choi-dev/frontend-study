package parsing_mail

import (
	"fmt"
	"io"
	"net/mail"
	"strings"
)

func PrintHeaderInfo(w io.Writer, header mail.Header) {
	// 단일 주소에 대해서만 파싱
	toAddress, err := mail.ParseAddress(header.Get("To"))
	if err == nil {
		fmt.Fprintf(w, "To: %s <%s>\n", toAddress.Name, toAddress.Address)
	}

	fromAddress, err := mail.ParseAddress(header.Get("From"))
	if err == nil {
		fmt.Fprintf(w, "From: %s <%s>\n", fromAddress.Name, fromAddress.Address)
	}

	fmt.Fprintf(w, "Subject: %s", header.Get("Subject"))

	if date, err := header.Date(); err == nil {
		fmt.Fprintf(w, "Date: %s\n", date.String())
	}

	fmt.Fprintln(w, strings.Repeat("=", 40))
	fmt.Fprintln(w)
}
