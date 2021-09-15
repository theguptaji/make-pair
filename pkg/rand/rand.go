package rand

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomPairs(members []string, maxMembers int) {
	var newList []string
	type team []string
	var teams []team
	//teams := []rune{'A', 'B', 'C', 'D'}
	rooms := []string{"Room B", "Room C", "Room D", "Room E", "Room F", "Room G"}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	memberCount := len(members)
	shuffledIndexes := r.Perm(memberCount)
	for _, v := range shuffledIndexes {
		newList = append(newList, members[v])
	}
	for i := 0; i < memberCount; i += maxMembers {
		if i+2*maxMembers <= memberCount {
			teams = append(teams, newList[i:i+maxMembers])
			//fmt.Println(newList[i : i+maxMembers])
		} else {
			teams = append(teams, newList[i:])
			//fmt.Println(newList[i:])
			break
		}
	}
	for i, v := range teams {
		fmt.Print(rooms[i] + " ")
		fmt.Println(v)
	}

}
