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
    "html/template"
)

type InputDetails struct {
    Abstract string
    Experience string
}

type ResumeText struct {
    Text string
}

type ExperienceText struct {
    Text string
}

func getRoot(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, nil)
}



func inputPage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("main.html"))

        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := InputDetails{
            Abstract:   r.FormValue("abstract"),
            Experience: r.FormValue("experience"),
        }
        a := details.Abstract
        e := details.Experience
        
        generateNew(w, r, a, e)
}

func generateNew(w http.ResponseWriter, r *http.Request, a string, e string) {
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
            pdf.MultiCell(0, 8, a, "", "J", false)
            pdf.CellFormat(85.00000000000000000, 0.00000000000000000, "", "", 0, "C", false, 0, "")
            pdf.MultiCell(30, 8, "EXPERIENCE", "", "C", false)
            pdf.MultiCell(0, 8, e, "", "J", false)

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

func generateExamplePDF(w http.ResponseWriter, r *http.Request) {
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
    http.HandleFunc("/example", generateExamplePDF)
    http.HandleFunc("/input", inputPage)

    err := http.ListenAndServe(":3333", nil)
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("server closed\n")
    } else if err != nil {
        fmt.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}