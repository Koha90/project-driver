package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/koha90/project-driver/internal/users"
	"github.com/koha90/project-driver/storage"
)

type APIServer struct {
	addr  string
	store storage.PostgresStorage
}

func NewAPIServer(addr string, store storage.PostgresStorage) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

func (s *APIServer) Run() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.HandleFunc("/", nil)

	apiR := chi.NewRouter()
	apiR.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	apiR.Delete("/user/{id}", makeHTTPHandleFunc(s.handleDeleteUser))

	r.Mount("/api", apiR)

	http.ListenAndServe(s.addr, r)
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetUser(w, r)
	case http.MethodPost:
		return s.handleCreateUser(w, r)
	case http.MethodPatch:
		return nil
	}

	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	user, err := s.store.GetUsers()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	req := new(users.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	user, err := users.NewUser(req.Username, req.Password)
	if err != nil {
		return err
	}

	if err := s.store.CreateUser(user); err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, user)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteUser(id); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func getID(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id: %s does not exists\n", idStr)
	}

	return id, nil
}
