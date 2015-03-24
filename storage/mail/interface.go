package mail

type Storage interface {
	Store(*Mail) error
}
