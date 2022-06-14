package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records, err
}

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func getIp() (net.IP, error) {
	interfaces, e := net.Interfaces()
	if e != nil {
		return nil, errors.New("get Interfaces error")
	}
	for _, i := range interfaces {
		// the flags value maybe 'pointtopoint', it also has a ip, filter it.
		if !strings.Contains(i.Flags.String(), "broadcast") {
			continue
		}

		addresses, e := i.Addrs()
		if e != nil {
			return nil, errors.New("get addresses failed")
		}
		for _, add := range addresses {
			var ip net.IP
			switch t := add.(type) {
			case *net.IPNet:
				ip = t.IP
			case *net.IPAddr:
				ip = t.IP
			}
			// filter loopback ip
			if !ip.IsLoopback() && ip.To4() != nil {
				// return the first occurrence ip
				return ip, nil
			}
		}
	}

	return nil, errors.New("find none ip")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	var password string
	ips, _ := getIp()
	//read mac address/ip
	macAddr, err := getMacAddr()
	checkErr(err)
	//read csv
	records, err := readCsvFile("./pass.csv")
	checkErr((err))
	for _, itr := range records {
		if contains(macAddr, itr[0]) {
			fmt.Println(itr)
			password = itr[1]
		} else if ips.String() == itr[0] {
			fmt.Println(itr)
			password = itr[1]
		}
	}
	fmt.Println(password)
	db, err := sql.Open("sqlite3", "./ac.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT OR REPLACE INTO settings(name, value) values(?,?)")

	checkErr(err)
	res, err := stmt.Exec("password", password)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	if affect == 1 {
		fmt.Println("change pasword success")
	} else {
		fmt.Println("Error")
	}

	db.Close()
}
