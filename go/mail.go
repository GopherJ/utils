package util

/**
 * mail.go
 * 2018-01-01 00:32:54
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "bytes"
	"crypto/tls"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender     string
	Receivers  []string
	Subject    string
	Body       []byte

    smtp      *Smtp
    auth       smtp.Auth
}

type Smtp struct {
	host string
	port string
}

func (s *Smtp) Addr()(string) {
	return s.host + ":" + s.port
}

func (mail *Mail) SetSmtp(host, port string)(*Mail) {
    mail.smtp = &Smtp{
        host,
        port,
    }
    return mail
}

func (mail *Mail) SetAuth(user, password string)(*Mail) {
    mail.auth = smtp.PlainAuth("", user, password, mail.smtp.host)
    return mail
}

func (mail *Mail) Data()([]byte) {
    b := new(bytes.Buffer)

    b.WriteString("From: ")
    b.WriteString(mail.Sender)
    b.WriteString("\r\n")

    b.WriteString("To: ")
    b.WriteString(strings.Join(mail.Receivers, ";"))
    b.WriteString("\r\n")

    b.WriteString("Subject: ")
    b.WriteString(mail.Subject)
    b.WriteString("\r\n")

    b.WriteString("\r\n")

    b.Write(mail.Body)

	return b.Bytes()
}


func (mail *Mail) Send()(error) {

	data := mail.Data()

	tlsconfig := &tls.Config{
        // client accepts any certificate
		InsecureSkipVerify  :   true,
		ServerName          :   mail.smtp.host,
	}

	conn, err := tls.Dial("tcp", mail.smtp.Addr(), tlsconfig)
	if err != nil {
        return err
	}

	client, err := smtp.NewClient(conn, mail.smtp.host)
	if err != nil {
        return err
	}

	if err = client.Auth(mail.auth); err != nil {
        return err
	}

	if err = client.Mail(mail.Sender); err != nil {
        return err
	}
	for _, receiver := range mail.Receivers {
		if err = client.Rcpt(receiver); err != nil {
            return err
		}
	}

	w, err := client.Data()
	if err != nil {
        return err
	}

	_, err = w.Write(data)
	if err != nil {
        return err
	}

	err = w.Close()
	if err != nil {
        return err
	}

	client.Quit()

    return nil
}
