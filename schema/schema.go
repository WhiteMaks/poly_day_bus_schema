package schema

const (
	PathMsAuthenticate     = "authenticate"
	PathMsDeveloperService = "developer-service"
)

const (
	PathEventAuthenticateUserCreated10 = PathMsAuthenticate + ".event.user-created-1.0"
)

const (
	PathSubscriptionDeveloperServiceAuthenticateUserCreated10 = PathMsDeveloperService + "." + PathEventAuthenticateUserCreated10
)

var Microservices = []string{
	PathMsAuthenticate,
	PathMsDeveloperService,
}

var Events = []string{
	PathEventAuthenticateUserCreated10,
}

var Subscribers = map[string][]string{
	PathEventAuthenticateUserCreated10: {
		PathSubscriptionDeveloperServiceAuthenticateUserCreated10,
	},
}
