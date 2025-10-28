package model

import "fmt"

type Model interface {
	Name() string
	Version() string
	Load() Model
	Save(model Model)
}