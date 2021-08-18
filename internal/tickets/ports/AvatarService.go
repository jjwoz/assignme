package ports

// Avatar is the concrete struct to include the file and name
type Avatar struct {
	file byte
	name string
}

//AvatarService  is an interface for retrieving avatar image
type AvatarService interface {
	RetrieveAvatar(path string) (byte, error)
}
