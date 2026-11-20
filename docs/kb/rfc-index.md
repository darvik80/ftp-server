# Полный индекс RFC (92)

Полное покрытие `docs/rfc-*.md`, упорядочено по номеру приоритета. Год, статус и
приоритет взяты дословно из метаданных страниц; теги — из `Ключевых положений`.
Метка и номер приоритета — независимые поля (см. [README](README.md), раздел «Известные
особенности»).

| Приоритет | RFC | Название | Год | Статус | Теги | Файл |
|---|---|---|---|---|---|---|
| 1 — Критично | [959](../rfc-959.md) | Протокол передачи файлов (FTP) | 1985 | Internet Standard | управляющее-соединение, соединение-данных, PORT, PASV, коды-ответов, команды, TYPE/FORMAT/MODE, REST, LIST/NLST, NVT-ASCII/CRLF, third-party-transfer, STD 9 | rfc-959.md |
| 2 — Критично | [1123](../rfc-1123.md) | Требования к интернет-хостам — приложения и поддержка | 1989 | Internet Standard | TYPE A/I, MODE S, PASV, PORT, тайм-ауты, multihomed, LIST/NLST, коды-ответов, STD 3, безопасность | rfc-1123.md |
| 3 — Критично | [2389](../rfc-2389.md) | Механизм согласования возможностей FTP | 1998 | Proposed Standard | FEAT, OPTS, многострочный 211, IANA-токены | rfc-2389.md |
| 4 — Критично | [2428](../rfc-2428.md) | Расширения FTP для IPv6 и NAT | 1998 | Proposed Standard | EPSV, EPRT, NAT/firewall, IPv6, коды 229/522, FEAT | rfc-2428.md |
| 5 — Критично | [3659](../rfc-3659.md) | Расширения FTP | 2007 | Proposed Standard | SIZE, MDTM, MLST/MLSD, REST STREAM, TVFS, OPTS, FEAT | rfc-3659.md |
| 6 — Высокий | [1579](../rfc-1579.md) | FTP, дружественный к межсетевым экранам | 1994 | Informational | PASV, EPSV, NAT/firewall, диапазон-портов-данных | rfc-1579.md |
| 7 — Высокий | [2577](../rfc-2577.md) | Соображения безопасности FTP | 1999 | Informational | анти-bounce, path-traversal, chroot/jail, anonymous, утечки-в-баннерах | rfc-2577.md |
| 8 — Высокий | [2228](../rfc-2228.md) | Расширения безопасности FTP | 1997 | Proposed Standard | AUTH, PROT/PBSZ, ADAT/MIC/CONF/ENC, CCC, коды безопасности | rfc-2228.md |
| 9 — Высокий | [4217](../rfc-4217.md) | Защита FTP с помощью TLS | 2005 | Proposed Standard | AUTH TLS, PROT/PBSZ, implicit-FTPS(990), TLS-data, политика-TLS, FEAT | rfc-4217.md |
| 10 — Высокий | [2640](../rfc-2640.md) | Интернационализация FTP | 1999 | Proposed Standard | LANG/UTF-8, FEAT(UTF8), OPTS | rfc-2640.md |
| 11 — Средний | [7151](../rfc-7151.md) | Команда HOST для виртуальных хостов FTP | 2014 | Proposed Standard | виртуальные-хосты(HOST), коды 220/504/530/550, FEAT | rfc-7151.md |
| 12 — Средний | [5797](../rfc-5797.md) | Реестр команд и расширений FTP | 2010 | Proposed Standard | IANA-реестр, FEAT/OPTS-токены | rfc-5797.md |
| 13 — Средний | [1635](../rfc-1635.md) | Как пользоваться анонимным FTP | 1994 | Informational | anonymous-доступ, chroot/jail, anonymous-политика | rfc-1635.md |
| 14 — Низкий | [6384](../rfc-6384.md) | FTP ALG для трансляции IPv6↔IPv4 | 2011 | Proposed Standard | ALG, NAT/firewall, EPSV, IPv6↔IPv4 | rfc-6384.md |
| 15 — Низкий | [1639](../rfc-1639.md) | Работа FTP с «длинными» адресами (FOOBAR) | 1994 | Experimental | LPRT/LPSV, HOST-PORT, IPv6-предки | rfc-1639.md |
| 16 — Справочно | [1545](../rfc-1545.md) | FOOBAR (ранняя редакция) | 1993 | Experimental | LPRT/LPSV, HOST-PORT, историческая-редакция | rfc-1545.md |
| 17 — Справочно | [775](../rfc-775.md) | Каталого-ориентированные команды FTP | 1980 | Unknown | команды MKD/RMD/PWD, LIST/NLST | rfc-775.md |
| 18 — Справочно | [640](../rfc-640.md) | Пересмотренные коды ответов FTP | 1974 | Unknown | коды-ответов | rfc-640.md |
| 19 — Справочно | [765](../rfc-765.md) | Спецификация FTP (предшественник) | 1980 | Unknown | управляющее-соединение, команды, предшественник-959 | rfc-765.md |
| 20 — Средний | [854](../rfc-854.md) | Спецификация протокола Telnet | 1983 | Internet Standard | Telnet-NVT, NVT-ASCII/CRLF, IAC | rfc-854.md |
| 21 — Справочно | [9141](../rfc-9141.md) | Обновление ссылок на FTP-сервис IETF | 2021 | Proposed Standard | URL-scheme (ftp:// → https), административное | rfc-9141.md |
| 22 — Низкий | [2585](../rfc-2585.md) | Операционные протоколы PKIX: FTP и HTTP | 1999 | Proposed Standard | PKIX/CRL-раздача, URL-scheme, TYPE binary | rfc-2585.md |
| 23 — Низкий | [4823](../rfc-4823.md) | FTP как транспорт для безопасного B2B-обмена | 2007 | Informational | S/MIME-профиль, EDI/B2B | rfc-4823.md |
| 24 — Справочно | [5024](../rfc-5024.md) | ODETTE File Transfer Protocol 2.0 | 2007 | Informational | OFTP/EDI, TLS, CMS-подписи | rfc-5024.md |
| 25 — Справочно | [2204](../rfc-2204.md) | ODETTE File Transfer Protocol | 1997 | Informational | OFTP/EDI | rfc-2204.md |
| 30 — Справочно | [949](../rfc-949.md) | Команда сохранения с уникальным именем | 1985 | Unknown | команда STOU | rfc-949.md |
| 31 — Справочно | [913](../rfc-913.md) | Простой протокол передачи файлов (SFTP) | 1984 | Historic | SFTP-1984, BFTP, OFTP-семейство | rfc-913.md |
| 32 — Справочно | [847](../rfc-847.md) | Сводка опросов Smallberg (Telnet/FTP/SMTP) | 1983 | Unknown | статистика-внедрения, Telnet-NVT, SMTP | rfc-847.md |
| 33 — Справочно | [751](../rfc-751.md) | Обзор FTP mail и MLFL | 1978 | Unknown | FTP-mail, MLFL | rfc-751.md |
| 34 — Справочно | [743](../rfc-743.md) | Расширение FTP: XRSQ/XRCP | 1977 | Unknown | FTP-mail (XRSQ/XRCP) | rfc-743.md |
| 35 — Справочно | [737](../rfc-737.md) | Расширение FTP: XSEN | 1977 | Unknown | FTP-mail (XSEN) | rfc-737.md |
| 36 — Справочно | [697](../rfc-697.md) | Команда CWD в FTP | 1975 | Unknown | команда CWD, доступ-к-каталогам | rfc-697.md |
| 37 — Справочно | [691](../rfc-691.md) | Ещё одна попытка уточнить FTP | 1975 | Unknown | print-файлы, TYPE | rfc-691.md |
| 38 — Справочно | [686](../rfc-686.md) | «Не трогай, пока работает» | 1975 | Unknown | версии-FTP, TYPE/print | rfc-686.md |
| 39 — Справочно | [683](../rfc-683.md) | FTPSRV — расширение Tenex для page-файлов | 1975 | Unknown | page-mode/TENEX | rfc-683.md |
| 40 — Справочно | [630](../rfc-630.md) | Коды ошибок FTP для надёжной почты | 1974 | Unknown | коды-ответов, FTP-mail | rfc-630.md |
| 41 — Справочно | [624](../rfc-624.md) | Комментарии к FTP | 1974 | Unknown | команды, дизайн-заметки | rfc-624.md |
| 42 — Справочно | [614](../rfc-614.md) | Ответ на RFC 607 | 1974 | Unknown | дискуссия, коды-ответов | rfc-614.md |
| 43 — Справочно | [607](../rfc-607.md) | Комментарии к FTP (ранние) | 1974 | Unknown | команды, дизайн-заметки | rfc-607.md |
| 44 — Справочно | [593](../rfc-593.md) | Изменение графика внедрения Telnet и FTP | 1973 | Unknown | график-внедрения | rfc-593.md |
| 45 — Справочно | [571](../rfc-571.md) | Проблема FTP на TENEX | 1973 | Unknown | TENEX-специфика | rfc-571.md |
| 46 — Справочно | [542](../rfc-542.md) | Протокол передачи файлов (1973) | 1973 | Unknown | управляющее-соединение, команды, коды-ответов | rfc-542.md |
| 47 — Справочно | [532](../rfc-532.md) | Server-FTP в UCSD-CC | 1973 | Unknown | реализация-сервера | rfc-532.md |
| 48 — Справочно | [520](../rfc-520.md) | Предложение File Access Protocol | 1973 | Unknown | файловый-доступ | rfc-520.md |
| 49 — Справочно | [506](../rfc-506.md) | Проблема именования команд FTP | 1973 | Unknown | именование-команд | rfc-506.md |
| 50 — Справочно | [505](../rfc-505.md) | Два решения проблемы доступа при передаче файлов | 1973 | Unknown | контроль-доступа | rfc-505.md |
| 51 — Справочно | [501](../rfc-501.md) | Прояснение «свободной» передачи файлов | 1973 | Unknown | free-file-transfer | rfc-501.md |
| 52 — Справочно | [487](../rfc-487.md) | Свободная передача файлов | 1973 | Unknown | free-file-transfer | rfc-487.md |
| 53 — Справочно | [486](../rfc-486.md) | Передача данных: повторный взгляд | 1973 | Unknown | DTP, соединение-данных | rfc-486.md |
| 54 — Справочно | [480](../rfc-480.md) | Хост-зависимые параметры FTP | 1973 | Unknown | хост-зависимые-параметры | rfc-480.md |
| 55 — Справочно | [479](../rfc-479.md) | Использование FTP журналом NIC | 1973 | Unknown | прикладной-кейс | rfc-479.md |
| 56 — Справочно | [478](../rfc-478.md) | Взаимодействие сервер–сервер FTP — II | 1973 | Unknown | third-party-transfer | rfc-478.md |
| 57 — Справочно | [475](../rfc-475.md) | FTP и сетевая почта | 1973 | Unknown | FTP-mail | rfc-475.md |
| 58 — Справочно | [468](../rfc-468.md) | Сжатие данных в FTP | 1973 | Unknown | MODE C, сжатие | rfc-468.md |
| 59 — Справочно | [463](../rfc-463.md) | Комментарии FTP и ответ на RFC 430 | 1973 | Unknown | дискуссия | rfc-463.md |
| 60 — Справочно | [458](../rfc-458.md) | Получение почты через FTP | 1973 | Unknown | FTP-mail | rfc-458.md |
| 61 — Справочно | [454](../rfc-454.md) | FTP — анонс встречи и новый черновик | 1973 | Unknown | организационное | rfc-454.md |
| 62 — Справочно | [448](../rfc-448.md) | Print-файлы в FTP | 1973 | Unknown | print-файлы | rfc-448.md |
| 63 — Справочно | [438](../rfc-438.md) | Взаимодействие сервер–сервер FTP | 1973 | Unknown | third-party-transfer, bounce-атака | rfc-438.md |
| 64 — Справочно | [430](../rfc-430.md) | Комментарии к протоколу передачи файлов | 1973 | Unknown | дискуссия | rfc-430.md |
| 65 — Справочно | [414](../rfc-414.md) | Статус FTP и дальнейшие комментарии | 1972 | Unknown | статус-внедрения | rfc-414.md |
| 66 — Справочно | [412](../rfc-412.md) | Документация пользовательского FTP | 1972 | Unknown | пользовательский-клиент | rfc-412.md |
| 67 — Справочно | [385](../rfc-385.md) | Комментарии к FTP (1972) | 1972 | Unknown | обновления RFC 354 | rfc-385.md |
| 68 — Справочно | [354](../rfc-354.md) | Протокол передачи файлов (1972) | 1972 | Unknown | управляющее-соединение, команды | rfc-354.md |
| 69 — Справочно | [327](../rfc-327.md) | Заметки воркшопа по Data/File Transfer | 1972 | Unknown | воркшоп DTP/FTP | rfc-327.md |
| 70 — Справочно | [310](../rfc-310.md) | Ещё один взгляд на протоколы передачи данных и файлов | 1972 | Unknown | DTP/FTP-анализ | rfc-310.md |
| 71 — Справочно | [309](../rfc-309.md) | Анонс воркшопа Data and File Transfer | 1972 | Unknown | воркшоп-анонс | rfc-309.md |
| 72 — Справочно | [294](../rfc-294.md) | Использование транзакции Set Data Type в FTP | 1972 | Unknown | TYPE/FORMAT/MODE | rfc-294.md |
| 73 — Справочно | [281](../rfc-281.md) | Предлагаемое дополнение к FTP | 1971 | Unknown | расширение-команд | rfc-281.md |
| 74 — Справочно | [269](../rfc-269.md) | Некоторый опыт передачи файлов | 1971 | Unknown | опыт-эксплуатации | rfc-269.md |
| 75 — Справочно | [265](../rfc-265.md) | Протокол передачи файлов (1971) | 1971 | Unknown | команды раннего FTP | rfc-265.md |
| 76 — Справочно | [264](../rfc-264.md) | Протокол передачи данных (DTP) | 1972 | Unknown | DTP, соединение-данных | rfc-264.md |
| 77 — Справочно | [250](../rfc-250.md) | Некоторые мысли о передаче файлов | 1971 | Unknown | эссе file transfer | rfc-250.md |
| 78 — Справочно | [238](../rfc-238.md) | Комментарии к предложениям DTP и FTP | 1971 | Unknown | ревью DTP/FTP | rfc-238.md |
| 79 — Справочно | [172](../rfc-172.md) | Протокол передачи файлов (июнь 1971) | 1971 | Unknown | команды раннего FTP | rfc-172.md |
| 80 — Справочно | [171](../rfc-171.md) | Протокол передачи данных (1971) | 1971 | Unknown | DTP, соединение-данных | rfc-171.md |
| 81 — Справочно | [163](../rfc-163.md) | Протоколы передачи данных | 1971 | Unknown | обзор data transfer | rfc-163.md |
| 82 — Справочно | [141](../rfc-141.md) | Комментарии к RFC 114 | 1971 | Unknown | рецензия RFC 114 | rfc-141.md |
| 83 — Справочно | [133](../rfc-133.md) | Передача файлов и восстановление после ошибок | 1971 | Unknown | REST, error recovery | rfc-133.md |
| 84 — Справочно | [114](../rfc-114.md) | Протокол передачи файлов (первый RFC) | 1971 | Unknown | историческая-точка-отсчёта | rfc-114.md |
| 85 — Справочно | [1350](../rfc-1350.md) | Протокол TFTP (редакция 2) | 1992 | Internet Standard | TFTP, UDP/69, STD 33 | rfc-1350.md |
| 86 — Справочно | [2347](../rfc-2347.md) | Расширение опций TFTP | 1998 | Draft Standard | TFTP, options | rfc-2347.md |
| 87 — Справочно | [2348](../rfc-2348.md) | Опция размера блока TFTP | 1998 | Draft Standard | TFTP, blksize | rfc-2348.md |
| 88 — Справочно | [2349](../rfc-2349.md) | Опции таймаута и размера передачи TFTP | 1998 | Draft Standard | TFTP, timeout/tsize | rfc-2349.md |
| 89 — Справочно | [3617](../rfc-3617.md) | Схема URI и применимость TFTP | 2003 | Informational | TFTP, tftp-URI, безопасность | rfc-3617.md |
| 90 — Справочно | [1986](../rfc-1986.md) | Эксперименты ETFTP для радиоканалов | 1996 | Experimental | NETBLT/ETFTP, радиоканал | rfc-1986.md |
| 91 — Справочно | [1440](../rfc-1440.md) | SIFT/UFT: передача по инициативе отправителя | 1993 | Experimental | SIFT/UFT | rfc-1440.md |
| 92 — Справочно | [1415](../rfc-1415.md) | Шлюз FTP–FTAM | 1993 | Historic | FTAM-шлюз | rfc-1415.md |
| 93 — Справочно | [1068](../rfc-1068.md) | Фоновая программа передачи файлов (BFTP) | 1988 | Unknown | BFTP, third-party-transfer | rfc-1068.md |
| 94 — Справочно | [5854](../rfc-5854.md) | Формат описания загрузок Metalink | 2010 | Proposed Standard | Metalink XML, хеши/зеркала | rfc-5854.md |
| 95 — Справочно | [6249](../rfc-6249.md) | Metalink/HTTP: зеркала и хеши | 2011 | Proposed Standard | Metalink HTTP-заголовки | rfc-6249.md |
| 96 — Справочно | [1738](../rfc-1738.md) | Унифицированные указатели ресурсов (URL) | 1994 | Proposed Standard | ftp-URL | rfc-1738.md |

Распределение статусов (92): Unknown 58, Proposed Standard 14, Informational 7,
Experimental 4, Internet Standard 4, Draft Standard 3, Historic 2.

Распределение меток: Критично 5 · Высокий 5 · Средний 4 · Низкий 4 · Справочно/история 74.

← [База знаний](README.md) · [Темы](topics.md) · [Связи RFC](relations.md) · [Требования](requirements.md)
