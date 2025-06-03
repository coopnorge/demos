package main

import (
	"log"

	"github.com/coopnorge/demos/bendik-ab-testing/hashing"
)

func main() {
	for range 2 {
		id := "my-id" // Tested to NOT be in the experiment at 1/7 rollout
		if IsInExperiment(id, 1, 7) {
			log.Printf("ID %q: Run some cool new feature", id)
		} else {
			log.Printf("ID %q: Run the old version of the code", id)
		}
	}
	for range 2 {
		id := "my-id-4" // Tested to  be in the experiment at 1/7 rollout
		if IsInExperiment(id, 1, 7) {
			log.Printf("ID %q: Run some cool new feature", id)
		} else {
			log.Printf("ID %q: Run the old version of the code", id)
		}
	}
}

// IsInExperiment returns true around rolloutBuckets/numBuckers of the time.
// 'id' can be any kind of ID: UserID, CoopID, OrderID, ProductID, DeviceID.
func IsInExperiment(id string, rolloutBuckets int32, numBuckets int32) bool {
	if rolloutBuckets >= numBuckets {
		return true
	}
	// Convert "id" to a random range in the [0-uin32.Max] range
	idInt32 := hashing.StringHashToUInt32(id)
	// Divide the number into buckets of approximately equal size
	mod := idInt32 % uint32(numBuckets)
	//
	return mod < uint32(rolloutBuckets)
}

/*

Pros:
- The user will get the same consistent experience every time the code is run.
- Easy to _increase_/_decrease_ the rollout, but not change number of buckets.

Cons:
- Increasing the number of buckets will redistribute the users.
	- A user may be in the experiment when the rollout is set to 1/7, but not when the rollout is set up 1/6.

*/
