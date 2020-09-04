package main

import (
	"encoding/json"
	"github.com/AlanW0ng/tiny-ip-database/ipdb"
	"io/ioutil"
	"log"
)

func update_test(ipdb *ipdb.IpDB, ip string) {
	log.Println(ipdb.Find(ip))
	log.Println(ipdb.Update(ip))
	log.Println(ipdb.Find(ip))
	//log.Println(ipdb.Dump("tmp.csv"))
}

func delete_test(ipdb *ipdb.IpDB, ip string) {
	log.Println(ipdb.Find(ip))
	log.Println(ipdb.Delete(ip))
	log.Println(ipdb.Find(ip))
	//log.Println(ipdb.Dump("tmp.csv"))
}

func match_test(ipdb *ipdb.IpDB, input, output string) {
	b, _ := ioutil.ReadFile(input)
	var ips []string
	err := json.Unmarshal(b, &ips)
	if err != nil {
		panic(err)
	}
	var bad_ips []string
	for _, ip := range ips {
		if ipdb.Find(ip) {
			bad_ips = append(bad_ips, ip)
		}
	}
	b, _ = json.Marshal(bad_ips)
	ioutil.WriteFile(output, b, 0644)
}

func main() {
	db1, db2 := new(ipdb.IpDB), new(ipdb.IpDB)
	ip := "1.1.1.1"
	dbFile := "tmp.csv"
	log.Println("find: ", db1.Find(ip))
	log.Println("update err: ", db1.Update(ip))
	log.Println("find after updating: ", db1.Find(ip))
	log.Println("dump err: ", db1.Dump(dbFile))
	log.Println("load err: ", db2.Load(dbFile))
	log.Println("find after loading: ", db2.Find(ip))
	log.Println("db2 count: ", db2.Count())
}
