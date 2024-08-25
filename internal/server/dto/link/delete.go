package linkdto

type DeleteQuery struct {
	// Name of the link to delete.
	Name string `form:"name,min=1,max=256"`
}
