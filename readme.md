# Instrucciones para ejecutar los microservicios

Este proyecto consta de dos microservicios: `accounts-service` y `transactions-service`, cada uno ejecutándose en su propio contenedor Docker. Ambos microservicios dependen de una instancia de MongoDB como base de datos.

## Requisitos previos

Asegúrate de tener instalado lo siguiente en tu sistema:

- Docker: (https://docs.docker.com/get-docker/)
- Git: (https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Pasos para ejecutar los microservicios

1. Clona este repositorio en tu máquina local:

   ```bash
   git clone https://github.com/lucaaramos/Tach.git

2. Navega hasta el repositorio

cd Tach

3. Crea un archivo .env en el directorio de accounts y transactions, y define las variables de entorno necesarias.

# Configuración para el servicio de cuentas
MONGO_URL=mongodb://mongodb:27017

4. Ejecuta el siguiente comando para construir las imagenes de los microservicios y levantar los contenedores

-- docker-compose up

5. Una vez que los contenedores estén en funcionamiento, podrás acceder a los microservicios utilizando las siguientes URLs:

accounts-service: http://localhost:8080
transactions-service: http://localhost:8081

6. Detener los microservicios:
Para detener los microservicios y eliminar los contenedores, puedes ejecutar el siguiente comando en el directorio del proyecto:

-- docker-compose down

