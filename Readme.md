import { CodeBlock } from 'react-code-blocks';

# GoChain

**GoChain** is a lightweight blockchain implementation written in Go, designed for learning and experimentation. This project demonstrates core blockchain concepts like block creation, proof-of-work, and data persistence using BoltDB.

## Project Structure

The project is organized into the following directories and files:

```plaintext
GoChain/
├── cmd/                    # Entry points for different applications
│   ├── test_app/           # A test application for demonstrating blockchain features
│   │   └── main.go         # Main entry point for the test application
│   └── user_cli_app/       # A CLI application for interacting with the blockchain
│       └── main.go         # Main entry point for the CLI application
├── internal/               # Core logic and internal modules
│   ├── blockchain/         # Blockchain implementation
│   │   ├── block.go        # Block structure and block-related logic
│   │   ├── blockchain.go   # Blockchain structure and management
│   │   └── proof.go        # Proof-of-Work (PoW) implementation
│   └── storage/            # Persistent storage for the blockchain
│       └── storage.go      # Functions for saving and loading the blockchain from BoltDB
├── .gitignore              # Git ignore file for unnecessary files
├── blockchain.db           # BoltDB file for storing the blockchain (auto-generated)
├── go.mod                  # Module file for dependency management
├── go.sum                  # Checksum file for dependencies
├── LICENSE                 # Project license
└── Readme.md               # Project documentation
```

## Key Components

### Blockchain (`internal/blockchain`)

- **Block**: Defines the structure of a block, including `Data`, `Hash`, `PrevHash`, and `Nonce`. Includes methods for calculating the block hash.
- **Blockchain**: A chain of blocks. Includes methods to add blocks and initialize the blockchain with a genesis block.
- **Proof of Work (PoW)**: Implements a proof-of-work algorithm for block validation using adjustable difficulty.

### Storage (`internal/storage`)

- Handles persistent storage of the blockchain using **BoltDB**.
- Supports saving the blockchain to a database and loading it back into memory.

### Applications (`cmd`)

- **Test App**: A sample application to demonstrate blockchain functionality.
- **User CLI App**: A command-line interface for users to interact with the blockchain (e.g., adding blocks and viewing the chain).

## Features

1. **Block Creation**: Create new blocks with data and link them using hashes.
2. **Proof of Work**: Ensures computational effort before adding blocks.
3. **Persistence**: Save and load the blockchain using a database.
4. **CLI Application**: Interact with the blockchain through a user-friendly CLI.

## How to Run

### Prerequisites

- Go 1.19 or later
- [BoltDB](https://github.com/boltdb/bolt)

### Running the Test Application

```bash
    cd cmd/test_app
    go run main.go
```

### Running the CLI Application

```bash
    cd cmd/user_cli_app
    go run main.go
```

Follow the prompts to add data to the blockchain or view the chain.

## Example Output

When running the CLI app:

```
Enter data for the block (or type 'exit' to stop): First Block
Enter data for the block (or type 'exit' to stop): Second Block
Enter data for the block (or type 'exit' to stop): exit

Data: Genesis
Hash: 4bf4fc9181b5c6d...ff
PrevHash:
Nonce: 17

Data: First Block
Hash: c5e4ac9189c8a8d...fd
PrevHash: 4bf4fc9181b5c6d...ff
Nonce: 45

Data: Second Block
Hash: 8f2ad881b9e8d6c...ee
PrevHash: c5e4ac9189c8a8d...fd
Nonce: 68

```

## License

This project is licensed under the [MIT License](LICENSE).
