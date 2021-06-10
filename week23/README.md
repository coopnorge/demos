
# Week 23 Engineering Demo

![Gif](https://media.giphy.com/media/gLcUG7QiR0jpMzoNUu/giphy.gif)

Welcome to yet another engineering demo session @CoopX. This week we have some great content to showcase!


## Table of content

- [Week 23 Engineering Demo](#week-23-engineering-demo)
  - [Table of content](#table-of-content)
  - [Python-typing](#python-typing)
    - [Demo-content](#demo-content)
    - [Installation](#installation)
  
## Python-typing

Tired of those runtime type errors when developing in python? Mypy, together with type hints, allows you to perform static type checking on your python code. 

### Demo-content

You can find plenty of examples under the `tests` folder and a configuration file to play with `mypy.ini`.  Tests include examples on generics too, so make sure to check it out!

### Installation 

You need Python 3.5 or later to run mypy.
Mypy can be installed using pip:

    $ python3 -m pip install -U mypy

If you want to run the latest version of the code, you can install from git:

    $ python3 -m pip install -U git+git://github.com/python/mypy.git


Now, if Python on your system is configured properly (else see
"Troubleshooting" below), you can type-check the [statically typed parts] of a
program like this:

    $ mypy PROGRAM

If you want an easier life, there are cool integration with populare IDEs, for example:
- VS Code: provides [basic integration](https://code.visualstudio.com/docs/python/linting#_mypy) with mypy.
