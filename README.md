# gitctx

A CLI tool for easily managing and switching between multiple Git contexts.

## 📖 About

`gitctx` is a command-line utility written in Go that allows you to store and switch between different Git configurations. It's particularly useful for developers who work on different projects requiring distinct Git identities.

## 🌟 Features

- Create and manage multiple Git contexts with specific parameters
- Define different identities (name, email) for each context
- Manage distinct SSH and GPG keys per context
- Apply a context globally or for a specific directory
- Share common configurations between contexts
- Intuitive command-line interface based on Cobra

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
gitctx current
```

## ⚙️ Configuration

Configurations are stored in `~/.gitctx/config.yaml` by default.

## 🤝 Contributing

Contributions are welcome! Feel free to open an issue or pull request.

## 🚀 Future Improvements

- **Templates management**: Add ability to create context templates for easier setup of similar configurations
- **Export/import functionality**: Commands to export and import contexts for sharing between machines
- **Auto-detection**: Feature to automatically detect appropriate context based on repository details
- **Integration with other tools**: Consider integration with key managers or tools like direnv
- **Hooks**: Add the ability to define hooks that run when changing contexts
- **Interactive interface**: Interactive command (using a library like promptui) for easier context creation and selection
- **Change history**: Keep a history of context changes to easily revert to previous contexts
- **Cloud synchronization**: Option to securely sync contexts (without sensitive keys) via cloud services
- **Alternative key management**: Instead of generating SSH/GPG keys directly, consider using system commands or allowing import of existing keys for better security

## 📄 License

MIT
