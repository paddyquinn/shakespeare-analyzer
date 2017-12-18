package dao

import "io"

type Interface interface {
  GetXMLStream(string) (io.ReadCloser, error)
}
