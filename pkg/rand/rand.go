package rand

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomPairs(members []string, maxMembers int) {
	var newList []string

	//teams := []rune{'A', 'B', 'C', 'D'}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	memberCount := len(members)
	shuffledIndexes := r.Perm(memberCount)
	for _, v := range shuffledIndexes {
		newList = append(newList, members[v])
	}
	for i := 0; i < memberCount; i += maxMembers {
		if i+2*maxMembers < memberCount {
			fmt.Println(newList[i : i+maxMembers])
		} else {
			fmt.Println(newList[i:])
			break
		}
	}

}
