bazel clean
bazel run //:gazelle

# Whenever we add new go dependencies we need to update the repos through bazel
bazel run //:gazelle-update-repos

bazel build //...
