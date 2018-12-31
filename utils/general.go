package utils

import (
	"log"
	"regexp"
)

func Errormessage(errCode string) string {
	var ErrorMessage = make(map[string]string)
	ErrorMessage["ERRLGN001"] = "Username not found."
	ErrorMessage["ERRLGN002"] = "Incorrect password."
	ErrorMessage["ERRLGN003"] = "Username, password not present."
	ErrorMessage["ERRLGN004"] = "Invalid Session token."
	ErrorMessage["ERRLGN005"] = "User logged out"
	ErrorMessage["ERRLGN06"] = "Deleted user account, cannot login"
	ErrorMessage["ERRLGN07"] = "Username, password empty"
	ErrorMessage["ERRADM001"] = "Required fields missing."
	ErrorMessage["ERRADM002"] = "Password, Confirm password not the same."
	ErrorMessage["ERRADM003"] = "Password hashing error."
	ErrorMessage["ERRADM004"] = "Admin user  Insert Failed"
	ErrorMessage["ERRADM005"] = "Duplicate user with same username"
	ErrorMessage["ERRDASH005"] = "Select a valid admin user to edit"
	ErrorMessage["ERRADM006"] = "Admin user  update Failed"
	ErrorMessage["ERRADM007"] = "Admin user listing Failed"
	ErrorMessage["ERRADM008"] = "Admin cannot delete self"
	ErrorMessage["ERRADM009"] = "Admin user delete failed"
	return ErrorMessage[errCode]
}

func Getqueryparams(queryObj map[string][]string, name string) string {
	if queryObj != nil && queryObj[name] != nil {
		value := queryObj[name][0]
		return value
	}
	return ""
}

func Notificationobj(queryObj map[string][]string) Message {
	msg := ""
	var viewObj = Message{}
	if queryObj != nil && queryObj["msg"] != nil {
		msg = queryObj["msg"][0]
		viewObj = Notificationobjfromstr(msg)
		return viewObj
	}
	return viewObj
}

func Notificationobjfromstr(msg string) Message {
	msgStr, cls := "", ""
	if msg != "" {
		msgStr = Errormessage(msg)
		cls = "alert-danger"
	}
	viewObj := Message{msg, msgStr, cls}
	log.Println("Notificationobjfromstr", viewObj)
	return viewObj
}

func getBrowser(uastr string) string {
	regexpStr := `(?i)(firefox|msie|chrome|safari)[/\s]([/\d.]+)`
	r, _ := regexp.Compile(regexpStr)
	match := r.FindString(uastr)
	return match
}
