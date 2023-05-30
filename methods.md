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