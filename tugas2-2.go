package main

import (
	"fmt"
)

func main() {
    // Data Uji
    // Data 1
    markWeight1 := 78.0  // kg
    markHeight1 := 1.69  // m
    johnWeight1 := 92.0  // kg
    johnHeight1 := 1.95  // m

    // Data 2
    // markWeight2 := 95.0  // kg
    // markHeight2 := 1.88  // m
    // johnWeight2 := 85.0   // kg
    // johnHeight2 := 1.76   // m

    // menghitung BMI (Body Mass Index)
    markBMI1 := calculateBMI(markWeight1, markHeight1)
    johnBMI1 := calculateBMI(johnWeight1, johnHeight1)

    // mencetak BMI Mark dan John
    fmt.Printf("BMI Mark: %.2f\n", markBMI1)
    fmt.Printf("BMI John: %.2f\n", johnBMI1)

    // mengecek apakah Mark memiliki BMI lebih tinggi dari John
    markHigherBMI := markBMI1 > johnBMI1

    // menampilkan hasil
    if markHigherBMI {
        fmt.Println("Mark memiliki BMI lebih tinggi dari John.")
    } else {
        fmt.Println("John memiliki BMI lebih tinggi dari Mark atau sama.")
    }
}

// menghitung BMI
func calculateBMI(weight, height float64) float64 {
    return weight / (height * height)
}
