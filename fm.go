package main

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"os"
)
type Net struct{
	name int
	CellList []int
}
type Cell struct{
	name int
	NetList []int
}
var (
	cellcount int
	degree float64
	cellmap map[int]*Cell
	netslice []Net
	bucketlist [][]int
	leftpart []*Cell
	rightpart []*Cell
)

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}

func LinesToCell(lines []string){
	var (
		netid, cellid int
		err error
	)
	for iter, line := range lines {
		netinfo := strings.Fields(line)
		if iter == 0 {
			degree, err = strconv.ParseFloat(netinfo[0], 64)
			if err != nil {fmt.Println(netinfo)}
			fmt.Println(degree)
		}
		for _, word := range netinfo {
			switch word[0] {
			case 'n':
				netid, err = strconv.Atoi(strings.Trim(word,"n"))
				if err != nil {fmt.Println(word)}
				var clist []int
				netslice = append(netslice, Net{netid,clist})
			case 'c':
				cellid, err = strconv.Atoi(strings.Trim(word,"c"))
				if err != nil {fmt.Println(word)}
				if curcell, ok := cellmap[cellid]; ok {
					curcell.NetList = append(curcell.NetList, netid)
					cellmap[cellid] = curcell
				} else {
					//Initial a cell
					nlist := []int{netid}
					cellmap[cellid] = &Cell{cellid, nlist}
					cellcount++
				}
				netslice[len(netslice)-1].CellList = append(netslice[len(netslice)-1].CellList, cellid)
			default :
			}
		}
	}
}
func InitialPartition(){
	for i, net := range netslice{
		if i < len(netslice)/2 {
			for _, cellid := range net.CellList{
				if float64(len(leftpart)+len(net.CellList)) < float64 (cellcount) * (1.0 - degree) {
					leftpart = append(leftpart, cellmap[cellid])
				} else {
					rightpart = append(rightpart, cellmap[cellid])
				}
			}
		} else {
			for _, cellid := range net.CellList{
				rightpart = append(rightpart, cellmap[cellid])
			}
		}
	}
	fmt.Println(cellcount)
	fmt.Println(len(leftpart), len(rightpart))
}

func main() {
	cellmap = make(map[int]*Cell)//Initial map
	// Loop over lines in file.
	lines := LinesInFile(`input_data/input_0.dat`)
	LinesToCell(lines)
	InitialPartition()
}
