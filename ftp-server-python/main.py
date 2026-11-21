import os
from pathlib import Path
import tomllib


def main() -> None:
    config_path = Path(__file__).with_name("config.toml")
    with config_path.open("rb") as config_file:
        config = tomllib.load(config_file)

    port = int(os.environ.get("FTP_PORT", config["port"]))
    if not 1 <= port <= 65535:
        raise ValueError("FTP_PORT must be between 1 and 65535")
    root = Path(os.environ.get("FTP_ROOT", config["root"]))

    print("ftp-server-python stub")
    print(f"configured port {port}  root {root}")
    print(
        "step 1 (RFC 959): USER/PASS/QUIT, TYPE/MODE/STRU, PORT/PASV, "
        "RETR/STOR, CWD/PWD/LIST, 1xx..5xx"
    )


if __name__ == "__main__":
    main()
