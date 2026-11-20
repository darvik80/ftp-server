# Правила подпроекта `ftp-server-cpp`

Роль агента и режим ответов задаёт корневой [AGENTS.md](../AGENTS.md). Здесь — только специфика C++.

## Стек и запуск

- C++23, Boost (`boost/1.91.0`) через Conan 2, сборка CMake через пресеты Conan.
- Установка зависимостей и сборка:
  `conan install . --build=missing -s compiler.cppstd=23` → `cmake --preset conan-release` → `cmake --build --preset conan-release`.
- Запуск: бинарник в `build/…` (у CMake с генератором VS — подкаталог `Release`). Порт/корень — `FTP_PORT`, `FTP_ROOT`.

## Структура

- `conanfile.py` — приложение (не библиотека), `cmake_layout`, `CMakeDeps` + `CMakeToolchain`, требование `boost/1.91.0`.
- `CMakeLists.txt` — `cxx_std_23`, `find_package(Boost REQUIRED CONFIG …)`, линковка `Boost::headers`/`Boost::system`.
- `src/main.cpp` — печатает версию стандарта, версию Boost и scope шага 1. Протокола нет.

## Конвенции

- Стандарт задаёт CMake (`cxx_std_23`), а не флаги в исходниках; `CMAKE_CXX_EXTENSIONS OFF`.
- Boost искать через CONFIG-конфиги пакета (`find_package(Boost … CONFIG)`), а не через legacy `FindBoost` — для новых Boost это путь к устаревшим правилам.
- Версии зависимостей — только в `conanfile.py`; в CMake их не дублировать.

## Особенности тулчейна на этой машине

- Профиль Conan по умолчанию — `msvc` 194 с `cppstd=14`, то есть для C++23 его надо переопределять (`-s compiler.cppstd=23`).
- MSVC 2022 покрывает C++23 частично; GCC 15.2 (MinGW, есть на машине) — ближе к полному C++23. Для кода, реально использующего C++23, профиль GCC предпочтительнее.
- `boost/1.92.0` уже выпущен upstream, но в conancenter максимум `1.91.0`: брать из реестра, пока нет цели собирать Boost из исходников.
- Boost собирается из исходников — первая установка занимает десятки минут, это нормально, а не зависшая сборка.

## Область реализации

Шаг 1 — [RFC 959](../docs/rfc-959.md), [RFC 1123](../docs/rfc-1123.md), [RFC 854](../docs/rfc-854.md); чек-лист — [../docs/kb/requirements.md](../docs/kb/requirements.md).
