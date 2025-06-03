package main

import (
	"log"
	"math/rand"
)

func main() {
	for range 100 {
		if IsInExperiment(10) {
			log.Println("Run some cool new feature")
		} else {
			log.Println("Run the old version of the code")
		}
	}
}

// IsInExperiment returns true around rolloutPercentage% of the time.
func IsInExperiment(rolloutPercentage int32) bool {
	percent := rand.Int31n(100)
	return percent < rolloutPercentage
}

/*

Pros:
- Easy to control the rollout percentage.

Cons:
- The same user will probably end up with different versions of the experiment at different times.

*/
