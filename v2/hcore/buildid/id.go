// Package buildid holds a per-build identity for donut-core artifacts (dylib/dll/so, DonutCli,
// tunnel service). The Flutter app never embeds this value: it asks the bundled core via
// `tunnel build-id` and compares that to the running service's build.id stamp.
//
// Set at link time in the donut-core repository (Makefile sets a random UUID when unset):
//
//	-X github.com/bobthebuilder1234232243/donutvpn-core/v2/hcore/buildid.ID=<uuid>
package buildid

// ID is overwritten by -ldflags -X. "unset" means the binary was linked without a build id.
var ID = "unset"

// Get returns the build identity.
func Get() string {
	if ID == "" {
		return "unset"
	}
	return ID
}

// IsSet reports whether this binary was built with a real build id.
func IsSet() bool {
	id := Get()
	return id != "" && id != "unset"
}
