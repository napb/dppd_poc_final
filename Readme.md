# Golang - Prometheus - Grafana PoC

Esta PoC se basa en generar metricas de tipo `count` para los endpoints de esta aplicacion Golang. 
Estas metricas seran consumidas por [Prometheus](https://prometheus.io) y mostradas en [Grafana](https://grafana.com). 
Las metricas de esta aplicacion, seran analizadas y procesadas por un modelo de prediccion se series 
de tiempo en Python con [SciKit-Learn](https://scikit-learn.org/stable/). Mostraran la serie de tiempo real,
la serie de tiempo ajustada y la prediccion de los valores a traves de Grafana. 

### Ejecucion de la aplicacion

Para ejecutar esta aplicacion, se debe ejecutar a traves de un terminal el comando
``` 
go run .
```

### Flujo de funcionamiento aplicacion

- Esta aplicacion expone el endpoint `/metrics` a traves del cual Prometheus hara scrapping a traves del endpoint `/metrics`
- Al invocar los endpoints estos generaran contadores en el formato de extraccion de Prometheus

  - Endpoint de simulacion de venta de TV
    ```
    curl --location --request GET 'localhost:8080/venta/tv'
    ```

  - Endpoint de simulacion de venta de computadores
    ```
    curl --location --request GET 'localhost:8080/venta/computador'
    ```

### Prometheus
Prometheus almacenara metricas de tipo `count` y `gauge` generadas por esta aplicacion. Para levantar prometheus