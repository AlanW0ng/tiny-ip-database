package ipdb

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func (ipdb *IpDB) Load(dbPath string) (err error) {
	var line []string
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return err
	}
	reader := csv.NewReader(strings.NewReader(string(b)))

	for {
		line, err = reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		start, end := line[0], line[1]
		ipSeg := IpSegment{
			Start: IPToUInt32(start),
			End:   IPToUInt32(end),
		}
		ipdb.IpSegments = append(ipdb.IpSegments, ipSeg)
	}

	sort.Sort(ipdb.IpSegments)
	return nil
}

func (ipdb *IpDB) Dump(dbPath string) (err error) {
	f, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, ipSeg := range ipdb.IpSegments {
		_, err = f.WriteString(fmt.Sprintf("%s,%s\n", UInt32ToIP(ipSeg.Start), UInt32ToIP(ipSeg.End)))
		if err != nil {
			return err
		}
	}
	return nil
}

func (ipdb *IpDB) Count() uint32 {
	if ipdb.count == 0 {
		total := uint32(0)
		for _, ipSeg := range ipdb.IpSegments {
			total += ipSeg.End - ipSeg.Start + 1
		}
		ipdb.count = total
	}
	return ipdb.count
}

func (ipdb *IpDB) Find(ip string) bool {
	intIp := IPToUInt32(ip)
	_, exists := ipdb.IpSegments.Find(intIp)
	return exists
}

func (ipdb *IpDB) Update(ip string) error {
	return ipdb.IpSegments.Update(IPToUInt32(ip))
}

func (ipdb *IpDB) Delete(ip string) error {
	return ipdb.IpSegments.Delete(IPToUInt32(ip))
}
