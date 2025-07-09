using Microsoft.AspNetCore.Mvc;
using ZplDesigner.Library;
using System.Text;
using BinaryKits.Zpl.Viewer;
using BinaryKits.Zpl.Viewer.Models;
using System.IO;
using System.Drawing.Imaging;
using System.Linq;

namespace ZplDesigner.WebAPI.Controllers;

[ApiController]
[Route("api/[controller]")]
public class ZplController : ControllerBase
{
    [HttpPost("render")]
    public async Task<IActionResult> RenderZpl([FromBody] ZplRenderRequest request)
    {
        try
        {
            IPrinterStorage printerStorage = new PrinterStorage();
            var drawer = new ZplElementDrawer(printerStorage);
            var analyzer = new ZplAnalyzer(printerStorage);
            var analyzeInfo = analyzer.Analyze(request.Zpl);
            if (analyzeInfo.LabelInfos.Count() == 0)
            {
                return BadRequest(new { success = false, message = "ZPL kodu işlenemedi veya boş." });
            }
            var imageData = drawer.Draw(analyzeInfo.LabelInfos[0].ZplElements);
            if (imageData == null || imageData.Length == 0)
            {
                return BadRequest(new { success = false, message = "ZPL render işlemi başarısız oldu." });
            }
            var base64Image = Convert.ToBase64String(imageData);
            return Ok(new {
                success = true,
                image = base64Image,
                message = "ZPL başarıyla PNG'ye dönüştürüldü."
            });
        }
        catch (Exception ex)
        {
            return BadRequest(new {
                success = false,
                message = $"Hata: {ex.Message}"
            });
        }
    }

    [HttpGet("test")]
    public IActionResult Test()
    {
        return Ok(new { 
            success = true, 
            message = "ZplDesigner Web API çalışıyor!",
            version = "1.0.0",
            framework = ".NET 8.0"
        });
    }
}

public class ZplRenderRequest
{
    public string Zpl { get; set; } = string.Empty;
    public double Width { get; set; } = 101.6;
    public double Height { get; set; } = 203.2;
    public int Dpi { get; set; } = 203;
} 