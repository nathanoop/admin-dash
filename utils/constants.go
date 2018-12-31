package utils

import "github.com/nathanoop/admin-dash/models"

const (
	/* APP CONSTANTS START */
	DEF_PAGE_COUNT = 10
	/* APP CONSTANTS END */

	/* SITE URLS  START */
	SITE_URL_INDEX                 = "/"
	SITE_URL_ADMIN_LOGIN           = "/login"
	SITE_URL_ADMIN_AUTHENTICATE    = "/authenticate"
	SITE_URL_ADMIN_DASHBOARD       = "/dashboard"
	SITE_URL_ADMIN_LOGOUT          = "/logout"
	SITE_URL_ADMIN_CREATE          = "/admin/create"
	SITE_URL_ADMIN_SAVE            = "/admin/save"
	SITE_URL_ADMIN_LIST            = "/admin/list"
	SITE_URL_ADMIN_EDIT            = "/admin/edit"
	SITE_URL_ADMIN_UPDATE          = "/admin/update"
	SITE_URL_ADMIN_CHANGE_USERNAME = "/admin/changeusername"
	SITE_URL_ADMIN_CHANGE_PASSWORD = "/admin/changepassword"
	SITE_URL_ADMIN_DELETE          = "/admin/delete"
	SITE_URL_PROFILE               = "/profile"
	SITE_URL_SETTING               = "/settings"
	/* SITE URLS  END */

	/* ERROR MESSAGE CODE START */
	ERR_LOGIN_INV_USR_PSS       = "ERRLGN001"
	ERR_LOGIN_INV_USR           = "ERRLGN002"
	ERR_LOGIN_INC_PSSWD         = "ERRLGN003"
	ERR_LOGIN_INV_TOKEN         = "ERRLGN004"
	ERR_LOGGED_OUT              = "ERRLGN005"
	ERR_LOGIN_DELETED_ACCNT     = "ERRLGN06"
	ERR_LOGIN_EMPTY_FIELDS      = "ERRLGN07"
	ERR_ADM_CREATE_REQ_FIELDS   = "ERRADM001"
	ERR_ADM_CREATE_PWD_ERR      = "ERRADM002"
	ERR_ADM_CREATE_PWD_HSH_ERR  = "ERRADM003"
	ERR_ADM_CREATE_INS_ERR      = "ERRADM004"
	ERR_ADM_CREATE_DUP_USER_ERR = "ERRADM005"
	ERR_DASHBRD_INV_ADMINID     = "ERRDASH005"
	ERR_ADM_UPDATE_UP_ERR       = "ERRADM006"
	ERR_ADMIN_LISTING           = "ERRADM007"
	ERR_ADM_DEL_SELF            = "ERRADM008"
	ERR_ADM_DELETE_ERR          = "ERRADM009"
	/* ERROR MESSAGE CODE END */
)

type (
	//Message Obect
	//to showerror,success message
	Message struct {
		Code    string
		Message string
		Class   string
	}
	//Admintoken Obect
	//has site token & login session object
	Admintoken struct {
		Token  string
		AdmObj models.Admin
	}

	//Page Obect
	//has page detailsfor List pagination
	Pageobj struct {
		From  int
		To    int
		Total int
		Pagesl int
		Pagesarr int[]
	}
)
