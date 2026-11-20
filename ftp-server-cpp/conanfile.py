from conan import ConanFile
from conan.tools.cmake import cmake_layout


class FtpServerCpp(ConanFile):
    name = "ftp_server_cpp"
    version = "0.1.0"
    package_type = "application"
    settings = "os", "compiler", "build_type", "arch"
    generators = "CMakeDeps", "CMakeToolchain"

    def requirements(self):
        # upstream релиз уже 1.92.0, но в conancenter максимум 1.91.0 — берём оттуда
        self.requires("boost/1.91.0")

    def layout(self):
        cmake_layout(self)
