DROP TABLE IF EXISTS user;
CREATE TABLE user (
  user_id           INT AUTO_INCREMENT NOT NULL,
  email             VARCHAR(128) NOT NULL,
  hashed_password   VARCHAR(255) NOT NULL,
  role              VARCHAR(128) NOT NULL,
  PRIMARY KEY (`user_id`)
);

DROP TABLE IF EXISTS passenger;
CREATE TABLE passenger (
  passenger_id  INT AUTO_INCREMENT NOT NULL,
  name          VARCHAR(128) NOT NULL,
  surname       VARCHAR(255) NOT NULL,
  email         VARCHAR(255) NOT NULL,
  PRIMARY KEY (`passenger_id`)
);


DROP TABLE IF EXISTS employee;
CREATE TABLE employee (
  employee_id   INT AUTO_INCREMENT NOT NULL,
  name          VARCHAR(128) NOT NULL,
  surname       VARCHAR(255) NOT NULL,
  email         VARCHAR(255) NOT NULL,
  title         VARCHAR(128) NOT NULL,
  salary        INT NOT NULL,
  PRIMARY KEY (`employee_id`)
);

DROP TABLE IF EXISTS plane;
CREATE TABLE plane (
  plane_id        INT AUTO_INCREMENT NOT NULL,
  airline_id      INT NOT NULL,
  code            VARCHAR(128) NOT NULL,
  capacity        INT NOT NULL,
  PRIMARY KEY (`plane_id`)
);

DROP TABLE IF EXISTS airline;
CREATE TABLE airline (
  airline_id      INT NOT NULL,
  name            VARCHAR(128) NOT NULL,
  PRIMARY KEY (`airline_id`)
);

DROP TABLE IF EXISTS flight;
CREATE TABLE flight (
  flight_id     INT AUTO_INCREMENT NOT NULL,
  plane_id      INT NOT NULL,
  dep_place     VARCHAR(128) NOT NULL,
  lan_place     VARCHAR(128) NOT NULL,
  dep_time      DATETIME NOT NULL,
  lan_time      DATETIME NOT NULL,
  PRIMARY KEY (`flight_id`)
);

DROP TABLE IF EXISTS ticket;
CREATE TABLE ticket (
  ticket_id         INT AUTO_INCREMENT NOT NULL,
  passenger_id      INT NOT NULL,
  flight_id         INT NOT NULL,
  seat              VARCHAR(128) NOT NULL,
  class             VARCHAR(128) NOT NULL,
  price             INT NOT NULL,
  PRIMARY KEY (`ticket_id`)
);

INSERT INTO passenger
    (name, surname, email)
VALUES
  ('Ken', 'Thompson', 'ken.thompson@gmail.com'),
  ('Rob', 'Pike', 'rob.pike@gmail.com'),
  ('Robert', 'Griesemer', 'robert.griesemer@gmail.com'),
  ('Russ', 'Cox', 'russ.cox@gmail.com'),
  ('William', 'Kennedy', 'bill.kennedy@ardanlabs.com'),
  ('Darcy', 'Hunt', 'd.hunt@randatmail.com'),
  ('Melissa', 'Myers', 'm.myers@randatmail.com'),
  ('Edith', 'West', 'e.west@randatmail.com'),
  ('Deanna', 'Hall', 'd.hall@randatmail.com'),
  ('Robert', 'Owens', 'r.owens@randatmail.com'),
  ('Abraham', 'Allen', 'a.allen@randatmail.com'),
  ('Maya', 'Taylor', 'm.taylor@randatmail.com'),
  ('Byron', 'Fowler', 'b.fowler@randatmail.com'),
  ('Fenton', 'Andrews', 'f.andrews@randatmail.com'),
  ('Arianna', 'Payne', 'a.payne@randatmail.com'),
  ('Anna', 'Gibson', 'a.gibson@randatmail.com'),
  ('Ned', 'Hawkins', 'n.hawkins@randatmail.com'),
  ('Derek', 'Morgan', 'd.morgan@randatmail.com'),
  ('Darcy', 'Riley', 'd.riley@randatmail.com'),
  ('Brooke', 'Morgan', 'b.morgan@randatmail.com'),
  ('Brianna', 'Fowler', 'b.fowler@randatmail.com'),
  ('Mary', 'Robinson', 'm.robinson@randatmail.com'),
  ('Melissa', 'Foster', 'm.foster@randatmail.com'),
  ('Michelle', 'Williams', 'm.williams@randatmail.com'),
  ('Rosie', 'Cameron', 'r.cameron@randatmail.com'),
  ('Kirsten', 'Johnson', 'k.johnson@randatmail.com'),
  ('Victoria', 'Brown', 'v.brown@randatmail.com'),
  ('John', 'Richards', 'j.richards@randatmail.com'),
  ('Daryl', 'Hunt', 'd.hunt@randatmail.com'),
  ('Chester', 'Cooper', 'c.cooper@randatmail.com'),
  ('Daisy', 'Davis', 'd.davis@randatmail.com'),
  ('Madaline', 'Morris', 'm.morris@randatmail.com'),
  ('Freddie', 'Rogers', 'f.rogers@randatmail.com'),
  ('Vanessa', 'Wilson', 'v.wilson@randatmail.com'),
  ('Robert', 'Clark', 'r.clark@randatmail.com'),
  ('Eddy', 'Barnes', 'e.barnes@randatmail.com'),
  ('Adrianna', 'Russell', 'a.russell@randatmail.com'),
  ('Sienna', 'Wilson', 's.wilson@randatmail.com'),
  ('Chelsea', 'Morgan', 'c.morgan@randatmail.com'),
  ('Daniel', 'Wilson', 'd.wilson@randatmail.com'),
  ('Lucy', 'Edwards', 'l.edwards@randatmail.com'),
  ('Adam', 'Wright', 'a.wright@randatmail.com'),
  ('Edgar', 'Higgins', 'e.higgins@randatmail.com'),
  ('Fenton', 'Williams', 'f.williams@randatmail.com'),
  ('Preston', 'Carroll', 'p.carroll@randatmail.com'),
  ('Adele', 'Andrews', 'a.andrews@randatmail.com'),
  ('Jessica', 'Williams', 'j.williams@randatmail.com'),
  ('Roman', 'Spencer', 'r.spencer@randatmail.com'),
  ('Jack', 'Johnson', 'j.johnson@randatmail.com'),
  ('Sawyer', 'Sullivan', 's.sullivan@randatmail.com'),
  ('Darcy', 'Brown', 'd.brown@randatmail.com'),
  ('Owen', 'Robinson', 'o.robinson@randatmail.com'),
  ('Nicole', 'Farrell', 'n.farrell@randatmail.com'),
  ('Michael', 'Stevens', 'm.stevens@randatmail.com'),
  ('Thomas', 'Howard', 't.howard@randatmail.com'),
  ('Thomas', 'Andrews', 't.andrews@randatmail.com'),
  ('Sofia', 'Hall', 's.hall@randatmail.com'),
  ('Charlotte', 'Grant', 'c.grant@randatmail.com'),
  ('Julia', 'Dixon', 'j.dixon@randatmail.com'),
  ('Aiden', 'Dixon', 'a.dixon@randatmail.com'),
  ('Maya', 'Morrison', 'm.morrison@randatmail.com'),
  ('Vanessa', 'Douglas', 'v.douglas@randatmail.com'),
  ('Miller', 'Sullivan', 'm.sullivan@randatmail.com'),
  ('Hailey', 'Foster', 'h.foster@randatmail.com'),
  ('Annabella', 'Barrett', 'a.barrett@randatmail.com'),
  ('Adele', 'Higgins', 'a.higgins@randatmail.com'),
  ('Lucas', 'Higgins', 'l.higgins@randatmail.com'),
  ('John', 'Miller', 'j.miller@randatmail.com'),
  ('Gianna', 'Watson', 'g.watson@randatmail.com'),
  ('Natalie', 'Nelson', 'n.nelson@randatmail.com'),
  ('Lyndon', 'Douglas', 'l.douglas@randatmail.com'),
  ('Emily', 'Fowler', 'e.fowler@randatmail.com'),
  ('Maria', 'Kelly', 'm.kelly@randatmail.com'),
  ('Aston', 'Mason', 'a.mason@randatmail.com'),
  ('Joyce', 'Johnson', 'j.johnson@randatmail.com'),
  ('Eddy', 'Carroll', 'e.carroll@randatmail.com'),
  ('Preston', 'Perry', 'p.perry@randatmail.com'),
  ('Kirsten', 'Martin', 'k.martin@randatmail.com'),
  ('John', 'Wilson', 'j.wilson@randatmail.com'),
  ('Chloe', 'Thomas', 'c.thomas@randatmail.com'),
  ('Eleanor', 'Evans', 'e.evans@randatmail.com'),
  ('Antony', 'Henderson', 'a.henderson@randatmail.com'),
  ('Audrey', 'Parker', 'a.parker@randatmail.com'),
  ('Frederick', 'Foster', 'f.foster@randatmail.com'),
  ('Joyce', 'Reed', 'j.reed@randatmail.com'),
  ('Isabella', 'Clark', 'i.clark@randatmail.com'),
  ('Jessica', 'Stewart', 'j.stewart@randatmail.com'),
  ('Heather', 'Stevens', 'h.stevens@randatmail.com'),
  ('Lucia', 'Harper', 'l.harper@randatmail.com'),
  ('Cadie', 'Bailey', 'c.bailey@randatmail.com'),
  ('Jack', 'Edwards', 'j.edwards@randatmail.com'),
  ('Sawyer', 'West', 's.west@randatmail.com'),
  ('Edgar', 'Campbell', 'e.campbell@randatmail.com'),
  ('Emma', 'Russell', 'e.russell@randatmail.com'),
  ('Adele', 'Nelson', 'a.nelson@randatmail.com'),
  ('Vincent', 'Ferguson', 'v.ferguson@randatmail.com'),
  ('Miranda', 'Spencer', 'm.spencer@randatmail.com'),
  ('Agata', 'Cooper', 'a.cooper@randatmail.com'),
  ('Tony', 'Murphy', 't.murphy@randatmail.com'),
  ('Patrick', 'Stewart', 'p.stewart@randatmail.com'),
  ('Mary', 'Kelley', 'm.kelley@randatmail.com'),
  ('Robert', 'Thompson', 'r.thompson@randatmail.com'),
  ('Cherry', 'Thomas', 'c.thomas@randatmail.com'),
  ('Alen', 'West', 'a.west@randatmail.com'),
  ('Annabella', 'Henderson', 'a.henderson@randatmail.com'),
  ('Antony', 'Cooper', 'a.cooper@randatmail.com'),
  ('Amber', 'Morgan', 'a.morgan@randatmail.com'),
  ('Freddie', 'Perry', 'f.perry@randatmail.com'),
  ('Martin', 'Harrison', 'm.harrison@randatmail.com'),
  ('Alina', 'Allen', 'a.allen@randatmail.com'),
  ('Emily', 'Hall', 'e.hall@randatmail.com'),
  ('Lucas', 'Alexander', 'l.alexander@randatmail.com'),
  ('Robert', 'Thomas', 'r.thomas@randatmail.com'),
  ('Connie', 'Morrison', 'c.morrison@randatmail.com'),
  ('Blake', 'Lloyd', 'b.lloyd@randatmail.com'),
  ('Edwin', 'Edwards', 'e.edwards@randatmail.com'),
  ('Luke', 'Smith', 'l.smith@randatmail.com'),
  ('Jasmine', 'Johnson', 'j.johnson@randatmail.com'),
  ('Haris', 'Stevens', 'h.stevens@randatmail.com'),
  ('Patrick', 'Russell', 'p.russell@randatmail.com'),
  ('Max', 'Lloyd', 'm.lloyd@randatmail.com'),
  ('Lilianna', 'Craig', 'l.craig@randatmail.com'),
  ('Agata', 'Davis', 'a.davis@randatmail.com'),
  ('Alissa', 'Rogers', 'a.rogers@randatmail.com'),
  ('Ellia', 'Kelley', 'e.kelley@randatmail.com'),
  ('Eleanor', 'Higgins', 'e.higgins@randatmail.com'),
  ('Joyce', 'Armstrong', 'j.armstrong@randatmail.com'),
  ('Honey', 'Higgins', 'h.higgins@randatmail.com'),
  ('Chester', 'Owens', 'c.owens@randatmail.com'),
  ('David', 'Ryan', 'd.ryan@randatmail.com'),
  ('Alina', 'Williams', 'a.williams@randatmail.com'),
  ('Cherry', 'Hamilton', 'c.hamilton@randatmail.com'),
  ('Mary', 'Williams', 'm.williams@randatmail.com'),
  ('Lily', 'Chapman', 'l.chapman@randatmail.com'),
  ('Lyndon', 'Henderson', 'l.henderson@randatmail.com'),
  ('Justin', 'Turner', 'j.turner@randatmail.com'),
  ('Lydia', 'Carroll', 'l.carroll@randatmail.com'),
  ('Amelia', 'Robinson', 'a.robinson@randatmail.com'),
  ('Dale', 'Johnston', 'd.johnston@randatmail.com'),
  ('Aiden', 'Ryan', 'a.ryan@randatmail.com'),
  ('Rubie', 'Cole', 'r.cole@randatmail.com'),
  ('Arianna', 'Nelson', 'a.nelson@randatmail.com'),
  ('Blake', 'Walker', 'b.walker@randatmail.com'),
  ('Ned', 'Brown', 'n.brown@randatmail.com'),
  ('Paul', 'Johnson', 'p.johnson@randatmail.com'),
  ('Alexia', 'Gray', 'a.gray@randatmail.com'),
  ('Jacob', 'Myers', 'j.myers@randatmail.com'),
  ('Sophia', 'Elliott', 's.elliott@randatmail.com'),
  ('Melanie', 'Ferguson', 'm.ferguson@randatmail.com'),
  ('Reid', 'Henderson', 'r.henderson@randatmail.com'),
  ('Preston', 'Hill', 'p.hill@randatmail.com'),
  ('Briony', 'Perry', 'b.perry@randatmail.com'),
  ('Roland', 'Craig', 'r.craig@randatmail.com'),
  ('Alissa', 'Nelson', 'a.nelson@randatmail.com'),
  ('Nicole', 'Harris', 'n.harris@randatmail.com'),
  ('Lucia', 'Stevens', 'l.stevens@randatmail.com'),
  ('Penelope', 'Watson', 'p.watson@randatmail.com'),
  ('Olivia', 'Foster', 'o.foster@randatmail.com'),
  ('Carl', 'Moore', 'c.moore@randatmail.com'),
  ('Oliver', 'Harper', 'o.harper@randatmail.com'),
  ('Freddie', 'Hill', 'f.hill@randatmail.com'),
  ('Mary', 'Carroll', 'm.carroll@randatmail.com'),
  ('Robert', 'Tucker', 'r.tucker@randatmail.com'),
  ('Robert', 'Taylor', 'r.taylor@randatmail.com'),
  ('Annabella', 'Russell', 'a.russell@randatmail.com'),
  ('Abraham', 'Howard', 'a.howard@randatmail.com'),
  ('Florrie', 'Hill', 'f.hill@randatmail.com'),
  ('Carlos', 'Hill', 'c.hill@randatmail.com'),
  ('Dominik', 'Howard', 'd.howard@randatmail.com'),
  ('Henry', 'Richardson', 'h.richardson@randatmail.com'),
  ('Sienna', 'Chapman', 's.chapman@randatmail.com'),
  ('Tara', 'Morris', 't.morris@randatmail.com'),
  ('Alissa', 'Edwards', 'a.edwards@randatmail.com'),
  ('Madaline', 'Adams', 'm.adams@randatmail.com'),
  ('Ashton', 'Armstrong', 'a.armstrong@randatmail.com'),
  ('Adrianna', 'Perry', 'a.perry@randatmail.com'),
  ('Hailey', 'Wilson', 'h.wilson@randatmail.com'),
  ('Carina', 'Martin', 'c.martin@randatmail.com'),
  ('Olivia', 'Hamilton', 'o.hamilton@randatmail.com'),
  ('Lenny', 'Watson', 'l.watson@randatmail.com'),
  ('Victoria', 'Reed', 'v.reed@randatmail.com'),
  ('Vincent', 'Johnson', 'v.johnson@randatmail.com'),
  ('Tara', 'Johnston', 't.johnston@randatmail.com'),
  ('Byron', 'Anderson', 'b.anderson@randatmail.com'),
  ('Clark', 'Bailey', 'c.bailey@randatmail.com'),
  ('Albert', 'Harper', 'a.harper@randatmail.com'),
  ('Byron', 'Ryan', 'b.ryan@randatmail.com'),
  ('Vanessa', 'Morrison', 'v.morrison@randatmail.com'),
  ('Edgar', 'Elliott', 'e.elliott@randatmail.com'),
  ('Tyler', 'Walker', 't.walker@randatmail.com'),
  ('Vanessa', 'Tucker', 'v.tucker@randatmail.com'),
  ('Owen', 'Evans', 'o.evans@randatmail.com'),
  ('Adrianna', 'Turner', 'a.turner@randatmail.com'),
  ('Ada', 'Allen', 'a.allen@randatmail.com'),
  ('Maximilian', 'Brown', 'm.brown@randatmail.com'),
  ('Adelaide', 'Morgan', 'a.morgan@randatmail.com'),
  ('Joyce', 'Henderson', 'j.henderson@randatmail.com'),
  ('Alford', 'Grant', 'a.grant@randatmail.com'),
  ('Daisy', 'Wilson', 'd.wilson@randatmail.com'),
  ('Paige', 'Armstrong', 'p.armstrong@randatmail.com'),
  ('Brad', 'Howard', 'b.howard@randatmail.com'),
  ('Honey', 'Carroll', 'h.carroll@randatmail.com'),
  ('Chester', 'Robinson', 'c.robinson@randatmail.com'),
  ('Lucy', 'Clark', 'l.clark@randatmail.com'),
  ('William', 'Watson', 'w.watson@randatmail.com');


INSERT INTO employee
    (name, surname, email, title, salary)
VALUES
    ('Enes', 'Toraman', 'enestoraman@yahoo.com', 'Pilot', 15000),
    ('Harry', 'Kane', 'harry.kane@gmail.com', 'Staff', 6000),
    ('Luka', 'Doncic', 'luka.doncic@gmail.com', 'Flight Attendant', 8000),
    ('Mete', 'Gazoz', 'mete.gazoz@gmail.com', 'Pilot', 14000),
    ('Cedi', 'Osman', 'cedi.osman@gmail.com', 'Staff', 7000),
    ('Joel', 'Embiid', 'joel.embiid@gmail.com', 'Flight Attendant', 12000),
    ('Wardell', 'Curry', 'wardell.curry@gmail.com', 'Pilot', 20000),
    ('Martin', 'Harrison', 'm.harrison@randatmail.com', 'Pilot', 15000),
    ('Alina', 'Allen', 'a.allen@randatmail.com', 'Staff', 6000),
    ('Emily', 'Hall', 'e.hall@randatmail.com', 'Flight Attendant', 8000),
    ('Lucas', 'Alexander', 'l.alexander@randatmail.com', 'Pilot', 14000),
    ('Robert', 'Thomas', 'r.thomas@randatmail.com', 'Staff', 7000),
    ('Connie', 'Morrison', 'c.morrison@randatmail.com', 'Flight Attendant', 12000),
    ('Blake', 'Lloyd', 'b.lloyd@randatmail.com', 'Pilot', 20000),
    ('Edwin', 'Edwards', 'e.edwards@randatmail.com', 'Pilot', 15000),
    ('Luke', 'Smith', 'l.smith@randatmail.com', 'Staff', 6000),
    ('Jasmine', 'Johnson', 'j.johnson@randatmail.com', 'Flight Attendant', 8000),
    ('Haris', 'Stevens', 'h.stevens@randatmail.com', 'Pilot', 14000),
    ('Patrick', 'Russell', 'p.russell@randatmail.com', 'Staff', 7000),
    ('Max', 'Lloyd', 'm.lloyd@randatmail.com', 'Flight Attendant', 12000),
    ('Lilianna', 'Craig', 'l.craig@randatmail.com', 'Pilot', 20000),
    ('Agata', 'Davis', 'a.davis@randatmail.com', 'Pilot', 15000),
    ('Alissa', 'Rogers', 'a.rogers@randatmail.com', 'Staff', 6000),
    ('Ellia', 'Kelley', 'e.kelley@randatmail.com', 'Flight Attendant', 8000),
    ('Eleanor', 'Higgins', 'e.higgins@randatmail.com', 'Pilot', 14000),
    ('Joyce', 'Armstrong', 'j.armstrong@randatmail.com', 'Staff', 7000),
    ('Honey', 'Higgins', 'h.higgins@randatmail.com', 'Flight Attendant', 12000),
    ('Chester', 'Owens', 'c.owens@randatmail.com', 'Pilot', 20000),
    ('David', 'Ryan', 'd.ryan@randatmail.com', 'Pilot', 15000),
    ('Alina', 'Williams', 'a.williams@randatmail.com', 'Staff', 6000),
    ('Cherry', 'Hamilton', 'c.hamilton@randatmail.com', 'Flight Attendant', 8000),
    ('Mary', 'Williams', 'm.williams@randatmail.com', 'Pilot', 14000),
    ('Lily', 'Chapman', 'l.chapman@randatmail.com', 'Staff', 7000),
    ('Lyndon', 'Henderson', 'l.henderson@randatmail.com', 'Flight Attendant', 12000),
    ('Justin', 'Turner', 'j.turner@randatmail.com', 'Pilot', 20000),
    ('Lydia', 'Carroll', 'l.carroll@randatmail.com', 'Pilot', 15000),
    ('Amelia', 'Robinson', 'a.robinson@randatmail.com', 'Staff', 6000),
    ('Dale', 'Johnston', 'd.johnston@randatmail.com', 'Flight Attendant', 8000),
    ('Aiden', 'Ryan', 'a.ryan@randatmail.com', 'Pilot', 14000),
    ('Rubie', 'Cole', 'r.cole@randatmail.com', 'Staff', 7000),
    ('Arianna', 'Nelson', 'a.nelson@randatmail.com', 'Flight Attendant', 12000),
    ('Blake', 'Walker', 'b.walker@randatmail.com', 'Pilot', 20000),
    ('Ned', 'Brown', 'n.brown@randatmail.com', 'Pilot', 15000),
    ('Paul', 'Johnson', 'p.johnson@randatmail.com', 'Staff', 6000),
    ('Alexia', 'Gray', 'a.gray@randatmail.com', 'Flight Attendant', 8000),
    ('Jacob', 'Myers', 'j.myers@randatmail.com', 'Pilot', 14000),
    ('Sophia', 'Elliott', 's.elliott@randatmail.com', 'Staff', 7000),
    ('Melanie', 'Ferguson', 'm.ferguson@randatmail.com', 'Flight Attendant', 12000),
    ('Reid', 'Henderson', 'r.henderson@randatmail.com', 'Pilot', 20000),
    ('Preston', 'Hill', 'p.hill@randatmail.com', 'Pilot', 15000),
    ('Briony', 'Perry', 'b.perry@randatmail.com', 'Staff', 6000),
    ('Roland', 'Craig', 'r.craig@randatmail.com', 'Flight Attendant', 8000),
    ('Alissa', 'Nelson', 'a.nelson@randatmail.com', 'Pilot', 14000),
    ('Nicole', 'Harris', 'n.harris@randatmail.com', 'Staff', 7000),
    ('Lucia', 'Stevens', 'l.stevens@randatmail.com', 'Flight Attendant', 12000),
    ('Penelope', 'Watson', 'p.watson@randatmail.com', 'Pilot', 20000),
    ('Olivia', 'Foster', 'o.foster@randatmail.com', 'Pilot', 15000),
    ('Carl', 'Moore', 'c.moore@randatmail.com', 'Staff', 6000),
    ('Oliver', 'Harper', 'o.harper@randatmail.com', 'Flight Attendant', 8000),
    ('Freddie', 'Hill', 'f.hill@randatmail.com', 'Pilot', 14000),
    ('Mary', 'Carroll', 'm.carroll@randatmail.com', 'Staff', 7000),
    ('Robert', 'Tucker', 'r.tucker@randatmail.com', 'Flight Attendant', 12000),
    ('Robert', 'Taylor', 'r.taylor@randatmail.com', 'Pilot', 20000),
    ('Annabella', 'Russell', 'a.russell@randatmail.com', 'Pilot', 15000),
    ('Abraham', 'Howard', 'a.howard@randatmail.com', 'Staff', 6000),
    ('Florrie', 'Hill', 'f.hill@randatmail.com', 'Flight Attendant', 8000),
    ('Carlos', 'Hill', 'c.hill@randatmail.com', 'Pilot', 14000),
    ('Dominik', 'Howard', 'd.howard@randatmail.com', 'Staff', 7000),
    ('Henry', 'Richardson', 'h.richardson@randatmail.com', 'Flight Attendant', 12000),
    ('Sienna', 'Chapman', 's.chapman@randatmail.com', 'Pilot', 20000),
    ('Tara', 'Morris', 't.morris@randatmail.com', 'Pilot', 15000),
    ('Alissa', 'Edwards', 'a.edwards@randatmail.com', 'Staff', 6000),
    ('Madaline', 'Adams', 'm.adams@randatmail.com', 'Flight Attendant', 8000),
    ('Ashton', 'Armstrong', 'a.armstrong@randatmail.com', 'Pilot', 14000),
    ('Adrianna', 'Perry', 'a.perry@randatmail.com', 'Staff', 7000),
    ('Hailey', 'Wilson', 'h.wilson@randatmail.com', 'Flight Attendant', 12000),
    ('Carina', 'Martin', 'c.martin@randatmail.com', 'Pilot', 20000),
    ('Olivia', 'Hamilton', 'o.hamilton@randatmail.com', 'Pilot', 15000),
    ('Lenny', 'Watson', 'l.watson@randatmail.com', 'Staff', 6000),
    ('Victoria', 'Reed', 'v.reed@randatmail.com', 'Flight Attendant', 8000),
    ('Vincent', 'Johnson', 'v.johnson@randatmail.com', 'Pilot', 14000),
    ('Tara', 'Johnston', 't.johnston@randatmail.com', 'Staff', 7000),
    ('Byron', 'Anderson', 'b.anderson@randatmail.com', 'Flight Attendant', 12000),
    ('Clark', 'Bailey', 'c.bailey@randatmail.com', 'Pilot', 20000),
    ('Albert', 'Harper', 'a.harper@randatmail.com', 'Pilot', 15000),
    ('Byron', 'Ryan', 'b.ryan@randatmail.com', 'Staff', 6000),
    ('Vanessa', 'Morrison', 'v.morrison@randatmail.com', 'Flight Attendant', 8000),
    ('Edgar', 'Elliott', 'e.elliott@randatmail.com', 'Flight Attendant', 10000),
    ('Tyler', 'Walker', 't.walker@randatmail.com', 'Flight Attendant', 8000),
    ('Vanessa', 'Tucker', 'v.tucker@randatmail.com', 'Flight Attendant', 10000),
    ('Owen', 'Evans', 'o.evans@randatmail.com', 'Flight Attendant', 8000),
    ('Adrianna', 'Turner', 'a.turner@randatmail.com', 'Flight Attendant', 8000),
    ('Ada', 'Allen', 'a.allen@randatmail.com', 'Flight Attendant', 10000),
    ('Maximilian', 'Brown', 'm.brown@randatmail.com', 'Flight Attendant', 8000),
    ('Adelaide', 'Morgan', 'a.morgan@randatmail.com', 'Flight Attendant', 10000),
    ('Joyce', 'Henderson', 'j.henderson@randatmail.com', 'Flight Attendant', 8000),
    ('Alford', 'Grant', 'a.grant@randatmail.com', 'Flight Attendant', 10000),
    ('Daisy', 'Wilson', 'd.wilson@randatmail.com', 'Flight Attendant', 8000),
    ('Paige', 'Armstrong', 'p.armstrong@randatmail.com', 'Flight Attendant', 10000),
    ('Brad', 'Howard', 'b.howard@randatmail.com', 'Flight Attendant', 8000),
    ('Honey', 'Carroll', 'h.carroll@randatmail.com', 'Flight Attendant', 10000),
    ('Chester', 'Robinson', 'c.robinson@randatmail.com', 'Flight Attendant', 8000),
    ('Lucy', 'Clark', 'l.clark@randatmail.com', 'Flight Attendant', 10000),
    ('William', 'Watson', 'w.watson@randatmail.com', 'Flight Attendant', 8000);

INSERT INTO plane
    (airline_id, code, capacity)
VALUES
    (1, 'UKET-12', 91),
    (1, 'SKTY-26', 78),
    (1, 'TOLK-82', 77),
    (1, 'SKFY-65', 72),
    (2, 'MGTH-19', 93),
    (2, 'PKTL-56', 95),
    (2, 'FKTT-71', 76),
    (2, 'MNGT-06', 71),
    (3, 'UFYK-57', 98),
    (3, 'FLVR-47', 95),
    (3, 'TKİX-78', 84),
    (3, 'PRCS-96', 88),
    (3, 'QVYA-34', 89),
    (4, 'CBMZ-34', 70),
    (4, 'CMZD-14', 87),
    (5, 'PZMB-61', 91),
    (5, 'QPRZ-89', 99),
    (5, 'AJQC-4', 78);

INSERT INTO airline
    (airline_id, name)
VALUES
    (1, 'THY'),
    (2, 'PEGASUS'),
    (3, 'SWISS'),
    (4, 'AIR FRANCE'),
    (5, 'FINNAIR');

INSERT INTO flight
    (plane_id, dep_place, lan_place, dep_time, lan_time)
VALUES
    (3, 'PRT', 'SVN', '2022-01-01 12:00:00', '2022-01-01 15:00:00'),
    (19, 'EST', 'MLT', '2022-01-02 13:00:00', '2022-01-02 15:00:00'),
    (28, 'CZE', 'MLT', '2022-01-03 14:00:00', '2022-01-03 18:00:00'),
    (16, 'TUR', 'PRT', '2022-01-04 10:00:00', '2022-01-04 14:00:00'),
    (20, 'IRL', 'LVA', '2022-01-05 17:00:00', '2022-01-05 20:00:00'),
    (19, 'FIN', 'TUR', '2022-01-06 19:00:00', '2022-01-06 23:00:00'),
    (5, 'EST', 'LUX', '2022-01-07 09:00:00', '2022-01-07 12:00:00'),
    (22, 'MLT', 'LVA', '2022-01-08 21:00:00', '2022-01-08 23:00:00'),
    (23, 'FRA', 'GBR', '2022-01-09 07:00:00', '2022-01-09 10:00:00'),
    (11, 'EST', 'LTU', '2022-01-10 11:00:00', '2022-01-10 12:30:00'),
    (16, 'BEL', 'ROU', '2022-01-11 22:00:00', '2022-01-11 23:30:00'),
    (16, 'CYP', 'GRC', '2022-01-12 01:00:00', '2022-01-12 04:00:00'),
    (10, 'CZE', 'IRL', '2022-01-13 19:00:00', '2022-01-13 21:00:00'),
    (7, 'SVK', 'GBR', '2022-01-14 15:00:00', '2022-01-14 18:00:00'),
    (26, 'POL', 'SVN', '2022-01-01 15:00:00', '2022-01-01 18:00:00'),
    (5, 'BGR', 'HRV', '2022-01-02 15:00:00', '2022-01-02 17:00:00'),
    (12, 'HRV', 'FRA', '2022-01-03 18:00:00', '2022-01-03 20:00:00'),
    (25, 'ESP', 'LTU', '2022-01-04 14:00:00', '2022-01-04 17:00:00'),
    (5, 'HRV', 'IRL', '2022-01-05 20:00:00', '2022-01-05 23:00:00'),
    (4, 'SVN', 'FRA', '2022-01-06 23:00:00', '2022-01-07 00:30:00'),
    (29, 'LVA', 'GRC', '2022-01-07 12:00:00', '2022-01-07 15:00:00'),
    (11, 'BGR', 'CZE', '2022-01-08 23:00:00', '2022-01-09 02:00:00'),
    (25, 'GBR', 'ESP', '2022-01-09 10:00:00', '2022-01-09 12:00:00'),
    (15, 'PRT', 'DEU', '2022-01-10 12:30:00', '2022-01-10 15:30:00'),
    (25, 'DNK', 'DEU', '2022-01-11 23:30:00', '2022-01-12 01:30:00'),
    (7, 'TUR', 'AUT', '2022-01-12 04:00:00', '2022-01-12 07:00:00'),
    (21, 'FRA', 'GBR', '2022-01-13 21:00:00', '2022-01-13 23:00:00'),
    (21, 'NLD', 'ROU', '2022-01-14 18:00:00', '2022-01-14 21:00:00');

INSERT INTO ticket
    (passenger_id, flight_id, seat, class, price)
VALUES
    (1, 7, '02A', 'B', 1029),
    (2, 10, '07B', 'E', 528),
    (3, 28, '16C', 'E', 726),
    (4, 14, '22D', 'E', 663),
    (5, 13, '02E', 'B', 1234),
    (6, 10, '18F', 'E', 538),
    (7, 25, '01A', 'B', 1096),
    (8, 24, '20B', 'E', 633),
    (9, 2, '03A', 'B', 1183),
    (10, 2, '19D', 'E', 701),
    (11, 26, '11E', 'E', 743),
    (12, 27, '17F', 'E', 652),
    (13, 16, '02B', 'B', 1137),
    (14, 1, '12B', 'E', 750),
    (15, 14, '11C', 'E', 513),
    (16, 17, '03D', 'B', 1051),
    (17, 24, '23E', 'E', 711),
    (18, 17, '02A', 'B', 1149),
    (19, 18, '07B', 'E', 700),
    (20, 2, '16C', 'E', 742),
    (21, 26, '22D', 'E', 613),
    (22, 14, '02E', 'B', 1152),
    (23, 11, '18F', 'E', 594),
    (24, 12, '01A', 'B', 1121),
    (25, 17, '20B', 'E', 702),
    (26, 4, '03A', 'B', 1077),
    (27, 2, '19D', 'E', 612),
    (28, 5, '11E', 'E', 652),
    (29, 7, '17F', 'E', 582),
    (30, 19, '02B', 'B', 1104),
    (31, 13, '12B', 'E', 584),
    (32, 7, '11C', 'E', 724),
    (33, 5, '03D', 'B', 1154),
    (34, 23, '23E', 'E', 604),
    (35, 5, '02A', 'B', 1139),
    (36, 17, '07B', 'E', 713),
    (37, 6, '16C', 'E', 703),
    (38, 21, '22D', 'E', 694),
    (39, 6, '02E', 'B', 1055),
    (40, 27, '18F', 'E', 665),
    (41, 23, '01A', 'B', 1094),
    (42, 10, '20B', 'E', 643),
    (43, 18, '03A', 'B', 1177),
    (44, 26, '19D', 'E', 548),
    (45, 6, '11E', 'E', 661),
    (46, 26, '17F', 'E', 599),
    (47, 18, '02B', 'B', 1209),
    (48, 6, '12B', 'E', 617),
    (49, 23, '11C', 'E', 715),
    (50, 21, '03D', 'B', 1105),
    (51, 20, '23E', 'E', 581),
    (52, 27, '02A', 'B', 1189),
    (53, 1, '07B', 'E', 607),
    (54, 8, '16C', 'E', 687),
    (55, 26, '22D', 'E', 729),
    (56, 5, '02E', 'B', 1233),
    (57, 4, '18F', 'E', 620),
    (58, 3, '01A', 'B', 1006),
    (59, 8, '20B', 'E', 715),
    (60, 2, '03A', 'B', 1084),
    (61, 3, '19D', 'E', 727),
    (62, 19, '11E', 'E', 730),
    (63, 23, '17F', 'E', 606),
    (64, 28, '02B', 'B', 1064),
    (65, 10, '12B', 'E', 520),
    (66, 26, '11C', 'E', 588),
    (67, 7, '03D', 'B', 1056),
    (68, 19, '23E', 'E', 728),
    (69, 12, '02A', 'B', 1218),
    (70, 8, '07B', 'E', 731),
    (71, 1, '16C', 'E', 690),
    (72, 22, '22D', 'E', 530),
    (73, 11, '02E', 'B', 1083),
    (74, 11, '18F', 'E', 738),
    (75, 27, '01A', 'B', 1217),
    (76, 22, '20B', 'E', 626),
    (77, 9, '03A', 'B', 1116),
    (78, 18, '19D', 'E', 727),
    (79, 2, '11E', 'E', 680),
    (80, 22, '17F', 'E', 690),
    (81, 14, '02B', 'B', 1187),
    (82, 13, '12B', 'E', 621),
    (83, 10, '11C', 'E', 724),
    (84, 12, '03D', 'B', 1138),
    (85, 5, '23E', 'E', 726),
    (86, 4, '02A', 'B', 1213),
    (87, 7, '07B', 'E', 580),
    (88, 13, '16C', 'E', 538),
    (89, 27, '22D', 'E', 519),
    (90, 22, '02E', 'B', 1164),
    (91, 16, '18F', 'E', 680),
    (92, 2, '01A', 'B', 1184),
    (93, 5, '20B', 'E', 693),
    (94, 12, '03A', 'B', 1143),
    (95, 13, '19D', 'E', 744),
    (96, 1, '11E', 'E', 532),
    (97, 10, '17F', 'E', 602),
    (98, 17, '02B', 'B', 1113),
    (99, 20, '12B', 'E', 731),
    (100, 4, '11C', 'E', 576),
    (101, 10, '03D', 'B', 1199),
    (102, 27, '23E', 'E', 581),
    (103, 7, '02A', 'B', 1003),
    (104, 18, '07B', 'E', 569),
    (105, 12, '16C', 'E', 589),
    (106, 28, '22D', 'E', 530),
    (107, 17, '02E', 'B', 1111),
    (108, 17, '18F', 'E', 542),
    (109, 16, '01A', 'B', 1242),
    (110, 20, '20B', 'E', 750),
    (111, 20, '03A', 'B', 1059),
    (112, 28, '19D', 'E', 673),
    (113, 1, '11E', 'E', 511),
    (114, 11, '17F', 'E', 644),
    (115, 6, '02B', 'B', 1095),
    (116, 21, '12B', 'E', 529),
    (117, 13, '11C', 'E', 740),
    (118, 9, '03D', 'B', 1220),
    (119, 20, '23E', 'E', 516),
    (120, 18, '02A', 'B', 1099),
    (121, 6, '07B', 'E', 643),
    (122, 9, '16C', 'E', 582),
    (123, 26, '22D', 'E', 552),
    (124, 20, '02E', 'B', 1136),
    (125, 23, '18F', 'E', 657),
    (126, 19, '01A', 'B', 1025),
    (127, 23, '20B', 'E', 568),
    (128, 12, '03A', 'B', 1074),
    (129, 12, '19D', 'E', 571),
    (130, 23, '11E', 'E', 698),
    (131, 28, '17F', 'E', 580),
    (132, 5, '02B', 'B', 1019),
    (133, 3, '12B', 'E', 712),
    (134, 6, '11C', 'E', 527),
    (135, 12, '03D', 'B', 1164),
    (136, 23, '23E', 'E', 569),
    (137, 2, '02A', 'B', 1081),
    (138, 20, '07B', 'E', 618),
    (139, 4, '16C', 'E', 729),
    (140, 26, '22D', 'E', 701),
    (141, 8, '02E', 'B', 1160),
    (142, 23, '18F', 'E', 704),
    (143, 16, '01A', 'B', 1116),
    (144, 15, '20B', 'E', 721),
    (145, 1, '03A', 'B', 1045),
    (146, 20, '19D', 'E', 722),
    (147, 15, '11E', 'E', 631),
    (148, 28, '17F', 'E', 550),
    (149, 5, '02B', 'B', 1030),
    (150, 13, '12B', 'E', 597),
    (151, 20, '11C', 'E', 552),
    (152, 27, '03D', 'B', 1162),
    (153, 24, '23E', 'E', 671),
    (154, 13, '02A', 'B', 1067),
    (155, 12, '07B', 'E', 588),
    (156, 10, '16C', 'E', 676),
    (157, 9, '22D', 'E', 575),
    (158, 7, '02E', 'B', 1122),
    (159, 22, '18F', 'E', 524),
    (160, 11, '01A', 'B', 1135),
    (161, 2, '20B', 'E', 741),
    (162, 12, '03A', 'B', 1167),
    (163, 5, '19D', 'E', 563),
    (164, 15, '11E', 'E', 630),
    (165, 6, '17F', 'E', 672),
    (166, 23, '02B', 'B', 1051),
    (167, 21, '12B', 'E', 657),
    (168, 26, '11C', 'E', 538),
    (169, 20, '03D', 'B', 1120),
    (170, 4, '23E', 'E', 560),
    (171, 24, '02A', 'B', 1049),
    (172, 14, '07B', 'E', 567),
    (173, 8, '16C', 'E', 741),
    (174, 7, '22D', 'E', 560),
    (175, 8, '02E', 'B', 1196),
    (176, 4, '18F', 'E', 706),
    (177, 23, '01A', 'B', 1180),
    (178, 13, '20B', 'E', 670),
    (179, 13, '03A', 'B', 1012),
    (180, 6, '19D', 'E', 683),
    (181, 24, '11E', 'E', 515),
    (182, 12, '17F', 'E', 745),
    (183, 24, '02B', 'B', 1151),
    (184, 28, '12B', 'E', 549),
    (185, 8, '11C', 'E', 702),
    (186, 6, '03D', 'B', 1020),
    (187, 19, '23E', 'E', 727),
    (188, 7, '02A', 'B', 1230),
    (189, 25, '07B', 'E', 711),
    (190, 17, '16C', 'E', 587),
    (191, 16, '22D', 'E', 507),
    (192, 11, '02E', 'B', 1065),
    (193, 15, '18F', 'E', 536),
    (194, 27, '01A', 'B', 1201),
    (195, 11, '20B', 'E', 655),
    (196, 11, '03A', 'B', 1023),
    (197, 17, '19D', 'E', 612),
    (198, 18, '11E', 'E', 553),
    (199, 5, '17F', 'E', 633),
    (200, 13, '02B', 'B', 1159),
    (201, 23, '12B', 'E', 608),
    (202, 8, '11C', 'E', 655),
    (203, 22, '03D', 'B', 1209),
    (204, 10, '23E', 'E', 501),
    (205, 22, '02A', 'B', 1035);
