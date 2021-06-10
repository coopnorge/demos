# ...

Setup tools:

```bash
python3 -m pip install --user --upgrade pipx
pipx install poetry
pipx upgrade poetry
```

Run:

```bash
poetry install
poetry run aucampia.template.poetry

poetry run python -m aucampia.template.poetry.cli
poetry run python -m aucampia.template.poetry.cli_click
poetry run python -m aucampia.template.poetry.cli_typer
```

invokes:
```
poetry run inv -l
poetry run inv -e validate
poetry run inv -e validate-fix

poetry run inv -e update
poetry run inv -e update-latest
poetry run inv -e update-latest --accept
poetry run inv -e install-editable
poetry run inv -e install-editable --escaped
```

Validate:

```bash
poetry run black ./src ./tests

poetry run black --check ./src ./tests
poetry run flake8 ./src ./tests
poetry run mypy --strict ./tests ./src
poetry run pytest ./tests

poetry run pylint ./src ./tests
poetry run pytest ./tests
poetry run mypy --strict ./src ./tests

poetry run tox
```

Publish

```bash
# Get and set token:
# https://pypi.org/manage/account/token/
poetry config pypi-token.pypi ...

poetry version minor
poetry version prerelease
find src/ -name '__init__.py' -print0 | xargs -0 sed -i -e 's|^.*__version__.*$|__version__ = "'"$(poetry version | gawk '{ print $2 }')"'"|g'
git status
git commit -m "preparing for versioning" .; git push
git tag "v$(poetry version | gawk '{ print $2 }')"
git push --tags
poetry publish --build
```

## Other stuff

```
poetry cache clear --all .
\rm -vr /home/iwana/.cache/pypoetry/virtualenvs/
find $(dirname $(poetry run which python))/../lib/*/site-packages/ | grep egg-link
echo "$(dirname -- "$(dirname -- "$(poetry run which python)")")"
\rm -vr "$(dirname -- "$(dirname -- "$(poetry run which python)")")"

find . -depth \( -name '*.egg-info' -o -name '__pycache__' -o -name '*.pyc' \) -print0 | xargs -0 rm -vr

poetry install -vvv
PYTHONPATH="$(pwd)/src/${PYTHONPATH:+:${PYTHONPATH}}" poetry run pylint src/
```

```
pip3 uninstall -y "$(poetry version | gawk '{ print $1 }')"
\rm -rv dist/; poetry build --format sdist && tar --wildcards -xvf dist/*.tar.gz -O '*/setup.py' > setup.py && pip3 install --prefix="${HOME}/.local/" --editable .

\rm -rv dist/;
poetry build --format sdist
tar -xvf --wildcards dist/*.tar.gz -O '*/setup.py' > setup.py
pip3 install --prefix="${HOME}/.local/" --editable .
```
