package auth

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getPermissionsByUserIDHandler := kithttp.NewServer(
		makeGetPermissionsEndpoint(s),
		getPermissionsByUserIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/permissions/{userId}", getPermissionsByUserIDHandler)

	return r
}
func getPermissionsByUserIDRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getPermissionsByUserId{
		UserID: chi.URLParam(r, "userId"),
	}, nil
}
