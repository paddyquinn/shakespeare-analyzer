package dao

import (
  "io"
  "net/http"
)

type Dao struct {}

func (d *Dao) GetXMLStream(link string) (io.ReadCloser, error) {
  response, err := http.Get(link)
  if err != nil {
    return nil, err
  }

  return response.Body, nil
}