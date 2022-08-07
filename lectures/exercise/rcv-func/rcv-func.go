//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Value uint

type EnergyStat struct {
	energy    Value
	maxEnergy Value
}

type HealthStat struct {
	health    Value
	maxHealth Value
}

type Player struct {
	name       string
	healthStat HealthStat
	energyStat EnergyStat
}

func (player *Player) healAll() {
	fmt.Println("Heal all.")
	player.healthStat.health = player.healthStat.maxHealth
	player.energyStat.energy = player.energyStat.maxEnergy
}

func (player Player) printStats() {
	fmt.Println("Name ->", player.name)
	fmt.Println("Energy ->", player.energyStat.energy)
	fmt.Println("Health ->", player.healthStat.health)
}

func main() {
	player := Player{
		name:       "Bob",
		healthStat: HealthStat{10, 100},
		energyStat: EnergyStat{10, 100},
	}
	player.printStats()
	player.healAll()
	player.printStats()
}
