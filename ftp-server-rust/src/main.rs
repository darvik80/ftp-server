use std::env;

fn main() {
    let port = env::var("FTP_PORT").unwrap_or_else(|_| "2121".into());
    let root = env::var("FTP_ROOT").unwrap_or_else(|_| "./ftp-data".into());

    println!("ftp-server-rust stub (edition 2024)");
    println!("listen :{port}  root {root}");
    println!("step 1 (RFC 959): USER/PASS/QUIT, TYPE/MODE/STRU, PORT/PASV, RETR/STOR, CWD/PWD/LIST, 1xx..5xx");
}
