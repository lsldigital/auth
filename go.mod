module go.lsl.digital/lardwaz/auth

go 1.12

require (
	github.com/fatih/structs v1.1.0
	github.com/sirupsen/logrus v1.4.2
	go.lsl.digital/lardwaz/sdk v0.1.0
	go.lsl.digital/passport/sdk v0.1.1

	golang.org/x/crypto v0.0.0-20190829043050-9756ffdc2472 // indirect
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 // indirect
	golang.org/x/text v0.3.0 // indirect
)

replace go.lsl.digital/lardwaz/sdk => ../lardwaz-sdk
