-- The idea for the project is pretty simple. It is a URL shortener like bitly. A user provedies a URL to shorten,
-- the app returns a short URL. When the short URL is used the server redirects the user to full URL and writes data 
-- to a database. Right now it only writes a number of redirects, but there can be more data that can be gathered from redirects.

--tables

--table:urlbase
--table where we store all the URLs that we generated and short URLs that relate to them. Values Updated.
create database bitme;

create user bituser password 'bit';

grant all privileges on database bitme to bituser;

CREATE schema bitme

CREATE TABLE bitme.urlbase (
    id uuid  PRIMARY KEY NOT NULL,
    fullurl text NOT NULL,
    shorturl varchar(10) NOT NULL
);

--table:urldata
--table that stores primary data for short URLs - number of time used and how active it is. Values updated.
CREATE TABLE bitme.urldata(
    url_id uuid NOT NULL,
    last_used timestamp PRIMARY KEY,
    number_of_uses int,    
    FOREIGN KEY (url_id) REFERENCES bitme.urlbase(id)
);

--table urlusedata
--table that stores all data every time the server is used. Create new line on every entry. The bigget table
CREATE TABLE bitme.urlusedata(
    url_id uuid  PRIMARY KEY NOT NULL,
    time_used timestamp,
    ip INET,
    FOREIGN KEY (url_id) REFERENCES bitme.urlbase(id),
    FOREIGN KEY (time_used) REFERENCES bitme.urldata(last_used)
);

