--query 1: get short url from admin url

select short_url from bitme.adminurl where admin_url = 'doturnspqo';

--query 2: get total num of uses from short url

select last_used, total_num_of_uses from bitme.urldata where short_url = 'udfhsdffe';

--query 3: get full url from short url

select full_url from bitme.urlbase where short_url = 'udfhsdffe';

--query 4: get ip_num_of_uses from user ip

select ip_num_of_uses from bitme.urlusedata where ip = '1.1.1.1';

--query 5 select all instances of short url use by all ips

select ip, ip_num_of_uses from bitme.urlusedata where short_url = 'udfhsdffe' order by ip_num_of_uses asc;