alter table t_cl_request_template
    add column author_url varchar(100);

update t_cl_request_template
set author_url = 'https://twitter.com/samdemaree'
where id = 1;

update t_cl_request_template
set author_url = 'https://snipe24.t.me/'
where id = 2;

update t_cl_request_template
set author_url = 'https://twitter.com/0xpolarzero'
where id = 3;

update t_cl_request_template
set author_url = 'https://twitter.com/MaxwellMelcher'
where id = 4;

update t_cl_request_template
set author_url = 'https://snipe24.t.me/'
where id = 5;

update t_cl_request_template
set author_url = 'https://twitter.com/PatrickAlphaC'
where id = 6;

update t_cl_request_template
set author_url = 'https://twitter.com/PatrickAlphaC'
where id = 7;

update t_cl_request_template
set author_url = 'https://twitter.com/ChainLinkGod'
where id = 8;

update t_cl_request_template
set author_url = 'https://twitter.com/ChainLinkGod'
where id = 9;

update t_cl_request_template
set author_url = 'https://twitter.com/mykcryptodev'
where id = 10;

update t_cl_request_template
set author_url = 'https://snipe24.t.me/'
where id = 11;

update t_cl_request_template
set author_url = 'https://moonthoon.io/'
where id = 12;

update t_cl_request_template
set author_url = 'https://twitter.com/0xpolarzero'
where id = 13;

update t_cl_request_template
set author_url = 'https://github.com/KuphJr'
where id = 14;