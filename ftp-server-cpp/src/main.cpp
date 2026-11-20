#include <boost/asio.hpp>
#include <boost/version.hpp>

#include <iostream>

int main() {
    std::cout << "ftp-server-cpp stub (C++" << __cplusplus / 100 % 100 << ")"
              << " boost " << BOOST_LIB_VERSION << "\n";
    std::cout << "Step 1 (RFC 959): control connection on port 21, USER/PASS/QUIT, TYPE, PORT/PASV, RETR/STOR\n";
    return 0;
}
