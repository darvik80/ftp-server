package dev.mix.ftp.server;

import org.springframework.boot.context.properties.ConfigurationProperties;

import java.nio.file.Path;

@ConfigurationProperties(prefix = "ftp")
public record FtpProperties(int port, Path root) {
}
