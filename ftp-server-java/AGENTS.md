# Правила подпроекта `ftp-server-java`

Роль агента и режим ответов задаёт корневой [AGENTS.md](../AGENTS.md). Здесь — только специфика Java.

## Стек и запуск

- Spring Boot 4.2.0 (`spring-boot-starter-parent`), Java 25 (baseline Boot 4 — Java 17).
- Сборка: Maven (`mvn`), Gradle на этой машине не установлен.
- Запуск: `mvn spring-boot:run`. Порт/корень — `FTP_PORT`, `FTP_ROOT` или `ftp.port`/`ftp.root` в `application.yml`.

## Структура

- `dev.mix.ftp.server` — `FtpServerApplication` (точка входа), `FtpProperties` (record, `@ConfigurationProperties("ftp")`), `FtpStartupReport` (`ApplicationRunner`, печатает scope шага 1).
- `src/test/…/FtpServerApplicationTests` — `@SpringBootTest`, проверка что контекст поднимается.
- Протокола нет: каркас стартует, печатает план и завершается.

## Конвенции

- Веб-стартера нет и не добавлять: FTP — не HTTP, `spring-boot-starter-web` здесь лишний.
- Значения порта и каталога — только через `FtpProperties`; в `application.yml` они читаются из окружения.
- Boot 4: код разбит на более узкие модули, JSON — Jackson 3, null-безопасность через JSpecify. При апгреде патч-линии ориентироваться на release notes, а не на привычки Boot 3.
- Версию Spring Boot не держать в двух местах: одна — в `<parent>`, остальное наследуется.

## Область реализации

Шаг 1 — [RFC 959](../docs/rfc-959.md), [RFC 1123](../docs/rfc-1123.md), [RFC 854](../docs/rfc-854.md); чек-лист — [../docs/kb/requirements.md](../docs/kb/requirements.md).

## На что смотреть в Java именно

- Управление сессией и жизненным циклом сокета не должно попадать в контейнер ради «красивого» DI: сервер слушает порт дольше, чем живёт любой запрос.
- `ServerSocket`/`Socket` — потоки и тайм-ауты ([RFC 1123](../docs/rfc-1123.md) требует устойчивости сессии) настраиваются явно, дефолты не подходят.
