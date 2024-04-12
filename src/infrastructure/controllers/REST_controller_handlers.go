package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"main/core/model"
	gen_api "main/go-server-generated/go"
)

func (me *ControllerREST) handleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		token := r.Header.Get("user_token")

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
	if !isAdminRole(r.Header.Get(RoleHeaderKey)) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var banner_id int32
	_, err := fmt.Sscanf(path.Base(r.RequestURI), "%d", &banner_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(InvalidURIDataErr))
		return
	}

	if !me.app.IsBannerExists(banner_id) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(NotExistsResourceErr))
		return
	}

	err = me.app.DeleteBanner(banner_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(InternalServerErr))
		return
	}
}

func (me *ControllerREST) handleBannerIdPatch(w http.ResponseWriter, r *http.Request) {
	if !isAdminRole(r.Header.Get(RoleHeaderKey)) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var banner_id int32
	_, err := fmt.Sscanf(path.Base(r.RequestURI), "%d", &banner_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(InvalidURIDataErr))
		return
	}

	var banner_new gen_api.BannerBody
	err = json.NewDecoder(r.Body).Decode(&banner_new)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(InvalidBodyDataErr))
		return
	}

	if !me.app.IsBannerExists(banner_id) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(NotExistsResourceErr))
		return
	}

	err = me.app.UpdatedBanner(banner_id, banner_new)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(InternalServerErr))
		return
	}
}

func (me *ControllerREST) handleBannerPost(w http.ResponseWriter, r *http.Request) {
	if !isAdminRole(r.Header.Get(RoleHeaderKey)) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var banner gen_api.BannerBody
	err := json.NewDecoder(r.Body).Decode(&banner)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(InvalidBodyDataErr))
		return
	}

	id, err := me.app.AddBanner(banner)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(InternalServerErr))
		return
	}

	res, _ := json.Marshal(gen_api.InlineResponse201{BannerId: int32(id)})
	w.Write(res)
}

func (me *ControllerREST) handleUserBannerGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (me *ControllerREST) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "{\"Nani\": \"Dattebayo!!!\"}")
}

func isBannerAllowedToRole(banner *gen_api.BannerIdBody, role string) bool {
	return !banner.IsActive && role == model.UserRole
}

func isAdminRole(role string) bool {
	return role == model.AdminRole
}
