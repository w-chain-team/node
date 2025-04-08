package rewards

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/umbracle/ethgo"

	"github.com/w-chain-team/node/command"
	"github.com/w-chain-team/node/command/helper"
	"github.com/w-chain-team/node/command/polybftsecrets"
	rootHelper "github.com/w-chain-team/node/command/rootchain/helper"
	sidechainHelper "github.com/w-chain-team/node/command/sidechain"
	"github.com/w-chain-team/node/consensus/polybft/contractsapi"
	"github.com/w-chain-team/node/contracts"
	"github.com/w-chain-team/node/helper/common"
	"github.com/w-chain-team/node/txrelayer"
	"github.com/w-chain-team/node/types"
)

var params withdrawRewardsParams

func GetCommand() *cobra.Command {
	unstakeCmd := &cobra.Command{
		Use:     "withdraw-rewards",
		Short:   "Withdraws pending rewards on child chain for given validator",
		PreRunE: runPreRun,
		RunE:    runCommand,
	}

	helper.RegisterJSONRPCFlag(unstakeCmd)
	setFlags(unstakeCmd)

	return unstakeCmd
}

func setFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&params.accountDir,
		polybftsecrets.AccountDirFlag,
		"",
		polybftsecrets.AccountDirFlagDesc,
	)

	cmd.Flags().StringVar(
		&params.accountConfig,
		polybftsecrets.AccountConfigFlag,
		"",
		polybftsecrets.AccountConfigFlagDesc,
	)

	cmd.MarkFlagsMutuallyExclusive(polybftsecrets.AccountDirFlag, polybftsecrets.AccountConfigFlag)
}

func runPreRun(cmd *cobra.Command, _ []string) error {
	params.jsonRPC = helper.GetJSONRPCAddress(cmd)

	return params.validateFlags()
}

func runCommand(cmd *cobra.Command, _ []string) error {
	outputter := command.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	validatorAccount, err := sidechainHelper.GetAccount(params.accountDir, params.accountConfig)
	if err != nil {
		return err
	}

	validatorAddr := validatorAccount.Ecdsa.Address()
	rewardPoolAddr := ethgo.Address(contracts.RewardPoolContract)

	txRelayer, err := txrelayer.NewTxRelayer(txrelayer.WithIPAddress(params.jsonRPC),
		txrelayer.WithReceiptTimeout(150*time.Millisecond))
	if err != nil {
		return err
	}

	encoded, err := contractsapi.RewardPool.Abi.Methods["pendingRewards"].Encode([]interface{}{validatorAddr})
	if err != nil {
		return err
	}

	response, err := txRelayer.Call(validatorAddr, rewardPoolAddr, encoded)
	if err != nil {
		return err
	}

	amount, err := common.ParseUint256orHex(&response)
	if err != nil {
		return err
	}

	encoded, err = contractsapi.RewardPool.Abi.Methods["withdrawReward"].Encode([]interface{}{})
	if err != nil {
		return err
	}

	txn := rootHelper.CreateTransaction(validatorAddr, &rewardPoolAddr, encoded, nil, false)

	receipt, err := txRelayer.SendTransaction(txn, validatorAccount.Ecdsa)
	if err != nil {
		return err
	}

	if receipt.Status != uint64(types.ReceiptSuccess) {
		return fmt.Errorf("withdraw transaction failed on block: %d", receipt.BlockNumber)
	}

	result := &withdrawRewardResult{
		ValidatorAddress: validatorAccount.Ecdsa.Address().String(),
		RewardAmount:     amount.Uint64(),
	}

	outputter.WriteCommandResult(result)

	return nil
}
