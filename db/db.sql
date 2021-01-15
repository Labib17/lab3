drop shema if exists paysystem;
 create schema if ot exists paysystem;
 use paysystem;
 
 create table account(
	id int [pk, not null] ,
    balance int [not null, defaul 0] ,
    last_operation datetime [null],
 );
 
 insert into account(balance, last_operation)
 values
	(2000, '2015-11-05 14:29:36.11'),
	(1000, '2019-09-29 15:14:18.32'),
    (5300, '2019-11-14 10:22:10.32'),
    (3450, '2020-12-12 18:11:24.45');

drop user if exists helen;
create user if not exists helen identified by 'x7x';
grant all privileges on  paysystem.* to helen;

delimiter $$
drop procedure if exists  transfer;
create procedure  transfer (Id_of_sender int, Id_of_recepient int, amount int,  last_operation datetime)

begin 

update account set balance  = Id_of_sender;
update accounts set balance  = Id_of_recepient;
update accounts set last_operation =  last_operation where id = Id_of_sender;
end; $$
