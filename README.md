# CSSKit (still WIP)

A tool for generating CSS utility classes.

## Missing features

- Support for template expressions.

Feel free to contribute.

But before you start working on a pull request 
for a feature that is not listed here, 
please create a proposal first.

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
className = keyword, { '-', keyword },
            [ '-', number, [ '%' ] ] .
```

## Compatibility

Compatibility with other technologies such as Astro, 
TypeScript and JSX is not within the scope of this project.
