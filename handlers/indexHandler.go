package handlers

/*
 Copyright (C) 2018 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2018 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	services "UlboraCmsAdmin/services"
	//"fmt"
	"net/http"
)

//HandleAdminIndex HandleAdminIndex
func (h *Handler) HandleAdminIndex(w http.ResponseWriter, r *http.Request) {
	h.Sess.InitSessionStore(w, r)
	session := h.getSession(w, r)
	loggedIn := session.Values["userLoggenIn"]
	token := h.getToken(w, r)
	//fmt.Print("loggedIn in index: ")
	// fmt.Print("token in index: ")
	// fmt.Println(token)
	if loggedIn == nil || !loggedIn.(bool) || token == nil {
		h.loginImplicit(w, r)
	} else {
		clientID := session.Values["clientId"].(string)
		var c services.ContentService
		c.ClientID = clientID
		c.APIClient = getGatewayAPIClient()
		c.APIKey = getGatewayAPIKey()
		c.Host = getContentHost()
		res := c.GetContentList(clientID)
		// fmt.Print("res in index: ")
		// fmt.Println(res)
		h.Templates.ExecuteTemplate(w, "index.html", &res)
	}
}
