package schema

var Permissions = map[string][]string{
	PathMsAuthenticate: {
		CommandIdentifyUser,
		CommandGetDevelopers,
	},
	PathMsGateway: {
		//authenticate
		CommandIdentifyUser,

		//developer-service
		CommandCreateRelease,
		CommandDeleteRelease,
		CommandUpdateRelease,
		CommandGetReleases,
	},
	PathMsDeveloperService: {
		CommandCreateRelease,
		CommandDeleteRelease,
		CommandUpdateRelease,
		CommandGetReleases,

		//authenticate
		CommandGetDevelopers,
	},
}
