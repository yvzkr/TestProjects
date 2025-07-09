# Barkod Test Rehberi

## 🎯 Barkod Test Yöntemleri

### 1. 📱 Kamera ile Test (Önerilen)

- `barcode-test.html` dosyasını tarayıcıda açın
- "Kamerayı Başlat" butonuna tıklayın
- ZPL-tool'da oluşturduğunuz barkodu kameraya gösterin
- Barkod okunduğunda veriler görünecek

### 2. 📱 Telefon Uygulamaları

**Android/iOS için ücretsiz barkod okuyucular:**

- **Google Lens** (Android/iOS)
- **QR & Barcode Scanner** (Android)
- **QR Code Reader** (iOS)
- **Barcode Scanner** (Android)

### 3. 🖥️ Online Barkod Test Araçları

- [Online Barcode Reader](https://www.online-barcode-reader.com/)
- [Barcode Scanner Online](https://barcode-scanner-online.com/)
- [QR Code Scanner Online](https://www.qr-code-scanner.org/)

## 🔧 ZPL-tool'da Barkod Test Adımları

### Adım 1: Barkod Oluşturma

1. ZPL-tool'u açın: `http://localhost:3000`
2. Editor'e barkod ZPL kodunu yapıştırın:

```zpl
^XA
^FO50,50^ADN,36,20^FDTest Barkod^FS
^FO50,100^BY3^BCN,100,Y,N,N^FD${BARCODE_DATA}^FS
^FO50,210^ADN,18,10^FD${BARCODE_DATA}^FS
^XZ
```

### Adım 2: Test Verisi Girme

1. Variables bölümünde:
   - `BARCODE_DATA`: `123456789`
   - `COMPANY_NAME`: `Test Şirketi`
   - `PRODUCT_NAME`: `Test Ürün`

### Adım 3: Barkod Test Etme

1. **Kamera ile test:**

   - `barcode-test.html` dosyasını açın
   - Kamerayı başlatın
   - ZPL-tool'daki preview'i telefon/tablet ile gösterin
   - Barkodu okutun

2. **Manuel test:**
   - ZPL-tool'da Variables bölümündeki veriyi kopyalayın
   - `barcode-test.html`'deki "Manuel Barkod Test" bölümüne yapıştırın
   - "Test Et" butonuna tıklayın

## 📋 Test Verileri

### Code 128 Test Verileri:

- `123456789`
- `ABC123DEF`
- `TEST123456`

### Code 39 Test Verileri:

- `ABC123`
- `123-ABC`
- `TEST-123`

### QR Code Test Verileri:

- `https://www.google.com`
- `Merhaba Dünya`
- `TEST123456`

### EAN-13 Test Verileri:

- `1234567890123` (13 hane)
- `1234567890128` (13 hane)

### UPC-A Test Verileri:

- `123456789012` (12 hane)

## ✅ Test Kontrol Listesi

- [ ] Barkod görsel olarak doğru görünüyor mu?
- [ ] Barkod okutulduğunda doğru veri çıkıyor mu?
- [ ] Barkod boyutu uygun mu?
- [ ] Barkod kalitesi yeterli mi?
- [ ] İnsan okunabilir metin doğru mu?

## 🚨 Yaygın Sorunlar ve Çözümler

### Problem: Barkod okunmuyor

**Çözüm:**

- Barkod boyutunu artırın (`^BY3` → `^BY5`)
- Barkod yüksekliğini artırın (`100` → `150`)
- Barkod kalitesini kontrol edin

### Problem: Yanlış veri okunuyor

**Çözüm:**

- Variables bölümündeki veriyi kontrol edin
- Barkod formatını kontrol edin
- Test verilerini doğrulayın

### Problem: Barkod çok büyük/küçük

**Çözüm:**

- `^BY` parametresini ayarlayın (1-10 arası)
- Barkod yüksekliğini ayarlayın
- Etiket boyutunu kontrol edin

## 🎯 Test Senaryoları

### Senaryo 1: Basit Code 128

```zpl
^XA
^FO50,50^BY3^BCN,100,Y,N,N^FD123456789^FS
^FO50,160^ADN,18,10^FD123456789^FS
^XZ
```

### Senaryo 2: QR Code ile URL

```zpl
^XA
^FO50,50^BQN,2,5^FDhttps://www.example.com^FS
^FO50,200^ADN,18,10^FDhttps://www.example.com^FS
^XZ
```

### Senaryo 3: Karma Barkod

```zpl
^XA
^FO50,50^ADN,36,20^FD${COMPANY_NAME}^FS
^FO50,100^BY3^BCN,80,Y,N,N^FD${BARCODE_DATA}^FS
^FO50,190^ADN,18,10^FD${BARCODE_DATA}^FS
^FO50,220^BQN,2,4^FD${QR_DATA}^FS
^FO50,320^ADN,12,8^FD${QR_DATA}^FS
^XZ
```

## 📱 Mobil Test İpuçları

1. **Telefon kamerası kullanın** - En pratik yöntem
2. **Farklı açılardan test edin** - 45° açıdan da okunmalı
3. **Farklı mesafelerden test edin** - 5-30 cm arası
4. **Farklı ışık koşullarında test edin**
5. **Birden fazla uygulama ile test edin**

## 🔍 Barkod Kalite Kontrolü

### Görsel Kontrol:

- [ ] Barkod çizgileri net ve düzgün
- [ ] Boşluklar (quiet zone) yeterli
- [ ] Boyut oranları doğru
- [ ] Renk kontrastı yeterli

### Fonksiyonel Kontrol:

- [ ] Barkod okunabiliyor
- [ ] Doğru veri çıkıyor
- [ ] Hata düzeltme çalışıyor
- [ ] Farklı cihazlarda çalışıyor

Bu rehber ile barkodlarınızı etkili bir şekilde test edebilir ve kalitelerini kontrol edebilirsiniz!
