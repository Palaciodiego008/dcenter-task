1. **CreateClient:**
   - Descripción: Crea un nuevo cliente y lo guarda en la base de datos.
   - Método HTTP: POST
   - Ruta: "/clients"
   - Parámetros de entrada:
     - Cuerpo de la solicitud: Objeto JSON que contiene los datos del cliente.
   - Respuesta HTTP:
     - Código 201 (Created) si el cliente se crea correctamente.
     - Código 400 (Bad Request) si hay un error al parsear los datos del cliente.
     - Código 500 (Internal Server Error) si hay un error al conectar a la base de datos o al crear la tabla "clients".

2. **GetClientByID:**
   - Descripción: Obtiene los datos de un cliente específico por su ID.
   - Método HTTP: GET
   - Ruta: "/clients/{id}"
   - Parámetros de entrada:
     - ID: ID del cliente.
   - Respuesta HTTP:
     - Código 200 (OK) si se obtienen los datos del cliente correctamente.
     - Código 404 (Not Found) si no se encuentra el cliente con el ID especificado.
     - Código 500 (Internal Server Error) si hay un error al conectar a la base de datos.

3. **UpdateClient:**
   - Descripción: Actualiza los datos de un cliente existente en la base de datos.
   - Método HTTP: PUT
   - Ruta: "/clients/{id}"
   - Parámetros de entrada:
     - ID: ID del cliente.
     - Cuerpo de la solicitud: Objeto JSON que contiene los datos actualizados del cliente.
   - Respuesta HTTP:
     - Código 200 (OK) si se actualizan los datos del cliente correctamente.
     - Código 400 (Bad Request) si hay un error al parsear los datos actualizados del cliente.
     - Código 404 (Not Found) si no se encuentra el cliente con el ID especificado.
     - Código 500 (Internal Server Error) si hay un error al conectar a la base de datos.

4. **DeleteClient:**
   - Descripción: Elimina un cliente existente de la base de datos.
   - Método HTTP: DELETE
   - Ruta: "/clients/{id}"
   - Parámetros de entrada:
     - ID: ID del cliente.
   - Respuesta HTTP:
     - Código 204 (No Content) si se elimina el cliente correctamente.
     - Código 404 (Not Found) si no se encuentra el cliente con el ID especificado.
     - Código 500 (Internal Server Error) si hay un error al conectar a la base de datos.

5. **CreateTruckDelivery:**
   - Descripción: Crea un nuevo plan de entrega de logística terrestre y lo guarda en la base de datos.
   - Método HTTP: POST
   - Ruta: "/truck-deliveries"
   - Parámetros de entrada:
     - Cuerpo de la solicitud: Objeto JSON que contiene los datos del plan de entrega de logística terrestre.
   - Respuesta HTTP:
     - Código 201 (Created) si el plan de entrega se crea correctamente.
     - Código 400 (Bad Request) si hay un error en la validación de los datos o en el formato de la placa del vehículo o el número de guía.
     - Código 500 (Internal Server Error) si hay un



     **CalcularDescuento:**
   - Descripción: Calcula el descuento para el precio de envío de un plan de entrega de logística terrestre en función de la cantidad de productos.
   - Parámetros de entrada:
     - truckDelivery: Objeto que representa el plan de entrega de logística terrestre.
   - Parámetros de salida:
     - N/A (el descuento se calcula directamente en el objeto truckDelivery).
   - Funcionalidad:
     - Verifica si la cantidad de productos en el plan de entrega es mayor a 10.
     - Si la condición es verdadera, calcula el descuento como el 5% del precio de envío.
     - Si la condición es falsa, el precio de envío se mantiene sin cambios.
     - Actualiza el valor del precio con descuento en la propiedad truckDelivery.DiscountedPrice del objeto truckDelivery.

  **Relación entre Clients y Deliveries:**

  En el contexto del sistema de gestión logística, existe una relación entre los clientes (Clients) y los planes de entrega (Deliveries). Esta relación se establece para asociar cada plan de entrega con el cliente correspondiente.

  Descripción de la relación:
  - Un cliente puede tener varios planes de entrega.
  - Cada plan de entrega está vinculado a un único cliente.

  Esta relación se representa a través de una relación de uno a muchos, donde un cliente puede tener múltiples planes de entrega, pero cada plan de entrega pertenece a un solo cliente.

  La relación entre Clients y Deliveries se establece mediante una clave foránea (foreign key) en la tabla de Deliveries, que hace referencia a la clave primaria del cliente en la tabla de Clients. Esta clave foránea permite vincular cada plan de entrega con el cliente al que pertenece.

  En el modelo de base de datos, la tabla de Clients contendría la información del cliente, como su ID, nombre, dirección, etc. Mientras que la tabla de Deliveries almacenaría los detalles de los planes de entrega, como el tipo de producto, cantidad, fechas, precios, etc., y también incluiría la clave foránea que referencia al ID del cliente al que pertenece el plan de entrega.

  En resumen, la relación entre Clients y Deliveries permite establecer la conexión entre los clientes y sus respectivos planes de entrega, lo que facilita la gestión y seguimiento de los envíos y servicios logísticos asociados a cada cliente.