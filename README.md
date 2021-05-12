# tempio
A template helper for docker images which using [text/template](https://golang.org/pkg/text/template/) golang engine with [sprig functions](http://masterminds.github.io/sprig/).

Support conf format:

- json

## Basic Example

```bash
tempio \
    -conf /config/coredns.json \
    -template /usr/share/tempio/corefile \
    -out /etc/corefile
```

## Pipe Example

```bash
echo '{"some": "value"}' | tempio \
    -template /usr/share/tempio/corefile \
    -out /etc/corefile
```

## No Outfile Example

_When  `-out` is not supplied, the result is printed to the console._

```bash
$ echo '{"some": "value"}' | tempio \
    -template /usr/share/tempio/corefile
> {"some": "value", "other": "value"}
```

## Add tempio to your GitHub actions

```yaml
- name: Setup tempio
  uses: home-assistant/tempio@main
- name: Run tempio
  run: tempio -help
```
