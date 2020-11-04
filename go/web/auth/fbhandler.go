package myfirebase

type FbHandler interface {
	GetUIDByToken(string) string
}
