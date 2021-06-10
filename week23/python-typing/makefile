# https://www.gnu.org/software/make/manual/make.html
SHELL=bash
.PHONY: default
default:

py_source=./src ./tests

.PHONY: test
test:
	poetry run pytest --cov-config=.coveragerc --cov=src -vv ./tests

.PHONY: validate-static
validate-static:
	poetry run isort --check $(py_source)
	poetry run black --check $(py_source)
	poetry run flake8 $(py_source)
	poetry run mypy --show-error-codes --show-error-context \
		$(py_source)


.PHONY: validate-types
validate-types:
	poetry run mypy --show-error-codes --show-error-context \
		$(py_source)



.PHONY: validate
validate: validate-static test
default: validate

.PHONY: validate-fix
validate-fix:
	poetry run isort $(py_source)
	poetry run black $(py_source)

.PHONY: install-editable
install-editable:
	rm -rv src/*.egg-info/
	## uninstall
	pip3 uninstall -y "$$(poetry version | gawk '{ print $$1 }')"
	## install
	rm -rv dist/
	poetry build --format sdist \
		&& tar --wildcards -xvf dist/*.tar.gz -O '*/setup.py' > setup.py \
		&& pip3 install --prefix="$${HOME}/.local/" --editable .

help:
	@printf "########################################################################\n"
	@printf "TARGETS:\n"
	@printf "########################################################################\n"
	@printf "%-32s%s\n" "help" "Show this output ..."
	@printf "%-32s%s\n" "validate" "Validate everything"
	@printf "%-32s%s\n" "validate-fix" "Fix auto-fixable validation errors"
	@printf "%-32s%s\n" "vendor" "Vendors dependencies"
	@printf "%-32s%s\n" "vendor-check" "Checks if vendoring is consistent"
	@printf "%-32s%s\n" "vendor-install" "Installs using vendored dependencies"
	@printf "########################################################################\n"
