insert into accounts values (2, 1, "uuid-gen-12345", "reseller", "My Name", "firstuser",md5("password"),"firstuser@email.org",now());
insert into hotspots values (1,2,"HSTest","HS for development",now());
insert into accounts values (3, 1, "uuid-gen-99999", "customer", "Customer Account", "cust1",md5("cust1"),"cust1@example.com",now());
insert into accounts_hotspots values (1, 3, 1);
insert into units values (1,1,'00-00-00-00-00-00','testunitname','testunitdesc','1234-uuid-aaaa','secret',NOW());
insert into users values (1, 1, 'First User', 'firstuser', 'firstpassword', 'first.user@nethserver.org', 'email', 0, 0, 0, NOW(), NOW() + INTERVAL 3600 DAY, NOW());
insert into hotspot_preferences values (1,1,"Acct-Session-Time","60");
insert into hotspot_preferences values (2,1,"Session-Timeout","61");
insert into hotspot_preferences values (3,1,"CoovaChilli-Bandwidth-Max-Up","100000");
insert into hotspot_preferences values (4,1,"CoovaChilli-Bandwidth-Max-Down","100001");
insert into subscriptions values (1, 2, 3, now(), date_add(now(), interval 365 day),  default);
