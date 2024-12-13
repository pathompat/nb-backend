-- create table user
CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    store_name VARCHAR(100) NOT NULL ,
    tier_id SMALLINT,
    role VARCHAR(20) NOT NULL DEFAULT 'CUSTOMER',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- create table schools
CREATE TABLE schools (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(11) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);

-- create table quotations
CREATE TABLE quotations (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    school_id INT NOT NULL,
    store_name VARCHAR(100) NOT NULL,
    school_name VARCHAR(100) NOT NULL,
    school_address VARCHAR(255) NOT NULL,
    school_telephone VARCHAR(11) NOT NULL,
    appointment_at TIMESTAMP,
    duedate_at TIMESTAMP NOT NULL,
    remark TEXT,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user_quotation FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_school FOREIGN KEY (school_id) REFERENCES schools(id)
);

-- create table quotation_items
CREATE TABLE quotation_items (
    id SERIAL NOT NULL  PRIMARY KEY,
    quotation_id INT NOT NULL,
    category VARCHAR(100) NOT NULL,
    plate VARCHAR(20),
    gram INT NOT NULL,
    color VARCHAR(20) NOT NULL,
    page INT NOT NULL,
    pattern VARCHAR(50) NOT NULL,
    has_reference BOOLEAN NOT NULL,
    quantity INT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    CONSTRAINT fk_quotation FOREIGN KEY (quotation_id) REFERENCES quotations(id)
);

-- create table productions
CREATE TABLE productions (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    school_id INT NOT NULL,
    quotation_id INT NOT NULL,
    remark TEXT,
    CONSTRAINT fk_user_production FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_school_production FOREIGN KEY (school_id) REFERENCES schools(id),
    CONSTRAINT fk_quotation_production FOREIGN KEY (quotation_id) REFERENCES quotations(id)
);

-- create table production_items
CREATE TABLE production_items (
    id SERIAL NOT NULL PRIMARY KEY,
    production_id INT NOT NULL,
    category VARCHAR(100) NOT NULL,
    plate VARCHAR(20) NOT NULL,
    gram INT NOT NULL,
    color VARCHAR(20) NOT NULL,
    page INT NOT NULL,
    pattern VARCHAR(50) NOT NULL,
    has_reference BOOLEAN NOT NULL,
    quantity INT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_production FOREIGN KEY (production_id) REFERENCES productions(id)
);

-- create table price references
CREATE TABLE price_references (
    id SERIAL NOT NULL PRIMARY KEY,
    tier_id SMALLINT NOT NULL,
    category VARCHAR(100) NOT NULL,
    plate VARCHAR(20) NOT NULL,
    gram INT NOT NULL,
    color VARCHAR(20) NOT NULL,
    page INT NOT NULL,
    pattern VARCHAR(50) NOT NULL,
    has_reference BOOLEAN NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);