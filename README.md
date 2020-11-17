### DevSpace Plugin Example

This repository is an example implementation of a [devspace](https://github.com/devspace-cloud/devspace) plugin written in golang. The plugin adds the following to devspace:
- 2 new commands: 
  - `devspace login` that fakes a login
  - `devspace list env` that prints the environment variables
- 2 new [predefined variables](https://devspace.sh/cli/docs/configuration/variables/basics#predefined-variables):
  - `EXAMPLE_USER` that is filled from the environment variable `USER`
  - `EXAMPLE_HOME` that is filled from the environment variable `HOME`
- 2 plugin hooks that change the following:
  - When executing `devspace print` all environment variables that are available to the plugin are printed to the console
  - When executing any non plugin command, the plugin will print the fake login message 

### Building the plugin

Clone the repository and compile the plugin via go (make sure you have go modules enabled):
```
# linux & mac
go build -mod vendor -o main

# windows
go build -mod vendor -o main.exe
```

### Adding the plugin to devspace

Install [devspace](https://github.com/devspace-cloud/devspace) and run:
```
devspace add plugin plugin.yaml
```

Make sure the plugin was installed successfully:
```
$ devspace list plugins

 Name                      Version   Commands   Vars  
 devspace-plugin-example   0.0.1     2          2     

```

### Removing the plugin from devspace

Simply run this to remove the plugin:
```
devspace remove plugin devspace-plugin-example
```

### Testing the plugin

#### Test the new plugin commands

Run the following in a shell:
```
$ devspace login
Logging in...
Successfully logged in!
```

Or the new list command:
```
$ devspace list env
Plugin will print the environment variables
TERM=xterm-256color
SHELL=/bin/bash
USER=fabiankramm
COMMAND_MODE=unix2003
...
```

#### Test the plugin hooks

Run the following command to execute the plugin hooks:
```
$ devspace print --config examples/devspace-print.yaml
Logging in...
Successfully logged in!
Plugin will print the environment variables
TERM=xterm-256color
...
DEVSPACE_PLUGIN_COMMAND=print
DEVSPACE_PLUGIN_COMMAND_LINE=devspace print [flags]
DEVSPACE_PLUGIN_OS_ARGS=["devspace","print","--config","examples/devspace-print.yaml"]
DEVSPACE_PLUGIN_CONFIG=version: v1beta9
images:
  default:
    image: myusername/devspace
deployments:
- name: quickstart
  helm:
    componentChart: true
    values:
      containers:
      - image: myusername/devspace
dev: {}

DEVSPACE_PLUGIN_COMMAND_FLAGS=["--config","examples/devspace-print.yaml"]

-------------------

Vars:
[info]   No vars found

-------------------

Loaded path: /Users/fabiankramm/Programmieren/go-workspace/src/github.com/devspace-cloud/devspace-plugin-example/examples/devspace-print.yaml

-------------------

version: v1beta9
images:
  default:
    image: myusername/devspace
deployments:
- name: quickstart
  helm:
    componentChart: true
    values:
      containers:
      - image: myusername/devspace
dev: {}
```

#### Test the plugin predefined variables

Run the following command to test the plugin predefined variables:

```
$ devspace --config examples/devspace-predefined-vars.yaml run hello
Logging in...
Successfully logged in!
Hello fabiankramm! Your home directory is in /Users/fabiankramm
```

What happened here? In the `devspace-predefined-vars.yaml`, the new variables `EXAMPLE_USER` and `EXAMPLE_HOME` are used within a devspace command. So after the root plugin hook (which prints `Logging in...` etc.) was executed, devspace will call the plugin with `plugin-binary print env USER` (EXAMPLE_USER) and  `plugin-binary print env HOME` (EXAMPLE_HOME) to get values for the variables. devspace then replaces the variables with these values and executes the command. 
