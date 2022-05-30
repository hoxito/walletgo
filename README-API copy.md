<a name="top"></a>
# StatsGo Service v0.1.0

Microservicio de Estadisticas

- [RabbitMQ_POST](#rabbitmq_post)
	- [Invalidar Token](#invalidar-token)
	
- [Estadisticas](#Estadisticas)
	- [Listar Ordenes](#Ordenes)
	- [Listar Usuarios](#Usuarios)
	- [Listar Productos](#Productos)
   
	


# <a name='rabbitmq_post'></a> RabbitMQ_POST

## <a name='invalidar-token'></a> Invalidar Token
[Back to top](#top)

<p>AuthService enviá un broadcast a todos los usuarios cuando un token ha sido invalidado. Los clientes deben eliminar de sus caches las sesiones invalidadas.</p>

	FANOUT auth/fanout





### Success Response

Mensaje

```
{
   "type": "logout",
   "message": "{Token revocado}"
}
```


# <a name='Estadisticas'></a> Estadisticas

## <a name='Ordenes'></a> Listar Ordenes
[Back to top](#top)

<p>Devuelve una lista segun los parametros de busqueda.</p>

	GET /v1/stats/{Query}



### Examples

{url}/v1/stats/?select=[estado,montomin,montomax,fechadesde,fechahasta]&orderby=fechadesde&items=1&pag=1

Header Autorización

```
Authorization=bearer {token}
```


### Success Response

Respuesta

```
HTTP/1.1 200 OK
  [{
       "id": "{Id Orden}",
       "estado": "{Estado de la orden}",
       "montomin": "{monto minimo total de la orden}",
       "montomax: "{monto maximo total de la orden}",
       "fechadesde": {fecha de creacion de la orden}
       "fechahasta": {fecha de finalizacion de la orden}

    }, ...]
```


### Error Response

401 Unauthorized

```
HTTP/1.1 401 Unauthorized
{
   "error" : "Unauthorized"
}
```
400 Bad Request

```
HTTP/1.1 400 Bad Request
{
   "messages" : [
     {
       "path" : "{Nombre de la propiedad}",
       "message" : "{Motivo del error}"
     },
     ...
  ]
}
```
500 Server Error

```
HTTP/1.1 500 Internal Server Error
{
   "error" : "Not Found"
}
```