# Demo of EditorConfig

[EditorConfig](https://editorconfig.org/) makes it easy to define how editors
should behave for code in a repository or folder.

This demo has two subdirectories:

- gitattr-fallback: This shows how editorconfig should be configured if
  gitattributes are left as is.
- gitattr-overruide: This shows how gitattributes should be configured if
  editorconfig manages encoding and line endings.

```bash
## Directory where gitattributes are overwritten
xxd gitattr-override/example.py
python gitattr-override/example.py
```


```bash
## Directory where parts managed by gitattributes are ignored
xxd gitattr-fallback/example.py
python gitattr-fallback/example.py
```
