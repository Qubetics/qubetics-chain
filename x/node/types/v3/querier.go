package v3

import (
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"

	base "github.com/qubetics/qubetics-blockchain/v2/types"
	v1base "github.com/qubetics/qubetics-blockchain/v2/types/v1"
)

func NewQueryNodeRequest(addr base.NodeAddress) *QueryNodeRequest {
	return &QueryNodeRequest{
		Address: addr.String(),
	}
}

func NewQueryNodesRequest(status v1base.Status, pagination *sdkquery.PageRequest) *QueryNodesRequest {
	return &QueryNodesRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryNodesForPlanRequest(id uint64, status v1base.Status, pagination *sdkquery.PageRequest) *QueryNodesForPlanRequest {
	return &QueryNodesForPlanRequest{
		Id:         id,
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
