# Запуск базы данных PostgreSQL (предполагается, что PostgreSQL установлен и запущен как служба)

# Функция для запуска Go сервиса
function Start-GoService {
    param (
        [string]$serviceName,
        [string]$servicePath
    )
    Write-Host "Starting $serviceName..."
    Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$servicePath'; go run cmd/main.go"
}

# Функция для запуска frontend
function Start-Frontend {
    param (
        [string]$frontendPath
    )
    Write-Host "Starting frontend..."
    Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$frontendPath'; npm run dev"
}

# Получаем абсолютный путь к корневой директории проекта
$rootPath = (Get-Location).Path

# Запускаем все сервисы параллельно
Start-GoService -serviceName "auth-service" -servicePath "$rootPath\auth-service"
Start-GoService -serviceName "course-service" -servicePath "$rootPath\course-service"
Start-GoService -serviceName "chat-service" -servicePath "$rootPath\chat-service"
Start-Frontend -frontendPath "$rootPath\frontend"

Write-Host "All services are starting..."
Write-Host "Frontend will be available at: http://localhost:3000"
Write-Host "Auth Service will be available at: http://localhost:8081"
Write-Host "Course Service will be available at: http://localhost:8080"
Write-Host "Chat Service will be available at: http://localhost:8082" 