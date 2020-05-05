package balance

type Balance interface {
	Add(...string) error
	Get(...string) (string, error)
}
