package ports

type AvatarService interface {
	retrieveAvatar(path string) (byte, error)
}
