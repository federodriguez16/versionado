# Control de Versiones de Productos

Una herramienta de línea de comandos (CLI) para mantener un listado actualizado de las versiones de productos instalados en diferentes clientes.

## Descripción

Este CLI se conecta a una base de datos donde se registran todos nuestros productos y sus respectivas versiones. Permite realizar operaciones básicas de CRUD (Crear, Leer, Actualizar, Borrar) sobre los clientes y gestionar las versiones de los productos que tienen asociados.

## Características Principales

- Crear Clientes: Añade nuevos clientes de forma individual o masivamente desde un archivo .csv.
- Listar Productos: Visualiza todos los productos y sus versiones para un cliente específico.
- Actualizar Versiones: Modifica la versión de uno o varios productos de un cliente.
- Borrar Clientes: Elimina un cliente y todos sus registros asociados.

## Tecnologías Utilizadas

- Go: Lenguaje de programación principal.
- GORM: ORM para la interacción con la base de datos.
- Cobra: Framework para la creación de la CLI.

## Instalacion y Uso

### Prerequisitos

- Para uso en producción: Se requiere una base de datos MariaDB accesible.
- Para pruebas (entorno de desarrollo): Se necesita tener Docker y Docker Compose instalados.

### Configuracion

Antes de ejecutar la aplicación, es necesario configurar las siguientes variables de entorno para la conexión a la base de datos:

``` bash
export VERSIONADO_DB_URL="tu-direccion-db:puerto"
export VERSIONADO_DB_USER="tu_usuario_db"
export VERSIONADO_DB_PASSWORD="tu_contraseña_db"
export VERSIONADO_DB_NAME="nombre_de_la_db"
```

### Despliegue

#### Entorno de Prueba (con Docker)

Este método es ideal para probar la aplicación de forma aislada.

1. Clona el repositorio:

``` bash
git clone git@github.com:federodriguez16/versionado.git
```

2. Levanta los servicios con Docker Compose:
```bash
docker-compose up -d
```

3. Accede al contenedor de la aplicación:
``` bash
docker-compose exec versionado sh
```

4. Una vez dentro, puedes ejecutar el CLI:
``` bash
./versionado --help
```

Contiene un archivo llamado `test.csv` para cargar algunos clientes iniciales.

#### Entorno de Produccion

1. Descarga la última versión del CLI desde la sección de Releases de este repositorio.

2. Descomprime el archivo descargado.

``` bash
tar -xzvf versionado_linux_amd64.tar.gz
```
3. Mueve el binario a una ruta incluida en tu PATH para poder ejecutarlo desde cualquier lugar, por ejemplo:
``` bash
sudo mv versionado /usr/local/bin/
```

### Uso Básico

Toda la ayuda sobre los comandos disponibles y sus opciones se encuentra integrada en el propio CLI. Para empezar, ejecuta:
``` bash
versionado --help
```