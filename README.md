## Go project with bazel
What is bazel ? Ref: https://docs.bazel.build/versions/master/build-ref.html

Using go rules: https://github.com/bazelbuild/rules_go

## Setup Bazel to use dependence from go.mod using [bazelisk](https://github.com/bazelbuild/bazelisk)

- Generate project with go mod

bazelisk run //:gazelle -- update-repos -from_file=go.mod

- Build package

bazelisk build //pkg/app:app

- Build project

bazelisk build //...

- Run package

bazelisk run //pkg/app:app

