-- Initial schema for IIOTBackend
-- This creates all the necessary tables for the EdgeX-compatible IoT backend

-- Core Metadata Tables
CREATE TABLE IF NOT EXISTS device_services (
    id UUID PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    base_address VARCHAR(500) NOT NULL,
    admin_state VARCHAR(50) DEFAULT 'UNLOCKED',
    labels JSONB DEFAULT '[]',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS device_profiles (
    id UUID PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    manufacturer VARCHAR(255),
    model VARCHAR(255),
    labels JSONB DEFAULT '[]',
    device_resources JSONB DEFAULT '[]',
    device_commands JSONB DEFAULT '[]',
    core_commands JSONB DEFAULT '[]',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS devices (
    id UUID PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    admin_state VARCHAR(50) DEFAULT 'UNLOCKED',
    operating_state VARCHAR(50) DEFAULT 'UP',
    protocols JSONB DEFAULT '{}',
    labels JSONB DEFAULT '[]',
    location VARCHAR(255),
    service_name VARCHAR(255) NOT NULL,
    profile_name VARCHAR(255) NOT NULL,
    auto_events JSONB DEFAULT '[]',
    tags JSONB DEFAULT '{}',
    properties JSONB DEFAULT '{}',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (service_name) REFERENCES device_services(name) ON DELETE CASCADE,
    FOREIGN KEY (profile_name) REFERENCES device_profiles(name) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS provision_watchers (
    id UUID PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    labels JSONB DEFAULT '[]',
    identifiers JSONB DEFAULT '{}',
    blocking_identifiers JSONB DEFAULT '{}',
    profile_name VARCHAR(255) NOT NULL,
    service_name VARCHAR(255) NOT NULL,
    admin_state VARCHAR(50) DEFAULT 'UNLOCKED',
    auto_events JSONB DEFAULT '[]',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (service_name) REFERENCES device_services(name) ON DELETE CASCADE,
    FOREIGN KEY (profile_name) REFERENCES device_profiles(name) ON DELETE CASCADE
);

-- Core Data Tables
CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    device_name VARCHAR(255) NOT NULL,
    profile_name VARCHAR(255) NOT NULL,
    source_name VARCHAR(255) NOT NULL,
    origin BIGINT NOT NULL,
    tags JSONB DEFAULT '{}',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS readings (
    id UUID PRIMARY KEY,
    event_id UUID NOT NULL,
    device_name VARCHAR(255) NOT NULL,
    resource_name VARCHAR(255) NOT NULL,
    profile_name VARCHAR(255) NOT NULL,
    value_type VARCHAR(50) NOT NULL,
    value TEXT,
    binary_value BYTEA,
    media_type VARCHAR(255),
    units VARCHAR(100),
    tags JSONB DEFAULT '{}',
    origin BIGINT NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);

-- Support Scheduler Tables
CREATE TABLE IF NOT EXISTS intervals (
    id UUID PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    start_time VARCHAR(255) NOT NULL,
    end_time VARCHAR(255),
    interval_time VARCHAR(255) NOT NULL,
    run_once BOOLEAN DEFAULT FALSE,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS interval_actions (
    id UUID PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    interval_name VARCHAR(255) NOT NULL,
    protocol VARCHAR(50) NOT NULL,
    host VARCHAR(255) NOT NULL,
    port INTEGER NOT NULL,
    path VARCHAR(500),
    parameters TEXT,
    http_method VARCHAR(10),
    address VARCHAR(500),
    publisher VARCHAR(255),
    target VARCHAR(255),
    user_name VARCHAR(255),
    password VARCHAR(255),
    topic VARCHAR(255),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (interval_name) REFERENCES intervals(name) ON DELETE CASCADE
);

-- Support Notifications Tables
CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY,
    slug VARCHAR(255) UNIQUE NOT NULL,
    sender VARCHAR(255) NOT NULL,
    category VARCHAR(100) NOT NULL,
    severity VARCHAR(50) NOT NULL,
    content TEXT NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'NEW',
    labels JSONB DEFAULT '[]',
    content_type VARCHAR(100) DEFAULT 'text/plain',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subscriptions (
    id UUID PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    receiver VARCHAR(255) NOT NULL,
    subscribed_categories JSONB DEFAULT '[]',
    subscribed_labels JSONB DEFAULT '[]',
    channels JSONB DEFAULT '[]',
    resend_limit INTEGER DEFAULT 0,
    resend_interval VARCHAR(255),
    admin_state VARCHAR(50) DEFAULT 'UNLOCKED',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Support Rules Engine Tables
CREATE TABLE IF NOT EXISTS rules (
    id UUID PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    enabled BOOLEAN DEFAULT TRUE,
    priority INTEGER DEFAULT 0,
    conditions JSONB DEFAULT '[]',
    actions JSONB DEFAULT '[]',
    tags JSONB DEFAULT '{}',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS pipelines (
    id UUID PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    enabled BOOLEAN DEFAULT TRUE,
    functions JSONB DEFAULT '[]',
    triggers JSONB DEFAULT '[]',
    targets JSONB DEFAULT '[]',
    tags JSONB DEFAULT '{}',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- System Management Tables
CREATE TABLE IF NOT EXISTS services (
    service_id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    version VARCHAR(100),
    status VARCHAR(50) DEFAULT 'UP',
    host VARCHAR(255) NOT NULL,
    port INTEGER NOT NULL,
    health_check VARCHAR(500),
    tags JSONB DEFAULT '[]',
    meta JSONB DEFAULT '{}',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_seen TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS service_configs (
    service_id UUID,
    config JSONB NOT NULL,
    version VARCHAR(100) DEFAULT '1.0',
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (service_id),
    FOREIGN KEY (service_id) REFERENCES services(service_id) ON DELETE CASCADE
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_devices_service_name ON devices(service_name);
CREATE INDEX IF NOT EXISTS idx_devices_profile_name ON devices(profile_name);
CREATE INDEX IF NOT EXISTS idx_devices_admin_state ON devices(admin_state);
CREATE INDEX IF NOT EXISTS idx_devices_operating_state ON devices(operating_state);

CREATE INDEX IF NOT EXISTS idx_events_device_name ON events(device_name);
CREATE INDEX IF NOT EXISTS idx_events_profile_name ON events(profile_name);
CREATE INDEX IF NOT EXISTS idx_events_created ON events(created);
CREATE INDEX IF NOT EXISTS idx_events_origin ON events(origin);

CREATE INDEX IF NOT EXISTS idx_readings_event_id ON readings(event_id);
CREATE INDEX IF NOT EXISTS idx_readings_device_name ON readings(device_name);
CREATE INDEX IF NOT EXISTS idx_readings_resource_name ON readings(resource_name);
CREATE INDEX IF NOT EXISTS idx_readings_created ON readings(created);
CREATE INDEX IF NOT EXISTS idx_readings_value_type ON readings(value_type);

CREATE INDEX IF NOT EXISTS idx_notifications_category ON notifications(category);
CREATE INDEX IF NOT EXISTS idx_notifications_severity ON notifications(severity);
CREATE INDEX IF NOT EXISTS idx_notifications_status ON notifications(status);
CREATE INDEX IF NOT EXISTS idx_notifications_created ON notifications(created);

CREATE INDEX IF NOT EXISTS idx_services_name ON services(name);
CREATE INDEX IF NOT EXISTS idx_services_status ON services(status);
CREATE INDEX IF NOT EXISTS idx_services_last_seen ON services(last_seen);

CREATE INDEX IF NOT EXISTS idx_rules_enabled ON rules(enabled);
CREATE INDEX IF NOT EXISTS idx_rules_priority ON rules(priority);

CREATE INDEX IF NOT EXISTS idx_pipelines_enabled ON pipelines(enabled);

-- Insert default data
INSERT INTO device_services (id, name, description, base_address, admin_state) VALUES 
('550e8400-e29b-41d4-a716-446655440000', 'device-virtual', 'Virtual device service for testing', 'http://localhost:59900', 'UNLOCKED')
ON CONFLICT (name) DO NOTHING;

INSERT INTO device_profiles (id, name, description, manufacturer, model, device_resources, device_commands, core_commands) VALUES 
('550e8400-e29b-41d4-a716-446655440001', 'Random-Integer-Device', 'Example device profile for random integer generation', 'IOTech', 'RND-INT-01', 
'[{"name":"RandomValue","description":"Random integer value","tag":"","properties":{"valueType":"Int32","readWrite":"R","units":"","minimum":"","maximum":"","defaultValue":"","mask":"","shift":"","scale":"","offset":"","base":"","assertion":"","mediaType":""},"attributes":{"min":"1","max":"100"}}]',
'[{"name":"RandomValue","isHidden":false,"readWrite":"R","resourceOperations":[{"deviceResource":"RandomValue","defaultValue":"","mappings":{}}]}]',
'[{"name":"RandomValue","get":true,"set":false,"path":"/api/v2/device/{deviceId}/RandomValue","url":"","parameters":[{"resourceName":"RandomValue","valueType":"Int32"}]}]')
ON CONFLICT (name) DO NOTHING;

-- Schema version tracking handled by migration runner
