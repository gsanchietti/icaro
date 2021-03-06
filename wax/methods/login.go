/*
 * Copyright (C) 2017 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Icaro project.
 *
 * Icaro is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * Icaro is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Icaro.  If not, see COPYING.
 *
 * author: Giacomo Sanchietti <giacomo.sanchietti@nethesis.it>
 */

package methods

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/nethesis/icaro/wax/utils"
)

func AuthAccept(c *gin.Context, prefs string) {
	c.String(http.StatusOK, "Auth: 1\n"+prefs)
}

func AuthReject(c *gin.Context, description string) {
	message := "Auth: 0"

	if len(description) > 0 {
		message = message + " \nReply-Message: " + description
	}

	c.String(http.StatusForbidden, message)
	c.Abort()
}

func autoLogin(c *gin.Context, unitMacAddress string, username string, userMac string, sessionId string) {
	isValid, user := utils.GetUserByMacAddressAndunitMacAddress(userMac, unitMacAddress)
	if !isValid {
		AuthReject(c, "user account not found")
		return
	}

	if !user.AutoLogin {
		AuthReject(c, "auto login not permitted")
		return
	}

	// check if user account is not expired
	if user.ValidUntil.Before(time.Now().UTC()) {
		AuthReject(c, "user account is expired")
		return
	}

	// extract preferences
	unit := utils.GetUnitByMacAddress(unitMacAddress)
	prefs := utils.GetHotspotPreferencesByKeys(
		unit.HotspotId,
		[]string{"Idle-Timeout", "Acct-Session-Time", "Session-Timeout"},
	)

	var outPrefs bytes.Buffer
	for _, pref := range prefs {
		outPrefs.WriteString(fmt.Sprintf("%s:%s\n", pref.Key, pref.Value))
	}

	outPrefs.WriteString(fmt.Sprintf("%s:%d\n", "CoovaChilli-Bandwidth-Max-Up", user.KbpsUp))
	outPrefs.WriteString(fmt.Sprintf("%s:%d\n", "CoovaChilli-Bandwidth-Max-Down", user.KbpsDown))

	// response to dedalo
	AuthAccept(c, outPrefs.String())
}

func Login(c *gin.Context, unitMacAddress string, username string, chapPass string, chapChal string, sessionId string) {
	// check if unit exists
	unit := utils.GetUnitByMacAddress(unitMacAddress)
	if unit.Id <= 0 {
		AuthReject(c, "unit not found")
		return
	}

	// check if user exists
	user := utils.GetUserByUsername(username)
	if user.Id <= 0 {
		AuthReject(c, "user not found")
		return
	}

	// check if user-sessions exists
	if user.AccountType != "email" && user.AccountType != "sms" {
		valid := utils.CheckUserSession(user.Id, sessionId)
		if !valid {
			AuthReject(c, "user-session not found")
			return
		}
	}

	// check if user credentials are valid
	if chapPass != utils.CalcUserDigest(user, chapChal) {
		AuthReject(c, "password mismatch")
		return
	}

	// check if user account is not expired
	if user.ValidUntil.Before(time.Now().UTC()) {
		AuthReject(c, "user account is expired")
		return
	}

	// check if hotspot belongs to a reseller with a valid plan
	if !utils.HotspotHasValidSubscription(user.HotspotId) {
		AuthReject(c, "reseller account is expired")
		return
	}

	// extract preferences
	prefs := utils.GetHotspotPreferencesByKeys(
		unit.HotspotId,
		[]string{"Idle-Timeout", "Acct-Session-Time", "Session-Timeout"},
	)
	var outPrefs bytes.Buffer
	for _, pref := range prefs {
		outPrefs.WriteString(fmt.Sprintf("%s:%s\n", pref.Key, pref.Value))
	}

	outPrefs.WriteString(fmt.Sprintf("%s:%d\n", "CoovaChilli-Bandwidth-Max-Up", user.KbpsUp))
	outPrefs.WriteString(fmt.Sprintf("%s:%d\n", "CoovaChilli-Bandwidth-Max-Down", user.KbpsDown))

	// response to dedalo
	AuthAccept(c, outPrefs.String())

}

func Logins(c *gin.Context) {

	service := c.Query("service")
	switch service {
	case "framed":
		unitMacAddress := c.Query("ap")
		user := c.Query("user")
		userMac := c.Query("mac")
		sessionId := c.Query("sessionid")
		autoLogin(c, unitMacAddress, user, userMac, sessionId)

	case "login":
		unitMacAddress := c.Query("ap")
		user := c.Query("user")
		chapPass := c.Query("chap_pass")
		chapChal := c.Query("chap_chal")
		sessionId := c.Query("sessionid")
		Login(c, unitMacAddress, user, chapPass, chapChal, sessionId)

	default:
		c.String(http.StatusNotFound, "Invalid login service: '%s'", service)
	}
}
