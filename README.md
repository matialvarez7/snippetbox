Proyecto Snippet box del libro Let's Go de Alex Edwars.
-------------------------------------------------------------------
Desarrollo del proyecto del libro Let's Go de Alex Edwards con la finalidad de incorporar conocimientos de desarrollo de aplicaciones web con el lenguaje Go.
En caso de clonarlo para poder utilizarlo deberán generar en su local los certificados TLS. Para esto deberán saber donde está instalado en su computadora el código fuente de Go para poder utilizar la herramienta generate_cert.go. 
Se recomienda generar un directorio para contener los certificados. En caso de haber instalado Go en linux el comando para generar debería ser:

 go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

