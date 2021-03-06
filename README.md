# Clocstat line count reports tool
This `clocstat` tool that reads a configuration file and generates a report of lines of code. The `clocstat` tool use [`cloc`](https://github.com/AlDanial/cloc) internally to count the lines of code.

## Getting started
### Prerequisites
- The `clocstat` tool use [`cloc`](https://github.com/AlDanial/cloc) internally to count the lines of code. Need to install `cloc` first, the installation guide can be found in [`cloc`](https://github.com/AlDanial/cloc) repository.
- There is a [github action](https://github.com/marketplace/actions/count-lines-of-code-cloc) that can be used if use want to integrate with CI 
### Installation
1. `cd` to the project root
2. `go install ./cmd/clocstat/`to install the `clocstat` to binary
### Running locally
- `make build` to build the project
- `./clocstat --config ./example.clocstat --verbose yes` to run the example config in this folder.

## Example

The following `clocstat` configuration:
```
# Project files that aren't tests
Business Logic: -match-f='.*(?<!test).go$' .

# Project test files
Tests: -match-f='.*test.go$' .

# Autogenerated business logic
Business Logic (Generated): internal/gen

# Hand coded business logic
Business Logic (Hand Coded): cmd internal/custom

# Autogenerated diagrams
Diagrams (Generated): -match-f='.*.svg$' docs/gen

# Hand coded diagrams
Diagrams (Hand Coded): -match-f='.*.svg$' docs/custom

# Compare the ratio of test and non-test files
!compare: Business Logic, Tests

# Compare the ratio of generated and hand coded business logic
!compare: Business Logic (Generated), Business Logic (Hand Coded)

# Compare the number of generated and hand coded diagrams
!compare: Diagrams (Generated), Diagrams (Hand Coded): files, files%
```

Generate a report in the following format:
```
                      files          code            code%
-----------------------------------------------------------
Business Logic        16             2400            80%
-----------------------------------------------------------
Tests                 1              600             20%
-----------------------------------------------------------


                                   files          code            code%
------------------------------------------------------------------------
Business Logic (Generated)         15             2220            86%
------------------------------------------------------------------------
Business Logic (Hand Coded)        1              180             14%
------------------------------------------------------------------------


                             files      files%
-----------------------------------------------
Diagrams (Generated)         20         100%
-----------------------------------------------
Diagrams (Hand Coded)        0          0%
-----------------------------------------------
```

## Usage

`clocstat` takes the following arguments.

#### Config

`config` is the location of the configuration file: `clocstat example.clocstat`. 

If no configuration file is present then `clocstat` looks for the `.clocstat` file in the current working directory.

#### Verbose

`--verbose` set to `yes` or `true` prints out the raw `cloc` output for each comparison.

## Config File Format

`clocstat` configuration files consist of `comment`, `count` and `compare` lines:

#### Comment

`comment` lines start with the `#` symbol and are ignored during processing:
```
# an example comment
```

Blank lines are also regarded as comments and ignored.

#### Count

`count` lines are used to calculate statistics about a set of files:
```
Tests: -match-f='.*test.go$' .
```

`count` lines are in the format: `identifier: command`.

`identifier` is the unique identifier of the count.

`command` are the arguments to use when calling `cloc`. For example, the above `Tests` count would generate the following call: `cloc -match-f='.*test.go$' .`

#### Compare

`compare` lines are in the format: `!compare: counts: selectors`.

`counts` is a comma-separated list of `count` identifiers.

`selectors` is a comma-separated list of items to compare between the two `counts`. The following values are valid: `files`, `files%`, `blank`, `blank%`, `comment`, `comment%`, `code`, `code%`. If no selectors are present then they default to: `files, code, code%`.

`clocstat` generates a table of data for each `compare` line that includes the data for each `selector`. 

For example: `!compare: Diagrams (Generated), Diagrams (Hand Coded): files, files%` generates the following table:

```
                             files      files%
-----------------------------------------------
Diagrams (Generated)         20         100%
-----------------------------------------------
Diagrams (Hand Coded)        0          0%
-----------------------------------------------
```

## Implementation details

`cloc` counts lines of code and separates the result by file type. If more than one file type is detected then `cloc` outputs a `SUM` row. If only a single file type is detected then `cloc` can be forced to output the `SUM` row by including the `--sum-one` option. The value that `clocstat` takes for each `selector` should be the `SUM` value.

