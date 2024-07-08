CREATE TABLE users(
    id varchar(20) PRIMARY KEY,
    username varchar(45) NOT NULL,
    password TEXT NOT NULL,
    role varchar(15) NOT NULL,
    first_name varchar(45) NOT NULL,
    last_name varchar(45) not null,
    email varchar(55) NOT NULL
);

CREATE TABLE lecturer(
    nidn varchar(15) primary key, 
    lecturer_name text not null
);

CREATE TABLE students_description(
    student_id varchar(10) PRIMARY KEY,
    first_name varchar(45) NOT NULL,
    last_name varchar(45) NOT NULL,
    email varchar(60) NOT NULL,
    major varchar(35) NOT NULL,
    gender varchar(40) NOT NULL,
    national_id varchar(30) NOT NULL,
    address TEXT NOT NULL,
    birth_date varchar(40) NOT NULL,
    message TEXT NOT NULL,
    message_skl TEXT,
    nidn_advisor_one varchar(18) NOT NULL,
    nidn_advisor_two varchar(18) NOT NULL,
    nidn_religion varchar(18) NOT NULL,
    verification_skl TEXT NOT NULL,
    birth_place varchar(30) NOT NULL,
    phone_number varchar(13) NOT NULL,
    telephone_number varchar(13),
    verification varchar(30) NOT NULL DEFAULT 'NOT_VERIFIED',
    credit_course INTEGER NOT NULL,
    gpa DOUBLE PRECISION NOT NULL,
    thesis_title TEXT NOT NULL,
    advisor TEXT NOT NULL,
    examiner TEXT NOT NULL,
    academic_year varchar(10) NOT NULL,
    semester varchar(7) NOT NULL,
    religion_advisor TEXT NOT NULL,
    graduate_date varchar(20),
    commencement_date varchar(20),
    CONSTRAINT fk_advisor_one FOREIGN KEY (nidn_advisor_one) REFERENCES lecturer(nidn),
    CONSTRAINT fk_advisor_two FOREIGN KEY (nidn_advisor_two) REFERENCES lecturer(nidn),
    CONSTRAINT fk_religion_advisor FOREIGN KEY (nidn_religion) REFERENCES lecturer(nidn)
);


CREATE TABLE graduate_certificate_form(
    id varchar(10) primary key,
    student_id varchar(10) NOT NULL,
    full_name TEXT NOT NULL,
    birth_place varchar(30) NOT NULL,
    birth_date varchar(40) NOT NULL,
    gender varchar(30) NOT NULL,
    address TEXT NOT NULL,
    major varchar(35) NOT NULL,
    phone_number varchar(13) NOT NULL,
    religion varchar(10) NOT NULL,
    gpa double precision NOT NULL,
    level varchar(30) NOT NULL,
    dad TEXT NOT NULL,
    mother TEXT NOT NULL,
    parent_telp varchar(15) NOT NULL,
    commencement_date varchar(20),
    parent_address TEXT NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students_description(student_id),
    UNIQUE (student_id)
);

CREATE TABLE documents_admin(
    id varchar(10) primary key,
    nidn varchar(15) not null,
    student_id varchar not null,
    advisor_assignment_letter text ,
    examiner_assignment_letter text,
    invitation text,
    temp_grad text,
    official_report text,
    FOREIGN KEY(student_id) REFERENCES students_description(student_id)
);



CREATE TABLE student_documents(
    id text primary key,
    id_student varchar(10) not null,
    birth_certificate text, 
    toeic_certificate text ,
    id_card text,
    article text,
    competency_certificate text,
    family_card text,
    student_card text, 
    thesis_file text, 
    temp_graduation_certificate text,
    validity_sheet text,
    photo text ,
    graduation_certificate text,
    FOREIGN KEY(id_student) REFERENCES students_description(student_id)
);

CREATE  TABLE "Session" (
    id serial primary key,
    username varchar(255)  NOT NULL,
    session_code varchar(10) NOT NULL
);


CREATE TABLE semester (
    id varchar(10) PRIMARY KEY,
    academic_year varchar(15) NOT NULL,
    semester varchar(20) NOT NULL,
    status varchar(20) NOT NULL
)



