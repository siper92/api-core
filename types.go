package api_base

type ToModelEntity[T interface{}] interface {
	ToModel() T
}

type RepoEntity interface {
	GetID() int64
	TableName() string
	IsFilterable(field string) bool
}

type EntityRepository[T RepoEntity] interface {
	New() T
	NewSlice() []T
	GetByID(id int64) (T, error)
	GetByField(field string, value interface{}) ([]T, error)
	Create(T) (T, error)
	Update(T) (bool, error)
	Delete(T) (bool, error)
}

type EmptyRepository[T RepoEntity] struct {
	value T
}

func NewEmptyRepository[T RepoEntity](typeVal T) *EmptyRepository[T] {
	return &EmptyRepository[T]{
		value: typeVal,
	}
}

var _ EntityRepository[RepoEntity] = (*EmptyRepository[RepoEntity])(nil)

func (u *EmptyRepository[T]) New() (result T) {
	return result
}

func (u *EmptyRepository[T]) NewSlice() (result []T) {
	return result
}

func (u *EmptyRepository[T]) GetByID(id int64) (T, error) {
	return u.New(), nil
}

func (u *EmptyRepository[T]) Create(t T) (T, error) {
	return t, nil
}

func (u *EmptyRepository[T]) Update(t T) (bool, error) {
	return true, nil
}

func (u *EmptyRepository[T]) Delete(t T) (bool, error) {
	return true, nil
}

func (u *EmptyRepository[T]) GetByField(field string, value interface{}) (result []T, err error) {
	return result, nil
}