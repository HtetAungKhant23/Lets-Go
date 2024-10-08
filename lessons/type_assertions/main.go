package main

import "fmt"

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e *email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s *sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i *invalid) cost() float64 {
	return 0.0
}

func getExpenseExport(e expense) (string, float64) {
	switch v := e.(type) { // using type switch
	case *email:
		return v.toAddress, v.cost()
	case *sms:
		return v.toPhoneNumber, v.cost()
	case *invalid:
		return "", v.cost()
	default:
		return "", 0.0
	}
}

// ----------------- Exercise Message Formatter -------------------

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

type Bold struct {
	message string
}

type Code struct {
	message string
}

func (p *PlainText) Format() string {
	return fmt.Sprintf("The message is %s", p.message)
}

func (b *Bold) Format() string {
	return fmt.Sprintf("*%s*", b.message)
}

func (c *Code) Format() string {
	return fmt.Sprintf("`%s`", c.message)
}

func sendMessage(formatter Formatter) string {
	return formatter.Format()
}

// ----------------- Exercise Process Notification -------------------

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (d *directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

func (g *groupMessage) importance() int {
	return g.priorityLevel
}

func (s *systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch m := n.(type) {
	case *directMessage:
		return m.senderUsername, m.importance()
	case *groupMessage:
		return m.groupName, m.importance()
	case *systemAlert:
		return m.alertCode, m.importance()
	default:
		return "", 0
	}
}

func main() {
	//s := sms{true, "this is sms body", "09123456780"}
	//st, co := getExpenseExport(&s)
	//fmt.Println(st, co)

	//p := PlainText{"Wu shu war!"}
	//p := Bold{"Wu shu war!"}
	//p := Code{"Wu shu war!"}
	//fmt.Println(sendMessage(&p))

	dm := directMessage{"hak", "heheharharhuu", 17, false}
	name, score := processNotification(&dm)
	fmt.Println(name, score)

}
