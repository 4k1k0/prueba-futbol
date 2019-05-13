# Calculador de salarios y bonos

Realizado por Rodrigo Patino utilizando Go en la versión

`go version go1.12.5 linux/amd64`

Construir con

`$ make`

Instalar con

`$ make install`

## Parámetros

Los parámetros aceptados son

 - equipo(string)
 - niveles(string)

Por defecto el programa utiliza

```
--equipo="resuelve.json"
--niveles="niveles.json"
```

## Output

El programa genera un archivo en `/tmp/` utilizando el mismo nombre de archivo del parametro `equipo` con la extensión `.log`

## Ejemplo

```
$ calculador --equipo="/home/wako/Documentos/Data/el-poderoso-atlas.json" --niveles="hardcore.json"
Archivo generado en /tmp/el-poderoso-atlas.log
```