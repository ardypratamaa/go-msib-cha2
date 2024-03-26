package main

import "fmt"

func main() {
    // Data Uji
    // Data 1
    //lumbaLumbaScores := []int{96, 108, 89}
    //koalaScores := []int{88, 91, 110}
    
    // Data Bonus 1
     lumbaLumbaScores := []int{97, 112, 101}
     koalaScores := []int{109, 95, 123}
    
    // Data Bonus 2
     //lumbaLumbaScores := []int{97, 112, 101}
     //koalaScores := []int{109, 95, 106}

    // menghitung skor rata-rata
    lumbaLumbaAverage := calculateAverage(lumbaLumbaScores)
    koalaAverage := calculateAverage(koalaScores)

    // menampilkan skor rata-rata
    fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", lumbaLumbaAverage)
    fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", koalaAverage)

    // cek skor minimum 100 untuk menentukan pemenang
    if lumbaLumbaAverage >= 100 && koalaAverage >= 100 {
        // Mcek hasil seri
        if lumbaLumbaAverage == koalaAverage {
            fmt.Println("Hasil seri! Tidak ada pemenang.")
        } else if lumbaLumbaAverage > koalaAverage {
            fmt.Println("Tim Lumba-lumba memenangkan trofi!")
        } else {
            fmt.Println("Tim Koala memenangkan trofi!")
        }
    } else {
        fmt.Println("Tidak ada pemenang. Salah satu atau kedua tim memiliki skor di bawah 100.")
    }
}

// menghitung rata-rata
func calculateAverage(scores []int) float64 {
    sum := 0
    for _, score := range scores {
        sum += score
    }
    return float64(sum) / float64(len(scores))
}
