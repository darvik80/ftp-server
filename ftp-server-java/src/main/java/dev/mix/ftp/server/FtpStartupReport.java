package dev.mix.ftp.server;

import org.springframework.boot.ApplicationRunner;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class FtpStartupReport {

    @Bean
    ApplicationRunner logStartup(FtpProperties props) {
        return args -> System.out.printf(
                """
                        ftp-server-java stub (Spring Boot)
                        listen :%d  root %s
                        step 1 (RFC 959): USER/PASS/QUIT, TYPE/MODE/STRU, PORT/PASV, RETR/STOR, CWD/PWD/LIST, 1xx..5xx%n""",
                props.port(), props.root());
    }
}
