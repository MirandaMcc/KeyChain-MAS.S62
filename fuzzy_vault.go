package main

import (
    "math"
    "fmt"
    "strings"
    "github.com/gonum/matrix/mat64"
)

var (
    degree int = 4
    t float64 = 10
    r float64 = 40
)

func GetCoefficients(word string) []float64 {

    word = strings.ToUpper(word)
    n := len(word) / degree
    if n < 1 {
        n = 1
    }
    var substring []string

    for i := 0; i < len(word); i += n {
        substring = append(substring, word[i:i+n])
    }
    var coeffs []float64

    for _, sub := range substring {
        num := 0.0
        for i, char := range sub {
            num += float64(int(char)) * math.Pow10(2 * i)
        }
        coeffs = append(coeffs, num)
    }
    return coeffs
}

func EvalAt(x float64, coeffs []float64) float64 {
    ret := 0.0

    for i, coefficient := range coeffs {
        ret += math.Pow(x, float64(i)) * coefficient
    }
    return ret
}

func Lock(secret string, template []float64) {
    var vault [][]float64
    coeffs := GetCoefficients(secret)

    maxY = math.Inf(-1)

    for _, point := range template {
        y := EvalAt(point, coeffs)
        if y > maxY {
            maxY = y
        }
        vault = append(vault, []float64{point, y})
    }

    maxX := MaxFloat64Slice(template)

    for i := t; i < r; i++ {
        xI = rand.Float64() * maxX * 1.1
        yI = rand.Float64() * maxY * 1.1
        vault = append(vault, []float64{xI, yI})
    }

    rand.Shuffle(len(vault),
        func(i, j int) {
            vault[i], vault[j] = vault[j], vault[i]
        }
    )

    return vault
}

func ApproxEqual(a , b, epsilon float64) bool {
    return math.Abs(a - b) < epsilon
}

func Unlock(template []float64, vault [][]float64) {
    func project(x float64) (float64, float64) {
        for _, point := range vault {
            if ApproxEqual(x, point[0], 0.001) {
                return x, point[1]
            }
        }
        return nil, nil
    }

    var tempX []float64
    var tempY []float64

    for _, val := range template {
        if x, y := project(val); x != nil {
            tempX = append(tempX, x)
            tempY = append(tempY, y)
        }
    }

    return polyfit(tempX, tempY)
}

func polyfit(X, Y []float64) []float64 {
    a := Vandermonde(X, degree)
    b := mat64.NewDense(len(Y), 1, Y)
    c := mat64.NewDense(degree+1, 1, nil)

    qr := new(mat64.QR)
    qr.Factorize(a)

    err := c.SolveQR(qr, false, b)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%.3f\n", mat64.Formatted(c))
    }
}

func Vandermonde(a []float64, degree int) *mat64.Dense {
    x := mat64.NewDense(len(a), degree+1, nil)
    for i := range a {
        for j, p := 0, 1.; j <= degree; j, p = j+1, p*a[i] {
            x.Set(i, j, p)
        }
    }
    return x
}

func Decode(coeffs float64) string {
    s := ""
    for _, c := range coeffs {
        num := int(math.Round(math.Pow(c, 3)))
        if num == 0 {
            continue
        }
        for num > 0 {
            s += strings.Lower(string(rune(num % 100)))
            num /= 100.0
        }
    }
    return s
}

func MaxFloat64Slice(s []float64) float64 {
    m := math.Inf(-1)
    for _, e := range s {
        if e > m {
            m = e
        }
    }
    return m
}

func main(){
    rand.Seed(42)
    fmt.Printf("You alread know what it is! \n")

    word := "hello"

    ret := GetCoefficients(word)

    fmt.Printf("%v\n", ret)

    fmt.Printf("%f\n", EvalAt(2.0, ret))
}