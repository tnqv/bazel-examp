load("@io_bazel_rules_go//go:def.bzl", "go_library")
go_library(
    name = "router",
    srcs = ["router.go"],
    importpath = "bazel_exam/pkg/app/router",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/shared/health:health",
        "@com_github_gin_gonic_gin//:go_default_library"
    ]
)