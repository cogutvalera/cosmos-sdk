package utils

import (
	"net/http"
	"net/url"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authctx "github.com/cosmos/cosmos-sdk/x/auth/client/context"
)

// WriteErrorResponse prepares and writes a HTTP error
// given a status code and an error message.
func WriteErrorResponse(w *http.ResponseWriter, status int, msg string) {
	(*w).WriteHeader(status)
	(*w).Write([]byte(msg))
}

// GenerateAndMarshallStdSignMsgJSON builds a StdSignMsg for the given
// messages and returns its JSON representation.
func GenerateAndMarshallStdSignMsgJSON(txCtx authctx.TxContext, msgs []sdk.Msg) (output []byte, err error) {
	stdMsg, err := txCtx.Build(msgs)
	if err != nil {
		return
	}
	return txCtx.Codec.MarshalJSON(stdMsg)
}

// HasGenerateOnlyArg returns whether a URL's query "generate-only" parameter is set to "true".
func HasGenerateOnlyArg(url *url.URL) bool {
	value := url.Query().Get(client.FlagGenerateOnly)
	return value == "true"
}
