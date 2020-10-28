# tempio
A template helper for docker images which using [text/template](https://golang.org/pkg/text/template/) golang engine.

Support conf format:

- json

Custom template commands

| Command                        | Description                                                                        |
|------------------------------- |----------------------------------------------------------------------------------- |
| **JoinString** .value ", "     | Join a array of strings with an seperator.                                         |

## Example

```bash
$ tempio \
    -conf /config/coredns.json \
    -template /usr/share/corefile.tempio \
    -out /etc/corefile
```
