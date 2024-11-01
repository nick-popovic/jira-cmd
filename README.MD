# Jira-Cmd
![Go Version](https://img.shields.io/badge/Go-1.23.2-blue)
![Bubble Tea](https://img.shields.io/badge/Bubble%20Tea-v1.1.1-brightgreen)
![Bubbles Version](https://img.shields.io/github/v/release/charmbracelet/bubbles?label=Bubbles) ![License](https://img.shields.io/badge/License-GPLv2-blue)

![Pre-Alpha](https://img.shields.io/badge/Status-Pre--Alpha-red)

## Introduction

Welcome to my project aimed at bringing Jira to the command line and TUI using Bubble Tea and Go. This project is currently in **beta** and is not ready for deployment.

## Project Goals

The primary goal of this version (v0) is to commit a primitive working version of the interface that allows users to enter commands. At this stage, the focus is on the basic functionality of command input and display.

## Future State

The future state of this project aims to provide a unified interface for Jira on the command line and TUI that allows the user to query visualize and add data.

## Current Features

- [x] Users can send queries to ChatGPT and the frontend will render the markdown responce
- [x] The UI renders screen size changes.
- [x] `esc` exits the application.
- [x] Interactive 'Loading ...' is now present so user knows something is running in backround ...
- [x] Added tabs at the top for horizontal navigation

# Demo

<img width="1200" src="./assets/demo.gif" />

## Known Issues

- **Bug**: Sometimes when the window is 'too small' the last line of rendered output isn't shown ... dont ask me why, I have no idea yet ...
- **Bug**: Icons not rendering in gif  above ... again dont ask me why, I have no idea yet ...


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
