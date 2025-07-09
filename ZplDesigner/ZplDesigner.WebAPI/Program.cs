using ZplDesigner.Library;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// Add CORS
builder.Services.AddCors(options =>
{
    options.AddPolicy("AllowAll", policy =>
    {
        policy.AllowAnyOrigin()
              .AllowAnyMethod()
              .AllowAnyHeader();
    });
});

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.UseCors("AllowAll");
app.UseAuthorization();
app.MapControllers();

// Add static files for web interface
app.UseStaticFiles();

// Serve the web interface
app.MapGet("/", async context =>
{
    context.Response.ContentType = "text/html";
    await context.Response.WriteAsync(@"
<!DOCTYPE html>
<html>
<head>
    <title>ZplDesigner Web API</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; background: #f5f5f5; }
        .container { max-width: 900px; margin: 0 auto; background: white; padding: 20px; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        textarea { width: 100%; height: 200px; margin: 10px 0; padding: 10px; border: 1px solid #ddd; border-radius: 4px; font-family: monospace; }
        button { padding: 12px 24px; background: #007bff; color: white; border: none; cursor: pointer; border-radius: 4px; font-size: 16px; }
        button:hover { background: #0056b3; }
        .result { margin-top: 20px; }
        img { max-width: 100%; border: 1px solid #ccc; border-radius: 4px; }
        .info { background: #e7f3ff; padding: 15px; border-radius: 4px; margin-bottom: 20px; }
    </style>
</head>
<body>
    <div class=""container"">
        <h1>🎨 ZplDesigner Web API</h1>
        <div class=""info"">
            <p><strong>✅ .NET Tabanlı Çözüm:</strong> ZPL kodunu PNG görseline dönüştürün. Tamamen yerel çalışır!</p>
            <p><strong>🔒 Veri Güvenliği:</strong> Tüm işlemler kendi sunucunuzda gerçekleşir.</p>
        </div>
        
        <h3>ZPL Kodu:</h3>
        <textarea id=""zplCode"" placeholder=""ZPL kodunuzu buraya yapıştırın..."">^XA
^FO50,50^ADN,36,20^FDTest Şirketi^FS
^FO50,100^ADN,24,12^FDTest Ürünü^FS
^FO50,150^ADN,18,10^FD99.99 TL^FS
^FO50,200^ADN,12,8^FD2025-07-09^FS
^XZ</textarea>
        
        <button onclick=""renderZPL()"">🖼️ PNG Oluştur</button>
        
        <div class=""result"" id=""result""></div>
    </div>

    <script>
        async function renderZPL() {
            const zplCode = document.getElementById('zplCode').value;
            const resultDiv = document.getElementById('result');
            
            resultDiv.innerHTML = '<div class=""info"">⏳ İşleniyor...</div>';
            
            try {
                const response = await fetch('/api/zpl/render', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        zpl: zplCode,
                        width: 101.6,
                        height: 203.2,
                        dpi: 203
                    })
                });
                
                const data = await response.json();
                
                if (data.success) {
                    resultDiv.innerHTML = '<div class=""info""><h3>✅ Başarılı!</h3><img src=""data:image/png;base64,' + data.image + '"" alt=""ZPL Render""></div>';
                } else {
                    resultDiv.innerHTML = '<div class=""info"" style=""background: #ffe6e6;""><h3>❌ Hata:</h3><p>' + data.message + '</p></div>';
                }
            } catch (error) {
                resultDiv.innerHTML = '<div class=""info"" style=""background: #ffe6e6;""><h3>❌ Hata:</h3><p>' + error.message + '</p></div>';
            }
        }
    </script>
</body>
</html>");
});

app.Run(); 