package hashing

import (
	"fmt"
	"hash/fnv"
)

// StringHashToUInt32 converts any input string into a 32-bit unsigned integer using FNV-1a.
// Note that is is not a cryptographically safe, but that is not required for experimenting.
// The hashing algorithm could be different from FNV-1a, as long as it consistently gives the same output number for the same input.
func StringHashToUInt32(input string) uint32 {
	hasher := fnv.New32a()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		// The 'Hash' interface from the stdlib implements the 'io.Writer' interface, which returns an error, but it is docunmented to never happen.
		// Not propagating an impossible error simplifies the signatures in this package.
		panic(fmt.Sprintf("The stdlib FNV-1a hash function failed with an error, but it is documented to never do that: %v", err))
	}
	hash := hasher.Sum32()
	return hash
}
