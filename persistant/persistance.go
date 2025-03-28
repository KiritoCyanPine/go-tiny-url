package persistant

type Persistant interface {
	Get(key string) (any, error)
	Set(key string, value any) error
}
