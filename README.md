# docker-kill
Kill a running container using SIGKILL or a specified signal

## Usage
```sh
docker-kill [OPTIONS] CONTAINER [CONTAINER, ...]
```

### Options
- `-signal`  Signal to send to the container (Default: KILL)


## Examples
```sh
docker-kill 8e2db5361c6
```

```sh
docker-kill 8e2db5361c6 77876745e5d6
```

```sh
docker-kill -signal HUP 8e2db5361c6
```


## Available signals
- `ABRT`
- `ALRM`
- `BUS`
- `CHLD`
- `CLD`
- `CONT`
- `FPE`
- `HUP`
- `ILL`
- `INT`
- `IO`
- `IOT`
- `KILL`
- `PIPE`
- `POLL`
- `PROF`
- `PWR`
- `QUIT`
- `SEGV`
- `STKFLT`
- `STOP`
- `SYS`
- `TERM`
- `TRAP`
- `TSTP`
- `TTIN`
- `TTOU`
- `UNUSED`
- `URG`
- `USR1`
- `USR2`
- `VTALRM`
- `WINCH`
- `XCPU`
- `XFSZ`
