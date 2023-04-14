package main

import (
	"io"
	"net/mail"
	"os"
	parsing_mail "parsing-mail"
	"strings"
)

const exampleMail string = `Date: Thu, 13 Apr 2023 08:00:00 -0700
From: Yoonjeong Choi <fake_sender@example.com>
To: YJ <fake_receiver@example.com>
Subject: Mail Parsing Test

This mail is for testing to parse mail content.
You can use any io.Writer to save the parsed result.
`

func main() {
	writer := os.Stdout
	reader := strings.NewReader(exampleMail)
	message, err := mail.ReadMessage(reader)
	if err != nil {
		panic(err)
	}

	parsing_mail.PrintHeaderInfo(writer, message.Header)

	if _, err := io.Copy(writer, message.Body); err != nil {
		panic(err)
	}
}
