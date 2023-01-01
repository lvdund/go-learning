// This is product
package builder

type House struct {
	windowType string
	floor      int
}

func (h *House) GetWindowType() string {
	return h.windowType
}

func (h *House) GetNumFoor() int {
	return h.floor
}
