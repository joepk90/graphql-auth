package auth

const (

	// actions
	createAction = "create"
	readAction   = "read"
	updateAction = "update"
	deletection  = "delete"

	// resources
	postResource = "post"
)

type ActionOnResource struct {
	ID       string
	Action   string
	Resource string
}

func CanCreatePost() ActionOnResource {
	return ActionOnResource{
		ID:       "*",
		Resource: postResource,
		Action:   createAction,
	}
}

func CanReadPost(id string) ActionOnResource {
	return ActionOnResource{
		ID:       id,
		Resource: postResource,
		Action:   readAction,
	}
}

func CanUpdatePost(id string) ActionOnResource {
	return ActionOnResource{
		ID:       id,
		Resource: postResource,
		Action:   updateAction,
	}
}

func CanDeletePost(id string) ActionOnResource {
	return ActionOnResource{
		ID:       id,
		Resource: postResource,
		Action:   deletection,
	}
}
