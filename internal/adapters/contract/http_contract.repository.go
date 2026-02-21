package adapters

import usecases "rungdee-apm-api/internal/usecases/contract"

type HttpContractHandler struct {
	contractUseCase usecases.ContractUseCase
}

func NewHttpContractHandler(usecase usecases.ContractUseCase) *HttpContractHandler {
	return &HttpContractHandler{contractUseCase: usecase}
}
