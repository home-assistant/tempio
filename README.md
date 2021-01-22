# tempio
A template helper for docker images which using [text/template](https://golang.org/pkg/text/template/) golang engine with [sprig functions](http://masterminds.github.io/sprig/).

Support conf format:

- json

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