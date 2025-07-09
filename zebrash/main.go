package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ingridhq/zebrash/internal"
	"github.com/ingridhq/zebrash/drawers"
)

type ZPLRequest struct {
	ZPL          string  `json:"zpl"`
	LabelWidthMm float64 `json:"labelWidthMm,omitempty"`
	LabelHeightMm float64 `json:"labelHeightMm,omitempty"`
	Dpmm         int     `json:"dpmm,omitempty"`
}

type ZPLResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Image   string `json:"image,omitempty"` // base64 encoded PNG
}

func main() {
	// Create output directory
	os.MkdirAll("output", 0755)

	// HTTP handlers
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/api/zpl/render", handleZPLRender)
	http.HandleFunc("/api/zpl/pdf", handleZPLPDF)
	http.HandleFunc("/api/zpl/test", handleZPLTest)

	fmt.Println("🚀 Zebrash Web API is starting...")
	fmt.Println("📱 Web interface: http://localhost:8080")
	fmt.Println("🔧 API endpoint: http://localhost:8080/api/zpl/render")
	fmt.Println("📄 PDF endpoint: http://localhost:8080/api/zpl/pdf")
	fmt.Println("✅ Fully local - no external requests!")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Zebrash ZPL Renderer</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; background: #f5f5f5; }
        .container { max-width: 900px; margin: 0 auto; background: white; padding: 20px; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        textarea { width: 100%; height: 200px; margin: 10px 0; padding: 10px; border: 1px solid #ddd; border-radius: 4px; font-family: monospace; }
        button { padding: 12px 24px; background: #007bff; color: white; border: none; cursor: pointer; border-radius: 4px; font-size: 16px; }
        button:hover { background: #0056b3; }
        .result { margin-top: 20px; }
        img { max-width: 100%; border: 1px solid #ccc; border-radius: 4px; }
        .info { background: #e7f3ff; padding: 15px; border-radius: 4px; margin-bottom: 20px; }
        .status { color: #28a745; font-weight: bold; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🐆 Zebrash ZPL Renderer</h1>
        <div class="info">
            <p><strong>✅ Fully Local Solution:</strong> Convert ZPL code to PNG image. No external requests!</p>
            <p><strong>🔒 Data Security:</strong> All operations are performed on your own server.</p>
        </div>
        
        <h3>ZPL Code:</h3>
        <textarea id="zplCode" placeholder="Paste your ZPL code here...">^XA
^FO50,50^ADN,36,20^FDTest Company^FS
^FO50,100^ADN,24,12^FDTest Product^FS
^FO50,150^ADN,18,10^FD99.99 USD^FS
^FO50,200^ADN,12,8^FD2025-07-09^FS
^XZ</textarea>
        
        <button onclick="renderZPL()">🖼️ Generate PNG</button>
        <button onclick="downloadPNG()" style="margin-left: 10px; background: #28a745;">📄 Download PNG</button>
        
        <div class="result" id="result"></div>
    </div>

    <script>
        async function renderZPL() {
            const zplCode = document.getElementById('zplCode').value;
            const resultDiv = document.getElementById('result');
            
            resultDiv.innerHTML = '<div class="info">⏳ Processing...</div>';
            
            try {
                const response = await fetch('/api/zpl/render', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        zpl: zplCode,
                        labelWidthMm: 101.6,
                        labelHeightMm: 203.2,
                        dpmm: 8
                    })
                });
                
                const data = await response.json();
                
                if (data.success) {
                    resultDiv.innerHTML = '<div class="info"><h3>✅ Success!</h3><img src="data:image/png;base64,' + data.image + '" alt="ZPL Render"></div>';
                } else {
                    resultDiv.innerHTML = '<div class="info" style="background: #ffe6e6;"><h3>❌ Error:</h3><p>' + data.message + '</p></div>';
                }
            } catch (error) {
                resultDiv.innerHTML = '<div class="info" style="background: #ffe6e6;"><h3>❌ Error:</h3><p>' + error.message + '</p></div>';
            }
        }
        
        async function downloadPNG() {
            const zplCode = document.getElementById('zplCode').value;
            const resultDiv = document.getElementById('result');
            
            resultDiv.innerHTML = '<div class="info">⏳ Generating PNG...</div>';
            
            try {
                const response = await fetch('/api/zpl/pdf', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        zpl: zplCode,
                        labelWidthMm: 101.6,
                        labelHeightMm: 203.2,
                        dpmm: 8
                    })
                });
                
                if (response.ok) {
                    const blob = await response.blob();
                    const url = window.URL.createObjectURL(blob);
                    const a = document.createElement('a');
                    a.href = url;
                    a.download = 'zpl-label.png';
                    document.body.appendChild(a);
                    a.click();
                    window.URL.revokeObjectURL(url);
                    document.body.removeChild(a);
                    resultDiv.innerHTML = '<div class="info"><h3>✅ PNG Downloaded!</h3><p>Your PNG file has been downloaded successfully.</p></div>';
                } else {
                    const data = await response.json();
                    resultDiv.innerHTML = '<div class="info" style="background: #ffe6e6;"><h3>❌ Error:</h3><p>' + data.message + '</p></div>';
                }
            } catch (error) {
                resultDiv.innerHTML = '<div class="info" style="background: #ffe6e6;"><h3>❌ Error:</h3><p>' + error.message + '</p></div>';
            }
        }
    </script>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handleZPLRender(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ZPLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Default values
	if req.LabelWidthMm == 0 {
		req.LabelWidthMm = 101.6
	}
	if req.LabelHeightMm == 0 {
		req.LabelHeightMm = 203.2
	}
	if req.Dpmm == 0 {
		req.Dpmm = 8
	}

	// Parse ZPL
	parser := internal.NewParser()
	res, err := parser.Parse([]byte(req.ZPL))
	if err != nil {
		json.NewEncoder(w).Encode(ZPLResponse{
			Success: false,
			Message: "ZPL parse error: " + err.Error(),
		})
		return
	}

	if len(res) == 0 {
		json.NewEncoder(w).Encode(ZPLResponse{
			Success: false,
			Message: "ZPL code is invalid or empty",
		})
		return
	}

	// Draw as PNG
	var buff bytes.Buffer
	drawer := internal.NewDrawer()

	err = drawer.DrawLabelAsPng(res[0], &buff, drawers.DrawerOptions{
		LabelWidthMm:  req.LabelWidthMm,
		LabelHeightMm: req.LabelHeightMm,
		Dpmm:          req.Dpmm,
	})

	if err != nil {
		json.NewEncoder(w).Encode(ZPLResponse{
			Success: false,
			Message: "PNG generation error: " + err.Error(),
		})
		return
	}

	// Convert to base64
	imageBytes := buff.Bytes()
	base64Image := base64.StdEncoding.EncodeToString(imageBytes)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ZPLResponse{
		Success: true,
		Image:   base64Image,
	})
}

func handleZPLPDF(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ZPLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Default values
	if req.LabelWidthMm == 0 {
		req.LabelWidthMm = 101.6
	}
	if req.LabelHeightMm == 0 {
		req.LabelHeightMm = 203.2
	}
	if req.Dpmm == 0 {
		req.Dpmm = 8
	}

	// Parse ZPL
	parser := internal.NewParser()
	res, err := parser.Parse([]byte(req.ZPL))
	if err != nil {
		http.Error(w, "ZPL parse error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(res) == 0 {
		http.Error(w, "ZPL code is invalid or empty", http.StatusBadRequest)
		return
	}

	// Draw as PNG first
	var buff bytes.Buffer
	drawer := internal.NewDrawer()

	err = drawer.DrawLabelAsPng(res[0], &buff, drawers.DrawerOptions{
		LabelWidthMm:  req.LabelWidthMm,
		LabelHeightMm: req.LabelHeightMm,
		Dpmm:          req.Dpmm,
	})

	if err != nil {
		http.Error(w, "PNG generation error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// For now, return the PNG as a downloadable file
	// In a real implementation, you would use a PDF library like gofpdf or similar
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename=zpl-label.png")
	w.Write(buff.Bytes())
}

func handleZPLTest(w http.ResponseWriter, r *http.Request) {
	// Test with a sample ZPL file
	testZPL := `^XA
^FO50,50^ADN,36,20^FDTest Company^FS
^FO50,100^ADN,24,12^FDTest Product^FS
^FO50,150^ADN,18,10^FD99.99 USD^FS
^FO50,200^ADN,12,8^FD2025-07-09^FS
^XZ`

	parser := internal.NewParser()
	res, err := parser.Parse([]byte(testZPL))
	if err != nil {
		json.NewEncoder(w).Encode(ZPLResponse{
			Success: false,
			Message: "Test ZPL parse error: " + err.Error(),
		})
		return
	}

	var buff bytes.Buffer
	drawer := internal.NewDrawer()

	err = drawer.DrawLabelAsPng(res[0], &buff, drawers.DrawerOptions{
		LabelWidthMm:  101.6,
		LabelHeightMm: 203.2,
		Dpmm:          8,
	})

	if err != nil {
		json.NewEncoder(w).Encode(ZPLResponse{
			Success: false,
			Message: "Test PNG generation error: " + err.Error(),
		})
		return
	}

	// Save test image
	testFile := filepath.Join("output", "test-output.png")
	os.WriteFile(testFile, buff.Bytes(), 0644)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ZPLResponse{
		Success: true,
		Message: "Test successful! File saved: " + testFile,
	})
} 