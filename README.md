# Cloud Infrastructure Management Solution in Golang

This project provides a Golang-based solution to manage cloud infrastructure operations, focusing on AWS EKS as the intended cloud platform. The solution adheres to best practices in coding, testing, and scalability, and includes error handling and logging.

## Features

1. **Creating Virtual Machines:**
   - Create virtual machines of specified capacity.

2. **Kubernetes Cluster Management:**
   - Provision a Kubernetes cluster of a specific flavor and version.
   - Upgrade an existing Kubernetes cluster to a new version.

3. **Public IP Management:**
   - Re-assign a public IP from one network resource to another.

4. **Firewall Management:**
   - Bring an existing resource under a firewall.

5. **SQL Server Management:**
   - Upgrade an SQL server from one version to another.
   - Migrate an SQL server from Region A to Region B.

6. **Virtual Machine Migration:**
   - Migrate a VM from one region to another.

7. **Infrastructure Migration:**
   - Migrate infrastructure components across regions.

## Implementation Highlights

### Infrastructure Migration
- **Migrate SQL Server Across Regions**:
  The system migrates an SQL server from Region A to Region B, ensuring data integrity and minimal downtime.

- **Migrate VM Across Regions**:
  The solution supports VM migration across regions, handling associated resource dependencies.

### Error Handling and Edge Cases
- Prevent duplicate resource creation by checking existing resource identifiers.
- Validate update requests to prevent invalid operations.
- Handle deletion requests for non-existent resources gracefully with appropriate error messages.

### Logging and Monitoring
- Logs all significant events, such as resource creation, updates, migrations, and errors.
- Uses structured logging for better observability and debugging.

### API Endpoints

#### Create Virtual Machine
```http
POST /api/v1/vm/create
{
  "name": "vm-1",
  "capacity": "t2.micro",
  "region": "us-west-1"
}
```

#### Create Kubernetes Cluster
```http
POST /api/v1/k8s/create
{
  "name": "cluster-1",
  "version": "1.32",
  "flavor": "eks"
}
```

#### Migrate SQL Server
```http
POST /api/v1/sql/migrate
{
  "server_id": "sql-123",
  "source_region": "us-east-1",
  "destination_region": "us-west-2"
}
```

#### Migrate VM
```http
POST /api/v1/vm/migrate
{
  "vm_id": "vm-123",
  "source_region": "us-east-1",
  "destination_region": "us-west-2"
}
```

## Example Usage

1. **Create a VM:**
   ```bash
   curl -X POST http://localhost:8080/api/v1/vm/create -d '{"name":"vm-1", "capacity":"t2.micro", "region":"us-west-1"}' -H "Content-Type: application/json"
   ```

2. **Migrate SQL Server:**
   ```bash
   curl -X POST http://localhost:8080/api/v1/sql/migrate -d '{"server_id":"sql-123", "source_region":"us-east-1", "destination_region":"us-west-2"}' -H "Content-Type: application/json"
   ```

3. **Migrate VM:**
   ```bash
   curl -X POST http://localhost:8080/api/v1/vm/migrate -d '{"vm_id":"vm-123", "source_region":"us-east-1", "destination_region":"us-west-2"}' -H "Content-Type: application/json"
   ```

## Testing

### Unit Tests
- Validate core functionalities, such as resource creation, updates, and migrations.
- Example:

  ```go
  func TestCreateVM(t *testing.T) {
      // Test logic for creating a VM
  }
  ```

### Integration Tests
- Test end-to-end API interactions.
- Example:

  ```go
  func TestMigrateSQLServer(t *testing.T) {
      // Test logic for migrating an SQL server
  }
  ```

## Logging and Error Handling
- Use structured logging for better traceability:
  ```go
  log.WithFields(log.Fields{
      "resource_id": "vm-123",
      "operation": "create",
  }).Info("Resource creation successful")
  ```

- Return appropriate HTTP status codes for errors:
  ```json
  {
      "error": "Resource not found",
      "code": 404
  }
  ```