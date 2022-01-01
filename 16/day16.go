package day16

import (
	"2021-aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	// RunA()
	RunB()
}

func RunA() {
	read()
}

func RunB() {
	packets := read()

	for i, packet := range packets {
		fmt.Println("i", i, "computed", compute(packet))
	}
}

func version(packet *Packet) (total int64) {
	if packet.subpackets != nil {
		for _, s := range packet.subpackets {
			total += version(s)
		}
	}

	return packet.version + total
}

func p(packet *Packet, indent int) (str string) {
	str = fmt.Sprintln("version", packet.version, "typeID", packet.typeID, "lt", packet.lengthTypeID, "literal", packet.literal)

	for _, subpacket := range packet.subpackets {
		for i := 0; i < indent; i++ {
			str += "-"
		}
		str += fmt.Sprintf("> %v", p(subpacket, indent+2))
	}

	return
}

type Packet struct {
	raw          string
	version      int64
	typeID       int64
	literal      int64
	lengthTypeID bool
	subpackets   []*Packet
}

func read() (packets []*Packet) {
	var err error

	data, err := utils.ReadFile("./16/data.txt")
	handle(err)

	lines := strings.Split(data, "\n")
	packets = make([]*Packet, len(lines))

	for i, line := range lines {
		// panic("unknown type id")
		fmt.Println("--- i", i, " ---")
		packets[i], _ = parse(toBits(line))
		// fmt.Print(p(packets[i], 2))
		fmt.Println("version total", version(packets[i]))
		fmt.Print("\n")
	}

	return
}

func compute(packet *Packet) (total int) {
	switch packet.typeID {
	case 0:
		for _, subpacket := range packet.subpackets {
			total += compute(subpacket)
		}
	case 1:
		total = 1
		for _, subpacket := range packet.subpackets {
			total *= compute(subpacket)
		}
	case 2:
		arr := make([]int, len(packet.subpackets))
		for i, subpacket := range packet.subpackets {
			arr[i] = int(compute(subpacket))
		}
		total, _ = utils.Min(arr)
	case 3:
		arr := make([]int, len(packet.subpackets))
		for i, subpacket := range packet.subpackets {
			arr[i] = int(compute(subpacket))
		}
		total, _ = utils.Max(arr)
	case 4:
		total = int(packet.literal)
	case 5:
		if compute(packet.subpackets[0]) > compute(packet.subpackets[1]) {
			total = 1
		} else {
			total = 0
		}
	case 6:
		if compute(packet.subpackets[0]) < compute(packet.subpackets[1]) {
			total = 1
		} else {
			total = 0
		}
	case 7:
		if compute(packet.subpackets[0]) == compute(packet.subpackets[1]) {
			total = 1
		} else {
			total = 0
		}
	default:
		panic("invalid type id")
	}

	return
}

func parse(bits string) (p_packet *Packet, remainingBits string) {
	var err error
	packet := initPacket(bits)

	switch packet.typeID {
	case 4:
		packet.literal, remainingBits = parseLiteral(packet.raw[6:])
	default:
		packet.lengthTypeID, err = strconv.ParseBool(packet.raw[6:7])
		handle(err)

		if packet.lengthTypeID {
			numSubpackets, err := strconv.ParseInt(packet.raw[7:18], 2, 64)
			handle(err)

			remainingBits = packet.raw[18:]
			packet.subpackets = make([]*Packet, numSubpackets)
			for i := 0; i < int(numSubpackets); i++ {
				fmt.Println(i, numSubpackets, len(remainingBits))
				packet.subpackets[i], remainingBits = parse(remainingBits)
			}
		} else {
			idx, err := strconv.ParseInt(packet.raw[7:22], 2, 64)
			handle(err)

			bits = packet.raw[22 : 22+idx]
			remainingBits = packet.raw[22+idx:]
			packet.subpackets = make([]*Packet, 0)
			for len(bits) > 0 {
				var subpacket *Packet
				subpacket, bits = parse(bits)
				packet.subpackets = append(packet.subpackets, subpacket)
			}
		}

	}

	p_packet = &packet
	return
}

func parseLiteral(s string) (literal int64, remainder string) {
	var k string
	var i int
	for i = 5; ; i = i + 5 {
		k += s[i-4 : i]

		// fmt.Println(k)

		if s[i-5:i-4] == "0" {
			break
		}
	}

	remainder = s[i:]
	literal, err := strconv.ParseInt(k, 2, 64)
	handle(err)

	return
}

func toBits(line string) (output string) {
	// fmt.Println(line, line[0:4])
	var s string
	for len(line) > 0 {
		len, _ := utils.Min([]int{len(line), 4})
		s, line = line[0:len], line[len:]
		// fmt.Println(s, line)
		i, err := strconv.ParseUint(s, 16, 64)
		handle(err)
		output += fmt.Sprintf("%0"+fmt.Sprint(len*4)+"b", i)
		// fmt.Println(output)
	}

	return
}

func initPacket(bits string) Packet {
	packet := Packet{raw: bits}
	v, err := strconv.ParseInt(packet.raw[0:3], 2, 64)
	handle(err)
	packet.version = v

	// s := strings.Split(packet.raw[3:6], "")
	// for i, j := 0, 2; i < j; i, j = i+1, j-1 {
	// 	s[i], s[j] = s[j], s[i]
	// }
	packet.typeID, err = strconv.ParseInt(packet.raw[3:6], 2, 64)
	handle(err)

	return packet
}
