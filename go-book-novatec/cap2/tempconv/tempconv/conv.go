package tempconv

// CToF converte uma temperatura Celsius para Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converte uma temperatura em Fahrenheit para Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converte uma temperatura em Celsius para Kelvin

func KToC(k Kelvin) Celsius { return Celsius(k - 273) }
