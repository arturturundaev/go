package controllerIntfc

type ControllerIntfc interface {
	ShowAll() map[string]interface{}
  Delete(id string)
}
