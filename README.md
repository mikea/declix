# declix

Declarative Unix (declix) is a configuration
format
that describes the state of Unix resources
and a command line tool that applies the
description to a running Unix system. 

Declix doesn't try to manage _full_ system
configuration but concerns itself only with
resources declared.

Declix can synchronize system state locally
or remotely using ssh. Bash and coreutils
are the only system dependencies required
to be present on a target.  

## Declix Configuration

Declix configuration is based on the pkl
configuration language. 