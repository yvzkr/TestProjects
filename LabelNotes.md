**Konu: ZPL Etiket Verilerinin Yerel Sunucularda (On-Premise) İşlenmesi İçin Açık Kaynaklı Sistem Alternatifleri**

Şirketimizin operasyonel süreçlerinde kullandığı ZPL tabanlı etiketlerin oluşturulması sırasında, mevcut durumda kullanılan harici web servisleri (örneğin Labelary.com) veri gizliliği ve bilgi güvenliği açısından risk teşkil etmektedir. Müşteri veya ürün verisi gibi hassas bilgilerin kontrolümüz dışındaki sunuculara gönderilmesini engellemek amacıyla, bu işlevselliği kendi altyapımızda barındıracak (on-premise) bir çözüme geçilmesi hedeflenmektedir.

Bu ihtiyaca yönelik yapılan araştırmalar sonucunda, ZPL verilerini PNG/PDF gibi formatlara dönüştürebilen, güvenilir ve açık kaynak kodlu aşağıdaki alternatifler tespit edilmiştir:

1.  **zpl-tool:** Labelary'ye benzer modern bir web arayüzü sunan, Docker ile kolayca kurulabilen ve canlı önizleme gibi gelişmiş özellikler barındıran komple bir çözümdür. Hızlıca devreye alınabilecek en yetenekli alternatiftir.

    - **Repo Linki:** [https://github.com/enoy19/zpl-tool](https://github.com/enoy19/zpl-tool)

    DeveloperNote:

    - Localde test ettim. Docker ile çalışıyor.
    - Docker ile çalışırken, docker-compose.yml dosyasında, binarykits-zpl servisini açık bırakıyorum.
    - SvelteKit ile yazılmış bir arayüzü var. Manüpüle yapılabiliyor(denedim).
    - repoyu forklamak yeterli, Fazla star yok ama düşündürdü.

2.  **Zebrash:** Doğrudan Labelary alternatifi olarak geliştirilmiş, yüksek performanslı bir Go kütüphanesidir. Mevcut sistemlerimize ZPL'den resim oluşturma yeteneğini bir mikroservis olarak eklemek için oldukça uygundur.

    - **Repo Linki:** [https://github.com/ingridhq/zebrash](https://github.com/ingridhq/zebrash)

3.  **ZplDesigner:** C#/.NET tabanlı bir ZPL işleme kütüphanesidir. Şirketimizdeki mevcut .NET tabanlı projelere kolayca entegre edilebilecek olması nedeniyle stratejik bir avantaj sunmaktadır.
    - **Repo Linki:** [https://github.com/IkeRolfe/ZplDesigner](https://github.com/IkeRolfe/ZplDesigner)

Bu projelerden birinin benimsenmesi, veri güvenliğimizi artırırken dışa bağımlılığı ortadan kaldıracak ve uzun vadede maliyet avantajı sağlayacaktır.
