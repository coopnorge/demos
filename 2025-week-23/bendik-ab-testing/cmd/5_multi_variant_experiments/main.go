package main

import (
	"fmt"
	"log"

	"github.com/coopnorge/demos/bendik-ab-testing/hashing"
)

func main() {
	const featureVariantRoundedCorners = "rounded-corners"
	const featureVariantSquareCorners = "square-corners"
	const featureVariantCircle = "circle"
	const featureVariantControl = "control"
	for range 2 {
		id := "my-id" // Tested to NOT be in any variant-group experiment at 14% rollout
		variant := GetExperimentVariant("my awesome experiment", id, []Variant{
			{featureVariantRoundedCorners, 5},
			{featureVariantSquareCorners, 5},
			{featureVariantCircle, 5},
			{featureVariantControl, 5},
		})
		switch variant {
		case featureVariantRoundedCorners:
			log.Printf("ID %q: Run the new feature with rounded corners", id)
		case featureVariantSquareCorners:
			log.Printf("ID %q: Run the new feature with square corners", id)
		case featureVariantCircle:
			log.Printf("ID %q: Run the new feature with circle corners", id)
		case featureVariantControl:
			log.Printf("ID %q: Run the old version of the code", id)
		default:
			log.Printf("ID %q: Not in any variant group", id)
		}
	}

	for range 2 {
		id := "my-id-23" // Tested to be in the square-corner variant.
		variant := GetExperimentVariant("my awesome experiment", id, []Variant{
			{featureVariantRoundedCorners, 5},
			{featureVariantSquareCorners, 5},
			{featureVariantCircle, 5},
			{featureVariantControl, 5},
		})
		switch variant {
		case featureVariantRoundedCorners:
			log.Printf("ID %q: Run the new feature with rounded corners", id)
		case featureVariantSquareCorners:
			log.Printf("ID %q: Run the new feature with square corners", id)
		case featureVariantCircle:
			log.Printf("ID %q: Run the new feature with circle corners", id)
		case featureVariantControl:
			log.Printf("ID %q: Run the old version of the code", id)
		default:
			log.Printf("ID %q: Not in any variant group", id)
		}
	}
}

// Variant represents a variant in an experiment.
type Variant struct {
	Name       string
	Percentage uint32
}

// GetExperimentVariant returns the variant assigned to the given id, or "" if unaffected.
// variants is a slice of tuples: (variantName, percentage)
func GetExperimentVariant(experimentKey string, id string, variants []Variant) string {
	// Create a composite key of both the experiment and the id.
	compositeKey := fmt.Sprintf("%s:%s", experimentKey, id)
	// Convert "experiment:id" to a random uint in the [0-uin32.Max] range
	idUInt32 := hashing.StringHashToUInt32(compositeKey)
	percentage := idUInt32 % uint32(100)

	var cumulative uint32
	for _, v := range variants {
		cumulative += v.Percentage
		if percentage < cumulative {
			return v.Name
		}
	}
	return "" // Not in any variant group
}

/*

Pros:
- The user will get the same consistent experience every time the code is run.
- The user-distribution can now be different per experiment.
- Supports multiple variants of the same experiment.

Cons:
- By changing the percentages, A user may move from one variant to another.
- Hard-coded limit to 100 "buckets".

*/
