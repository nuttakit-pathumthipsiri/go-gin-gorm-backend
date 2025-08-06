# API Testing Guide

## Topics API

### 1. Create Topic

```bash
POST /topics
Content-Type: application/json

{
  "name": "ยา"
}
```

### 2. Get All Topics

```bash
GET /topics
```

### 3. Get Topic by ID

```bash
GET /topics/1
```

### 4. Update Topic

#### Update with all fields (optional)
```bash
PUT /topics/1
Content-Type: application/json

{
  "name": "ยาแก้ปวด",
  "order": 2
}
```

#### Update with only name (order remains unchanged)
```bash
PUT /topics/1
Content-Type: application/json

{
  "name": "ยาแก้ปวด"
}
```

#### Update with only order (name remains unchanged)
```bash
PUT /topics/1
Content-Type: application/json

{
  "order": 2
}
```

#### Update with duplicate order (automatic resolution)
```bash
PUT /topics/1
Content-Type: application/json

{
  "order": 1
}
```

**Response when duplicate order is resolved:**
```json
{
  "message": "Order was automatically adjusted due to conflict",
  "original_order": 1,
  "resolved_order": 3,
  "topic": {
    "id": 1,
    "name": "ยาแก้ปวด",
    "order": 3,
    "created_by": "admin",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_by": "admin",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 5. Delete Topic

```bash
DELETE /topics/1
```

## Topic Details API

### 1. Create Topic Detail

```bash
POST /topics/1/details
Content-Type: application/json

{
  "name": "ยาแก้ปวด"
}
```

### 2. Get All Details by Topic ID

```bash
GET /topics/1/details
```

### 3. Get Detail by ID

```bash
GET /details/1
```

### 4. Update Topic Detail

#### Update with all fields (optional)
```bash
PUT /details/1
Content-Type: application/json

{
  "name": "ยาแก้ปวดหัว",
  "order": 2
}
```

#### Update with only name (order remains unchanged)
```bash
PUT /details/1
Content-Type: application/json

{
  "name": "ยาแก้ปวดหัว"
}
```

#### Update with only order (name remains unchanged)
```bash
PUT /details/1
Content-Type: application/json

{
  "order": 2
}
```

#### Update with duplicate order (automatic resolution)
```bash
PUT /details/1
Content-Type: application/json

{
  "order": 1
}
```

**Response when duplicate order is resolved:**
```json
{
  "message": "Order was automatically adjusted due to conflict",
  "original_order": 1,
  "resolved_order": 3,
  "detail": {
    "id": 1,
    "topic_id": 1,
    "name": "ยาแก้ปวดหัว",
    "order": 3,
    "created_by": "admin",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_by": "admin",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 5. Delete Topic Detail

```bash
DELETE /details/1
```

## Notes

- `created_at` and `updated_at` fields are automatically managed by the database
- These fields will be included in the response but should not be sent in requests
- The database will automatically set `created_at` when creating new records
- The database will automatically update `updated_at` when updating existing records
- **New Feature**: `name` and `order` fields in update requests are now optional
- **New Feature**: When updating with a duplicate order, the system automatically resolves it by finding the next available order number
- **New Feature**: The response includes information about order resolution when conflicts occur
