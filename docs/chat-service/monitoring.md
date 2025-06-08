# Руководство по мониторингу Chat Service

## Обзор

В этом документе описаны подходы и инструменты для мониторинга Chat Service, включая метрики, логирование и алерты.

## Метрики

### Prometheus

#### Конфигурация
```yaml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'chat-service'
    static_configs:
      - targets: ['localhost:8082']
    metrics_path: '/metrics'
```

#### Основные метрики

1. HTTP метрики:
```go
var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path", "status"},
    )

    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "HTTP request duration in seconds",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "path"},
    )
)
```

2. WebSocket метрики:
```go
var (
    wsConnectionsTotal = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "ws_connections_total",
            Help: "Total number of WebSocket connections",
        },
    )

    wsMessagesTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "ws_messages_total",
            Help: "Total number of WebSocket messages",
        },
        []string{"type"},
    )
)
```

3. База данных метрики:
```go
var (
    dbOperationsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "db_operations_total",
            Help: "Total number of database operations",
        },
        []string{"operation", "collection"},
    )

    dbOperationDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "db_operation_duration_seconds",
            Help:    "Database operation duration in seconds",
            Buckets: prometheus.DefBuckets,
        },
        []string{"operation", "collection"},
    )
)
```

### Grafana

#### Дашборды

1. Общий обзор:
```json
{
  "dashboard": {
    "title": "Chat Service Overview",
    "panels": [
      {
        "title": "HTTP Requests",
        "type": "graph",
        "datasource": "Prometheus",
        "targets": [
          {
            "expr": "rate(http_requests_total[5m])",
            "legendFormat": "{{method}} {{path}}"
          }
        ]
      },
      {
        "title": "WebSocket Connections",
        "type": "gauge",
        "datasource": "Prometheus",
        "targets": [
          {
            "expr": "ws_connections_total"
          }
        ]
      }
    ]
  }
}
```

2. Производительность:
```json
{
  "dashboard": {
    "title": "Performance Metrics",
    "panels": [
      {
        "title": "Response Time",
        "type": "graph",
        "datasource": "Prometheus",
        "targets": [
          {
            "expr": "rate(http_request_duration_seconds_sum[5m]) / rate(http_request_duration_seconds_count[5m])",
            "legendFormat": "{{method}} {{path}}"
          }
        ]
      },
      {
        "title": "Database Operations",
        "type": "graph",
        "datasource": "Prometheus",
        "targets": [
          {
            "expr": "rate(db_operations_total[5m])",
            "legendFormat": "{{operation}} {{collection}}"
          }
        ]
      }
    ]
  }
}
```

## Логирование

### ELK Stack

#### Конфигурация Logstash
```yaml
input {
  beats {
    port => 5044
  }
}

filter {
  if [type] == "chat-service" {
    json {
      source => "message"
    }
    date {
      match => [ "timestamp", "ISO8601" ]
    }
  }
}

output {
  elasticsearch {
    hosts => ["localhost:9200"]
    index => "chat-service-%{+YYYY.MM.dd}"
  }
}
```

#### Структура логов

1. HTTP логи:
```go
func logHTTPRequest(r *http.Request, status int, duration time.Duration) {
    logger.Info("http_request",
        zap.String("method", r.Method),
        zap.String("path", r.URL.Path),
        zap.Int("status", status),
        zap.Duration("duration", duration),
        zap.String("ip", r.RemoteAddr),
        zap.String("user_agent", r.UserAgent()),
    )
}
```

2. WebSocket логи:
```go
func logWebSocketEvent(conn *websocket.Conn, eventType string, data interface{}) {
    logger.Info("ws_event",
        zap.String("type", eventType),
        zap.Any("data", data),
        zap.String("user_id", getUserID(conn)),
        zap.String("chat_id", getChatID(conn)),
    )
}
```

3. Ошибки:
```go
func logError(err error, context map[string]interface{}) {
    logger.Error("error",
        zap.Error(err),
        zap.Any("context", context),
        zap.String("stack", string(debug.Stack())),
    )
}
```

## Алерты

### AlertManager

#### Конфигурация
```yaml
global:
  resolve_timeout: 5m

route:
  group_by: ['alertname']
  group_wait: 30s
  group_interval: 5m
  repeat_interval: 4h
  receiver: 'slack-notifications'

receivers:
  - name: 'slack-notifications'
    slack_configs:
      - channel: '#alerts'
        send_resolved: true
```

#### Правила алертов

1. Высокая нагрузка:
```yaml
groups:
  - name: chat-service
    rules:
      - alert: HighRequestRate
        expr: rate(http_requests_total[5m]) > 1000
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: High request rate
          description: Request rate is above 1000 requests per second

      - alert: HighWebSocketConnections
        expr: ws_connections_total > 10000
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: High WebSocket connections
          description: Number of WebSocket connections is above 10000
```

2. Ошибки:
```yaml
groups:
  - name: chat-service
    rules:
      - alert: HighErrorRate
        expr: rate(http_requests_total{status=~"5.."}[5m]) / rate(http_requests_total[5m]) > 0.05
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: High error rate
          description: Error rate is above 5%

      - alert: DatabaseErrors
        expr: rate(db_operations_total{operation="error"}[5m]) > 0
        for: 1m
        labels:
          severity: critical
        annotations:
          summary: Database errors
          description: Database operations are failing
```

## Мониторинг системы

### Node Exporter

#### Метрики
- CPU использование
- Использование памяти
- Использование диска
- Сетевая активность

### cAdvisor

#### Метрики
- Использование контейнеров
- Сетевые метрики
- Файловая система

## Визуализация

### Grafana Dashboards

1. Системные метрики:
```json
{
  "dashboard": {
    "title": "System Metrics",
    "panels": [
      {
        "title": "CPU Usage",
        "type": "graph",
        "datasource": "Prometheus",
        "targets": [
          {
            "expr": "100 - (avg by (instance) (irate(node_cpu_seconds_total{mode=\"idle\"}[5m])) * 100)"
          }
        ]
      },
      {
        "title": "Memory Usage",
        "type": "graph",
        "datasource": "Prometheus",
        "targets": [
          {
            "expr": "node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes"
          }
        ]
      }
    ]
  }
}
```

2. Бизнес метрики:
```json
{
  "dashboard": {
    "title": "Business Metrics",
    "panels": [
      {
        "title": "Active Users",
        "type": "stat",
        "datasource": "Prometheus",
        "targets": [
          {
            "expr": "ws_connections_total"
          }
        ]
      },
      {
        "title": "Messages per Second",
        "type": "graph",
        "datasource": "Prometheus",
        "targets": [
          {
            "expr": "rate(ws_messages_total[5m])"
          }
        ]
      }
    ]
  }
}
```

## Рекомендации

### 1. Настройка мониторинга
- Настройте базовые метрики
- Добавьте бизнес-метрики
- Настройте алерты
- Регулярно проверяйте дашборды

### 2. Обслуживание
- Регулярно обновляйте конфигурации
- Проверяйте работоспособность алертов
- Очищайте старые логи
- Оптимизируйте запросы

### 3. Улучшения
- Добавляйте новые метрики
- Улучшайте визуализацию
- Настраивайте новые алерты
- Оптимизируйте производительность 