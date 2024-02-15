package main

import "fmt"

func main() {
	var point int64 = 6

	switch point {
	case 8:
		fmt.Println("awesome")
	case 6:
		fmt.Println("good")
	default:
		fmt.Println("not bad")
	}
	// switch dengan banyak perbandingan
	switch point {
	case 8:
		fmt.Println("awesome")
	case 7, 6, 5, 4:
		fmt.Println("good")
	default:
		fmt.Println("not bad")

	}

	// switch dengan case tanda kurawal {}
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// swotch dengan gaya ala if else
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better")
		}
	}

	// penggunaan falltrhough di switch
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("Awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need learn vary hard")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better")
		}
	}

	// switch pada kondisi bersarang
	var nilai int64 = 10

	if nilai >= 7 {
		switch {
		case nilai == 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("awesome")
		}
	} else {
		if nilai == 6 {
			fmt.Println("not bad")
		} else if nilai == 5 {
			fmt.Println("keep try")
		} else {
			fmt.Println("you can do it")
			if nilai == 0 {
				fmt.Println("you must be keep learn")
			}
		}
	}
}
