# version-badge

![](./docs/icons/go.svg)
[![](./docs/icons/go-badge.svg)](https://github.com/essentialkaos/go-badge)
[![](./docs/icons/depexplorer.svg)](https://github.com/ArtARTs36/depexplorer)

version-badge - is console app and GitHub Action for generate SVG badges with version of language or dependency

Supported dependency managers:
* go.mod
* Composer
* NPM

Supported fonts:
* Arial
* Comic Sans MS
* Verdana

## Usage

### Generate badge with language version

```
./version-badge lang.svg
```

### Generate badge with dependency version

```
./version-badge dep.svg --object=<dependency-name>
```

### Generate badge from specific dependencies file

```
./version-badge lang.svg --from=go.mod1
```

### Generate badge with Comic Sans MS

```
./version-badge lang.svg --font=comic-sans-ms
```
