package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type Usr struct {
	id    int
	email string
	logs  []logItem
}

func (u Usr) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i, item.action, item.timestamp)
	}

	return out
}

func main() {
	/*go func() {
		time.Sleep(time.Second)
		fmt.Println("Competitive ninja with delay!")
	}()
	go fmt.Println("Competitive ninja!")
	go fmt.Println("Competitive ninja!")

	time.Sleep(2 * time.Second)

	fmt.Println("Not a competitive ninja!")*/

	/*	u := Usr{
			id:    1,
			email: "jakhninja.go",
			logs: []logItem{
				{actions[0], time.Now()},
				{actions[2], time.Now()},
				{actions[4], time.Now()},
				{actions[1], time.Now()},
				{actions[3], time.Now()},
				{actions[2], time.Now()},
			},
		}

		fmt.Println(u.getActivityInfo())*/
	rand.Seed(time.Now().Unix())

	users := generateUsers(1000)

	for _, user := range users {
		saveUserInfo(user)
	}

}

func saveUserInfo(user Usr) error {
	fmt.Printf("WRITING FILE USER ID: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(user.getActivityInfo())
	return err
}

func generateUsers(count int) []Usr {
	users := make([]Usr, count)

	for i := 0; i < count; i++ {
		users[i] = Usr{
			id:    i + 1,
			email: fmt.Sprintf("jakh%d@ninja.go"),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}

	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			timestamp: time.Now(),
			action:    actions[rand.Intn(len(actions)-1)],
		}
	}
	return logs
}
