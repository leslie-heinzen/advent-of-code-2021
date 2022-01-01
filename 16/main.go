package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type packet struct {
	versionId    int64
	typeId       int64
	lengthTypeId int
	value        int64
	packetLength int
	subPackets   []packet
}

func (p packet) val() int64 {
	var val int64

	switch p.typeId {
	case 0:
		for _, sp := range p.subPackets {
			val += sp.val()
		}
	case 1:
		val = 1
		for _, sp := range p.subPackets {
			val *= sp.val()
		}
	case 2:
		val = -1
		for _, sp := range p.subPackets {
			v := sp.val()
			if v < val || val == -1 {
				val = v
			}
		}
	case 3:
		val = -1
		for _, sp := range p.subPackets {
			v := sp.val()
			if v > val || val == -1 {
				val = v
			}
		}
	case 4:
		val = p.value
	case 5:
		a := p.subPackets[0].val()
		b := p.subPackets[1].val()
		if a > b {
			val = 1
		}
	case 6:
		a := p.subPackets[0].val()
		b := p.subPackets[1].val()
		if a < b {
			val = 1
		}
	case 7:
		a := p.subPackets[0].val()
		b := p.subPackets[1].val()
		if a == b {
			val = 1
		}
	}

	return val
}

func main() {
	input, _ := os.ReadFile("input.txt")

	var sb strings.Builder
	for _, r := range string(input) {
		num, _ := strconv.ParseInt(string(r), 16, 64)

		sb.WriteString(fmt.Sprintf("%04b", num))
	}

	binaryStr := sb.String()
	packets := readPackets(binaryStr)
	versionSum := getVersionSums(packets)

	res := packets[0].val()
	fmt.Print(versionSum)
	fmt.Print("\n")
	fmt.Print(res)
}

func readPackets(binaryStr string) []packet {
	v, t := binaryStr[:3], binaryStr[3:6]
	versionId, _ := strconv.ParseInt(v, 2, 64)
	typeId, _ := strconv.ParseInt(t, 2, 64)
	packets := []packet{}

	if typeId == 4 {
		// value packet
		packet := readValuePacket(versionId, typeId, binaryStr)
		packets = append(packets, packet)
	} else {
		// operator packet
		lHeaderPos := 7
		lengthTypeId, _ := strconv.Atoi(binaryStr[6:lHeaderPos])
		var packetLength int
		var spBinaryStr string
		var subs []packet

		if lengthTypeId == 0 { // num of bits in sub-packets
			from := lHeaderPos + 15
			packetLength = len(binaryStr[:from])
			spLen, _ := strconv.ParseInt(binaryStr[lHeaderPos:from], 2, 64)

			for spLen > 0 {
				spBinaryStr = binaryStr[from:]
				subPackets := readPackets(spBinaryStr)
				packetLengthSeen := getTotalPacketLength(subPackets)
				subs = append(subs, subPackets...)
				spLen -= int64(packetLengthSeen)
				from += packetLengthSeen
			}
		} else { // num of sub-packets
			spPos := lHeaderPos + 11
			packetLength = len(binaryStr[:spPos])
			spNum, _ := strconv.ParseInt(binaryStr[lHeaderPos:spPos], 2, 64)

			for spNum > 0 {
				spBinaryStr = binaryStr[spPos:]
				subPackets := readPackets(spBinaryStr)
				packetLengthSeen := getTotalPacketLength(subPackets)
				subs = append(subs, subPackets...)
				spNum--
				spPos += packetLengthSeen
			}
		}

		newPacket := packet{
			versionId:    versionId,
			typeId:       typeId,
			lengthTypeId: lengthTypeId,
			packetLength: packetLength,
			subPackets:   subs,
		}

		packets = append(packets, newPacket)
	}

	return packets
}

func readValuePacket(versionId int64, typeId int64, binaryStr string) packet {
	input := binaryStr[6:]
	values := ""
	var valueDec int64
	var bitsRead int
	for i := range input {
		if i > 0 && i%5 == 0 {
			values += input[i-4 : i]
			bitsRead += 5
			if input[i-5] == '0' {
				break
			}
		}
	}

	v, _ := strconv.ParseInt(values, 2, 64)
	valueDec = v

	return packet{
		versionId:    versionId,
		typeId:       typeId,
		value:        valueDec,
		packetLength: len(binaryStr[0:6]) + bitsRead,
	}
}

func getTotalPacketLength(packets []packet) int {
	sum := 0

	for _, packet := range packets {
		sum += packet.packetLength
		subSum := 0
		if len(packet.subPackets) > 0 {
			subSum += getTotalPacketLength(packet.subPackets)
		}
		sum += subSum
	}

	return sum
}

func getVersionSums(packets []packet) int64 {
	var total int64

	for _, p := range packets {
		total += p.versionId
		if len(p.subPackets) > 0 {
			subSum := getVersionSums(p.subPackets)
			total += subSum
		}
	}

	return total
}
