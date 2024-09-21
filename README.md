# CSSKit (still WIP)

A simple, dependency-free tool for generating CSS utility classes.

## Missing features

- Support for template expressions.
- Support for script tags in HTML.
- Ignoring HTML comments.
- Tests.

Feel free to contribute.

But before you start working on a pull request 
for a feature that is not listed here, 
please create a proposal first.

## Installation

Prerequisites:
- [Go 1.23.1](https://go.dev/doc/install)

Install CSSKit from source:

```bash
go install github.com/igormichalak/csskit/cmd/csskit@latest
```

If the `csskit` command can't be found, make sure that the `$GOBIN` (or `$GOPATH/bin`) directory is added to your system PATH.

Here's how to find the location of the binary:
```bash
go env GOBIN
```
```bash
echo "$(go env GOPATH)/bin"
```

More info: https://go.dev/wiki/GOPATH

## Usage

Only JavaScript, HTML and Go template files can be scanned for class names.

```bash
csskit -o outfile.css infile1.js infile2.html ...
```

Supported extensions: `.js`, `.html`, `.gohtml`.

## Grammar

```ebnf
letter    = 'a' ... 'z' .
digit     = '0' ... '9' .
keyword   = letter, { letter } .
number    = digit, { digit }, [ '.', digit, { digit } ] .
unit      = 'px' | '%' | 'vw' | 'vh'
          | 'rad' | 'deg' | 'ms' | 's' .
className = keyword, { '-', keyword },
            [ '-', number, [ unit ] ] .
```

## Compatibility

Compatibility with other technologies such as Astro, 
TypeScript and JSX is not within the scope of this project.
