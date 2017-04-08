package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Session interface {
	Get(string) interface{}
	Set(string, interface{})
	Delete(string)
	Save()
	Clear()
}

type session struct {
	r       *http.Request
	s       *sessions.Session
	written bool
}

func NewSession(r *http.Request, s *sessions.Session) Session {
	return &session{
		r: r,
		s: s,
	}
}

func (s *session) Get(key string) interface{} {
	return nil
}

func (s *session) Set(key string, value interface{}) {
}

func (s *session) Delete(key string) {
	s.written = true
}

func (s *session) Save() {
	s.written = true
}

func (s *session) Clear() {
}
