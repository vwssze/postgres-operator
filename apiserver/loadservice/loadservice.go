package loadservice

/*
Copyright 2017-2019 Crunchy Data Solutions, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/crunchydata/postgres-operator/apiserver"
	msgs "github.com/crunchydata/postgres-operator/apiservermsgs"
	//"github.com/gorilla/mux"
	"net/http"
)

// LoadHandler ...
// pgo load --selector=name=mycluster --load-config=./sample-load-config.json
func LoadHandler(w http.ResponseWriter, r *http.Request) {

	log.Infoln("loadservice.LoadHandler called")

	var request msgs.LoadRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	err := apiserver.Authn(apiserver.LOAD_PERM, w, r)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var resp msgs.LoadResponse
	if request.ClientVersion != msgs.PGO_VERSION {
		resp = msgs.LoadResponse{}
		resp.Status = msgs.Status{Code: msgs.Error, Msg: apiserver.VERSION_MISMATCH_ERROR}
	} else {
		resp = Load(&request)
	}

	json.NewEncoder(w).Encode(resp)
}
