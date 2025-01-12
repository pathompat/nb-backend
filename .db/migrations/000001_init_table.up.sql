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
    deleted_at TIMESTAMP DEFAULT NULL,
    UNIQUE(username)
);

-- create table schools
CREATE TABLE schools (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    contact_name VARCHAR(255) DEFAULT NULL,
    address VARCHAR(255) DEFAULT NULL,
    telephone VARCHAR(11) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);

-- create table schools
CREATE TABLE categories (
    id SERIAL NOT NULL PRIMARY KEY,
    key VARCHAR(100) NOT NULL,
    name_th VARCHAR(100) NOT NULL,
    name_en VARCHAR(100) NOT NULL,
    list_gram TEXT[] DEFAULT ARRAY[]::TEXT[],
    list_page TEXT[] DEFAULT ARRAY[]::TEXT[],
    list_pattern TEXT[] DEFAULT ARRAY[]::TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    UNIQUE(key),
    UNIQUE(name_th),
    UNIQUE(name_en)
);

-- create table quotations
CREATE TABLE quotations (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    school_id INT NOT NULL,
    store_name VARCHAR(100) NOT NULL,
    school_name VARCHAR(100) NOT NULL,
    school_address VARCHAR(255) DEFAULT NULL,
    school_telephone VARCHAR(11) DEFAULT NULL,
    school_contact_name VARCHAR(255) DEFAULT NULL,
    appointment_at TIMESTAMP DEFAULT NULL,
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
    category_id INT NOT NULL,
    plate VARCHAR(20),
    gram INT NOT NULL,
    color VARCHAR(20) DEFAULT NULL,
    page INT NOT NULL,
    pattern VARCHAR(50) NOT NULL,
    printed_content VARCHAR(100) DEFAULT NULL,
    quotation_config_ids INT[] NOT NULL DEFAULT ARRAY[]::INT[],
    has_reference BOOLEAN NOT NULL DEFAULT FALSE,
    quantity INT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    charge DOUBLE PRECISION NOT NULL,
    CONSTRAINT fk_quotation FOREIGN KEY (quotation_id) REFERENCES quotations(id),
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- create quotation additional list
CREATE TABLE quotation_additional_lists (
    id SERIAL NOT NULL PRIMARY KEY,
    quotation_id INT NOT NULL,
    quotation_config_id INT NOT NULL,
    quotation_config_key VARCHAR(100) NOT NULL,
    value DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_quotation FOREIGN KEY (quotation_id) REFERENCES quotations(id),
);

-- create quotation config
CREATE TABLE quotation_configs (
    id SERIAL NOT NULL PRIMARY KEY,
    category_id INT NOT NULL,
    tier_ids SMALLINT[] NOT NULL DEFAULT ARRAY[]::SMALLINT[],
    key VARCHAR(100) NOT NULL,
    value DOUBLE PRECISION NOT NULL,
    description TEXT DEFAULT NULL,
    label VARCHAR(100) DEFAULT NULL,
    unit VARCHAR(100) NOT NULL DEFAULT 'BAHT',
    level VARCHAR(100) NOT NULL DEFAULT 'quotation_additional_lists',
    comparator VARCHAR(50) DEFAULT NULL,
    compare_value INT DEFAULT NULL,
    has_fixed_charge BOOLEAN NOT NULL DEFAULT FALSE,
    fixed_charge_price DOUBLE PRECISION NOT NULL DEFAULT 0,
    type VARCHAR(20) NOT NULL DEFAULT 'CHARGES',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id),
    UNIQUE(key)
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
    CONSTRAINT fk_quotation_production FOREIGN KEY (quotation_id) REFERENCES quotations(id),
    UNIQUE(quotation_id)
);

-- create table production_items
CREATE TABLE production_items (
    id SERIAL NOT NULL PRIMARY KEY,
    production_id INT NOT NULL,
    category_id INT NOT NULL,
    plate VARCHAR(20) NOT NULL,
    gram INT NOT NULL,
    color VARCHAR(20) DEFAULT NULL,
    page INT NOT NULL,
    pattern VARCHAR(50) NOT NULL,
    printed_content VARCHAR(100) DEFAULT NULL,
    quotation_config_ids INT[] DEFAULT ARRAY[]::INT[],
    has_reference BOOLEAN NOT NULL DEFAULT FALSE,
    quantity INT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_production FOREIGN KEY (production_id) REFERENCES productions(id),
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- create table price references
CREATE TABLE price_references (
    id SERIAL NOT NULL PRIMARY KEY,
    tier_id SMALLINT NOT NULL,
    category_id INT DEFAULT NULL,
    gram INT NOT NULL,
    page INT NOT NULL,
    color VARCHAR(20) DEFAULT NULL,
    pattern TEXT[] NOT NULL DEFAULT ARRAY[]::TEXT[],
    price DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id)
);