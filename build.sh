#!/bin/bash
# build.sh

# usage:
#     build.sh [SUBCMD]
#
#     SUBCMD:
#         build: build packages outputs to ./bin/
#         clean: remove build files (`bin/*` and `go.sum`)
#         exec [cmd]: execute any commands
#         test [dir]: run go test

SUBCMD="$1"
shift
ARGS="$@"

CMD_NAME="vorbispak"

EXEC_UID=${SUDO_UID-$(id -u)}
EXEC_GID=${SUDO_GID-$(id -g)}
EXEC_HOME="~${SUDO_USER-${USER}}"

# go command by docker.
go() {
  docker run -i --rm \
    -v $(eval echo "${EXEC_HOME}")/go:/go \
	-v $(realpath "${HOME}")/.ssh:/root/.ssh \
	-v $(realpath "${PWD}"):${PWD} \
	-w $(realpath "${PWD}") \
	golang:latest \
	sh -c "export GOFLAGS='-buildvcs=false' PATH=\$PATH:/go/bin ; go $@ ; echo \$? > /tmp/EXITCODE ; chown -R ${EXEC_UID} ./ ; chgrp -R ${EXEC_UID} ./ ; cat /tmp/EXITCODE" \
  | tr -d '\r'
}

exec_go_with_stdoutput() {
	ARGS="$@"

	tmpfile="/tmp/${RANDOM}"
	STDOUT=$(go "${ARGS}" 2>${tmpfile})
	STDERR=$(cat ${tmpfile})
	rm -f ${tmpfile}

	# ignore exitcode
	[[ "${STDOUT}" =~ ^[0-9]*$ ]] || ( [[ "$(echo "${STDOUT}" | tail -n 2 | head -n +1 )" = "" ]] && echo "${STDOUT}" | head -n -2 || echo "${STDOUT}" | head -n -1 )
	# stderr if exists
	[[ "${STDERR}" != "" ]] && echo "${STDERR}" >&2

	# exitcode is the lastline of stdout
	exit $(echo "${STDOUT}" | tail -n 1)
}

case ${SUBCMD} in
	"build" )
		exec_go_with_stdoutput 'get -u ./... && go mod tidy && go build -o bin/ ./...'
	;;

    "try" )
        ./build.sh build && ./bin/try ${ARGS}
    ;;

    "run" )
        ./build.sh build && ./bin/${CMD_NAME} ${ARGS}
    ;;

	"clean" )
		rm -rf go.sum bin/*
	;;

	"exec" )
		exec_go_with_stdoutput "${ARGS}"
	;;

	"test" )
		exec_go_with_stdoutput 'get -u ./... && go mod tidy && go test -v $(go list -m)/'${ARGS}
	;;
esac
