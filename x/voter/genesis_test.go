package voter_test

import (
	"testing"

	keepertest "github.com/iroan/voter/testutil/keeper"
	"github.com/iroan/voter/testutil/nullify"
	"github.com/iroan/voter/x/voter"
	"github.com/iroan/voter/x/voter/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PollList: []types.Poll{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PollCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VoterKeeper(t)
	voter.InitGenesis(ctx, *k, genesisState)
	got := voter.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PollList, got.PollList)
	require.Equal(t, genesisState.PollCount, got.PollCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
