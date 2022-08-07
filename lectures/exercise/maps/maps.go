//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// * Create a function to print server status, including:
//   - Number of servers
//   - Number of servers for each status (Online, Offline, Maintenance, Retired)
func printServerStatus(servers map[string]int) {
	fmt.Println("------------------------------------")
	fmt.Println("Number of servers:", len(servers))
	status := make(map[int]int)
	for _, value := range servers {
		switch value {
		case Online:
			status[Online] += 1
		case Offline:
			status[Offline] += 1
		case Maintenance:
			status[Maintenance] += 1
		case Retired:
			status[Retired] += 1
		default:
			panic("Unhandled server status")
		}
	}
	fmt.Println("Total online:", status[Online])
	fmt.Println("Total offline:", status[Offline])
	fmt.Println("Total maintenance:", status[Maintenance])
	fmt.Println("Total retired:", status[Retired])
	fmt.Println("------------------------------------")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	//* Store the existing slice of servers in a map
	//* Default all the servers to `Online`
	serverStatus := make(map[string]int)
	for _, value := range servers {
		serverStatus[value] = Online
	}
	//* Perform the following status changes and display server info:
	//  - display server info
	printServerStatus(serverStatus)
	//  - change `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired
	//  - change `aiur` to `Offline`
	serverStatus["aiur"] = Offline
	//  - display server info
	printServerStatus(serverStatus)
	//  - change all servers to `Maintenance`
	for key := range serverStatus {
		serverStatus[key] = Maintenance
	}
	//  - display server info
	printServerStatus(serverStatus)
}
