package notification

// Notification system sends out a datadog event with provided text.
// Emails will be triggered through this to the provided email IDs in config
// This process already runs in a pool, so please do not fire a go-routine

// notification -> pool object for job dispatcher

import (
	"common/notification/datadog"
	"fmt"
	//"github.com/jabong/floRest/src/common/utils/logger"
	"os"
)

var notification pool
var dd datadog.Datadog

//var address string
var appName string
var hostname string

// InitNotifpool -> Initializes the notification pool object
func InitNotifpool() {

	APIKey := os.Getenv("APIKEY")
	AppKey := os.Getenv("APPKEY")
	dd = datadog.Init(APIKey, AppKey)
	//NotifAddr := os.Getenv("NotifAddr")
	//address = NotifAddr
	fmt.Println(APIKey)

	fmt.Println(AppKey)
	notification = newWorker("NotificationService", 4, 100)
	notification.startWorkers(notify)
	appName = "Lalit-my-app"
	hostname, _ = os.Hostname()

}

// SendNotification sends a datadog notification
func SendNotification(title, text string, tags []string, alertType string) {
	if alertType != datadog.INFO && alertType != datadog.WARNING && alertType != datadog.ERROR {
		alertType = datadog.INFO
	}

	//fmt.Println("In send Notification")
	tags = append(tags, appName)
	text = text + "\n\nHostname: " + hostname //+ "\nNotifying:\n" + address
	notif := make(map[string]interface{}, 0)
	notif["title"] = title
	notif["text"] = text
	notif["tags"] = tags
	notif["alertType"] = alertType
	//fmt.Println("before startjob send Notification")
	notification.startJob(notif)
	//fmt.Println("ENd send Notification")
}

// SendNotificationSimple sends a simple datadog notification
func SendNotificationSimple(title, text string) {
	tags := []string{appName}
	text = text + "\n\nHostname: " + hostname //+ "\nNotifying:\n" + address
	notif := make(map[string]interface{}, 0)
	notif["title"] = title
	notif["text"] = text
	notif["tags"] = tags
	notif["alertType"] = datadog.INFO
	notification.startJob(notif)
}

// worker fires the datadog event
func notify(data interface{}) {
	m := data.(map[string]interface{})
	title := m["title"].(string)
	text := m["text"].(string)
	tags := m["tags"].([]string)
	alertType := m["alertType"].(string)
	//logger.Info(fmt.Sprintf("Notification fired: %s", title))
	//fmt.Println("Error in notification.go notify")
	_ = dd.PostEvent(title, text, tags, alertType)
}

// pool struct for notification pool
type pool struct {
	name     string
	poolSize int
	channel  chan interface{}
}

// NewWorker -> Creates a new worker and returns a pool object
func newWorker(name string, poolSize int, queueMax int) pool {
	pool := pool{
		name:     name,
		poolSize: poolSize,
		channel:  make(chan interface{}, queueMax),
	}
	return pool
}

// StartWorkers -> Starts workers based on pool size
func (p pool) startWorkers(job func(interface{})) {
	for w := 0; w <= p.poolSize; w++ {
		go p.worker(job)
	}
}

//type job1 func(interface{})

// StartJob -> Use this function to start a new job, send data in interface
func (p pool) startJob(jobName interface{}) {
	//fmt.Println("In start Job")
	p.channel <- jobName
	close(p.channel)
	//p.worker(job1(p))
	//fmt.Println("In start Job")
}

// worker -> runs the given job with the interface in channel as argument
func (p pool) worker(job func(interface{})) {
	for j := range p.channel {
		func(name string) {
			defer recoverHandler(name)
			job(j)
		}(p.name)
	}
}
