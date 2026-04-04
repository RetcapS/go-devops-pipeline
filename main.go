package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Go DevOps Hesap Makinesi</title>
    <style>
        body { font-family: sans-serif; text-align: center; margin-top: 100px; background-color: #f4f4f9; }
        .container { background: white; padding: 30px; border-radius: 10px; box-shadow: 0px 4px 10px rgba(0,0,0,0.1); display: inline-block; }
        input, select, button { padding: 10px; margin: 5px; font-size: 16px; }
        .result { margin-top: 20px; font-size: 24px; color: #28a745; }
    </style>
</head>
<body>
    <div class="container">
        <h2>🚀 Go Bulut Hesap Makinesi</h2>
        <form method="POST" action="/">
            <input type="number" step="any" name="num1" required placeholder="1. Sayı">
            <select name="operator">
                <option value="+">+</option>
                <option value="-">-</option>
                <option value="*">x</option>
                <option value="/">/</option>
            </select>
            <input type="number" step="any" name="num2" required placeholder="2. Sayı">
            <button type="submit">Hesapla</button>
        </form>
        {{if .HasResult}} <div class="result">Sonuç: {{.Result}}</div> {{end}}
        {{if .Error}} <div style="color:red">{{.Error}}</div> {{end}}
    </div>
</body>
</html>
`

type PageData struct {
	HasResult bool
	Result    float64
	Error     string
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	if r.Method == http.MethodPost {
		r.ParseForm()
		num1, _ := strconv.ParseFloat(r.FormValue("num1"), 64)
		num2, _ := strconv.ParseFloat(r.FormValue("num2"), 64)
		op := r.FormValue("operator")
		data.HasResult = true
		switch op {
		case "+": data.Result = num1 + num2
		case "-": data.Result = num1 - num2
		case "*": data.Result = num1 * num2
		case "/": 
			if num2 != 0 { data.Result = num1 / num2 } else { data.Error = "Sıfıra bölünmez!"; data.HasResult = false }
		}
	}
	t, _ := template.New("c").Parse(htmlTemplate)
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", calculateHandler)
	fmt.Println("Web sunucusu 8081 portunda basliyor...")
	http.ListenAndServe(":8081", nil)
}
