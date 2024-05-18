package main

import (
    "bytes"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "github.com/krazybee/gofpdf"
)
func getRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got / request\n")
    io.WriteString(w, "This is a PDF generator!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got /hello request\n")
    io.WriteString(w, "Hello, HTTP!\n")
}

func generatePDF(w http.ResponseWriter, r *http.Request) {
    var b bytes.Buffer
    t, _ := os.Open("./static/landscape.png")
    pw := io.Writer(&b)
    pr := io.Reader(&b)

    pdf := gofpdf.New("P", "mm", "A4", "")
    fontsize := 12.0

    font := "Arial"

    pdf.SetFont("Arial", "", 6)

    pdf.AddPage()
    pdf.SetMargins(5, 5, 5)
    pdf.SetAutoPageBreak(true, 34)

    pdf.SetFont(font, "B", fontsize)

    var opt = gofpdf.ImageOptions{ImageType: "PNG",ReadDpi: true}

    pdf.RegisterImageOptionsReader("profile", opt, t)
    pdf.ImageOptions("profile", 70, 0, 70, 0, true, opt, 0, "")
    pdf.Ln(4)
    pdf.CellFormat(85.00000000000000000, 0.00000000000000000, "", "", 0, "C", false, 0, "")
    pdf.MultiCell(30, 8, "RESUME", "", "C", false)
    pdf.CellFormat(85.00000000000000000, 0.00000000000000000, "", "", 0, "C", false, 0, "")
    pdf.MultiCell(30, 8, "Experience", "", "C", false)

    // Show the pdf on localhost web page.
    err := pdf.Output(pw)
    if err != nil {
        fmt.Println(err)
        return
    }
    w.Header().Set("Content-Type", "application/pdf")
    resPDF, _ := ioutil.ReadAll(pr)
    w.Write(resPDF)
    fmt.Println("PDF served successfully")
}

func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/pdf", generatePDF)

    err := http.ListenAndServe(":3333", nil)
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("server closed\n")
    } else if err != nil {
        fmt.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}