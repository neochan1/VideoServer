package apis

import (
	"VideoServer/apis/database"
	"VideoServer/apis/models"
	"VideoServer/apis/session"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &models.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, models.ErrorRequestBodyParseFailed)
		return
	}

	if err := database.AddUserCredential(ubody.Username, ubody.Password); err != nil {
		sendErrorResponse(w, models.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &models.Signedup{Success: true, SessionID: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, models.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}