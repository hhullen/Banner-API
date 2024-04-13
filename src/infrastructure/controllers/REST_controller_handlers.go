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
			"[API] %s %s %s %s",
			r.Method,
			r.RequestURI,
			r.Host,
			time.Since(start),
		)
	})
}

func (me *ControllerREST) handleBannerGet(w http.ResponseWriter, r *http.Request) {
	featureId, err := readQueryInteger(r, "feature_id")
	if err != nil {
		featureId = -1
	}

	tagId, err := readQueryInteger(r, "tag_id")
	if err != nil {
		tagId = -1
	}

	if featureId == -1 && tagId == -1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(AbsenceOfQueryParam))
		return
	}

	limit, err := readQueryInteger(r, "limit")
	if err != nil {
		limit = -1
	}

	offset, err := readQueryInteger(r, "offset")
	if err != nil {
		offset = 0
	}

	includeDeactivated := isAdminRole(r.Header.Get(RoleHeaderKey))
	banners, err := me.app.GetAllBannersByFilters(featureId, tagId, limit, offset, includeDeactivated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(InternalServerErr))
		return
	}
	if len(banners) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResp, err := json.Marshal(banners)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(InternalServerErr))
		return
	}
	w.Write(jsonResp)
}

func readQueryInteger(r *http.Request, name string) (num int32, err error) {
	_, err = fmt.Sscanf(r.URL.Query().Get(name), "%d", &num)
	return
}

func (me *ControllerREST) handleBannerIdDelete(w http.ResponseWriter, r *http.Request) {
	if !isAdminRole(r.Header.Get(RoleHeaderKey)) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var bannerId int32
	_, err := fmt.Sscanf(path.Base(r.RequestURI), "%d", &bannerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(InvalidURIDataErr))
		return
	}

	if !me.app.IsBannerExists(bannerId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(NotExistsResourceErr))
		return
	}

	err = me.app.DeleteBanner(bannerId)
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

	var bannerId int32
	_, err := fmt.Sscanf(path.Base(r.RequestURI), "%d", &bannerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(InvalidURIDataErr))
		return
	}

	var bannerNew gen_api.BannerBody
	err = json.NewDecoder(r.Body).Decode(&bannerNew)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(InvalidBodyDataErr))
		return
	}

	if !me.app.IsBannerExists(bannerId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(NotExistsResourceErr))
		return
	}

	err = me.app.UpdatedBanner(bannerId, bannerNew)
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
	featureId, err := readQueryInteger(r, "feature_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(AbsenceOfQueryParam))
		return
	}

	tagId, err := readQueryInteger(r, "tag_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(AbsenceOfQueryParam))
		return
	}
	includeDeactivated := isAdminRole(r.Header.Get(RoleHeaderKey))
	banner, err := me.app.GetSpecificBanner(featureId, tagId, includeDeactivated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(InternalServerErr))
		return
	}
	if banner == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResp, err := json.Marshal(banner)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(InternalServerErr))
		return
	}
	w.Write(jsonResp)
}

func (me *ControllerREST) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "{\"Nani\": \"Dattebayo!!!\"}")
}

func isAdminRole(role string) bool {
	return role == model.AdminRole
}
