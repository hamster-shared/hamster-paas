package utils

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(toEmail, requestId, result, error string) {
	from := os.Getenv("FROM_EMIAl")
	password := os.Getenv("EMAIL_PASSWORD")
	to := toEmail
	subject := "Test email from Golang with HTML content"
	body := "<html>\n<head>\n\t<meta charset=\"UTF-8\">\n\t<title>Email Template</title>\n\t<style type=\"text/css\">\n\t\tbody {\n\t\t\tbackground-color: #f6f6f6;\n\t\t\tmargin: 0;\n\t\t\tpadding: 0;\n\t\t\tfont-family: Arial, sans-serif;\n\t\t\tfont-size: 14px;\n\t\t\tline-height: 1.5;\n\t\t\tcolor: #333333;\n\t\t}\n\t\t.container {\n\t\t\tmax-width: 600px;\n\t\t\tmargin: 0 auto;\n\t\t\tbackground-color: #ffffff;\n\t\t\tpadding: 20px;\n\t\t\tborder-radius: 10px;\n\t\t\tbox-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);\n\t\t}\n\t\th1 {\n\t\t\tfont-size: 28px;\n\t\t\tfont-weight: bold;\n\t\t\tmargin-top: 0;\n\t\t\tmargin-bottom: 20px;\n\t\t\ttext-align: center;\n\t\t\tcolor: #333333;\n\t\t}\n\t\tp {\n\t\t\tmargin: 0 0 20px;\n\t\t}\n\t</style>\n</head>\n<body>\n\t<div class=\"container\">\n\t\t<h1>请求结果</h1>\n\t\t<div>\n\t\t    <p>RequestId:</p>\n\t\t    <p>" + requestId + "</p>\n\t\t</div>\n\t\t<div>\n\t\t    <p>Result:</p>\n\t\t    <p>" + result + "</p>\n\t\t</div>\n\t\t<div>\n\t\t    <p>Error:</p>\n\t\t    <p>" + error + "</p>\n\t\t</div>\n\t</div>\n</body>"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-type: text/html\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Println(err.Error())
		log.Fatal(err)
	}
	log.Println("Email sent successfully!")
}
