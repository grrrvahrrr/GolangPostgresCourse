with ins AS(
    Insert into bitme.urlbase values('udfhsdffe','http://google.com/')

    RETURNING short_url
)

insert into bitme.urlusedata(short_url, ip, last_used, ip_num_of_uses)
SELECT short_url, '1.1.1.1', '01.01.2001 00:00:00', '1'
FROM ins;

Insert into bitme.adminurl (admin_url, short_url)
SELECT 'doturnspqo', short_url
FROM bitme.urlbase;

insert into bitme.urldata(short_url, last_used, total_num_of_uses)
SELECT short_url, last_used, SUM(ip_num_of_uses) as total_num_of_uses
From bitme.urlusedata
Group by short_url;