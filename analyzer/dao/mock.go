package dao

import (
  "io"

  "github.com/stretchr/testify/mock"
)

type MockDao struct {
  mock.Mock
}

func (m *MockDao) GetXMLStream(link string) (io.ReadCloser, error) {
  args := m.Called(link)
  readCloser, _ := args.Get(0).(io.ReadCloser)
  return readCloser, args.Error(1)
}