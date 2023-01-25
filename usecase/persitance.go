package usecase

type Persistance interface {
	StoreString(key string, value string) error
	FetchString(key string) (string, error)
	StoreList(key string, value []string) error
	FetchList(key string) ([]string, error)
}
