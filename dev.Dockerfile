FROM tradingai/bazel:latest AS build

Label maintainer="zmjhacker@gmail.com"

ENV PROJECT_PATH=/go/src/github.com/tradingai/tweb

WORKDIR ${PROJECT_PATH}

ENTRYPOINT ["bash", "entrypoint.sh"]
