package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	if IsInExperiment() {
		log.Println("Run some cool new feature")
	} else {
		log.Println("Run the old version of the code")
	}
}

// IsInExperiment returns true if the feature flag is enabled
func IsInExperiment() bool {
	str, ok := os.LookupEnv("FEATURE_FLAG_MY_COOL_FEATURE_ENABLED")
	if !ok {
		return false
	}
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return b
}

/*

Pros:
- We already use environment variables for many things in our services, so reusing the same feature for toggling a feature on/off is easy.
- We can test out a new feture in e.g. Staging before running in production.

Cons:
- If something goes wrong in "my cool feature", then it affects all users.

*/
