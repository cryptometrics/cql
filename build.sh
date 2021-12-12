bazel clean
bazel run //:gazelle
bazel run //:gazelle-update-repos
bazel build //...
