# 🚀 Uçtan Uca DevOps CI/CD Pipeline & Go Web Uygulaması

Bu proje, Go (Golang) ile geliştirilmiş web tabanlı bir hesap makinesinin sıfırdan AWS bulut ortamına tam otomatik bir şekilde taşınmasını sağlayan **Uçtan Uca (End-to-End) CI/CD Pipeline** mimarisidir. 

Projenin amacı; yazılan kodun donanımdan bağımsız çalışabilmesini (Docker), altyapının kodla yönetilmesini (Terraform) ve dağıtım süreçlerinin insan eli değmeden otomatize edilmesini (Jenkins & Ansible) sağlamaktır.

## 🛠 Kullanılan Teknolojiler

* **Backend:** Go (Golang)
* **Konteyner Mimari:** Docker & Docker Hub
* **CI/CD Otomasyonu:** Jenkins (Pipeline-as-Code)
* **IaC (Altyapı Kodlama):** Terraform
* **Konfigürasyon Yönetimi:** Ansible
* **Bulut Sağlayıcı:** AWS (EC2 & Security Groups)

---

## 🏗 Sistem Mimarisi ve Pipeline Akışı

1. **Geliştirme:** Go ile geliştirilen web uygulaması GitHub'a itilir.
2. **Sürekli Entegrasyon (CI - Jenkins):** * Jenkins kodu GitHub'dan çeker.
   * `Dockerfile` kullanılarak uygulama konteynerize edilir.
   * Oluşturulan imaj, Jenkins üzerinden güvenli bir şekilde **Docker Hub'a** (`retcaps/huseyin-go-app:latest`) itilir.
3. **Altyapı Hazırlığı (IaC - Terraform):** AWS üzerinde projenin çalışacağı EC2 sunucusu ve 8081 portuna izin veren güvenlik grupları Terraform ile ayağa kaldırılır.
4. **Sürekli Dağıtım (CD - Ansible):** Ansible, hedef AWS sunucusuna bağlanır ve Docker Hub'daki en güncel imajı çekerek canlı ortama alır.

---

## 📸 Projeden Görüntüler

### ⚙️ Jenkins CI Pipeline Aşamaları
Jenkins'in kodu çekme, test etme, Docker imajı oluşturma ve Docker Hub'a gönderme aşamalarının (Pipeline) başarıyla tamamlandığı ekran.

![Jenkins Pipeline Başarısı](jenkins-basari.png)
*(NOT: Yukarıdaki satıra Jenkins arayüzünde aşamaların yeşil renkte bittiğini gösteren ekran görüntüsünü koymalısın.)*

### ☁️ AWS Altyapısı ve Sunucu Detayları
Terraform tarafından otomatik olarak oluşturulan AWS EC2 bulut sunucumuz.

![AWS EC2 Instance](aws-sunucu.png)
*(NOT: Yukarıdaki satıra az önce bana attığın, "Go-App-Produ..." yazan ve durumun yeşil renkle "Çalışıyor" olduğunu gösteren o harika AWS tablosunun resmini koymalısın.)*

### 🐳 Docker Hub İmaj Deposu
Jenkins tarafından otomatik olarak derlenip buluta fırlatılan uygulamanın en güncel Docker imajı.

![Docker Hub Registry](docker-hub-imaj.png)
*(NOT: Yukarıdaki satıra Docker Hub sitesine girip retcaps/huseyin-go-app imajını gösteren ekran görüntüsünü koymalısın.)*

### 🚀 Canlı Web Uygulaması
Pipeline'ın son adımı olarak AWS sunucumuzda `8081` portunda ayağa kalkan web tabanlı Go uygulamamız.

![Canlı Uygulama](web-uygulama.png)
*(NOT: Yukarıdaki satıra tarayıcıdan kendi IP adresine (http://63.183.0.0:8081) girip hesap makinesini çalıştırırken aldığın ekran görüntüsünü koymalısın.)*

---

## 💻 Kurulum ve Çalıştırma

Projeyi kendi ortamınızda test etmek isterseniz aşağıdaki adımları izleyebilirsiniz:

**1. AWS Altyapısını Kurma (Terraform)**
```bash
cd devops-terraform
terraform init
terraform apply -auto-approve
