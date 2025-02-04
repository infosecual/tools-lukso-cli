# LUKSO CLI

> ⚠️ DO NOT USE IN PRODUCTION, SCRIPTS ARE NOT DEPLOYED YET.

The LUKSO CLI is a command line tool to install, manage and set up validators of different clients for the LUKSO Blockchain. For more information and tutorial, please check the [LUKSO Tech Docs](https://docs.lukso.tech/).

## Features

- 🧰 Installation of Execution, Consensus, and Validator Clients
- 📀 Running a node as a validator
- 📑 Accessing various client logs

## Supported EVM Clients

> WIP: More clients will be added

The LUKSO CLI is able to install multiple clients for running the node.

- Execution Clients: [Geth](https://geth.ethereum.org/)
- Consensus Clients: [Prysm](https://github.com/prysmaticlabs/prysm)
- Validator Client for Staking: [Prysm](https://docs.prylabs.network/docs/how-prysm-works/prysm-validator-client)

## Setting up the Node

Process of setting up the node using the LUSKO CLI.

### Installing the LUKSO CLI

- Download and execute the LUKSO CLI installation script
- Running this script will install the full LUKSO CLI on Mac and Linux
- Installation directory: `/usr/local/bin/lukso`

```sh
# Might need admin access by typing `sudo` in front of the command
$ curl https://install.lukso.network | sh
```

### Initialise the Working Directory

1. Create and move into a working directory for your node's data

```sh
# Exchange [folder_name] with the folder name you want
$ mkdir [folder_name] && cd ./[folder_name]
```

2. Initialize the working directory

```sh
# Downloads all network configs from https://github.com/lukso-network/network-configs
$ lukso init
```

### Installing the Clients

3. Install choosen LUKSO clients into the working directory

```sh
# Installing Execution, Consensus, and Validator Client
# Might need admin access by typing `sudo` in front of the command
$ lukso install
```

### Starting the Clients

Please refer to the `start` command below for more information.

## Working Directories's Structure

As the LUKSO CLI is able to manage multiple clients for multiple blockchain networks in one folder, the structure of the node is set up in a generic way.

- When initializing the node (with `lukso init`), a global configuration folder is created, which holds shared and unique client information for each type of network.
- When executing commands, directories for the associated network type will be created accordingly.

Network Types: `mainnet`, `testnet`, `devnet`

> Even if multiple networks are set up, only one can be active at the time

```
lukso-node
│
├───configs                                 // Configuration
│   └───[network_type]                      // Network's Config Data
│       ├───shared
|       |   ├───genesis.json                // Genesis JSON Data
|       |   ├───genesis.ssz                 // Genesis Validator File
|       |   └───config.yaml                 // Global Client Config
│       ├───geth                            // Config for Geth Client
│       ├───prysm                           // Config for Prysm Client
│
├───[network_type]-keystore                 // Network's Validator Data
│   ├───keys                                // Encrypted Private Keys
│   ├───...                                 // Files for Signature Creation
|   ├───pubkeys.json                        // Validator Public Keys
|   ├───deposit_data.json                   // Deposit JSON for Validators
|   └───node_config.yaml                    // Node Configuration File
|
├───[network_type]-wallet                   // Network's Transaction Data
|
├───[network_type]-data                     // Network's Blockchain Data
│   ├───consensus                           // Storage of used Consensus Client
│   ├───execution                           // Storage of used Execution Client
│   └───validator                           // Storage of Validator Client
│
├───[network_type]-logs                     // Network's Logged Data
|
└───cli-config.yaml                         // Global CLI Configuration
```

## Available Commands

| Command            | Description                                                                              |
| ------------------ | ---------------------------------------------------------------------------------------- |
| `install`          | Installs choosen clients (Execution, Consensus, Validator) and their binary dependencies |
| `init`             | Initializes the working directory, it's structure, and network configuration             |
| `start`            | Starts all or specific clients and connects to the specified network                     |
| `stop`             | Stops all or specific clients that are currently running                                 |
| `log`              | Listens and saves all log events from a specific client in the current terminal window   |
| `status`           | Shows the client processes that are currently running                                    |
| `reset`            | Resets all or specific client data directories and logs excluding the validator keys     |
| `validator import` | Import the validator keys in the wallet                                                  |
| `version`          | Display the version of the LUKSO CLI that is currently installed                         |
| `help`, `h`        | Shows the full list of commands, global options, and their usage                         |

## Global Help Flag

| Flag       | Description                                                   |
| ---------- | ------------------------------------------------------------- |
| --help, -h | Can be added before or after a command to show it's help page |

## Examples and Explanations

For almost each command in the list, options can be added after it to modify or specify certain behavior.
Below, you can find examples and options tables for all available commands.

> Options containting [string] expects a string input in quotes.

> Options containting [int] expects an int input without quotes.

### `install`

#### How to install the clients

```sh
# User is able to select its Consensus and Execution clients.
# Detects pre-installed clients and will ask for overrides
$ lukso install

# Installs clients and agrees with Terms & Conditions automatically
$ lukso install --agree-terms
```

#### Options for `install`

| Option        | Description                               |
| ------------- | ----------------------------------------- |
| --agree-terms | Automatically accept Terms and Conditions |

### `start`

#### How to start the clients

```sh
# Starts your node and connects to LUKSO mainnet
# Uses the default config files from configs/mainnet folder
$ lukso start

# Starts your node and connects to mainnet as a validator
$ lukso start --validator

# Starts your node and connects to the LUKSO testnet
$ lukso start --testnet

# Starts your node and connects to testnet as a validator
$ lukso start --testnet --validator
```

#### How to start a node using config files

```sh
# Geth Configutation
# Change [config] to the name of your configuration file
$ lukso start --geth-config "./[config].toml"

# Prysm Configutation
# Change [config] to the name of your configuration file
# Change [custom_bootnode] to the bootnode's name
$ lukso start --prysm-config "./[config].yaml" --geth-bootnodes "[custom_bootnode]"

```

#### How to set up and customize a log folder

```sh
# Setting up a custom log directory
# Change [folder path] to a static or dynamic directory path
$ lukso start --log-folder "[folder_path]"
```

#### Options for `start`

| Option                               | Description                                                              |
| ------------------------------------ | ------------------------------------------------------------------------ |
| **NETWORK**                          |                                                                          |
| --mainnet                            | Starts the LUKSO node with mainnet data (default) (./configs/mainnet)    |
| --testnet                            | Starts the LUKSO node with testnet data (./configs/tesnet)               |
| --devnet                             | Starts the LUKSO node with devnet data (./configs/devnet)                |
| **VALIDATOR**                        |                                                                          |
| --validator                          | Starts the validator client                                              |
| --validator-keys [string]            | Directory of the validator keys (default: "./\[network_type\]-keystore") |
| --validator-wallet-password [string] | Location of password file that you used for generated validator keys     |
| --validator-config [string]          | Path to prysms validator.yaml config file                                |
| --transaction-fee-recipient [string] | The address that receives block reward from transactions (required for --validator flag)     |
| --genesis-json [string]              | The path to genesis JSON file                                    |
| --genesis-ssz [string]               | The path to genesis SSZ file                                     |
| --no-slasher                         | Disables slasher                                                         |
| **CLIENT OPTIONS**                   |                                                                          |
| --log-folder [string]                | Sets up a custom log directory (default: "./\[network_type\]-logs")      |
| --geth-config [string]               | Defines the path to geth TOML config file                                |
| --erigon-config [string]               | Defines the path to erigon TOML config file                            |
| --prysm-config [string]              | Defines the path to prysm YAML config file                               |
| --geth-[command]                     | The `command` will be passed to the geth client. [See the client docs for details](https://geth.ethereum.org/docs/fundamentals/command-line-options)  |
| --ergion-[command]                     | The `command` will be passed to the erigon client. [See the client docs for details](https://github.com/ledgerwatch/erigon#readme)  |
| --prysm-[command]                     | The `command` will be passed to the prysm client. [See the client docs for details](https://docs.prylabs.network/docs/prysm-usage/parameters)  |
| --lighthouse-[command]                     | The `command` will be passed to the lighthouse client. [See the client docs for details](https://lighthouse-book.sigmaprime.io/run_a_node.html)  |


### `stop`

#### How to stop all or specific clients

```sh
# Stops all running node clients
$ lukso stop

# Only stops the validator client
$ lukso stop --validator

# Only stops the execution client
$ lukso stop --execution

# Only stops the consensus client
$ lukso stop --consensus
```

#### Options for `stop`

| Option      | Description                |
| ----------- | -------------------------- |
| --validator | Stops the validator client |
| --execution | Stops the execution client |
| --consensus | Stops the consensus client |

### `log`

#### How to view logs of the clients

```sh
# Displays and saves the logs of the mainnet's consensus client
$ lukso log consensus

# Displays and saves the logs of the devnet's execution client
$ lukso log execution --devnet

# Displays and saves the logs of the testnet's validator
$ lukso log validator --testnet
```

#### Options for `log`

| Option    | Description                                                       |
| --------- | ----------------------------------------------------------------- |
| --mainnet | Logs the mainnet client (default) (./mainnet-logs/[client_type]/) |
| --testnet | Logs the testnet client (./testnet-logs/[client_type]/)           |
| --devnet  | Logs the devnet client (./devnet-logs/[client_type]/)             |

### `status`

#### How to check the status of the node

```sh
# Shows which client processes are currently running
$ lukso status
```

### `reset`

#### How to reset the node's data directory

```sh
# Resets LUKSO's mainnet data directory
$ lukso reset

# Resets LUKSO's testnet data directory
$ lukso reset --testnet

# Resets LUKSO's devnet data directory
$ lukso reset --devnet
```

#### Options for `reset`

| Option              | Description                 |
| ------------------- | --------------------------- |
| --mainnet [default] | Resets LUKSO's mainnet data |
| --testnet           | Resets LUKSO's testnet data |
| --devnet            | Resets LUKSO's devnet data  |

### `version`

#### How to check the version of the LUKSO CLI

```sh
# Displays the currently installed version of the LUKSO CLI
$ lukso version
```

### `help`

In addition to the help command, the global help flag can be used to generate help pages for commands

#### How to retrieve the help page in the CLI

```sh
# Displays the global help page
$ lukso help

# Displays the help page of the start command
$ lukso start --help

# Displays the help page of the start command
$ lukso start -h
```

## Running a Validator

Validator keys can be generated using:

- CLI: [tools-key-gen-cli](https://github.com/lukso-network/tools-key-gen-cli)
- GUI: [wagyu-key-gen](https://github.com/lukso-network/tools-wagyu-key-gen)

> Both of them will generate the validator keys directory.

After generating the validator keys, they can be imported into the LUKSO CLI. To fill the validator keys with funds to participate on the LUKSO Blockchain, you must use the [LUKSO Launchpad](https://deposit.mainnet.lukso.network) to send LYXe from your wallet to the generated keys.

#### Genesis Amounts

- All genesis validators will be prompted to vote for the initial token supply of LYX
- The initial token supply will determine how much LYX the Foundation will receive
- More details at: https://deposit.mainnet.lukso.network

#### Validator Stake

- Genesis Validators need to have at least 32 LYXe per validator
- Validators also need some ETH on their wallet to pay for gas expenses

### `validator import`

Import existing EIP-2335 keystore files (such as those generated by the [Wagyu Keygen](https://github.com/lukso-network/tools-wagyu-key-gen)) into Prysm.

#### How to import validator keys

```sh
# Regular import process
# You will be asked for password and key directory
lukso validator import

# Import skipping generated questions
lukso validator import --keys-dir "./myDir"
```

#### Options for `validator import`

| Option              | Description                                                                |
| ------------------- | -------------------------------------------------------------------------- |
| **SHORTCUT**        |                                                                            |
| --keys-dir [string] | Directory of the validator keys (default: "./\[network_type\]-keystore")   |
| **NETWORK**         |                                                                            |
| --mainnet           | Will import the keys for mainnet [default] (default: "./mainnet-keystore") |
| --testnet           | Will import the keys for testnet (default: "./testnet-keystore")           |
| --devnet            | Will import the keys for devnet (default: "./devnet-keystore")             |

For specific validator options, please visit the [Prysm Validator Specification](https://docs.prylabs.network/docs/wallet/nondeterministic). All flags and their parameters will be passed to the client. This can be useful to configure additional features like the validator's graffiti, extended logging, or connectivity options.

### Starting the Validator

#### How to start your validator (keys & tx fee recipient)

When you use `--validator`, the `--transaction-fee-recipient` flag is required.

```sh
# Specify the transaction fee recipient, also known as coinbase
# It is the address where the transactions fees are sent to
$ lukso start --validator --transaction-fee-recipient "0x12345678..."
```

If no `--validator-keys` is defined (example above), the CLI will look in the default directory: `./[network_type]-keystore`. If you want to provide a specific keystore directory, you can use `--validator-keys`:

```sh
# Validator keys
# Command split across multiple lines for readability
# Change [file_name] with the your password text file's name
$ lukso start --validator \
--transaction-fee-recipient "0x12345678..." \
--validator-keys "./custom-keystore-dir-path"
```

#### How to start a genesis node

```sh
# Command split across multiple lines for readability
# Make sure that both SSZ and JSON files are placed correctly
$ lukso start \
--genesis-ssz "./config/mainnet/shared/genesis.ssz" \
--genesis-json "./config/mainnet/shared/genesis.json"
```

## Uninstalling

The LUKSO CLI and downloaded clients are located within the binary folder of the user's system directory.
It can be removed at any time. All node data is directly located within the working directory.

```sh
# Make sure to stop the node
$ lukso stop

# Uninstall the LUKSO CLI
$ rm -rf /usr/local/bin/lukso

# Uninstall Geth Execution Client
$ rm -rf /usr/local/bin/geth

# Uninstall Prysm Consensus Client
$ rm -rf /usr/local/bin/prysm

# Remove the node data
# Make sure to backup your keys first
$ rm -rf ~/myNodeFolder
```

## Contributing

If you want to contribute to this repository, please check [`CONTRIBUTING.md`](./CONTRIBUTING.md).
