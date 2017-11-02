// Sample golang mongo
package main
import (
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Counters struct {
	Id int `json: "id"`
	Count int `json: "count"`
	Time time.Time `json: "time"`
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	var num = 0
	sess, _ := mgo.Dial("database")
	defer sess.Close()
	sess.SetMode(mgo.Monotonic, true)
	conn := sess.DB("gomong").C("counters")
	cnow := Counters{}
	err := conn.Find(bson.M{"id": 1}).One(&cnow)
	if err != nil {
		conn.Insert(&Counters{Id: 1, Count: num, Time: time.Now()})
	}else{
		num = cnow.Count
		num++
		// Update
		col := bson.M{"id": 1}
		chg := bson.M{"$set": bson.M{"count": num, "time": time.Now()}}
		conn.Update(col, chg)
	}
	fmt.Fprintf(w, "Halaman telah dibuka sebanyak %+v kali.", num)
}

func main() {
	http.HandleFunc("/", reqHandler)
	http.ListenAndServe(":8080", nil)
}