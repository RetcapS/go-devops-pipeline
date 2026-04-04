provider "aws" {
  region = "eu-central-1"
}

# 1. Ansible'ın sızması için oluşturduğumuz SSH anahtarını AWS'ye tanıtıyoruz
resource "aws_key_pair" "deployer" {
  key_name   = "devops-ansible-key"
  public_key = file("~/.ssh/aws_ansible_key.pub")
}

# 2. Güvenlik Duvarı Kuralları (Portları açıyoruz)
resource "aws_security_group" "allow_web_ssh" {
  name        = "allow_web_ssh"
  description = "SSH ve Go uygulamasi icin portlari ac"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 8081
    to_port     = 8081
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# 3. En güncel Ubuntu imajını otomatik buluyoruz
data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"]

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
  }
}

# 4. Asıl Sunucumuzu (EC2) yaratıyoruz
resource "aws_instance" "app_server" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.micro"
  key_name      = aws_key_pair.deployer.key_name
  vpc_security_group_ids = [aws_security_group.allow_web_ssh.id]

  tags = {
    Name = "Go-App-Production-Server"
  }
}

# Kurulum bitince bize sunucunun IP adresini ekrana yazdıracak
output "sunucu_ip_adresi" {
  value = aws_instance.app_server.public_ip
}
