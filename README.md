# CSSKit

A tool for generating CSS utility classes.

## Missing features

- Support for template expressions.

Feel free to contribute.

Before you start working on a pull request 
for a feature that is not listed here, 
please create a proposal first.

Compatibility with other technologies such as TypeScript 
and JSX is not within the scope of this project.

## Grammar

```ebnf
letter    = 'a' ... 'z' .
digit     = '0' ... '9' .
keyword   = letter, { letter } .
number    = digit, { digit } .
className = keyword, { '-', keyword },
            [ '-', number, [ '%' ] ] .
```
