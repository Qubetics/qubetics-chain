package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/qubetics/qubetics-blockchain/v2/x/deposit/keeper"
	v1 "github.com/qubetics/qubetics-blockchain/v2/x/deposit/types/v1"
)

var (
	_ v1.QueryServiceServer = (*queryServer)(nil)
)

// maxDepositsPageLimit defines an upper bound for pagination limit to
// prevent excessive memory/CPU/network usage on large queries.
const maxDepositsPageLimit uint64 = 1000

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v1.QueryServiceServer {
	return &queryServer{k}
}

func (q *queryServer) QueryDeposit(c context.Context, req *v1.QueryDepositRequest) (*v1.QueryDepositResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryDeposit(ctx, req)
}

func (q *queryServer) QueryDeposits(c context.Context, req *v1.QueryDepositsRequest) (*v1.QueryDepositsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Enforce a maximum page size to avoid resource exhaustion.
	if req.Pagination == nil {
		req.Pagination = &sdkquery.PageRequest{}
	}
	if req.Pagination.Limit == 0 || req.Pagination.Limit > maxDepositsPageLimit {
		req.Pagination.Limit = maxDepositsPageLimit
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryDeposits(ctx, req)
}
