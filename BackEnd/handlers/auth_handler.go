package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"
	"uva-devtest/jwtauth"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
	"uva-devtest/restapi/operations/auth"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

const BearerCookieName string = "Bearer-Cookie"
const ReauthCookieName string = "ReAuth-Cookie"

const AuthHours int = 1
const ReauthHours int = 48

func CreateCookie(name string, token string, maxage int) *http.Cookie {

	cookie := &http.Cookie{
		Name:     name,
		Value:    token,
		Path:     "/",
		HttpOnly: true,                    // Evita ataques XSS
		Secure:   true,                    // Fuerza HTTPS
		MaxAge:   maxage,                  //Poner fin en 24h
		SameSite: http.SameSiteStrictMode, // Evita ataques XSRF
	}
	/*path := "/"
	maxage := "86400"
	samesite := "strict"
	cookie := fmt.Sprintf("%s=%s; Path=%s; Secure; SameSite=%s; HttpOnly; Max-Age=%s ", name, token, path, samesite, maxage)*/
	return cookie
}

func CreateDeprecatedCookie(name string) *http.Cookie {
	/*path := "/"
	samesite := "strict"
	cookie := fmt.Sprintf("%s=; Path=%s; Secure; SameSite=%s; HttpOnly; expires=Thu, 01 Jan 1970 00:00:00 GMT", name, path, samesite)*/
	return CreateCookie(name, "", 1)
}

func hoursToSeconds(hours int) int {
	return hours * 3600
}

func CreateJWTWrapper(u dao.User, expirationHours int64) jwtauth.JwtWrapper {
	var wrap jwtauth.JwtWrapper
	wrap.SecretKey = *u.Pwhash
	wrap.Issuer = "DevTest"
	wrap.ExpirationHours = expirationHours
	return wrap
}

func CreateJWT(u dao.User, expirationHours int64) (string, error) {
	wrap := CreateJWTWrapper(u, expirationHours)
	signedToken, err := wrap.GenerateToken(u.Email.String())
	log.Println(wrap.SecretKey, u.Email.String(), signedToken)
	return signedToken, err
}

func GetJWTModelUserCookies(cookieSlice []string, expectedName string, expectedHours int64) (*models.User, error) {
	for _, cookie := range cookieSlice {
		cookieName := cookie[0:14]
		var err error
		if expectedName == cookieName {
			token := cookie[14:]
			var email string
			email, err = jwtauth.GetEmailToken(token)
			if err == nil {
				var db *sql.DB
				db, err = dbconnection.ConnectDb()
				if err == nil {
					var u *dao.User
					u, err = dao.GetUserEmail(db, email)
					if u != nil || err == nil {
						wrap := CreateJWTWrapper(*u, expectedHours)
						_, err = wrap.ValidateToken(token)
						if err == nil {
							mu := dao.ToModelUser(u)
							return mu, err
						}
					}
				}
			}
		}
		log.Println("Cookie incorrecta: ", err)
	}
	errMsg := strings.Join([]string{"no se puede leer la cookie", expectedName}, " ")
	return nil, errors.New(errMsg)
}

// BearerAuth gets the model User for the token, if valid JWT
func BearerAuth(cookies string) (*models.User, error) {
	expectedName := strings.Join([]string{BearerCookieName, "="}, "")
	cookieSlice := strings.Split(cookies, ";")
	mu, err := GetJWTModelUserCookies(cookieSlice, expectedName, int64(AuthHours))
	if err == nil && mu != nil {
		return mu, err
	}
	return nil, err
}

func ReAuth(cookies string) (*models.User, error) {
	expectedName := strings.Join([]string{BearerCookieName, "="}, "")
	cookieSlice := strings.Split(cookies, ";")
	mu, err := GetJWTModelUserCookies(cookieSlice, expectedName, int64(AuthHours))
	if err == nil && mu != nil {
		return mu, err
	}

	expectedName = strings.Join([]string{ReauthCookieName, "="}, "")
	mu, err = GetJWTModelUserCookies(cookieSlice, expectedName, int64(ReauthHours))
	if err == nil && mu != nil {
		return mu, err
	}
	return nil, err
}

func Logout(auth.LogoutParams) middleware.Responder {
	bcookie := CreateDeprecatedCookie(BearerCookieName)
	rcookie := CreateDeprecatedCookie(ReauthCookieName)
	return auth.NewLogoutOK().WithAuth(bcookie).WithReAuth(rcookie)
}

func Relogin(params auth.ReloginParams, u *models.User) middleware.Responder {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var du *dao.User
		du, err = dao.GetUserUsername(db, *u.Username)
		if err == nil && du != nil {
			var btoken, rtoken string
			btoken, err = CreateJWT(*du, int64(AuthHours))
			if err == nil {
				rtoken, err = CreateJWT(*du, int64(ReauthHours))
				if err == nil {
					bcookie := CreateCookie(BearerCookieName, btoken, hoursToSeconds(AuthHours))
					recookie := CreateCookie(ReauthCookieName, rtoken, hoursToSeconds(ReauthHours))
					return auth.NewReloginOK().WithAuth(bcookie).WithReAuth(recookie)
				}
			}
		}
	}
	return auth.NewReloginInternalServerError()
}

func CloseSessions(params auth.CloseSessionsParams, u *models.User) middleware.Responder {
	if userOrAdmin(params.Username, u) {
		p := params.Password
		db, err := dbconnection.ConnectDb()
		if err == nil {
			ud, _ := dao.GetUserUsername(db, params.Username)
			oldHash := []byte(*ud.Pwhash)
			oldHashSt := string(oldHash)
			var newHashSt string
			if bcrypt.CompareHashAndPassword(oldHash, []byte(*p.Password)) == nil {
				for {
					bytes, errBcrypt := bcrypt.GenerateFromPassword([]byte(*p.Password), Cost)
					if errBcrypt != nil {
						log.Print("Error en encriptacion CloseSessions", errBcrypt)
						return auth.NewCloseSessionsInternalServerError()
					}
					newHashSt = string(bytes)
					if oldHashSt != newHashSt {
						break
					}
				}
				err = dao.PutPasswordUsername(db, params.Username, newHashSt)
				if err != nil {
					log.Println("Error al modificar la contraseña con el mismo valor en CloseSessions: ", err)
					return auth.NewCloseSessionsInternalServerError()
				}

				return auth.NewCloseSessionsOK()
			}
			return auth.NewCloseSessionsForbidden() //O forbidden?
		}
	}
	return auth.NewCloseSessionsInternalServerError()
}