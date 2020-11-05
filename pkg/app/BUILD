load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_test")
go_binary(
    name = "app", #  the name for Bazel, for deps and Bazel resolver
    srcs = ["main.go"], # entry points for modules, will be like artifacts
    importpath = "bazel_exam/pkg/app", # how we should import this module in our applications
    deps = [
        "//pkg/app/router:router",
        "@org_mongodb_go_mongo_driver//mongo:go_default_library",
        "@org_mongodb_go_mongo_driver//mongo/options:go_default_library",
        "@org_mongodb_go_mongo_driver//mongo/readpref:go_default_library",
    ], # an array of dependencies
)

go_test(
    name = "mainTest", # for the run test
    srcs = ["main_test.go"], # entry points for tests
)