# WebGoKit

WebGoKit es una herramienta GUI diseñada para facilitar el desarrollo de aplicaciones web con Go. Su enfoque principal es proporcionar una experiencia de desarrollo ágil y eficiente, especialmente orientada a la creación de módulos JavaScript o WebAssembly utilizando Go.

## Características principales

- Interfaz gráfica intuitiva desarrollada con Fyne
- Configuración simplificada para una rápida puesta en marcha
- Soporte para desarrollo de módulos JavaScript y WebAssembly con Go

## Motivación

WebGoKit nace de la necesidad de simplificar el proceso de desarrollo web en Go, eliminando la complejidad asociada a extensos archivos de configuración. Con WebGoKit, los desarrolladores pueden iniciar sus proyectos rápidamente, seleccionando las carpetas de entrada y salida directamente desde la interfaz gráfica.

Esta herramienta busca optimizar el flujo de trabajo, permitiendo a los desarrolladores enfocarse en la codificación y la lógica de negocio, en lugar de perder tiempo en configuraciones complejas.

## Características pendientes

- [x] creación de interfaz con [Fyne](https://fyne.io/)
- [x] agregar persistencia en la configuración con [bbolt](https://github.com/etcd-io/bbolt) ruta .git/webgokit.db
- [x] Soporte para hot-reloading con [fsnotify](https://github.com/fsnotify/fsnotify)
- [ ] Agregar navegador web automatizado para desarrollo con [rod](https://github.com/go-rod/rod)
- [ ] Creación de servidor web automático
- [ ] Compilación con [WebAssembly](https://es.wikipedia.org/wiki/WebAssembly) (WASM)
- [ ] Agregar compilación con [esbuild](https://esbuild.github.io/)
- [ ] omitir errores de lectura en config según estado de la app (primer uso) 
