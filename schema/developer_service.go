package schema

const (
	PathMsDeveloperService = "developer-service"
)

const (
	CommandCreateRelease = "create-release"
	CommandDeleteRelease = "delete-release"
	CommandUpdateRelease = "update-release"
	CommandGetReleases   = "get-releases"
)

const (
	CommandCreateRelease10 = CommandCreateRelease + "-" + Version10
	CommandDeleteRelease10 = CommandDeleteRelease + "-" + Version10
	CommandUpdateRelease10 = CommandUpdateRelease + "-" + Version10
	CommandGetReleases10   = CommandGetReleases + "-" + Version10
)

const (
	PathSubscriptionDeveloperServiceAuthenticateUserCreated10 = PathMsDeveloperService + "." + PathEventAuthenticateUserCreated10
)
