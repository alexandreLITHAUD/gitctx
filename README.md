# gitctx 👽

A CLI tool for easily managing and switching between multiple Git contexts.

![Build](https://github.com/alexandreLITHAUD/Own-Git/actions/workflows/launch-CI.yaml/badge.svg)
![Tests](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/alexandreLITHAUD/3aff3ab94739bdcdd6a9640f0150eeda/raw/gitctx-tests.json)
![Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/alexandreLITHAUD/3aff3ab94739bdcdd6a9640f0150eeda/raw/gitctx-coverage.json)

## 📖 About

`gitctx` is a command-line utility written in Go that allows you to store and switch between different Git configurations. It's particularly useful for developers who work on different projects requiring distinct Git identities.

Link of the documentation : https://alexandrelithaud.github.io/gitctx

## 🌟 Features

- Create and manage multiple Git contexts with specific parameters
- Define different identities (name, email) for each context
- Manage distinct SSH and GPG keys per context
- Apply a context globally or for a specific directory
- Share common configurations between contexts
- Intuitive command-line interface based on Cobra
- Fetch common .gitconfig location for adding as contexts

## 📁 Context Storage

- Contexts are stored in `~/.gitctx/`.
- Each file represents a full Git configuration (`gitconfig`) and is named by the user.
- An internal metadata file tracks:
  - The **current** context.
  - The **last used** context (used for quick switching).

All context files are plain `.gitconfig` files and can be edited manually if needed.

## 🚀 Installation

```bash
# Via Go
go install github.com/your-username/gitctx@latest

# Or download a pre-compiled binary from the releases page
```

## 🔧 Usage

### Create a new context

```bash
gitctx create my-project --name "John Doe" --email "john@example.com" --ssh-key ~/.ssh/id_rsa_project --gpg-key ABC123
```

### List available contexts

```bash
gitctx list
```

### Use a context globally

```bash
gitctx use-context my-project
```

### Use a context for the current directory only

```bash
gitctx use-context my-project --local
```

### Set common configurations

```bash
gitctx common set --editor "vim" --diff-tool "meld"
```

### Create a context using common parameters

```bash
gitctx create other-project --name "John Doe" --email "john@example.com" --use-common
```

### Show current context

```bash
gitctx show
```

## ⚙️ Configuration

Configurations are stored in `~/.gitctx/config.yaml` by default.

## 🤝 Contributing

Contributions are welcome! Feel free to open an issue or pull request.

## 🚀 Future Improvements

### 🔄 Git Hook Integration

- Add a `post-commit` hook that logs the context used at the time of commit.
- Or a `prepare-commit-msg` hook to prefix commits with context info (e.g., `[Work]`, `[Personal]`).

### 🔐 Encrypted SSH Key Storage !!!

- Option to encrypt SSH keys inside context files using tools like Age or SOPS.
- Keep the private key encrypted until apply. When done remove the one of the other context, decrypt, and add it to SSH agent. So not scyptes totally only when not used.
- Useful for syncing across machines safely.

### ☁️ Remote Sync (future idea)

- Sync contexts between devices via:
  - Git repository (with encrypted files).
  - Dotfiles manager.
  - Cloud storage (Dropbox, etc.).

### 🧠 Smart Context Detection

- Detect potential mistakes (e.g., wrong email for a repo).
- Suggests or applies a context based on the Git remote (e.g., `github.com/company` → use work context).

### 🧩 LazyGit Plugin Support

- Create a plugin or wrapper for [LazyGit](https://github.com/jesseduffield/lazygit) to:
  - Display current context in the UI.
  - Switch contexts from within LazyGit.
  - Warn when a context mismatch is detected.

## 📄 License

MIT
