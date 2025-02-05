package main

import "fmt"

// FrequencyTracker structure with two maps:
// - one to track the frequency of each number
// - another to track how many numbers have a specific frequency
type FrequencyTracker struct {
	numberFrequency map[int]int
	frequencyCount  map[int]int
}

// Constructor to initialize the FrequencyTracker
func Constructor() FrequencyTracker {
	return FrequencyTracker{
		numberFrequency: make(map[int]int),
		frequencyCount:  make(map[int]int),
	}
}

// Adds a number to the data structure
func (ft *FrequencyTracker) add(number int) {
	// Remove old frequency if the number already exists
	oldFreq := ft.numberFrequency[number]
	if oldFreq > 0 {
		ft.frequencyCount[oldFreq]--
	}

	// Increment the frequency of the number
	newFreq := oldFreq + 1
	ft.numberFrequency[number] = newFreq

	// Update the frequency count map
	ft.frequencyCount[newFreq]++
}

// Deletes one occurrence of a number
func (ft *FrequencyTracker) deleteOne(number int) {
	oldFreq := ft.numberFrequency[number]
	if oldFreq == 0 {
		return // No such number to delete
	}

	// Remove the number from the frequency count map
	ft.frequencyCount[oldFreq]--

	// Decrement the frequency of the number
	newFreq := oldFreq - 1
	if newFreq > 0 {
		ft.numberFrequency[number] = newFreq
		ft.frequencyCount[newFreq]++
	} else {
		// If frequency goes to zero, remove it from the map
		delete(ft.numberFrequency, number)
	}
}

// Checks if there is any number with the given frequency
func (ft *FrequencyTracker) hasFrequency(frequency int) bool {
	return ft.frequencyCount[frequency] > 0
}

func main() {
	ft := Constructor()
	ft.add(1)
	ft.add(1)
	ft.add(2)
	ft.deleteOne(1)

	fmt.Println(ft.hasFrequency(2)) // true (because 1 appears twice)
	fmt.Println(ft.hasFrequency(1)) // true (because 2 appears once)
	fmt.Println(ft.hasFrequency(3)) // false (no number appears 3 times)
}
