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
	return s.s.Values[key]
}

func (s *session) Set(key string, value interface{}) {
	s.s.Values[key] = value
}

func (s *session) Delete(key string) {
	delete(s.s.Values, key)
	s.written = true
}

func (s *session) Save() {
	s.written = true
}

func (s *session) Clear() {
	for key := range s.s.Values {
		delete(s.s.Values, key)
	}
	s.written = true
}
