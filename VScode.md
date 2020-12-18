# VS code note


### Mac developing toole
`xcode-select --install`

### Embed gitbash into VScode
Open `Open Settings(JSON)` and put below. git-bash will pop up out of the frame. Using
bash.exe which is under bin folder inside git.

`terminal.integrated.shell.windows="/path/to/git/bash.exe"`.

### Keyboard binding

##### toggle size of terminal and editor
Open `Open Keyboard Shortcut` and put below
```
    { "key": "ctrl+alt+m", "command": "workbench.action.toggleMaximizedPanel",      "when": "!terminalFocus" },
    { "key": "ctrl+`",     "command": "-workbench.action.terminal.toggleTerminal",  "when": "!terminalFocus" },
    { "key": "ctrl+alt+m", "command": "workbench.action.terminal.toggleTerminal",   "when": "terminalFocus" },
    { "key": "ctrl+left",  "command": "workbench.action.navigateBack"       },
    { "key": "ctrl+right", "command": "workbench.action.navigateForward"    }
```