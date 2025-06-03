package main

import (
	"fmt"
	"log"

	"github.com/coopnorge/demos/bendik-ab-testing/hashing"
)

func main() {
	for range 2 {
		id := "my-id" // Tested to NOT be in the experiment at 14% rollout
		if IsInExperiment("my awesome experiment", id, 14) {
			log.Printf("ID %q: Run some cool new feature", id)
		} else {
			log.Printf("ID %q: Run the old version of the code", id)
		}
	}

	for range 2 {
		id := "my-id-23" // Tested to be in the experiment at 14% rollout
		if IsInExperiment("my awesome experiment", id, 14) {
			log.Printf("ID %q: Run some cool new feature", id)
		} else {
			log.Printf("ID %q: Run the old version of the code", id)
		}
	}

	for range 2 {
		id := "my-id-23" // Tested to NOT be in the OTHER experiment at 14% rollout
		if IsInExperiment("some OTHER experiment", id, 14) {
			log.Printf("ID %q: Run some cool new feature", id)
		} else {
			log.Printf("ID %q: Run the old version of the code", id)
		}
	}
}

// IsInExperiment returns true around rolloutPercentage% of the time.
// 'id' can be any kind of ID: UserID, CoopID, OrderID, ProductID, DeviceID.
func IsInExperiment(experimentKey string, id string, rolloutPercentage uint32) bool {
	if rolloutPercentage >= 100 {
		return true
	}
	// Create a composite key of both the experiment and the id.
	compositeKey := fmt.Sprintf("%s:%s", experimentKey, id)
	// Convert "id" to a random range in the [0-uin32.Max] range
	idUInt32 := hashing.StringHashToUInt32(compositeKey)
	// Divide the number into buckets of approximately equal size
	percentage := idUInt32 % uint32(100)
	return percentage < rolloutPercentage
}

/*

Pros:
- The user will get the same consistent experience every time the code is run.
- Easy to change the rollout, avoids the problem of changing sizes of the buckets by using percentage.
- The user-distribution can now be different per experiment.

Cons:
- Hard-coded limit to 100 "buckets".
- No support for multi-variant experiments

*/
