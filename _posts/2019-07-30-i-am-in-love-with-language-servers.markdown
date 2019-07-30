---
heroimg: /img/hero/strawberries.jpg
img: /img/gianarb.png
layout: post
title:  "I am in love with language servers"
date:   2019-07-30 08:08:27
categories: [post]
tags: [oss, language server, lsp, vim, go]
summary: "Language Servers are a nice way to reuse common features required by
editors such as auto complete, formatting, go to definition. This article is a
an open letter to share my love for this project with everybody"
changefreq: daily
---
Hello everybody! I am writing this article because I had a chat with a friend of
mine [@wdalmut](https://twitter.com/walterdalmut). He is a busy businessman and
vimmer like me.

This article is a quick and practical way to understand why language servers are
fantastic! Because they are!

When I started to use vim, I was developing almost all the time with PHP. PHP is
a tricky language, and back in the day, YouCompleteMe was the way to go to have
some autocomplete. However, as I said, PHP was not an excellent language for
that because the number of files is enormous, and to load all of them to suggest
functions and methods is tricky. Probably it is still like that.

Compared with a couple of years ago, we have more IDE and editors: Atom, VSCode,
Sublime, and many more. All of them to be successful requires the same features:

* Syntax highlight
* autocomplete
* Formatting

You can see the language serves as a protocol to abstract and reuse those
features, and many more such as go to definition, find all references, show
documentation.  Vim is almost like WordPress; there is a plugin for everything;
for example, there is an excellent vim-go plugin to make vim to work smart with
golang. The problem is works for vim and as I said, almost all the editors
need the set of shared features just listed to be usable on a daily base.

The community that builds a language has a lexer a parser, and it can traverse
the AST for the language that it develops. It has the knowledge and all the
building blocks to provide a tool usable by different clients. The way for them
to build something reusable is a language server. The clients are different
editors.

This story is real, and the Golang community develops
[gopls](https://github.com/golang/go/wiki/gopls) (it stays for go please), the
Golang language server. I use it with vim, and as a client, I use
[vim-coc](https://github.com/neoclide/coc.nvim)

vim-go >1.20 works with gopls as well, you need to set it explicitly:

```
let g:go_def_mode='gopls'
let g:go_info_mode='gopls'
```

This article expresses my love for language server, not for Go or vim-go or vim!
Even if I love all of them!

We spent a good amount of time to achieve developer happiness and to boost our
productivity.

There are more tools and developers around here. The killer feature for LSP is
its ability to create communities and to give us the ability to share reusable
code.

Other than the gopls I also use
[sourcegraph/javascript-typescript-langserver](https://github.com/sourcegraph/javascript-typescript-langserver)
for JavaScript and TypeScript and [rust-lang/rls-vscode](https://github.com/rust-lang/rls-vscode) for rust.

As you can see, rls-vscode looks from the name a VSCode project, but only
because also VSCode supports the language server protocol!

Thanks sourcegraph, microsoft and everybody behind the LSP effort!
