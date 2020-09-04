package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"tiny-ip-database/ipdb"
)

func update_test(ipdb *ipdb.IpDB) {
	ip := "1.26.15.40"
	fmt.Println(ipdb.Find(ip))
	fmt.Println(ipdb.Update(ip))
	fmt.Println(ipdb.Find(ip))
	fmt.Println(ipdb.Dump("b.csv"))
}

func delete_test(ipdb *ipdb.IpDB) {
	ip := "223.247.95.200"
	fmt.Println(ipdb.Find(ip))
	fmt.Println(ipdb.Delete(ip))
	fmt.Println(ipdb.Find(ip))
	fmt.Println(ipdb.Dump("b.csv"))
}

func match_test(ipdb *ipdb.IpDB) {
	b, _ := ioutil.ReadFile("/tmp/ips_to_match.json")
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
	ioutil.WriteFile("/tmp/match_ips.json", b, 0644)
}

func main() {
	ipdb := new(ipdb.IpDB)
	ipdb.Load("a.csv")
	fmt.Println(ipdb.Count())
}
