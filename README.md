# Same Text To Many (GoLang)

## Description

Application write same text and changes variables to values that you wrote in config file

## Run Application

```bash
# Build project
make build
# Run Project
./app [configFile] [contentFile] [outputFile]
```

- configFile - Path to file with configs
- contentFile - Path to file with content
- outputFile - Output content

## Example files

### configFile

```sh
name;age # Name of Variable
Ivan;20  # Variables to Change
John;42
Fred;38
```

### contentFile

```plain
# {name} - variable
Hello {name}. I'm {age} years old.
```

### outputFile

```plain
Hello Ivan. I'm 20 years old.
Hello John. I'm 42 years old.
Hello Fred. I'm 38 years old.
```
