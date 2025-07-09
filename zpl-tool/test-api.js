const axios = require('axios');

// zpl-tool API endpoint
const ZPL_TOOL_URL = 'http://localhost:3000';

// Test ZPL template with variables
const testZplTemplate = `^XA
^FO50,50^ADN,36,20^FD\${COMPANY_NAME}^FS
^FO50,100^ADN,24,12^FD\${PRODUCT_NAME}^FS
^FO50,150^ADN,18,10^FD\${PRICE}^FS
^FO50,200^ADN,12,8^FD\${DATE}^FS
^XZ`;

// Test data
const testData = {
    COMPANY_NAME: 'Test Şirketi',
    PRODUCT_NAME: 'Test Ürünü',
    PRICE: '99.99 TL',
    DATE: '2025-07-09'
};

// @ts-ignore
async function testZplTool() {
    try {
        console.log('🚀 zpl-tool API Test Başlatılıyor...');
        
        // 1. Template oluşturma (web arayüzü üzerinden yapılır)
        console.log('📝 ZPL Template:');
        console.log(testZplTemplate);
        
        // 2. Değişken değerlerini göster
        console.log('\n📊 Test Değişkenleri:');
        console.log(JSON.stringify(testData, null, 2));
        
        // 3. BinaryKits.Zpl API'sini test et (port 7763)
        console.log('\n🖼️ BinaryKits.Zpl API Test:');
        const binaryKitsResponse = await axios.post('http://localhost:7763/api/zpl/render', {
            zpl: testZplTemplate.replace(/\$\{(\w+)\}/g, (match, variable) => {
                return testData[variable] || match;
            })
        }, {
            responseType: 'arraybuffer'
        });
        
        console.log('✅ BinaryKits.Zpl API çalışıyor!');
        console.log(`📏 Response size: ${binaryKitsResponse.data.length} bytes`);
        
        // 4. Kullanım talimatları
        console.log('\n📋 Kullanım Talimatları:');
        console.log('1. Web arayüzü: http://localhost:3000');
        console.log('2. Template editöründe ZPL kodunu yapıştırın');
        console.log('3. Variables bölümünde değişkenleri tanımlayın');
        console.log('4. Live preview otomatik olarak güncellenecek');
        console.log('5. Printers bölümünden yazıcı ekleyebilirsiniz');
        
    } catch (error) {
        console.error('❌ Hata:', error.message);
    }
}

testZplTool(); 