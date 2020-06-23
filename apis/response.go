package apis

import (
	"encoding/json"
	"io"
	"net/http"
	"VideoServer/apis/models"
)

func sendErrorResponse(w http.ResponseWriter, errResp models.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}


