A great okr platform needs great tools; we're excited to introduce Gtmhub CLI, our next generation command line experience for Gtmhub.

# Prerequisites
The only thing you need to successfully run the gtmhub-cli tool is an active subscription with [the best okr platform.](https://gtmhub.com)
# Installation

### Mac

```bash
brew tap gtmhub/homebrew-gtmhub
brew install gtmhub-cli
```
### Windows

```bash
choco install gtmhub-cli
```
### Linux
```bash
curl https://github.com/gtmhub/gtmhub-cli/blob/master/distribution/linux/install.sh | bash
```

# Usage
```bash
 gtmhub [command] [subcommand] {parameters}
```

### supported commands 
1. login - authorizes the tool to access your gtmhub instance and perform actions on your behalf.
2. logout - removes the integration with your account
3. status - displays items that need your attention
4. update - lets you update your key results
5. get - a group of commands that allows you to get various items from gtmhub
   * lists - lets you iteract with gtmhub lists
    * krs - lets you get information about your key results
    
# Exit codes
For scripting purposes, we output certain exit codes for differing scenarios.

|Exit Code   |Scenario   |
|---|---|
|0  |Command ran successfully.   |
|1   |Generic error; most likely the gtmhub server returned an error code. Enable debug for more information   |
|2   |Parser error; check input to command line.   |