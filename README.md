# goparsec

A toy parser combinator library, heavily inspired by haskell's [parsec](https://github.com/haskell/parsec) and alternative typeclass.

---------------------------------------

## Primitive combinators
  * `And(...Combinator)` - sequence of combinators.
  * `Or(...Combinator)` - attempt to parse using combinators until some success.
  * `Count2(n, m, Combinator)` - parse using combinator atleast **n** times, but not more than **m** times.
  * `RuneFunc(func(rune) bool)` - parse input using given rune function.
  * `Try(Combinator)` - ignore number of bytes parsed.
    * By default parser **does not** backtrack!
  * `EOF()` - parse for end of input.

## Additional combinators

Following combinators are based on primitive ones.

  * `Some(Combinator)` - parse using combinator *once or more*.
  * `More(Combinator)` - parse using combinator *zero times or more*.
  * `Opt(Combinator)` - parse using combinator *zero times or once*.
  * `Count(n, Combinator)` - parse using combinator exactly **n** times.
  * `Rune(rune)` - parse rune.
  * `Str(string)` - parse sequence of runes in given string.

---------------------------------------

## Examples

See `./examples` folder for examples how to use this library.
