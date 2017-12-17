package config

// inspired by
// 1) https://stackoverflow.com/questions/11354518/golang-application-auto-build-versioning
// 2) https://github.com/golang/go/wiki/GcToolchainTricks#including-build-information-in-the-executable

// Version of the app (uses semantic versioning)
var Version string

// BuildDate represents the date of the current build
var BuildDate string
