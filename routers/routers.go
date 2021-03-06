package routers

import (
	"github.com/atmiguel/cerealnotes/handlers"
	"github.com/atmiguel/cerealnotes/paths"
	"net/http"
)

func SetRoutes() {
	// static files
	{
		staticDirectoryName := "static"
		staticDirectoryPaddedWithSlashes := "/" + staticDirectoryName + "/"

		fileServer := http.FileServer(http.Dir(staticDirectoryName))

		http.Handle(
			staticDirectoryPaddedWithSlashes,
			http.StripPrefix(staticDirectoryPaddedWithSlashes, fileServer))
	}

	// templates
	http.HandleFunc(paths.LoginOrSignupPath, handlers.HandleLoginOrSignupRequest)

	// forms
	http.HandleFunc("/user", handlers.HandleUserRequest)
	http.HandleFunc("/session", handlers.HandleSessionRequest)

	// requires Authentication
	handleAuthenticated(paths.HomePath, handlers.HandleHomeRequest)
}

func handleAuthenticated(
	pattern string,
	handlerFunc handlers.AuthentictedRequestHandlerType,
) {
	http.HandleFunc(pattern, handlers.AuthenticateOrRedirectToLogin(handlerFunc))
}
