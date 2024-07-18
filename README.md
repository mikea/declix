# declix

Declarative Linux (declix) is a file format
that describes the state of Linux resources
and a command line tool that applies the
description to a running Linux system.

Declix doesn't try to manage _full_ system
configuration but concerns itself only with
resources declared. Thus you can use Declix
to manage only the part of the system, to
take over existing one or live side-by-side
with other configuration management systems.

Declix can synchronize system state locally
or remotely using ssh. Bash and coreutils
are the only system dependencies required
to be present on a target computer. Declix
doesn't have any persistent state. It sends
all resources needed to bring the target
up-to-date and sends them over ssh connection,
thus allowing deployments in a restricted
environment.
