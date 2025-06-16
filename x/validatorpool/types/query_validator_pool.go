package types

// QueryValidatorPoolRequest is the request type for the Query/ValidatorPool RPC method
type QueryValidatorPoolRequest struct{}

// QueryValidatorPoolResponse is the response type for the Query/ValidatorPool RPC method
type QueryValidatorPoolResponse struct {
	ValidatorPool *ValidatorPool `protobuf:"bytes,1,opt,name=validator_pool,json=validatorPool,proto3" json:"validator_pool,omitempty"`
}
