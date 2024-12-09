# Simulador de Estacionamiento

Este proyecto es un simulador interactivo de estacionamiento desarrollado en Go utilizando la biblioteca Fyne.

## Estructura del Proyecto

### 📁 assets/
Esta carpeta contiene todos los recursos gráficos utilizados en la aplicación, como imágenes y archivos de estilo. Se separa de la lógica del código para mantener la claridad y la modularidad en el proyecto.

### 📁 src/
Esta carpeta es el núcleo del código fuente del simulador. Dentro de ella, se encuentran tres subcarpetas que organizan las distintas partes de la aplicación:

#### 📁 models/
Esta carpeta alberga los modelos de datos, que encapsulan la lógica del negocio y las entidades del sistema. 

- `Observer.go`: Contiene la lógica de actualización del observador y cómo se manejan los espacios del estacionamiento.
- `Parking_lot.go`: Maneja la lógica de los espacios ocupados del estacionamiento y cómo los carros salen del estacionamiento. También se encarga de la suscripción y notificación al observador.
- `Vehicle.go`: Define la estructura del carro y los datos necesarios para las demás funciones.

#### 📁 scenes/
Implementa las escenas de la aplicación, que son responsables de la presentación y disposición de los elementos en la interfaz gráfica.

- `Simulator.go`: Contiene la lógica del simulador junto con la lógica de la interfaz hecha con la librería Fyne.

#### 📁 views/
Contiene los componentes de la interfaz de usuario que permiten a los observadores interactuar con el sistema.

- `Interface.go`: Maneja la lógica de la interfaz respecto a la actualización del observador cuando los espacios están ocupados o vacíos.
- `Main.go`: Inicia la animación del simulador y maneja los 100 hilos conforme se van generando los carros.

## Cómo Ejecutar

Para compilar y ejecutar el proyecto:

```sh
go run main.go
