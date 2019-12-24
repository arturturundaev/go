package intfc
import (
    model "../model"
)

type Intfc interface {
    ShowAll() map[string]model.Entity
}