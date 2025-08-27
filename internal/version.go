package internal

const (
	versionOffsetMajor uint64 = 48
	versionOffsetMinor uint64 = 32
	versionOffsetPatch uint64 = 16
)

// GetVersionV2 encodes the major, minor, and patch version numbers into a single uint64,
// using bit-shifting to create a compact representation.
//
// This format is used to construct the VersionV2 field in the Mumble protocol, allowing
// for efficient comparison and transmission of version information.
//
// Parameters:
//   - major: The major version number.
//   - minor: The minor version number.
//   - patch: The patch version number.
//
// Returns:
//   - A uint64 representing the combined version, encoded with offsets.
func GetVersionV2(major uint64, minor uint64, patch uint64) uint64 {
	return major<<versionOffsetMajor | minor<<versionOffsetMinor | patch<<versionOffsetPatch
}
