package engine

import "context"

type userKey struct{}

func WithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey{}, user)
}

func UserFromContext(ctx context.Context) string {
	return ctx.Value(userKey{}).(string)
}

func HandleUSER(ctx context.Context, args []string) Reply {
	if len(args) == 0 {
		return Reply{Code: FTPCodeSyntaxErrorCommandNotRecognized, Messages: []string{"Syntax error, command unrecognized"}}
	}
	user := args[0]
	ctx = WithUser(ctx, user)
	return Reply{Code: FTPCodeCommandCompleted, Messages: []string{"User name accepted, need password"}}
}

func HandlePASS(ctx context.Context, args []string) Reply {
	if len(args) == 0 {
		return Reply{Code: FTPCodeSyntaxErrorCommandNotRecognized, Messages: []string{"Syntax error, command unrecognized"}}
	}

	return Reply{Code: FTPCodeCommandCompleted, Messages: []string{"User name accepted, need password"}}
}
