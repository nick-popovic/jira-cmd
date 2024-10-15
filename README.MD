# Jira-Cmd
![Go Version](https://img.shields.io/badge/Go-1.23.2-blue)
![Bubble Tea](https://img.shields.io/badge/Bubble%20Tea-v1.1.1-brightgreen)
![Bubbles Version](https://img.shields.io/github/v/release/charmbracelet/bubbles?label=Bubbles) ![License](https://img.shields.io/badge/License-GPLv2-blue)

![Pre-Alpha](https://img.shields.io/badge/Status-Pre--Alpha-red)

## Introduction

Welcome to the first commit of my project aimed at bringing Jira to the command line and TUI using Bubble Tea and Go. This project is currently in **beta** and is not ready for deployment.

## Project Goals

The primary goal of this version (v0) is to commit a primitive working version of the interface that allows users to enter commands. At this stage, the focus is on the basic functionality of command input and display.

## Future State

The future state of this project aims to provide a unified interface for Jira on the command line and TUI that allows the user to query visualize and add data.

## Current Features

- [x] Users can enter commands (that don't do anything).
- [x] Commands are displayed above the input area.
- [x] The UI renders screen size changes.
- [x] `esc` exits the application.
- [~] `> clear` clears the UI text (see 'Known Issues' Below).

## Known Issues

- **Clear Command Bug**: When the user executes the `clear` command, the UI fails to re-render the first line visually. However, the command history data is correctly emptied in the backend.

## Debugging Configuration

We have set up two debugging configurations in `.vscode/launch.json`:

1. **Regular Debug**: This configuration is used for standard debugging.
    ```json
    {
        "name": "Regular Debug",
        "type": "go",
        "request": "launch",
        "mode": "auto",
        "program": "${workspaceFolder}",
        "env": {},
        "args": [],
        "buildFlags": "",
        "dlvToolPath": "${env:HOME}/go/bin/dlv"
    }
    ```

2. **BubbleTea Debug**: This configuration is used for debugging Bubble Tea applications. It requires a Delve server running in headless mode.
    ```json
    {
        "name": "BubbleTea Debug",
        "type": "go",
        "request": "attach",
        "mode": "remote",
        "remotePath": "${workspaceFolder}",
        "port": 2345,
        "host": "127.0.0.1",
        "apiVersion": 2
    }
    ```
