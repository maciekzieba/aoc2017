package input

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/maciekzieba/aoc2017/day20/domain"
)

func LoadData() []domain.Particle {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var particles []domain.Particle

	i := 0
	for scanner.Scan() {
		particle := parseLine(scanner.Text(), i)
		particle.ID = i
		particles = append(particles, particle)
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading standard input:", err)
		os.Exit(0)
	}
	return particles
}

func parseLine(line string, ID int) domain.Particle {
	exp := regexp.MustCompile("\\<(.*?)\\>")
	match := exp.FindAllStringSubmatch(line, -1)
	if len(match) == 3 {
		position := parseCoordinate(match[0][1])
		velocity := parseCoordinate(match[1][1])
		acceleration := parseCoordinate(match[2][1])
		return domain.NewParticle(position, velocity, acceleration, ID)
	}
	fmt.Printf("file parse error line: %s\n", line)
	panic("panic")
}

func parseCoordinate(entry string) domain.Coordinate {
	parts := strings.Split(entry, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])

	return domain.Coordinate{
		X: x,
		Y: y,
		Z: z,
	}
}
