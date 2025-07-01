package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "378282246310005"
	str2 := "5555555555554444"
	str3 := "371449635398431"
	str4 := "5105105105105100"
	str5 := "5610591081018250"
	str6 := "6011000990139424"

	re1 := regexp.MustCompile("^3[4][0-9]{13}$")
	re2 := regexp.MustCompile("^5[1-5][0-9]{14}$")

	fmt.Printf("Credit card : %v :%v\n", str1, re1.MatchString(str1))
	fmt.Printf("\n Credit card : %v :%v\n", str2, re1.MatchString(str2))
	fmt.Printf("\n Credit card : %v :%v\n", str3, re1.MatchString(str3))
	fmt.Printf("\n Credit card : %v :%v\n", str4, re1.MatchString(str4))
	fmt.Printf("\n Credit card : %v :%v\n", str5, re1.MatchString(str5))
	fmt.Printf("\n Credit card : %v :%v\n", str6, re1.MatchString(str6))

	//space betweeen amex to mastercard
	fmt.Println("\n\n\nfor Mastercard")
	fmt.Printf("\n Credit card : %v :%v\n", str1, re2.MatchString(str1))
	fmt.Printf("\n Credit card : %v :%v\n", str2, re2.MatchString(str2))
	fmt.Printf("\n Credit card : %v :%v\n", str3, re2.MatchString(str3))
	fmt.Printf("\n Credit card : %v :%v\n", str4, re2.MatchString(str4))
	fmt.Printf("\n Credit card : %v :%v\n", str5, re2.MatchString(str5))
	fmt.Printf("\n Credit card : %v :%v\n", str6, re2.MatchString(str6))
}
