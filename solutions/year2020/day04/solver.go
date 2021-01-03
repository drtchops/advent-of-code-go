package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

var validColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

type passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func (p *passport) Populated() bool {
	return p.BirthYear != "" &&
		p.IssueYear != "" &&
		p.ExpirationYear != "" &&
		p.Height != "" &&
		p.HairColor != "" &&
		p.EyeColor != "" &&
		p.PassportID != ""
}

func (p *passport) Valid() bool {
	if !p.Populated() {
		return false
	}

	if len(p.BirthYear) != 4 || len(p.IssueYear) != 4 || len(p.ExpirationYear) != 4 || len(p.Height) < 4 || len(p.HairColor) != 7 || len(p.PassportID) != 9 {
		return false
	}

	if y, err := strconv.ParseInt(p.BirthYear, 10, 64); err != nil || y < 1920 || y > 2002 {
		return false
	}

	if y, err := strconv.ParseInt(p.IssueYear, 10, 64); err != nil || y < 2010 || y > 2020 {
		return false
	}

	if y, err := strconv.ParseInt(p.ExpirationYear, 10, 64); err != nil || y < 2020 || y > 2030 {
		return false
	}

	heightUnits := p.Height[len(p.Height)-2:]
	heightValue := p.Height[:len(p.Height)-2]
	if heightUnits == "cm" {
		if cm, err := strconv.ParseInt(heightValue, 10, 64); err != nil || cm < 150 || cm > 193 {
			return false
		}
	} else if heightUnits == "in" {
		if in, err := strconv.ParseInt(heightValue, 10, 64); err != nil || in < 59 || in > 76 {
			return false
		}
	} else {
		return false
	}

	for i, c := range p.HairColor {
		if i == 0 && c != '#' {
			return false
		}
		if i > 0 && !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}

	colorValid := false
	for _, color := range validColors {
		if color == p.EyeColor {
			colorValid = true
			break
		}
	}
	if !colorValid {
		return false
	}

	if _, err := strconv.ParseInt(p.PassportID, 10, 64); err != nil {
		return false
	}

	return true
}

func parse(input string) []passport {
	passports := make([]passport, 0)
	p := passport{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			passports = append(passports, p)
			p = passport{}
			continue
		}

		pairs := strings.Split(line, " ")
		for _, pair := range pairs {
			keyValue := strings.Split(pair, ":")
			key := keyValue[0]
			value := keyValue[1]

			switch key {
			case "byr":
				p.BirthYear = value
			case "iyr":
				p.IssueYear = value
			case "eyr":
				p.ExpirationYear = value
			case "hgt":
				p.Height = value
			case "hcl":
				p.HairColor = value
			case "ecl":
				p.EyeColor = value
			case "pid":
				p.PassportID = value
			case "cid":
				p.CountryID = value
			}
		}
	}
	passports = append(passports, p)

	return passports
}

func (s *Solver) SolveA(input string) string {
	passports := parse(input)
	var valid int64
	for _, p := range passports {
		if p.Populated() {
			valid++
		}
	}
	return fmt.Sprint(valid)
}

func (s *Solver) SolveB(input string) string {
	passports := parse(input)
	var valid int64
	for _, p := range passports {
		if p.Valid() {
			valid++
		}
	}
	return fmt.Sprint(valid)
}
