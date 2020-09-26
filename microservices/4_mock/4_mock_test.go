package __mock

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type Kitten struct {
	ID int
}

type Store interface {
	Search(name string) []Kitten
}

type MockStore struct {
	mock.Mock
}

func (s *MockStore) Search(name string) []Kitten {
	args := s.Mock.Called(name)
	return args.Get(0).([]Kitten)
}

func Test_StoreImpl(t *testing.T) {
	store := MockStore{}
	kittens := []Kitten{
		{
			ID: 7,
		},
	}
	store.On("Search", "normal").Return(kittens)

	result := store.Search("normal")
	assert.Equal(t, result, kittens)
}
