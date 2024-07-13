package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/Anacardo89/tpsi25_blog.git/auth"
	"github.com/Anacardo89/tpsi25_blog.git/db"
	"github.com/Anacardo89/tpsi25_blog.git/logger"
	"github.com/Anacardo89/tpsi25_blog.git/rabbit"
)

type RegisterData struct {
	Email    string `json:"email"`
	Subject  string `json:"subject"`
	MailBody string `json:"mail_body"`
}

func isValidInput(input string) bool {
	if strings.Contains(input, ";") {
		return false
	}
	return true
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	// Parse Form
	err := r.ParseForm()
	if err != nil {
		logger.Error.Println(err.Error())
		fmt.Fprintln(w, err.Error())
		return
	}
	var userReg = &auth.User{
		UserName:  r.FormValue("user_name"),
		UserEmail: r.FormValue("user_email"),
		UserPass:  r.FormValue("user_password"),
	}
	pass2 := r.FormValue("user_password2")
	if userReg.UserPass != pass2 {
		fmt.Fprintln(w, "Password strings don't match")
		return
	}
	if !isValidInput(userReg.UserName) || !isValidInput(userReg.UserEmail) || !isValidInput(userReg.UserPass) {
		fmt.Fprintln(w, "Invalid character in form")
		return
	}

	// Check if UserName or Email in use
	dbUser := &db.User{}
	err = db.Dbase.QueryRow(db.SelectUserByName,
		userReg.UserName).
		Scan(dbUser.UserName)
	if err == nil {
		fmt.Fprintln(w, "User alreay exists")
		return
	}
	err = db.Dbase.QueryRow(db.SelectUserByEmail,
		userReg.UserEmail).
		Scan(dbUser.UserEmail)
	if err == nil {
		fmt.Fprintln(w, "Email already exists")
		return
	}

	// Password Hashing
	userReg.HashedPass, err = auth.HashPassword(userReg.UserPass)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	// Send Regiter Mail to Queue
	regData := RegisterData{
		Email: userReg.UserEmail,
	}
	mailSubject, err := template.New("registerSubject").Parse(registerSubject)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	mailBody, err := template.New("registerBody").Parse(registerBody)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	var mbuf bytes.Buffer
	err = mailSubject.Execute(&mbuf, userReg)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	regData.Subject = mbuf.String()
	mbuf.Reset()
	err = mailBody.Execute(&mbuf, userReg)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	regData.MailBody = mbuf.String()
	data, err := json.Marshal(regData)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	err = rabbit.RabbitMQ.MQSendRegMail(data)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	// Insert User in DB
	_, err = db.Dbase.Exec(db.InsertUser,
		userReg.UserName, userReg.UserEmail, userReg.HashedPass, 1)
	if err != nil {
		logger.Error.Println(err.Error())
		fmt.Fprintln(w, err.Error())
		return
	}

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}
