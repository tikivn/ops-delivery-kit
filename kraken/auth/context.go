package auth

import "context"

var (
	uidKey = uidKeyType{}
)

type UID struct {
	ID       string `json:"id"`
	Trn      string `json:"trn"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (u UID) IsZero() bool {
	return u.ID == "" && u.Username == "" && u.Trn == "" && u.Email == ""
}

type uidKeyType struct{}

func ContextWithUID(ctx context.Context, uid UID) context.Context {
	return context.WithValue(ctx, uidKey, uid)
}

func UIDFromContext(ctx context.Context) UID {
	val := ctx.Value(uidKey)
	if uid, ok := val.(UID); ok {
		return uid
	}

	return UID{}
}
