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
		CommandCreateEntity,
		CommandDeleteEntity,
		CommandUpdateEntity,
		CommandGetEntities,
	},
	PathMsDeveloperService: {
		CommandCreateRelease,
		CommandDeleteRelease,
		CommandUpdateRelease,
		CommandGetReleases,
		CommandCreateEntity,
		CommandDeleteEntity,
		CommandUpdateEntity,
		CommandGetEntities,

		//authenticate
		CommandGetDevelopers,
	},
}
