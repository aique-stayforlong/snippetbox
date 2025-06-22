# Let's GO

Ejemplo práctico y pequeños apuntes tomados durante el seguimiento del libro [Let's GO de Alex Edwards](https://lets-go-further.alexedwards.net/) para
aprender programación Web en este lenguaje.

## Servemux

### Configuración por defecto

Como buena práctica, nunca se debería utilizar el *servemux* por defecto por los siguientes motivos:

- Es menos explícito y por tanto menos legible.

- Se puede acceder mediante una variable global, lo que permite ser accedido y modificado en cualquier punto del
- programa o por librerías de terceros. Dificulta el mantenimiento y supone un agujero en la seguridad.

### Conflicto de rutas

Ante la definición de dos rutas que encajan con una URL concreta, tendrá más preferencia aquella que sea más
restrictiva.

Esto permite definir las rutas sin ningún orden en particular. No obstante, en el diseño de las rutas de una API,
debería minimizarse o eliminar si es posible aquellas rutas que puedan solaparse.

La misma política rige aquellas rutas donde se declara el método HTTP para acceder a ellas. En caso de solapamiento, la
ruta más restrictiva es la que tiene mayor prioridad.

## Estructura de directorios

La carpeta denominada `internal` tiene un especial comportamiento en las aplicaciones GO. Tan sólo es accesible desde la
carpeta padre, sin importar la visibilidad que posean los ficheros que se encuentran dentro de ella.

## FileServer

Esta herramienta, además de permitir servir los ficheros estáticos de la web, proporciona otras utilidades como el
escapado de rutas URL.

En cuanto al rendimiento, esta herramienta servirá los ficheros directamente de la RAM, evitando los lentos accesos a
disco.

Si se quiere evitar que esta herramienta permita acceder a los directorios, será necesario incluir un fichero
`index.html` en blanco. De esta forma, al acceder a un directorio se servirá ese fichero en lugar de mostrar la
estructura de directorios.

Para crear un `index.html` en todos los subdirectorios de una ruta, se puede utilizar el comando:

```bash
$ find {ruta} -type d -exec touch {}/index.html \;
```