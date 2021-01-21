# tempio
A template helper for docker images which using [text/template](https://golang.org/pkg/text/template/) golang engine.

Support conf format:

- json

Custom template commands

| Command                        | Description                                                                        |
|------------------------------- |----------------------------------------------------------------------------------- |
| **JoinString** .value ", "     | Join a array of strings with an seperator.                                         |
| **CalcAdd** a b    | Returns the sum of a and b.                                         |
| **CalcSubtract** a b    | Returns the difference of a from b.                                         |
| **CalcMultiply** a b    | Returns the product of a and b.                                         |
| **CalcDivide** a b    | Returns the division of a from b.                                         |

## Basic Example

```bash
tempio \
    -conf /config/coredns.json \
    -template /usr/share/corefile.tempio \
    -out /etc/corefile
```

## Pipe Example

```bash
echo '{"some": "value"}' | tempio \
    -template /usr/share/corefile.tempio \
    -out /etc/corefile
```

## Add tempio to your GitHub actions

```yaml
- name: Setup tempio
  uses: home-assistant/tempio@main
- name: Run tempio
  run: tempio -help
```