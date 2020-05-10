package pb_hello

import (
	"encoding/json"
	"net/http"
	pbgo "rpc/Protobuf/pbgo/plugins"
	. "rpc/Protobuf/pbgo/util"

	"github.com/julienschmidt/httprouter"
)

type HelloServiceInterface interface {
	Hello(in *String, out *String) error
}

func HelloServiceHandler(svc HelloServiceInterface) http.Handler {
	var router = httprouter.New()
	_handle_HelloService_get(router, svc)
	return router
}

func _handle_HelloService_get(router *httprouter.Router, svc HelloServiceInterface) {
	router.Handle("GET", "/hello/:value", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var protoReq, protoReply String

		err := pbgo.PopulaterFieldFromPath(&protoReq, fieldPath, ps.ByName("value"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := svc.Hello(&protoReq, &protoReply); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&protoReply); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
