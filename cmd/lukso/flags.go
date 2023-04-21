package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	// geth related flag names
	gethTagFlag        = "geth-tag"
	gethCommitHashFlag = "geth-commit-hash"
	gethDatadirFlag    = "geth-datadir"
	gethConfigFileFlag = "geth-config"
	genesisJsonFlag    = "genesis-json"

	// erigon related flag names
	erigonTagFlag        = "erigon-tag"
	erigonConfigFileFlag = "erigon-config"
	erigonDatadirFlag    = "erigon-datadir"

	// Prysm related flag names
	prysmTagFlag             = "prysm-tag"
	prysmGenesisStateFlag    = "genesis-ssz"
	prysmChainConfigFileFlag = "prysm-chain-config"
	prysmConfigFileFlag      = "prysm-config"
	prysmDatadirFlag         = "prysm-datadir"
	noSlasherFlag            = "no-slasher"

	// lighthouse related flag names
	lighthouseTagFlag = "lighthouse-tag"

	// Validator related flag names
	validatorTagFlag                = "validator-tag"
	validatorDatadirFlag            = "validator-datadir"
	validatorWalletPasswordFileFlag = "validator-wallet-password"
	validatorWalletDirFlag          = "validator-wallet-dir"
	validatorConfigFileFlag         = "validator-config"
	validatorChainConfigFileFlag    = "validator-chain-config"

	// shared flags
	transactionFeeRecipientFlag = "transaction-fee-recipient"
	logFolderFlag               = "log-folder"

	// non-specific flags
	mainnetFlag   = "mainnet"
	testnetFlag   = "testnet"
	devnetFlag    = "devnet"
	validatorFlag = "validator"
	consensusFlag = "consensus"
	executionFlag = "execution"

	agreeTermsFlag = "agree-terms"

	// flag defaults used in different contexts
	executionMainnetDatadir = "./mainnet-data/execution"
	executionTestnetDatadir = "./testnet-data/execution"
	executionDevnetDatadir  = "./devnet-data/execution"

	consensusMainnetDatadir = "./mainnet-data/consensus"
	consensusTestnetDatadir = "./testnet-data/consensus"
	consensusDevnetDatadir  = "./devnet-data/consensus"

	validatorMainnetDatadir = "./mainnet-data/validator"
	validatorTestnetDatadir = "./testnet-data/validator"
	validatorDevnetDatadir  = "./devnet-data/validator"

	mainnetLogs = "./mainnet-logs"
	testnetLogs = "./testnet-logs"
	devnetLogs  = "./devnet-logs"

	configsRootDir = "./configs"

	// client flags - same values as dependencies but moved to flags for better code readability
	gethFlag       = gethDependencyName
	erigonFlag     = erigonDependencyName
	prysmFlag      = prysmDependencyName
	lighthouseFlag = lighthouseDependencyName

	mainnetConfig = configsRootDir + "/mainnet"
	testnetConfig = configsRootDir + "/testnet"
	devnetConfig  = configsRootDir + "/devnet"

	mainnetKeystore = "./mainnet-keystore"
	testnetKeystore = "./testnet-keystore"
	devnetKeystore  = "./devnet-keystore"

	// layers
	executionLayer = "execution"
	consensusLayer = "consensus"
	validatorLayer = "validator"

	// structure inside configs/selected-network directory.
	// we will select directory based on provided flag, by concatenating config path + file path
	genesisStateFilePath = "shared/genesis.ssz"
	chainConfigYamlPath  = "shared/config.yaml"
	gethTomlPath         = "geth/geth.toml"
	erigonTomlPath       = "erigon/erigon.toml"
	genesisJsonPath      = "shared/genesis.json"
	prysmYamlPath        = "prysm/prysm.yaml"
	validatorYamlPath    = "prysm/validator.yaml"

	// validator tool related flags
	validatorKeysFlag = "validator-keys"
)

var (
	mainnetEnabledFlag = &cli.BoolFlag{
		Name:  mainnetFlag,
		Usage: "Run for mainnet",
		Value: true,
	}
	testnetEnabledFlag = &cli.BoolFlag{
		Name:  testnetFlag,
		Usage: "Run for testnet",
		Value: false,
	}
	devnetEnabledFlag = &cli.BoolFlag{
		Name:  devnetFlag,
		Usage: "Run for devnet",
		Value: false,
	}

	consensusSelectedFlag = &cli.BoolFlag{
		Name:  consensusFlag,
		Usage: "Run for consensus",
		Value: false,
	}
	executionSelectedFlag = &cli.BoolFlag{
		Name:  executionFlag,
		Usage: "Run for execution",
		Value: false,
	}
	validatorSelectedFlag = &cli.BoolFlag{
		Name:  validatorFlag,
		Usage: "Run for validator",
		Value: false,
	}

	networkFlags = []cli.Flag{
		mainnetEnabledFlag,
		testnetEnabledFlag,
		devnetEnabledFlag,
	}

	validatorImportFlags = []cli.Flag{
		&cli.StringFlag{
			Name:   validatorWalletDirFlag,
			Usage:  "Selected wallet",
			Hidden: true,
		},
	}

	installFlags []cli.Flag
	updateFlags  []cli.Flag
	stopFlags    = []cli.Flag{
		executionSelectedFlag,
		consensusSelectedFlag,
		validatorSelectedFlag,
	}

	startFlags = []cli.Flag{
		&cli.BoolFlag{
			Name:  validatorFlag,
			Usage: "Run LUKSO node with validator",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  noSlasherFlag,
			Usage: "Disables slasher",
			Value: false, // default is true, we change it to false only when running validator
		},
		&cli.StringFlag{
			Name:  transactionFeeRecipientFlag,
			Usage: "Address to receive fees from blocks",
		},
		&cli.StringFlag{
			Name:  logFolderFlag,
			Usage: "Directory to output logs into",
			Value: mainnetLogs,
		},
	}
	logsFlags  []cli.Flag
	resetFlags []cli.Flag
	appFlags   = []cli.Flag{
		&cli.BoolFlag{
			Name:  agreeTermsFlag,
			Usage: "Accept terms of use. Default: false",
			Value: false,
		},
	}

	// GETH FLAGS
	// DOWNLOAD
	gethDownloadFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  gethTagFlag,
			Usage: "A tag of geth you would like to run",
			Value: "1.11.5",
		},
		&cli.StringFlag{
			Name:  gethCommitHashFlag,
			Usage: "A hash of commit that is bound to given release tag",
			Value: "a38f4108",
		},
	}
	// UPDATE
	gethUpdateFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  gethTagFlag,
			Usage: "A tag of geth you would like to run",
			Value: "1.11.5",
		},
	}
	// START
	gethStartFlags = []cli.Flag{
		&cli.StringFlag{
			Name:   gethDatadirFlag,
			Usage:  "Geth datadir",
			Value:  executionMainnetDatadir,
			Hidden: true,
		},
		&cli.StringFlag{
			Name:  gethConfigFileFlag,
			Usage: "Path to geth.toml config file",
			Value: mainnetConfig + "/" + gethTomlPath,
		},
		&cli.StringFlag{
			Name:  genesisJsonFlag,
			Usage: "Path to genesis.json file",
			Value: mainnetConfig + "/" + genesisJsonPath,
		},
	}
	// LOGS
	gethLogsFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  logFolderFlag,
			Usage: "Directory to output logs into",
			Value: "./mainnet-logs",
		},
	}
	// RESET
	gethResetFlags = []cli.Flag{
		&cli.StringFlag{
			Name:   gethDatadirFlag,
			Usage:  "geth datadir",
			Value:  executionMainnetDatadir,
			Hidden: true,
		},
	}

	// ERIGON
	// DOWNLOAD
	erigonDownloadFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  erigonTagFlag,
			Usage: "Tag for erigon",
			Value: "2.42.0",
		},
	}
	// START
	erigonStartFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  erigonConfigFileFlag,
			Usage: "Path to erigon.toml config file",
			Value: mainnetConfig + "/" + erigonTomlPath,
		},
		&cli.StringFlag{
			Name:  erigonDatadirFlag,
			Usage: "Erigon datadir",
			Value: executionMainnetDatadir,
		},
	}

	// PRYSM FLAGS
	// DOWNLOAD
	prysmDownloadFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  prysmTagFlag,
			Usage: "Tag for prysm",
			Value: "v4.0.1",
		},
	}
	// UPDATE
	prysmUpdateFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  prysmTagFlag,
			Usage: "Tag for prysm",
			Value: "v4.0.1",
		},
	}
	// START
	prysmStartFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  prysmGenesisStateFlag,
			Usage: "Path to genesis.ssz file",
			Value: mainnetConfig + "/" + genesisStateFilePath,
		},
		&cli.StringFlag{
			Name:   prysmDatadirFlag,
			Usage:  "Prysm datadir",
			Value:  consensusMainnetDatadir,
			Hidden: true,
		},
		&cli.StringFlag{
			Name:   prysmChainConfigFileFlag,
			Usage:  "Path to chain config file",
			Value:  mainnetConfig + "/" + chainConfigYamlPath,
			Hidden: true,
		},
		&cli.StringFlag{
			Name:  prysmConfigFileFlag,
			Usage: "Path to prysm.yaml config file",
			Value: mainnetConfig + "/" + prysmYamlPath,
		},
	}
	// LOGS
	prysmLogsFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  logFolderFlag,
			Usage: "Directory to output logs into",
			Value: "./mainnet-logs",
		},
	}
	// RESET
	prysmResetFlags = []cli.Flag{
		&cli.StringFlag{
			Name:   prysmDatadirFlag,
			Usage:  "prysm datadir",
			Value:  consensusMainnetDatadir,
			Hidden: true,
		},
	}

	// LIGHTHOUSE
	// DOWNLOAD
	lighthouseDownloadFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  lighthouseTagFlag,
			Usage: "Tag for lighthouse",
			Value: "v4.1.0",
		},
	}

	// VALIDATOR
	// DOWNLOAD
	validatorDownloadFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  validatorTagFlag,
			Usage: "Tag for validator binary",
			Value: "v4.0.1",
		},
	}
	// UPDATE
	validatorUpdateFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  validatorTagFlag,
			Usage: "Tag for validator binary",
			Value: "v4.0.1",
		},
	}
	// START
	validatorStartFlags = []cli.Flag{
		&cli.StringFlag{
			Name:   validatorDatadirFlag,
			Usage:  "Validator datadir",
			Value:  validatorMainnetDatadir,
			Hidden: true,
		},
		&cli.StringFlag{
			Name:  validatorKeysFlag,
			Usage: "Location of generated wallet",
			Value: mainnetKeystore,
		},
		&cli.StringFlag{
			Name:  validatorWalletPasswordFileFlag,
			Usage: "Location of the password file that you used to generate keys",
			Value: "",
		},
		&cli.StringFlag{
			Name:  validatorConfigFileFlag,
			Usage: "Path to prysm.yaml config file",
			Value: mainnetConfig + "/" + validatorYamlPath,
		},
		&cli.StringFlag{
			Name:   validatorChainConfigFileFlag,
			Usage:  "Prysm chain config file path",
			Value:  chainConfigYamlPath,
			Hidden: true,
		},
	}
	// LOGS
	validatorLogsFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  logFolderFlag,
			Usage: "Directory to output logs into",
			Value: mainnetLogs,
		},
	}
	// RESET
	validatorResetFlags = []cli.Flag{
		&cli.StringFlag{
			Name:   validatorDatadirFlag,
			Usage:  "Validator datadir",
			Value:  validatorMainnetDatadir,
			Hidden: true,
		},
	}
)

func (dependency *ClientDependency) PassStartFlags(ctx *cli.Context) (startFlags []string) {
	name := dependency.name
	args := ctx.Args()
	argsLen := args.Len()
	flagsToSkip := []string{
		validatorFlag,
		gethConfigFileFlag,
		prysmConfigFileFlag,
		validatorConfigFileFlag,
		validatorWalletPasswordFileFlag,
	}

	for i := 0; i < argsLen; i++ {
		skip := false
		arg := args.Get(i)
		for _, flagToSkip := range flagsToSkip {
			if arg == fmt.Sprintf("--%s", flagToSkip) {
				skip = true
			}
		}
		if skip {
			continue
		}

		if strings.HasPrefix(arg, fmt.Sprintf("--%s", name)) {
			if i+1 == argsLen {
				startFlags = append(startFlags, removePrefix(arg, name))

				return
			}

			// we found a flag for our client - now we need to check if it's a value or bool flag
			nextArg := args.Get(i + 1)
			if strings.HasPrefix(nextArg, "--") { // we found a next flag, so current one is a bool
				startFlags = append(startFlags, removePrefix(arg, name))

				continue
			}

			startFlags = append(
				startFlags,
				fmt.Sprintf("%s=%s", removePrefix(arg, name), nextArg),
			)
		}
	}

	return
}

func removePrefix(arg, name string) string {
	prefix := fmt.Sprintf("--%s-", name)

	arg = strings.TrimPrefix(arg, prefix)

	return fmt.Sprintf("--%s", strings.Trim(arg, "- "))
}

func prepareGethStartFlags(ctx *cli.Context) (startFlags []string, isCorrect bool) {
	isCorrect = true
	if !flagFileExists(ctx, gethConfigFileFlag) {
		isCorrect = false

		return
	}

	startFlags = clientDependencies[gethDependencyName].PassStartFlags(ctx)
	startFlags = append(startFlags, fmt.Sprintf("--config=%s", ctx.String(gethConfigFileFlag)))

	return
}

func prepareErigonStartFlags(ctx *cli.Context) (startFlags []string, isCorrect bool) {
	isCorrect = true
	if !flagFileExists(ctx, erigonConfigFileFlag) {
		isCorrect = false

		return
	}

	startFlags = clientDependencies[erigonDependencyName].PassStartFlags(ctx)
	startFlags = append(startFlags, fmt.Sprintf("--config=%s", ctx.String(erigonConfigFileFlag)))

	return
}

func preparePrysmStartFlags(ctx *cli.Context) (startFlags []string, err error) {
	genesisExists := flagFileExists(ctx, prysmGenesisStateFlag)
	prysmConfigExists := flagFileExists(ctx, prysmConfigFileFlag)
	chainConfigExists := flagFileExists(ctx, prysmChainConfigFileFlag)
	if !genesisExists || !prysmConfigExists || !chainConfigExists {
		err = errFlagPathInvalid

		return
	}

	logFilePath, err := prepareTimestampedFile(ctx.String(logFolderFlag), prysmDependencyName)
	if err != nil {
		return
	}

	startFlags = clientDependencies[prysmDependencyName].PassStartFlags(ctx)
	startFlags = append(startFlags, fmt.Sprintf("--log-file=%s", logFilePath))

	// terms of use already accepted during installation
	startFlags = append(startFlags, "--accept-terms-of-use")
	startFlags = append(startFlags, fmt.Sprintf("--config-file=%s", ctx.String(prysmConfigFileFlag)))

	if ctx.String(transactionFeeRecipientFlag) != "" {
		startFlags = append(startFlags, fmt.Sprintf("--suggested-fee-recipient=%s", ctx.String(transactionFeeRecipientFlag)))
	}
	if ctx.String(prysmGenesisStateFlag) != "" {
		startFlags = append(startFlags, fmt.Sprintf("--genesis-state=%s", ctx.String(prysmGenesisStateFlag)))
	}
	if ctx.String(prysmChainConfigFileFlag) != "" {
		startFlags = append(startFlags, fmt.Sprintf("--chain-config-file=%s", ctx.String(prysmChainConfigFileFlag)))
	}

	isSlasher := !ctx.Bool(noSlasherFlag)
	isValidator := ctx.Bool(validatorFlag)
	if isSlasher && isValidator {
		startFlags = append(startFlags, "--slasher")
	}

	return
}

func prepareLighthouseStartFlags(ctx *cli.Context) (startFlags []string, isCorrect bool) {
	isCorrect = true
	if !flagFileExists(ctx, gethConfigFileFlag) {
		isCorrect = false

		return
	}

	startFlags = clientDependencies[gethDependencyName].PassStartFlags(ctx)
	startFlags = append(startFlags, fmt.Sprintf("--config=%s", ctx.String(gethConfigFileFlag)))

	return
}

func prepareValidatorStartFlags(ctx *cli.Context) (startFlags []string, err error) {
	validatorConfigExists := flagFileExists(ctx, validatorConfigFileFlag)
	chainConfigExists := flagFileExists(ctx, prysmChainConfigFileFlag)
	validatorKeysExists := flagFileExists(ctx, validatorKeysFlag)
	if !validatorConfigExists || !chainConfigExists {
		err = errFlagPathInvalid

		return
	}
	if !validatorKeysExists {
		err = errValidatorNotImported

		return
	}

	validatorPasswordPath := ctx.String(validatorWalletPasswordFileFlag)
	if validatorPasswordPath == "" {
		var password []byte
		fmt.Print("\nWallet password flag not found: Please enter your wallet password: ")
		password, err = term.ReadPassword(0)
		fmt.Println("")

		if err != nil {
			log.Errorf("Couldn't read password: %v", err)

			return
		}

		passwordPath := ctx.String(validatorKeysFlag) + "/pass.txt"
		err = os.WriteFile(passwordPath, password, configPerms)
		if err != nil {
			log.Errorf("Couldn't create password file: %v", err)

			return
		}

		log.Infof("Password file created in %s", passwordPath)
		err = ctx.Set(validatorWalletPasswordFileFlag, passwordPath)
		if err != nil {
			return
		}
	}

	startFlags = clientDependencies[validatorDependencyName].PassStartFlags(ctx)

	logFilePath, err := prepareTimestampedFile(ctx.String(logFolderFlag), validatorDependencyName)
	if err != nil {
		return
	}

	// terms of use already accepted during installation
	startFlags = append(startFlags, "--accept-terms-of-use")
	startFlags = append(startFlags, fmt.Sprintf("--config-file=%s", ctx.String(validatorConfigFileFlag)))
	startFlags = append(startFlags, fmt.Sprintf("--log-file=%s", logFilePath))
	startFlags = append(startFlags, fmt.Sprintf("--suggested-fee-recipient=%s", ctx.String(transactionFeeRecipientFlag)))
	if ctx.String(validatorKeysFlag) != "" {
		startFlags = append(startFlags, fmt.Sprintf("--wallet-dir=%s", ctx.String(validatorKeysFlag)))
	}
	if ctx.String(validatorWalletPasswordFileFlag) != "" {
		startFlags = append(startFlags, fmt.Sprintf("--wallet-password-file=%s", ctx.String(validatorWalletPasswordFileFlag)))
	}
	if ctx.String(validatorChainConfigFileFlag) != "" {
		startFlags = append(startFlags, fmt.Sprintf("--chain-config-file=%s", ctx.String(validatorChainConfigFileFlag)))
	}
	return
}
