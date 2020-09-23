[![Netlify Status](https://api.netlify.com/api/v1/badges/32b3e3fa-01c0-49f3-9e5c-bfa5bf0f5e3e/deploy-status)](https://app.netlify.com/sites/dazzling-neumann-227af1/deploys)

This repository contains [gianarb.it](https://gianarb.it) and tools used to
manage my site.

## build it locally

I use Nix and nix-shell to start Jekyll:

```
$ nix-shell
Configuration file: /Users/gianarb/git/gianarb.github.com/_config.yml
            Source: /Users/gianarb/git/gianarb.github.com
       Destination: /Users/gianarb/git/gianarb.github.com/_site
 Incremental build: disabled. Enable with --incremental
      Generating...
```

## Deploy live

I use Netlify to deploy this website and to manage the TLS certificate.

At every commit to `master` the website gets built and deployed. Easy like that.
