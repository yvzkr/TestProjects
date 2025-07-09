# ZPL Barcode Test Project

This project is a comprehensive test environment that includes various tools for testing and developing ZPL (Zebra Programming Language) barcodes.

## 📋 Project Contents

This project includes the following components:

- **zpl-tool**: Modern web interface for ZPL barcode creation and testing tool
- **zebrash**: High-performance Go-based ZPL processing library
- **ZplDesigner**: C#/.NET based ZPL processing and design tool
- **Barcode Test Tools**: HTML-based camera barcode testing tool
- **Test Data**: Test files for various barcode formats

## 🚀 Installation and Setup

### Requirements

- Docker and Docker Compose
- Modern web browser
- Camera (for barcode testing)

### 1. ZPL-Tool Installation

ZPL-Tool provides a modern web interface similar to Labelary:

```bash
cd zpl-tool
docker-compose up -d
```

**Access:**

- Web Interface: http://localhost:3000
- Binarykits API: http://localhost:7763

### 2. Zebrash Installation

Go-based high-performance ZPL processing:

```bash
cd zebrash
docker-compose up -d
```

**Access:**

- API: http://localhost:8080

### 3. ZplDesigner Installation

C#/.NET based ZPL design tool:

```bash
cd ZplDesigner
docker-compose -f docker-compose.zpldesigner.yml up -d
```

**Access:**

- API: http://localhost:8081

## 🧪 Barcode Testing Methods

### 1. Camera Testing (Recommended)

1. Open `barcode-test.html` file in your browser
2. Click "Start Camera" button
3. Show the barcode created in ZPL-tool to the camera
4. When the barcode is read, the data will be displayed

### 2. Manual Testing

1. Copy the data from Variables section in ZPL-tool
2. Paste it into the "Manual Barcode Test" section in `barcode-test.html`
3. Click "Test" button

### 3. Testing with Phone Applications

**Free barcode readers for Android/iOS:**

- Google Lens (Android/iOS)
- QR & Barcode Scanner (Android)
- QR Code Reader (iOS)
- Barcode Scanner (Android)

## 📝 Usage Examples

### Creating Barcodes with ZPL-Tool

1. Go to http://localhost:3000
2. Paste ZPL code into the editor:

```zpl
^XA
^FO50,50^ADN,36,20^FDTest Barcode^FS
^FO50,100^BY3^BCN,100,Y,N,N^FD${BARCODE_DATA}^FS
^FO50,210^ADN,18,10^FD${BARCODE_DATA}^FS
^XZ
```

3. Enter test data in Variables section:
   - `BARCODE_DATA`: `123456789`
   - `COMPANY_NAME`: `Test Company`

### Zebrash API Usage

```bash
curl -X POST http://localhost:8080/convert \
  -H "Content-Type: application/json" \
  -d '{
    "zpl": "^XA^FO50,50^BY3^BCN,100,Y,N,N^FD123456789^FS^XZ",
    "format": "png"
  }'
```

### ZplDesigner API Usage

```bash
curl -X POST http://localhost:8081/api/zpl/convert \
  -H "Content-Type: application/json" \
  -d '{
    "zpl": "^XA^FO50,50^BY3^BCN,100,Y,N,N^FD123456789^FS^XZ"
  }'
```

## 🧪 Test Data

### Code 128 Test Data:

- `123456789`
- `ABC123DEF`
- `TEST123456`

### Code 39 Test Data:

- `ABC123`
- `123-ABC`
- `TEST-123`

### QR Code Test Data:

- `https://www.google.com`
- `Hello World`
- `TEST123456`

## 🔧 Development

### Local Development Environment

```bash
# ZPL-Tool development
cd zpl-tool
npm install
npm run dev

# Zebrash development
cd zebrash
go mod download
go run main.go

# ZplDesigner development
cd ZplDesigner
dotnet restore
dotnet run --project ZplDesigner.WebAPI
```

### Development with Docker

```bash
# Start all services
docker-compose -f zpl-tool/docker-compose.yml up -d
docker-compose -f zebrash/docker-compose.yml up -d
docker-compose -f ZplDesigner/docker-compose.zpldesigner.yml up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## 📁 File Structure

```
TestProjects/
├── zpl-tool/           # Modern ZPL web tool
├── zebrash/            # Go-based ZPL processing
├── ZplDesigner/        # C#/.NET ZPL design tool
├── barcode-test.html   # Barcode testing tool
├── test-barcodes.zpl   # Test ZPL files
├── BARKOD_TEST_REHBERI.md  # Detailed test guide
└── LabelNotes.md       # Project notes
```

## 🚨 Troubleshooting

### Common Issues:

1. **Port conflict**: Use different ports
2. **Docker connection issue**: Restart Docker service
3. **Camera access**: Use HTTPS or localhost
4. **Barcode not reading**: Increase barcode size

### Log Checking:

```bash
# ZPL-Tool logs
docker-compose -f zpl-tool/docker-compose.yml logs

# Zebrash logs
docker-compose -f zebrash/docker-compose.yml logs

# ZplDesigner logs
docker-compose -f ZplDesigner/docker-compose.zpldesigner.yml logs
```

## 📚 Additional Resources

- [ZPL Command Reference](https://www.zebra.com/content/dam/zebra/software/en/printers/link-os/zpl/zpl-zbi2-pm-en.pdf)
- [Barcode Testing Tools](https://www.online-barcode-reader.com/)
- [ZPL-Tool GitHub](https://github.com/enoy19/zpl-tool)
- [Zebrash GitHub](https://github.com/ingridhq/zebrash)
- [ZplDesigner GitHub](https://github.com/IkeRolfe/ZplDesigner)

## 🤝 Contributing

1. Fork the project
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project uses open source software. Each component has its own license terms.

---

**Note**: This project provides a comprehensive environment for testing and developing ZPL barcodes. All processing is done locally for security and data privacy.
