-- create table user
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    store_name VARCHAR(100),
    tier_id SMALLINT,
    `role` VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create table schools
CREATE TABLE schools (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255),
    telephone VARCHAR(11),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);

-- create table quotations
CREATE TABLE quotations (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    school_id INT NOT NULL,
    store_name VARCHAR(100),
    school_name VARCHAR(100),
    school_address VARCHAR(255),
    school_telephone VARCHAR(11),
    is_instant BOOLEAN,
    appointment_at TIMESTAMP,
    duedate_at TIMESTAMP,
    `status` VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_quotation FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_school FOREIGN KEY (school_id) REFERENCES schools(id)
);

-- create table quotation_items
CREATE TABLE quotation_items (
    id SERIAL PRIMARY KEY,
    quotation_id INT NOT NULL,
    product VARCHAR(100),
    plate VARCHAR(20),
    gram INT,
    color INT,
    `page` INT,
    `line` VARCHAR(50),
    has_ref BOOLEAN,
    quantity INT,
    price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_quotation FOREIGN KEY (quotation_id) REFERENCES quotations(id)
);

-- create table productions
CREATE TABLE productions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    school_id INT NOT NULL,
    quotation_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_production FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_school_production FOREIGN KEY (school_id) REFERENCES schools(id),
    CONSTRAINT fk_quotation_production FOREIGN KEY (quotation_id) REFERENCES quotations(id)
);

-- create table production_items
CREATE TABLE production_items (
    id SERIAL PRIMARY KEY,
    prod_id INT NOT NULL,
    product VARCHAR(100),
    plate VARCHAR(20),
    gram INT,
    color INT,
    `page` INT,
    `line` VARCHAR(50),
    has_ref BOOLEAN,
    `status` VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_production FOREIGN KEY (prod_id) REFERENCES productions(id)
);

-- create table price references
CREATE TABLE price_references (
    id SERIAL PRIMARY KEY,
    tier_id SMALLINT NOT NULL,
    product VARCHAR(100),
    gram INT,
    color INT,
    `page` INT,
    `line` VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);