# demo.tape — run: vhs demo.tape
# Generates a demo GIF for the README
# Requires: https://github.com/charmbracelet/vhs

Output demo.gif

Set FontSize 14
Set Width 1200
Set Height 600
Set Theme "Catppuccin Mocha"
Set Padding 20

Type "mcp version"
Enter
Sleep 2s

Type "mcp search database"
Enter
Sleep 2s

Type "mcp info sqlite"
Enter
Sleep 2s

Type "mcp doctor"
Enter
Sleep 3s
