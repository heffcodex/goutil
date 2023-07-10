# A bunch of useful structures, constants and helpers for Go

![Test Status](https://img.shields.io/github/actions/workflow/status/heffcodex/goutil/checks.yml?label=test)
![Codecov](https://img.shields.io/codecov/c/github/heffcodex/goutil)
![License](https://img.shields.io/github/license/heffcodex/goutil)

## Installation

```bash
go get github.com/heffcodex/goutil/v2
```

## Contents
Most of the methods over standard types (primitives, arrays, maps) are generic.

- [uarray](https://github.com/heffcodex/goutil/tree/master/uarray) - array helpers:
  - convert (from map)
  - map
  - merge
  - in-place reverse
  - clustering
  - diffing
  - filtering
  - searching
- [uconst](https://github.com/heffcodex/goutil/tree/master/uconst) - useful constants:
  - sizes
- [umap](https://github.com/heffcodex/goutil/tree/master/umap) - map helpers:
  - convert (from array)
- [umath](https://github.com/heffcodex/goutil/tree/master/umath) - math helpers:
  - generic min
  - generic max
  - generic abs
- [umime](https://github.com/heffcodex/goutil/tree/master/umime) - MIME helpers:
  - stream MIME validation
  - MIME-aware file extension replacement
- [usync](https://github.com/heffcodex/goutil/tree/master/usync) - thread-safe types:
  - generic syncmap
- [utime](https://github.com/heffcodex/goutil/tree/master/utime) - time.Time wrapper:
  - marshalling/unmarshalling:
    - supported formats:
      - binary
      - json
      - protobuf
      - sql
      - text
    - custom layout (via global variables)
    - using local time by default
  - day/week/month shift methods
  - custom start-of-week day (via global variable)
  - month name translation (currently only for Russian):
    - singular
    - prepositional
- [utype](https://github.com/heffcodex/goutil/tree/master/utype) - generic type interfaces:
  - integers
  - chars
  - comparables
- [uvalue](https://github.com/heffcodex/goutil/tree/master/uvalue) - value helpers:
  - conditional pipe
  - references
  - generic IsZero

## Examples

Sorry, there are no such things (yet?)