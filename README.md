This repository contains [gianarb.it](https://gianarb.it) and tools used to
manage my site.

## Static website

The website currently uses Jekyll. I am thinking about the possibility to move
to Hugo because I don't know Ruby and I don't want to learn it only to manage my
website. I know Go very well and Hugo is popular enough.

Anyway, I use Docker to run my site and there is an utility make target that
serves the website via `jekyll serve` running inside a Docker container.

```console
make
```

### SCSS

I use bootstrap and SASS to manage the CSS.

```console
make build
```

To install dependencies via npm, copy files from `node_modules` to a location
where Jekyll can pick them up.

```
make sass
```

Generates a CSS file from `scss/custom.scss`.
