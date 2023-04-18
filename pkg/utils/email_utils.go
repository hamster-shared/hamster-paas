package utils

import (
	"log"
	"net/smtp"
	"os"
	"time"
)

func SendEmail(toEmail, requestId, result, requestName, errorInfo string) {
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	to := toEmail
	timeData := time.Now()
	timeFormat := timeData.Format("2006-01-02 15:04:05")
	subject := "Test Results"
	flag := "Success"
	if errorInfo != "" {
		flag = "Failed"
	}
	body := "<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Test Results</title>\n    <style>\n\n        .content {\n            background-color: white;\n            padding: 20px;\n        }\n\n        .content .head {\n            font-family: BinancePlex, Arial, PingFangSC-Regular, 'Microsoft YaHei', sans-serif;\n            font-size: 20px;\n            font-weight: 900;\n            line-height: 25px;\n            text-align: left;\n            color: #000000;\n        }\n\n        .content .label {\n            font-size: 16px;\n            font-weight: 800;\n            line-height: 20px;\n            color: #000000;\n        }\n\n        .footer {\n            position: absolute;\n            bottom: 0;\n            right: 0;\n            padding: 10px;\n            font-size: 12px;\n            color: #000000;\n            text-align: right;\n        }\n\n        .footer a {\n            color: #000000;\n            text-decoration: none;\n        }\n\n        .wrapper {\n            position: relative;\n            min-height: 100%;\n        }\n\n        .content-wrapper {\n            padding-bottom: 60px;\n        }\n    </style>\n</head>\n<body>\n    <div class=\"wrapper\">\n<div class=\"content-wrapper\">\n            <div class=\"content\">\n                <div class=\"head\">Test Results</div>\n                <p>Your request test results are as follows:</p>\n                <table>\n                    <tr>\n                        <td class=\"label\">Request Name:</td>\n                        <td>" + requestName + "</td>\n                    </tr>\n                    <tr>\n                        <td class=\"label\">Request ID:</td>\n                        <td>" + requestId + "</td>\n                    </tr>\n                    <tr>\n                        <td class=\"label\">Send Time:</td>\n                        <td>" + timeFormat + "</td>\n                    </tr>\n                    <tr>\n                        <td class=\"label\">Test Results:</td>\n                        <td>" + flag + "</td>\n                    </tr>\n                    <tr>\n                        <td class=\"label\">Reason:</td>\n                        <td>" + errorInfo + "</td>\n                    </tr>\n                </table>\n            </div>\n        </div>\n        <div class=\"footer\">\n            <p>Hamster team</p>\n            <p><a href=\"http://www.hamsternet.io\">www.hamsternet.io</a></p>\n        </div>\n    </div>\n</body>\n</html>\n"
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
