package clique

type ValidatorListProvider interface {
	GetValidatorList() []string
}

type CycleHandler struct {
	listProvider      ValidatorListProvider
	validatorList     []string
	cycleDuration     int
	CurrentCycleBlock int // mod cycleDuration
}

func (h *CycleHandler) StartCycle() {
	h.validatorList = h.listProvider.GetValidatorList()
	h.cycleDuration = len(h.validatorList)
}

func (h *CycleHandler) FinishCycle() {

}

func (h *CycleHandler) GetCurrentValidator() {

}

func (h *CycleHandler) setCurrentValidator() {

}
