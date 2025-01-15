
CREATE TABLE admins (
    id UUID PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


-- CREATE TABLE petrol_types (
--     id UUID PRIMARY KEY,
--     name VARCHAR(100) NOT NULL,
--     created_at TIMESTAMP DEFAULT NOW(),
--     updated_at TIMESTAMP DEFAULT NOW()
-- );

CREATE TABLE departments (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    region  VARCHAR(100) NOT NULL,
    address VARCHAR(100) NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE cars (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    state_number VARCHAR(100) NOT NULL,
    year VARCHAR(100) NOT NULL,
    ban_status VARCHAR(100),
    tech_condition VARCHAR(100),
    defect TEXT,
    address VARCHAR(100) NOT NULL,
    department_id UUID REFERENCES departments("id"),
    petrol_name VARCHAR(100) NOT NULL,
    petrol NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE petrol_history (
    id UUID PRIMARY KEY,
    car_id UUID REFERENCES cars("id"),
    car_name VARCHAR(100) NOT NULL,
    car_model VARCHAR(100) NOT NULL,
    car_state_number VARCHAR(100) NOT NULL,
    car_year VARCHAR(100) NOT NULL,
    car_ban_status VARCHAR(100),
    car_tech_condition VARCHAR(100),
    car_defect TEXT,
    car_address VARCHAR(100) NOT NULL,
    car_department_id UUID REFERENCES departments("id"),
    car_petrol_name VARCHAR(100) NOT NULL,
    car_remaining_petrol NUMERIC NOT NULL,
    car_added_petrol NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
