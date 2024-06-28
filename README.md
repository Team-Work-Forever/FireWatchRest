# FireWatch 🔥
Esta aplicação tem como objetivo gerir queimadas em todo o país: Portugal.

## Como iniciar:

### Clonar o repositório:

```bash
git clone https://github.com/Team-Work-Forever/FireWatchRest.git firewatch
```

### Naveguar para o diretório do projeto:
```bash
cd firewatch
```

### Colocar ficheiro .env:

```bash
cp .env.example .env
```

#### Exemplo de .env:
```env
POSTGRES_PASSWORD=password
POSTGRES_USER=sacanaArmado_sa
POSTGRES_DB=verceldb
POSTGRES_HOST=superHost
POSTGRES_PORT=5432

FIRE_WATCH_API_PORT=3000

JWT_AUDIENCE=Fire Watch
JWT_ISSUER=firewatch.com
JWT_ACCESS_EXPIRED=15
JWT_REFRESH_EXPIRED=2
JWT_SECRET=E9097979ABF51C113F7772C086E569E7192AABEFBA9050D962EC312BB27DB89A

SMTP_HOST_EMAIL=no-replay@firewatch.io
SMTP_HOST=sandbox.smtp.mailtrap.io
SMTP_PORT=25
SMTP_HOST_USER=sacanaArmado_sa
SMTP_HOST_PASSWORD=superPassword

BLOB_ACCESS_KEY=BLOB_ACCESS_KEY
BLOB_PROJECT_KEY=E9097979ABF51C113F7772C086E569E7192AABEFBA9050D962EC312BB27DB89A
BLOB_PUBLIC_URL=BLOB_PUBLIC_URL
BLOB_S3_URL=BLOB_S3_URL
BLOB_REGION=us-west-1

REDIS_USER=sacanaArmado_sa
REDIS_PASSWD=E9097979ABF51C113F7772C086E569E7192AABEFBA9050D962EC312BB27DB89A
REDIS_HOST=superHost
REDIS_PORT=6379
REDIS_DB=0POSTGRES_PASSWORD=password
```

### Correr o programa (instala as dependências automaticamente):

```bash
go build cmd/fireWatch/main.go  
```

```bash
./main
```

## Endpoints da API
Para visualizar os endpoints da API pode utilizar o:

### Swagger  [URL](https://firewatchrest.onrender.com/swagger/index.html) 👈

![image](https://github.com/Team-Work-Forever/FireWatchRest/assets/74202840/09db6606-c1bb-435e-a777-d3dd711998e6)
