package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

const c = 3e8
var frequency float64
var distance float64
var rxpower float64
var txpower float64
var txgain float64
var rxgain float64
var wavelength float64

func Usage() {
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "======================================\n")
	fmt.Fprintf(os.Stderr, "Calculate path loss according to Friis\n")
	fmt.Fprintf(os.Stderr, "======================================\n\n")
	fmt.Fprintf(os.Stderr, "Either distance or received power must be given.\n\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
}

func init() {
	flag.Usage = Usage
	flag.Float64Var(&frequency, "f", 5.8e9, "Frequency in Hz (default 5.8GHz)")
	flag.Float64Var(&distance, "d", math.NaN(), "Distance in meters (no default)")
	flag.Float64Var(&txpower, "tx", 27, "Transmitted power in decibel-milliwatts (default: 27dBm)")
	flag.Float64Var(&rxpower, "rx", math.NaN(), "Received power in decibel-milliwatts (no default)")
	flag.Float64Var(&txgain, "gt", 20, "Transmit antenna gain in decibels (default 20dB)")
	flag.Float64Var(&rxgain, "gr", 20, "Receive antenna gain in decibels (default 20dB)")
}

func main() {
	flag.Parse()


	if math.IsNaN(rxpower) && math.IsNaN(distance) {
		Usage()
		os.Exit(1)
	}

	wavelength = c / frequency
	if math.IsNaN(distance) {
		distance = wavelength / (4 * math.Pi * math.Pow(10, (rxpower - txpower - txgain - rxgain)/20))
	}
	loss := 20 * math.Log10(wavelength / (4 * math.Pi * distance))
	rxpower = txpower + txgain + rxgain + loss

	fmt.Printf("frequency:\t%.02f Hz\n", frequency)
	fmt.Printf("wavelength:\t%.04f m\n", wavelength)
	fmt.Printf("distance:\t%.02f m\n", distance)
	fmt.Printf("transmit power:\t%.02f dBm\n", txpower)
	fmt.Printf("transmit gain:\t%.02f dB\n", txgain)
	fmt.Printf("receive gain:\t%.02f dB\n", rxgain)
	fmt.Printf("path loss:\t%.02f dB\n", loss)
	fmt.Printf("receive power:\t%0.02f dBm\n", rxpower)
}
