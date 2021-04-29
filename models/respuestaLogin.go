package models

// RespuestaLogin tiene el token que se devuelven con el Login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
