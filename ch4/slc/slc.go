package main

import "fmt"

func main() {

	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "march",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	fmt.Printf("cap(months):%d\n", cap(months))
	fmt.Printf("len(months):%d\n", len(months))

	fmt.Printf("cap(summer):%d\n", cap(summer))
	fmt.Printf("len(summer):%d\n", len(summer))

	fmt.Printf("cap(Q2):%d\n", cap(Q2))
	fmt.Printf("len(Q2):%d\n", len(Q2))

	endlessSummer := summer[:4]
	fmt.Println(endlessSummer)



}
