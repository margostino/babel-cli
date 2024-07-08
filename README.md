# 🛠️ Babel CLI

(This project is under Babel Foundation initiative. You can read the manifest [here](https://github.com/margostino/babel-foundation))

A command-line interface for performing specific tasks such as manual indexing, searching, and other operations on Babel data.

<p align="center">
  <img src="https://github.com/margostino/babel-foundation/blob/master/assets/babel-architecture.png?raw=true" alt="Babel Foundation Architecture"/>
</p>

## Features

- **Manual Indexing**: Index your data manually from the command line.
- **Search**: Perform searches on indexed data.
- **Task Automation**: Automate various tasks related to Babel data management.

## Commands

### Indexing

```bash
babel-cli index [options]
```

### TODO

- [x] CRUD commands (basic)
- [x] Store assets in SQLite
- [x] Initialize database and home directory
- [ ] Tag implementation
- [ ] Sync with remote (Git)
- [ ] BabelQL implementation
- [ ] Workflow implementation: inbox -> project -> area -> resource -> archive
- [ ] Share implementation
- [ ] Bot implementation with LLM integration and quick assets (auto sync)
- [ ] Search implementation (local and remote)
  - By similarity
  - By tags
  - By parameters (name, category, etc.)
- [ ] Graph model implementation (Graphlite, BabelDB)
