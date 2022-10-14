# My first project
 frontend is hard:(



Создание необходимых бд
```
create table order_details(
id_of_person text,
id_of_product int
);
select * from order_details;
truncate order_details;


create table Shop(
id serial,
price numeric(15),
 amount integer,
  name text,
   description text,
    category text,
    picture text
)
select * from Shop;
truncate Shop;

create table admin(
    login varchar(20),
    password varchar(70)
)
select * from admin
truncate admin;

```




триггер-логгер для таблицы Shop
```

CREATE TABLE logger_for_Shop
(
  txt text,
  added timestamp without time zone,
  MyNew text,
  MyOld text
)

CREATE OR REPLACE FUNCTION add_to_logger() RETURNS TRIGGER AS $$
DECLARE
    mstr varchar(30);
    astr varchar(100);
    newid varchar(20);
    newprice varchar(20);
    newamount varchar(20);
    newname varchar(100);
    newdescription varchar(500);
    newcategory varchar(100);
    newpicture varchar(110);
    oldid varchar(20);
    oldprice varchar(20);
    oldamount varchar(20);
    oldname varchar(100);
    olddescription varchar(500);
    oldcategory varchar(100);
    oldpicture varchar(110);
    retstr varchar(300);
    evrythingnew varchar(1400);
    evrythingold varchar(1400);
BEGIN
    IF    TG_OP = 'INSERT' THEN
        astr = NEW.id;
        mstr := 'Add new product ';
        newid:=NEW.id;
        newprice:=NEW.price;
        newamount:=NEW.amount;
        newname:=NEW.name;
        newdescription:=NEW.description;
        newcategory:=NEW.category;
        newpicture:=NEW.picture;
        retstr := mstr || astr;
        evrythingnew:=newid || ' ' || newprice || ' ' ||newamount || ' ' ||newname || ' ' ||newdescription || ' ' ||newcategory || ' ' ||newpicture;
        INSERT INTO logger_for_Shop(txt,added,MyNew) values (retstr,NOW(),evrythingnew);
        RETURN NEW;
    ELSIF TG_OP = 'UPDATE' THEN
        astr = NEW.id;
        mstr := 'Update product ';
        newid:=NEW.id;
        newprice:=NEW.price;
        newamount:=NEW.amount;
        newname:=NEW.name;
        newdescription:=NEW.description;
        newcategory:=NEW.category;
        newpicture:=NEW.picture;
        evrythingnew:=newid || ' ' || newprice || ' ' ||newamount || ' ' ||newname || ' ' ||newdescription || ' ' ||newcategory || ' ' ||newpicture;
        oldid:=OLD.id;
        oldprice:=OLD.price;
        oldamount:=OLD.amount;
        oldname:=OLD.name;
        olddescription:=OLD.description;
        oldcategory:=OLD.category;
        oldpicture:=OLD.picture;
        evrythingold:=oldid || ' ' || oldprice || ' ' || oldamount || ' ' ||oldname || ' ' ||olddescription || ' ' ||oldcategory || ' ' ||oldpicture;
        retstr := mstr || astr;
        INSERT INTO logger_for_Shop(txt,added,MyNew,MyOld) values (retstr,NOW(),evrythingnew,evrythingold);
        RETURN NEW;
    ELSIF TG_OP = 'DELETE' THEN
        astr = OLD.id;
        mstr := 'Remove product ';
        oldid:=OLD.id;
        oldprice:=OLD.price;
        oldamount:=OLD.amount;
        oldname:=OLD.name;
        olddescription:=OLD.description;
        oldcategory:=OLD.category;
        oldpicture:=OLD.picture;
        evrythingold:=oldid || ' ' || oldprice || ' ' || oldamount || ' ' ||oldname || ' ' ||olddescription || ' ' ||oldcategory || ' ' ||oldpicture;
        retstr := mstr || astr;
        INSERT INTO logger_for_Shop(txt,added,MyOld) values (retstr,NOW(),evrythingold);
        RETURN OLD;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER logging
AFTER INSERT OR UPDATE OR DELETE ON Shop FOR EACH ROW EXECUTE PROCEDURE add_to_logger ();

```


триггер-логгер для таблицы Order

```

create table logs_for_order(
    operation text,
    added timestamp without time zone
)





CREATE OR REPLACE FUNCTION add_to_logger_order() RETURNS TRIGGER AS $$
DECLARE
    mstr varchar(30);
    astr varchar(300);
    retstr varchar(254);
BEGIN
    IF    TG_OP = 'INSERT' THEN
        astr = NEW.Id_of_person;
        mstr := 'Add new order ';
        retstr := mstr || astr;
        INSERT INTO logs_for_order(operation,added) values (retstr,NOW());
        RETURN NEW;
    ELSIF TG_OP = 'DELETE' THEN
        astr = OLD.Id_of_person;
        mstr := 'Remove order ';
        retstr := mstr || astr;
        INSERT INTO logs_for_order(operation,added) values (retstr,NOW());
        RETURN OLD;
    END IF;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER logging_for_order
AFTER INSERT OR UPDATE OR DELETE ON order_details FOR EACH ROW EXECUTE PROCEDURE add_to_logger_order();

```
перед запуском в main.go необходимо добавить следующий код

```
adm := database.Admin{"2mfree@mail.ru", "localhost"}
adm.InsertData(database.Dv, database.DBCon)
```

это ваш логин/пароль в панель администрации на сайте
через панель администрации проще всего добавить товары



или это можно сделать из бд, но чтобы выводило картинку, нужно иметь cat.jpeg в папке images

```
insert into Shop(price,amount,name,description,category,picture) values(1000,10,'Название 1','Описание 1','Категория 1','cat.jpeg')
```



В программе авторизация проходит через JWT токен, который хранится в куках. в JWT токене хранится IsAdmin который отвечает за разграничение доступа.
Так же в таблице Order, JWT Токен является ключем, по которому добавляются/удаляются данные в корзине.
Большая часть операций над бд(удаление/добавление/изменение) работают в транзакциях(уровень транзакции не изменял, так что стандартный Read Commited).
Использую стандартный драйвер для бд lib/pq, валидация данных выполняется через библиотеку go-playground/validator/v10. Роутер- gorilla/mux. Сессии-gorilla/sessions.

P.S. картинки только jpeg/png






