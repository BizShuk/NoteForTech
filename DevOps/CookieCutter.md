# CookieCutter

## Requirement

- gradle, check [here](https://gradle.org/install/)
- JDK
- ~/.cookiecutterrc , check [example](#rc-example)

## How to use

`cookiecutter [-v] (local folder|git repo)`

### rc example

```yaml
default_context: # value for inject
  aws_account_id: 1211111
abbreviations:
  gn: https://github.com/{0}.git
cookiecutters_dir: "e:\path\to\template\dir"
replay_dir: "e:\path\to\replay\dir"
```

### issues: subprocess: CalledProcessError: for git clone

1. `git config core.longpaths=true`
2. Naming Files, Paths, and Namespaces for windows 10
3. Windows 10 "Enable NTFS long paths policy" option missing
