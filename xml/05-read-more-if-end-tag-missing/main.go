// Working but the peek function with a fixed nr is not optimal
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fh, err := os.Open("ardrone3.xml")
	if err != nil {
		log.Println("Error: failed opening file: ", err)
	}

	br := bufio.NewReader(fh)

	for {
		s, err := readBlock(br)
		if err != nil {
			log.Println("Error: ", err)
			break
		}
		fmt.Println(s)
	}

}

//readBlock will check for ending brackets,and if no such
// exist on the same line it will combine the line with
// the next one, and returned the combined result.
func readBlock(r *bufio.Reader) (string, error) {
	var lString string
	var l []byte
	var err error
	var tmpString string

	fmt.Println("--------------------------------------------------------------------")
	for {

		l, _, err = r.ReadLine()
		if err != nil {
			log.Println("Error: failed reading line: ", err)
			return "", err
		}
		//fmt.Printf("===== l =====:%v\n", string(l))
		//fmt.Printf("===== l =====:%v\n", l)

		peek, _ := r.Peek(4)
		peekString := strings.TrimSpace(string(peek))
		lString = strings.TrimSpace(string(l))
		//fmt.Printf("== lString ==:%v\n", lString)

		startOK := strings.HasPrefix(lString, "<")
		simpleEndOK := strings.HasSuffix(lString, ">")
		peekStartOK := strings.HasPrefix(string(peekString), "<")
		_ = fmt.Sprintln(startOK, simpleEndOK, peekStartOK)

		// In the IF below just specify when you want to run once, or break out.
		// Breaking out means read no more lines this run.
		//
		if (startOK && simpleEndOK) || //check if line contains "<" and ">", indicating complete tag line.
			peekStartOK { //check if next line contains "<", indicating new bracket on next line
			tmpString = fmt.Sprintf("%v %v", tmpString, lString)

			// If the finnished line don't have any brackets at all, we assume it is
			// a description, so we add a new tags called <description> & </description>.
			//
			tmpString = strings.TrimSpace(tmpString)
			startOK := strings.HasPrefix(tmpString, "<")
			simpleEndOK := strings.HasSuffix(tmpString, ">")
			if !startOK && !simpleEndOK {
				tmpString = fmt.Sprintf("<description>%v</description>", tmpString)
			}

			break
		}
		tmpString = fmt.Sprintf("%v %v", tmpString, lString)
	}
	fmt.Println("--------------------------------------------------------------------")
	return tmpString, nil
}
