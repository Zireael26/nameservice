package nameservice_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zireael26/nameservice/testutil/keeper"
	"github.com/zireael26/nameservice/testutil/nullify"
	"github.com/zireael26/nameservice/x/nameservice"
	"github.com/zireael26/nameservice/x/nameservice/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NameserviceKeeper(t)
	nameservice.InitGenesis(ctx, *k, genesisState)
	got := nameservice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
