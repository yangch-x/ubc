package ctxdata

import (
	"context"
)

type ctxKey string

// CtxKeyJwtUserId get uid from ctx
const (
	CtxKeyJwtUserEmail ctxKey = "jwtUserEmail"
	CtxKeyUserId       ctxKey = "jwtUserId"
	CtxKeyUserRole     ctxKey = "jwtUserRole"
)

func GetAccountFromCtx(ctx context.Context) string {
	if uid, ok := ctx.Value(string(CtxKeyUserId)).(string); ok {
		return uid
	}
	return ""
}

func GetUserRoleFromCtx(ctx context.Context) string {
	if uid, ok := ctx.Value(string(CtxKeyUserRole)).(string); ok {
		return uid
	}
	return ""
}

func GetEmailFromCtx(ctx context.Context) string {
	if email, ok := ctx.Value(string(CtxKeyJwtUserEmail)).(string); ok {
		return email
	}
	return ""
}
