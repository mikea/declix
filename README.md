# Declix

Declarative Unix (declix) is both:

- a [pkl](https://pkl-lang.org/)-based
  configuration format that describes the
  state of various Unix resources.
- a command line tool that applies the
  description to a running Linux system.

Declix doesn't try to manage _full_ system
configuration but concerns itself only with
resources declared. Thus you can use Declix
to manage only the part of the system, to
take over existing one or live side-by-side
with other configuration management systems.

Declix can synchronize system state locally
or remotely using ssh. Bash is the only
system dependencies required to be present
on a target.

Visit [docs/tutorial.md](docs/tutorial.md)
for a taste of declix operations.

## Features

- Over-ssh remote management with minimum
  target requirements
- Stateless operations
- Partial system management
- Powerful configuration language

Currently supported resources:

- Files (present/missing)
- `.deb` packages (present/missing)
- `apt` packages (present/missing)

## Requirements

There are two machines invloved in declix operations
with different requirements:

- driver machine: this is the machine that has access
  to the configuration, all the necessary resources
  and runs `declix` binary to manage the target machine.
- target machine: the machine that is being managed.

Sometimes driver and target machine could be the same,
but that is not very typical.

### Driver Requirements

- [pkl-cli](https://pkl-lang.org/main/current/pkl-cli/index.html#installation)
  is installed, and found in `PATH` under `pkl` name.
- go

### Target Requirements

- ssh+scp access with public key
- passwordless sudo
- bash

To manage certain kinds of resources additional tools
must be present on the target:

- Files: `sha256sum` from `coreutils`
- Deb packages: `dpkg`
- Apt packages: `dpkg`, `apt`

Many of them come preinstalled on most systems, others
could be installed using `declix` itself.

## Declix Configuration

Declix configuration is based on the
[pkl](https://pkl-lang.org/) configuration
language. It fully defines the schema for
each managed resource and ensures that declaration
files have no errors and are consistent with
the current declix version.

Declix doesn't concern itself with where
configuration comes from: it could be written
manually, generated by various tools or even
downloaded from the internet.

