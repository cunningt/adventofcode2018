package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var immunegroups map[int]Group
var infectiongroups map[int]Group

type Attack struct {
	army           string
	uid            int
	initiative     int
	targetuid      int
	effectivepower int
	multiplier     int
}

type Group struct {
	uid            int
	attacked       bool
	effectivepower int
	units          int
	hitpoints      int
	weaknesses     map[string]bool
	immunities     map[string]bool
	attackdamage   int
	attacktype     string
	initiative     int
}

func parseGroup(point string) *Group {

	var g *Group
	g = nil
	re := regexp.MustCompile("([0-9]+) units each with ([0-9]+) hit points with an attack that does ([0-9]+) ([^\\s]+) damage at initiative ([0-9]+)")
	if re.MatchString(point) {
		g = new(Group)
		matches := re.FindAllStringSubmatch(point, -1)
		g.units, _ = strconv.Atoi(matches[0][1])
		g.hitpoints, _ = strconv.Atoi(matches[0][2])
		g.attackdamage, _ = strconv.Atoi(matches[0][3])
		g.attacktype = matches[0][4]
		g.initiative, _ = strconv.Atoi(matches[0][5])
		g.weaknesses = make(map[string]bool)
		g.immunities = make(map[string]bool)
		return g
	}

	re = regexp.MustCompile("([0-9]+) units each with ([0-9]+) hit points \\(([^\\)]+)\\) with an attack that does ([0-9]+) ([^\\s]+) damage at initiative ([0-9]+)")
	weakre := regexp.MustCompile("weak to ([^\\;]+)")
	immunere := regexp.MustCompile("immune to ([^\\;]+)")
	if re.MatchString(point) {
		g = new(Group)
		matches := re.FindAllStringSubmatch(point, -1)
		g.units, _ = strconv.Atoi(matches[0][1])
		g.hitpoints, _ = strconv.Atoi(matches[0][2])
		characteristics := matches[0][3]
		g.weaknesses = make(map[string]bool)
		g.immunities = make(map[string]bool)

		if weakre.MatchString(characteristics) {
			weakmatches := weakre.FindAllStringSubmatch(characteristics, -1)
			splitstring := weakmatches[0][1]
			splits := strings.Split(splitstring, ", ")
			for i := 0; i < len(splits); i++ {

				g.weaknesses[splits[i]] = true
			}
		}

		if immunere.MatchString(characteristics) {
			immunematches := immunere.FindAllStringSubmatch(characteristics, -1)
			splitstring := immunematches[0][1]
			splits := strings.Split(splitstring, ", ")
			for i := 0; i < len(splits); i++ {
				g.immunities[splits[i]] = true
			}
		}

		g.attackdamage, _ = strconv.Atoi(matches[0][4])
		g.attacktype = matches[0][5]
		g.initiative, _ = strconv.Atoi(matches[0][6])
		return g

	}
	return g
}

func printGroup(g *Group) {
	fmt.Printf("%d : Group %d contains %d units\n", g.effectivepower, g.uid, g.units)
}

type kv struct {
	Key    Group
	Value  int
	Value2 int
}

type av struct {
	Key    Attack
	Value  int
	Value2 int
}

func sortAttacks(a []Attack) []Attack {
	var ss []av
	for i := 0; i < len(a); i++ {
		v := a[i]
		ss = append(ss, av{v, v.initiative, v.effectivepower})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Value2 > ss[j].Value2
		} else {
			return ss[i].Value > ss[j].Value
		}
	})

	var copy []Attack
	for _, kv := range ss {
		copy = append(copy, kv.Key)
	}
	return copy
}

func sortEffectivePower(g map[int]Group) []Group {
	var ss []kv
	for _, v := range g {
		v.effectivepower = v.units * v.attackdamage
		v.attacked = false
		ss = append(ss, kv{v, v.effectivepower, v.initiative})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Value2 > ss[j].Value2
		} else {
			return ss[i].Value > ss[j].Value
		}
	})

	var copy []Group
	for _, kv := range ss {
		copy = append(copy, kv.Key)
	}
	return copy
}

func simulate(turns int) {
	for turn := 0; turn < turns; turn++ {
		immgroups := sortEffectivePower(immunegroups)
		infgroups := sortEffectivePower(infectiongroups)

		fmt.Println("Immune System:")
		total := 0
		for i := 0; i < len(immgroups); i++ {
			if immgroups[i].units > 0 {
				total += immgroups[i].units
				printGroup(&immgroups[i])
			}
		}
		fmt.Printf("Total armies : %d\n", total)

		total = 0
		fmt.Println("Infection System:")
		for i := 0; i < len(infgroups); i++ {
			if infgroups[i].units > 0 {
				total += infgroups[i].units
				printGroup(&infgroups[i])
			}
		}
		fmt.Printf("Total armies : %d\n", total)

		var attacks []Attack

		// Targeting
		for i := 0; i < len(immgroups); i++ {
			maxgroupid := -1
			maxgroupdamage := -1
			multiplier := -1
			maxindex := -1
			maxeffpower := -1

			if immgroups[i].units > 0 {
				for j := 0; j < len(infgroups); j++ {
					if infgroups[j].units == 0 {
						continue
					}
					if infgroups[j].attacked == false {
						attacktype := immgroups[i].attacktype
						attackmultiplier := 0
						attack := 0

						if infgroups[j].weaknesses[attacktype] {
							attack = immgroups[i].effectivepower * 2
							attackmultiplier = 2
						} else if infgroups[j].immunities[attacktype] {
							attack = 0
							attackmultiplier = 0
						} else {
							attack = immgroups[i].effectivepower
							attackmultiplier = 1
						}

						fmt.Printf("Immune System group %d would deal defending group %d %d damage\n", immgroups[i].uid, infgroups[j].uid, attack)

						if (attack != 0) && (attack > maxgroupdamage || (attack == maxgroupdamage && infgroups[j].effectivepower > maxeffpower)) {
							maxgroupid = infgroups[j].uid
							maxgroupdamage = attack
							maxindex = j
							multiplier = attackmultiplier
							maxeffpower = infgroups[j].effectivepower
						}
					}
				}

				if maxindex != -1 {
					// Add the attack to the list
					a := new(Attack)
					a.army = "immune"
					a.initiative = immgroups[i].initiative
					a.effectivepower = immgroups[i].effectivepower
					a.targetuid = maxgroupid
					a.uid = immgroups[i].uid
					a.multiplier = multiplier
					attacks = append(attacks, *a)
					infgroups[maxindex].attacked = true
				}
			}
		}

		for i := 0; i < len(infgroups); i++ {
			maxgroupid := -1
			maxgroupdamage := -1
			maxindex := -1
			multiplier := -1
			maxeffpower := -1

			if infgroups[i].units > 0 {
				for j := 0; j < len(immgroups); j++ {
					if immgroups[j].units == 0 {
						continue
					}
					if immgroups[j].attacked == false {
						attacktype := infgroups[i].attacktype
						attackmultiplier := 0
						attack := 0

						if immgroups[j].weaknesses[attacktype] {
							attack = infgroups[i].effectivepower * 2
							attackmultiplier = 2
						} else if immgroups[j].immunities[attacktype] {
							attack = 0
							attackmultiplier = 0
						} else {
							attack = infgroups[i].effectivepower
							attackmultiplier = 1
						}

						fmt.Printf("Infection System group %d would deal defending group %d %d damage\n", infgroups[i].uid, immgroups[j].uid, attack)

						if (attack != 0) && (attack > maxgroupdamage || (attack == maxgroupdamage && immgroups[j].effectivepower > maxeffpower)) {
							maxgroupid = immgroups[j].uid
							maxgroupdamage = attack
							maxindex = j
							maxeffpower = immgroups[j].effectivepower
							multiplier = attackmultiplier
						}
					}
				}

				// Add the attack to the list
				if maxindex != -1 {
					a := new(Attack)
					a.army = "infection"
					a.initiative = infgroups[i].initiative
					a.effectivepower = infgroups[i].effectivepower
					a.targetuid = maxgroupid
					a.multiplier = multiplier
					a.uid = infgroups[i].uid
					immgroups[maxindex].attacked = true
					attacks = append(attacks, *a)
				}
			}
		}

		// Attack
		fmt.Println()
		attacks = sortAttacks(attacks)
		for i := 0; i < len(attacks); i++ {
			if attacks[i].army == "immune" {
				g := immunegroups[attacks[i].uid]
				d := infectiongroups[attacks[i].targetuid]

				effectivepower := g.units * g.attackdamage * attacks[i].multiplier
				var unitskilled int = effectivepower / (d.hitpoints)
				//fmt.Printf("effective power %d hitpoints %d\n", effectivepower, d.hitpoints)
				if d.units > unitskilled {
					d.units = d.units - unitskilled
				} else {
					unitskilled = d.units
					d.units = 0
				}
				infectiongroups[attacks[i].targetuid] = d

				fmt.Printf("%d Immune System group %d attacks defending group %d, killing %d of %d units\n", attacks[i].initiative, attacks[i].uid, attacks[i].targetuid, unitskilled, d.units)

			} else if attacks[i].army == "infection" {
				g := infectiongroups[attacks[i].uid]
				d := immunegroups[attacks[i].targetuid]

				effectivepower := g.units * g.attackdamage * attacks[i].multiplier
				var unitskilled int = effectivepower / (d.hitpoints)
				//fmt.Printf("effective power %d hitpoints %d\n", effectivepower, d.hitpoints)
				if d.units > unitskilled {
					d.units = d.units - unitskilled
				} else {
					unitskilled = d.units
					d.units = 0
				}
				immunegroups[attacks[i].targetuid] = d

				fmt.Printf("%d Infection group %d attacks defending group %d, killing %d of %d units\n", attacks[i].initiative, attacks[i].uid, attacks[i].targetuid, unitskilled, d.units)

			} else {
				fmt.Printf("ERROR!  Invalid attack group %s\n", attacks[i])
			}
		}

		fmt.Println()

	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	immunegroups = make(map[int]Group)
	infectiongroups = make(map[int]Group)

	infectflag := false
	infcounter := 1
	immcounter := 1
	for scanner.Scan() {
		group := scanner.Text()

		if strings.Contains(group, "Infection:") {
			infectflag = true
		}

		g := parseGroup(group)
		if g != nil {
			if infectflag {
				g.uid = infcounter
				infectiongroups[infcounter] = *g
				infcounter++
			} else {
				g.uid = immcounter
				immunegroups[immcounter] = *g
				immcounter++
			}
		}
	}

	fmt.Printf("Infection group size : %d\n", len(infectiongroups))
	fmt.Printf("Immune group size : %d\n", len(immunegroups))
	// Part 1
	//simulate(2000)

	// Part 2 : give a boost
	for k, v := range immunegroups {
		v.attackdamage += 42
		immunegroups[k] = v
	}
	simulate(4000)

}
