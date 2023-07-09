package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	membershiptypes "github.com/noria-net/module-membership/x/membership/types"
)

const UpgradeName = "add_membership_module"

func (app WasmApp) RegisterUpgradeHandlers() {

	app.UpgradeKeeper.SetUpgradeHandler(
		UpgradeName,
		func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

			// set genesis state for the new module
			// newParams := membershiptypes.NewParams(...)

			// newGenesis := membershiptypes.GenesisState{
			// 	...
			// }
			// encoded, err := app.appCodec.MarshalJSON(&newGenesis)
			// if err != nil {
			// 	return nil, err
			// }

			// fromVM[membershiptypes.ModuleName] = adminmodule.AppModule{}.ConsensusVersion()
			// module := app.ModuleManager.Modules[membershiptypes.ModuleName].(adminmodule.AppModule)
			// module.InitGenesis(ctx, app.appCodec, encoded)

			return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
		},
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				membershiptypes.ModuleName,
			},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
