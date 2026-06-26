package lasagna

// TODO: define the 'OvenTime' constant
// lasagnaVal := lasagnaOptions {
//     expectedTime: ovenTime,
//     actualTime: 0,
//     layers: 0
// }
const ovenTime = 40

type lasagnaOptions struct {
    expectedTime int
    actualTime int
    layers int
}


// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(opts lasagnaOptions) int {
	if opts.actualTime < opts.expectedTime {
        return opts.expectedTime - opts.actualTime
    }
    
    return 0
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	panic("PreparationTime not implemented")
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	panic("ElapsedTime not implemented")
}
