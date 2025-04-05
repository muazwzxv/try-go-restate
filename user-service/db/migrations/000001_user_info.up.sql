create table user (
	id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	uuid VARCHAR(36) NOT NULL UNIQUE,
  name VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL, -- should be unique, enforce at code level

	status VARCHAR(36) NOT NULL,

	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by VARCHAR(36),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_by VARCHAR(36)
);
