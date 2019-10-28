# SQLBoiler non deterministic generation issue

## The issues
1. The issue is that sqlboiler does not generate deterministic code
if executed against different postgresql engine versions. The generated code varies on different
postgresql versions which can cause various issues/bugs which are hard to find
2. SQLBoiler uses Golang maps which does not provide deterministic order of the keys
    which results in non-deterministic code generation.

## Version affected
* `SQLBoiler 3.6.0`

## Environment setup

### Requirements
* `make`
* `sqlboiler 3.6.0`
* `postgresql-11.5`
* `postgresql-9.5`
* `go 1.12.+`

## Test

### Test with postgresql-11.5
* Make sure you start from clean slate - `make db.drop`
* Setup the database - `make db.setup`
* Generate the schema models - `make generate`
* Postgresql version issue - Open `./schema/received_messages.go` and find `FindReceivedMessage`.
    Arguments of this function should be `FindReceivedMessage(exec boil.Executor, attendeeID string, messageID string)`
* Save `sqlboiler.sum` file somewhere else for later comparison

### Test with postgresql-9.5
* Make sure you start from clean slate - `make db.drop`
* Setup the database - `make db.setup`
* Generate the schema models - `make generate`
* Postgresql version issue - Open `./schema/read_messages.go` and find `FindReceivedMessage`.
    Arguments of this function should be `FindReceivedMessage(exec boil.Executor, messageID string, attendeeID string)`
* Now compare current `sqlboiler.sum` file with the one from the test with postgresql-11.5
* You can also check `received_messages-9.5.go` in the root directory which is generated with `postgresql-9.5` and you can compare it with the `postgresql-11.5` generated file.

### Why this happens
1. Different postgresql engines can export tables' columns in different order but sqlboiler actually counts
    on tables' column order to generate proper schema models.
2. SQLBoiler uses Golang maps to manage information about tables and columns but counts on the order of the tables' columns
    to generate proper schema models. Since the order of the key within the maps in Golang is not deterministic it can cause
    different columns order on each schema model generation on the same environment as well as across different environments.

## Solution proposal
Use stable sorting algorithm in sqlboiler generation related commands to ensure deterministic generations. This way
the generated code with be platform/version agnostic.
