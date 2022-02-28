-- The idea for the project is pretty simple. It is a URL shortener like bitly. A user provedies a URL to shorten,
-- the app returns a short URL. When the short URL is used the server redirects the user to full URL and writes data 
-- to a database. Right now it only writes a number of redirects, but there can be more data that can be gathered from redirects.

--tables

--table:urlbase
--table where we store all the URLs that we generated and short URLs that relate to them. Values Updated.
create database bitmedb;

create user bituser password 'bit';

grant all privileges on database bitmedb to bituser;

\c bitmedb bituser;

CREATE schema bitme;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE bitme.urlbase (
    short_url varchar(10) PRIMARY KEY,
    full_url text NOT NULL,   
    url_id uuid NOT NULL   
);

--table:urldata
--table that stores primary data for short URLs - number of time used and how active it is. Values updated.
CREATE TABLE bitme.urldata(
    short_url varchar(10) PRIMARY KEY,    
    last_used timestamp NOT NULL,
    total_num_of_uses int NOT NULL,    
    FOREIGN KEY (short_url) REFERENCES bitme.urlbase(short_url),
    constraint total_num_of_uses_check check (total_num_of_uses >= 0)
);

--table:adminurl
--table that stores urls for getting info by admins
CREATE TABLE bitme.adminurl (
    admin_url varchar(10) PRIMARY KEY,
    short_url varchar(10) NOT NULL,      
    admin_id uuid NOT NULL,
    FOREIGN KEY (short_url) REFERENCES bitme.urlbase(short_url)
);

--table urlusedata
--table that stores uses of short url from certain ip adresses
CREATE TABLE bitme.urlusedata(
    short_url varchar(10) PRIMARY KEY,
    ip INET NOT NULL,
    last_used timestamp NOT NULL,
    ip_num_of_uses int NOT NULL,
    FOREIGN KEY (short_url) REFERENCES bitme.urlbase(short_url),
    constraint ip_num_of_uses_check check (ip_num_of_uses >= 0)
);

