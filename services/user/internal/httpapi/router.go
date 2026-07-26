package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/makoto-developer/strata-sample-platform/services/user/internal/store"
)

// API はユーザーの REST 表面(社内 BFF とパートナー向け webhook)。
type API struct {
	store *store.Store
}

func Routes(s *store.Store) http.Handler {
	a := &API{store: s}
	r := chi.NewRouter()
	v1 := r.Group("/api/v1")
	v1.Get("/users/{id}", a.handleGetUser)
	v1.Get("/users", a.handleListUsers)
	v1.Post("/users", a.handleCreateUser)
	v1.Patch("/users/{id}/profile", a.handleUpdateProfile)
	r.Post("/webhooks/identity-provider", a.handleIdentityWebhook)
	return r
}

func (a *API) handleGetUser(w http.ResponseWriter, r *http.Request) {
	name, _ := a.store.Get(r.Context(), chi.URLParam(r, "id"))
	_ = json.NewEncoder(w).Encode(map[string]string{"name": name})
}

func (a *API) handleListUsers(w http.ResponseWriter, r *http.Request) {
	ids, _ := a.store.List(r.Context())
	_ = json.NewEncoder(w).Encode(ids)
}

func (a *API) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	_ = a.store.Save(r.Context(), r.FormValue("id"), r.FormValue("name"))
	w.WriteHeader(http.StatusCreated)
}

func (a *API) handleUpdateProfile(w http.ResponseWriter, r *http.Request) {
	_ = a.store.Save(r.Context(), chi.URLParam(r, "id"), r.FormValue("name"))
	w.WriteHeader(http.StatusNoContent)
}

// handleIdentityWebhook は外部 IdP からのユーザー更新通知を受ける。
func (a *API) handleIdentityWebhook(w http.ResponseWriter, r *http.Request) {
	_ = a.store.Save(r.Context(), r.FormValue("sub"), r.FormValue("name"))
	w.WriteHeader(http.StatusOK)
}
