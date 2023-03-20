// docker-bake.hcl
target "docker-metadata-action" {}

target "build" {
  inherits = ["docker-metadata-action"]
  context = "./"
  dockerfile = "contrib/executor/init/build/agent/Dockerfile"
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}