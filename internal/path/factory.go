package path

import "github.com/hashicorp/vault/sdk/framework"

func GetPaths() []*framework.Path {
	return []*framework.Path{
		getPath(&appPathConfig{}),
		getPath(&signPathConfig{}),
	}
}

func getPath(c config) *framework.Path {
	return &framework.Path{
		Pattern:        c.getPattern(),
		HelpSynopsis:   c.getHelpSynopsis(),
		Fields:         c.getFields(),
		ExistenceCheck: c.getExistenceFunc(),
		//Callbacks:      c.getCallbacks(),
		Operations: c.getOperations(),
	}
}
