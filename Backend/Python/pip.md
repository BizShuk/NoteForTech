# PIP

## How to install with pip and modules to use

`pip install --install-option="--prefix='E:\\writable\\pip'"` .
almost equal
`python setup.py install --prefix=E:\\writable\\pip`

`pip install --upgrade --user setuptools`
`pip install --upgrade --user cookiecutter`

## environment variables

- `PYTHONPATH`, for Python to lookup modules. Usually, it looks for `site-packages` folder.
Example: `export PYTHONPATH="/e/writable/Python38/site-packages:/c/Program\ Files/Python38/Lib/site-packages:/c/Users/tliu/AppData/Roaming/Python/Python38/site-packages"

- `PYTHONUSERBASE`, if `--uesr` is given, it will look given folder to install modules.
Example: `export PYTHONUSERBASE="E:\\wirtable"`, it will install modules into `E:\writable\Python38\Scripts` and `E:\writable\Python38\site-packages`(later one can be specified in to PYTHONPATH ahead)
