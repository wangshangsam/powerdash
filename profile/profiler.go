// package profile provides profilers that can extract power and energy
// related information of the host machine.
package profile

// Item is the type of attributes (power, energy, etc.) to be profiled.
type Item int

const (
	// PowerNow represents the current machine power (µW) seen by the
	// power supply.
	PowerNow = iota
	// EnergyNow represents the current energy (µWh) remaining.
	EnergyNow
	// VoltageNow represents the current voltage (µV) seen by the power
	// supply.
	VoltageNow
)

// Profiler defines the Probe() method that any specific Profiler
// implementation should implement.
type Profiler interface {
	Probe(items []Item) map[Item]uint64
}
