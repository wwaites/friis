# Calculate path loss according to Friis

Either distance or received power must be given.

    -d=NaN: Distance in meters (no default)
    -f=5.8e+09: Frequency (default 5.8GHz)
    -gr=20: Receive antenna gain in decibels (default 20dB)
    -gt=20: Transmit antenna gain in decibels (default 20dB)
    -rx=NaN: Received power in decibel-milliwatts (no default)
    -tx=27: Transmitted power in decibel-milliwatts (default: 27dBm)

Example:

    % friis -d 1e4
    frequency:      5800000000.00 Hz
    wavelength:     0.0517 m
    distance:       10000.00 m
    transmit power: 27.00 dBm
    transmit gain:  20.00 dB
    receive gain:   20.00 dB
    path loss:      -127.71 dB
    receive power:  -60.71 dBm
