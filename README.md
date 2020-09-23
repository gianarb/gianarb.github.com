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
