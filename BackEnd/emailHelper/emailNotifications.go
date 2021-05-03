// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas

// Package emailHelper provides functions to send proper emails to users
package emailHelper

import (
	"fmt"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
)

func generateEmailBodyUserInvitedTest(username string, testid int64, address string, frontEnd string, m *models.Message) ([]byte, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var test *dao.Test
		test, err = dao.GetTest(db, testid)
		mensInvitacion := ""
		if m != nil && m.Body != nil {
			mensInvitacion = *m.Body
		}
		if err == nil && test != nil {
			msg := []byte("To: " + address + "\r\n" +
				"Subject: [NO RESPONDER] Invitado al test: " + *test.Title + "\r\n" +
				"\r\n" +
				"Hola " + username + ", has sido invitado a realizar el test " + *test.Title + ".\r\n" +
				"Puedes acceder mediante la web, o pulsando el siguiente enlace: " + frontEnd + "/pt/" + fmt.Sprint(testid) + "\r\n" +
				"\r\n" +
				mensInvitacion)
			return msg, nil
		}
	}
	return nil, err
}

// SendEmailUserInvitedToTest sends an email to the user with username username,
// notifying them that they have been invited to a test
func SendEmailUserInvitedToTest(username string, testid int64, m *models.Message) {
	address, err := getEmailFromUsername(username)
	if err == nil {
		emailInfo, err := getOwnEmailInfo()
		if err == nil {
			emailBody, err := generateEmailBodyUserInvitedTest(username, testid, address, emailInfo.FrontEndUrl, m)
			if err == nil {
				sendEmail(emailBody, address)
			}
		}
	}
}

func generateEmailBodyUserTeamInvitedTest(username string, testid int64, address string, frontEnd string, teamname string, m *models.Message) ([]byte, error) {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var test *dao.Test
		test, err = dao.GetTest(db, testid)
		mensInvitacion := ""
		if m != nil && m.Body != nil {
			mensInvitacion = *m.Body
		}
		if err == nil && test != nil {
			msg := []byte("To: " + address + "\r\n" +
				"Subject: [NO RESPONDER] Invitado al test: " + *test.Title + "\r\n" +
				"\r\n" +
				"Hola " + username + ", has sido invitado a realizar el test " + *test.Title + ", debido a pertenecer al equipo " + teamname + ".\r\n" +
				"Puedes acceder mediante la web, o pulsando el siguiente enlace: " + frontEnd + "/pt/" + fmt.Sprint(testid) + "\r\n" +
				"\r\n" +
				mensInvitacion)
			return msg, nil
		}
	}
	return nil, err
}

// SendEmailUserInvitedToTest sends an email to the user with username username,
// notifying them that they have been invited to a test
func SendEmailUserTeamInvitedToTest(username string, testid int64, teamname string, m *models.Message) {
	address, err := getEmailFromUsername(username)
	if err == nil {
		emailInfo, err := getOwnEmailInfo()
		if err == nil {
			emailBody, err := generateEmailBodyUserTeamInvitedTest(username, testid, address, emailInfo.FrontEndUrl, teamname, m)
			if err == nil {
				sendEmail(emailBody, address)
			}
		}
	}
}
