# Cloud-Native Server Simulation

This project demonstrates a basic HTTP server designed with cloud-native best practices, including readiness and liveness probes, and graceful shutdown capabilities. The application is ideal for deployment in environments such as Kubernetes, where health checks and smooth shutdowns are critical.

---

## **Features**

- **Readiness Probe (`/ready`)**: Indicates whether the server is ready to handle incoming requests. Returns a `503 Service Unavailable` status during shutdown to signal that the server should stop receiving traffic.
- **Liveness Probe (`/health`)**: Confirms if the server process is alive. Continues to return `200 OK` unless the server encounters a fatal error.
- **Graceful Shutdown**:
  - Listens for termination signals (`SIGTERM`, `SIGINT`, `SIGQUIT`).
  - Simulates cleanup tasks (e.g., closing database connections, flushing logs, or releasing resources).
  - Ensures cleanup tasks complete before the server shuts down.

---

## **Endpoints**

1. **Simulation Page (`/`)**:
   - Provides a simple HTML page to demonstrate the readiness and liveness probes.

2. **Readiness Probe (`/ready`)**:
   - Checks if the server is ready to handle requests.
   - Returns:
     - `200 OK` when ready.
     - `503 Service Unavailable` during shutdown to signal that the server is no longer ready to serve traffic.

3. **Liveness Probe (`/health`)**:
   - Confirms if the server process is alive and functioning.
   - Returns:
     - `200 OK` unless the server encounters a fatal issue.

---

## **Why Simulate Delay During Shutdown?**

### **Purpose of Delayed Readiness**
During server shutdown, the readiness probe immediately starts returning `503 Service Unavailable` to signal load balancers or orchestrators (like Kubernetes) to stop routing new traffic to the server. However, the liveness probe continues to return `200 OK` while the server performs cleanup tasks.

This behavior ensures:
- **Active Traffic Draining**: Existing connections are gracefully completed without abrupt termination.
- **Proper Resource Cleanup**: Tasks like closing database connections, persisting in-memory data, and logging are allowed to finish without rushing.
- **Clear Separation of Health States**:
  - **Readiness**: Indicates whether the server is prepared to handle new requests.
  - **Liveness**: Confirms that the server is still running during the shutdown process.

By simulating a delay (`8 seconds` in this case), we mimic real-world scenarios where cleanup tasks take time to complete. This delay helps illustrate the distinction between a "healthy but not ready" state versus an "unhealthy" state, which is crucial in cloud-native architectures.

---

## **Requirements**

- **Go**: Version 1.20 or higher.
- **Make**: For building and running the application.

---

## **Build and Run**

## **Makefile Commands**

### **Build**

Compiles the Go application into an executable server.

```bash
make build
```

---

### **Run**

Starts the server on [http://localhost:8080](http://localhost:8080).

```bash
make run
```

---

### **Clean**

Removes the generated executable.

```bash
make clean
```

---

## **Example Logs During Shutdown**

When you press `Ctrl+C` to stop the server, you'll observe the following sequence of events in the logs:

### **Readiness Probe Returns 503**

The `/ready` endpoint signals to orchestrators and load balancers that the server is no longer ready to accept traffic:

```plaintext
Server is shutting down.
Performing cleanup tasks...
```

---

### **Cleanup Task Completion**

The cleanup tasks (e.g., closing resources or connections) complete successfully:

```plaintext
Cleanup completed.
```

---

### **Server Shutdown**

The server shuts down gracefully after completing the cleanup:

```plaintext
Shutting down gracefully...
```
