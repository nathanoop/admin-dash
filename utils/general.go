package utils

func getErrorMessageResponse(errCode string) string {
	var ErrorMessage = make(map[string]string)
	ErrorMessage["ERRLGN001"] = "Username not found."
	ErrorMessage["ERRLGN002"] = "Incorrect password."
	ErrorMessage["ERRLGN003"] = "Username, password not present."
	ErrorMessage["ERRLGN004"] = "Invalid Session token."
	ErrorMessage["ERRLGN005"] = "User logged out"
	ErrorMessage["ERRLGN06"] = "Deleted user account, cannot login"
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
