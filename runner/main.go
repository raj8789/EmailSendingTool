package main

import (
	"log"
	"net/smtp"
	"github.com/spf13/viper"
)

func main() {
	// Set up authentication information.

	// Set up the configuration file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Add the current directory as the config search path

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	from := viper.GetString("from")
	password := viper.GetString("password")
	host := viper.GetString("host")
	toread := viper.GetString("to")
	subject := viper.GetString("subject")
	body := viper.GetString("body")
	auth := smtp.PlainAuth("", from, password, host)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{toread}
	msg := []byte("To:" + toread + "\r\n" +
		"Subject:" + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
	err = smtp.SendMail(host+":587", auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Email Sent Successfully")
	}
}
