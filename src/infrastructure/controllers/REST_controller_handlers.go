package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"main/core/model"
	gen_api "main/go-server-generated/go"
)

func (me *ControllerREST) handleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		token := r.Header.Get("user_token")
		if token == "" {
			token = r.URL.Query().Get("user_token")
		}

		role := me.app.CheckRoleByToken(token)
		if role == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		r.Header.Set(RoleHeaderKey, role)
		next.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.Host,
			time.Since(start),
		)
	})
}

func (me *ControllerREST) handleBannerGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (me *ControllerREST) handleBannerIdDelete(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(RoleHeaderKey) != model.AdminRole {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

func (me *ControllerREST) handleBannerIdPatch(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(RoleHeaderKey) != model.AdminRole {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

func (me *ControllerREST) handleBannerPost(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(RoleHeaderKey) != model.AdminRole {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var banner gen_api.BannerBody
	err := json.NewDecoder(r.Body).Decode(&banner)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(gen_api.InlineResponse400{Error_: InvalidBodyErr})
		w.Write(res)
		return
	}

	id, err := me.app.AddBanner(banner)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(gen_api.InlineResponse400{Error_: InternalServerErr})
		w.Write(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(gen_api.InlineResponse201{BannerId: int32(id)})
	w.Write(res)
}

func (me *ControllerREST) handleUserBannerGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (me *ControllerREST) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "{\"what\": \"Dattebayo!!!\"}")
}
