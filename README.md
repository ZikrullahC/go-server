# Basit HTTP Sunucusu

Bu proje, Go dilinde yazılmış basit bir HTTP sunucusudur. Statik dosyaları sunmak, form verilerini işlemek ve basit bir 'hello' endpoint'i sağlamak için kullanılır.

## Başlarken

Bu talimatlar, projenizi yerel makinenizde çalıştırmak ve geliştirmek için bir kopya almanıza yardımcı olacaktır.

### Önkoşullar

Bu projeyi çalıştırmak için Go'nun yüklü olması gerekmektedir.

### Kurulum

Projeyi yerel makinenize klonlayın:

```bash
https://github.com/ZikrullahC/go-server.git
```

### Çalıştırma

Sunucuyu başlatmak için:

```bash
go run main.go
```

Varsayılan olarak, sunucu 8080 numaralı portta çalışacaktır.

## API Kullanımı

API, aşağıdaki endpoint'leri sunar:

- `GET /hello`: Basit bir selamlama mesajı döndürür.
- `POST /form`: Gönderilen form verilerini alır ve işler.

## Yapılandırma

Sunucu portunu değiştirmek için, `main.go` dosyasındaki `http.ListenAndServe` fonksiyonunu düzenleyin:

```go
http.ListenAndServe(":8080", nil)
