target "default" {
  dockerfile = "Dockerfile.build"
  target     = "artifact"
  output     = ["type=local,dest=./bin"]
}