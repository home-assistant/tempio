# tempio
A template helper for docker images which using [text/template](https://golang.org/pkg/text/template/) golang engine with [sprout functions](https://docs.atom.codes/sprout).

Support conf format:

- json

## Template functions

Templates can use all functions provided by [sprout](https://docs.atom.codes/sprout), the maintained successor of [sprig](http://masterminds.github.io/sprig/). Both naming schemes work:

- The new sprout names, e.g. `toUpper`, `base64Encode`, `toCamelCase`
- The legacy sprig names, e.g. `upper`, `b64enc`, `camelcase`

Existing templates written for sprig keep working unchanged. The legacy names are deprecated upstream and may be removed in a future major version of tempio, so prefer the sprout names for new templates.

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
