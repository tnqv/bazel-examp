load("@io_bazel_rules_go//go:def.bzl", "go_library")
go_library(
    name = "health",
    srcs = ["health.go"],
    importpath = "bazel_exam/pkg/shared/health",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_gin_gonic_gin//:go_default_library"
    ]
)