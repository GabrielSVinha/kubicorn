load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cluster.go",
        "interface.go",
        "machine.go",
    ],
    importpath = "k8s.io/kube-deploy/ext-apiserver/pkg/client/informers_generated/externalversions/cluster/v1alpha1",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/apis/cluster/v1alpha1:go_default_library",
        "//pkg/client/clientset_generated/clientset:go_default_library",
        "//pkg/client/informers_generated/externalversions/internalinterfaces:go_default_library",
        "//pkg/client/listers_generated/cluster/v1alpha1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/watch:go_default_library",
        "//vendor/k8s.io/client-go/tools/cache:go_default_library",
    ],
)
