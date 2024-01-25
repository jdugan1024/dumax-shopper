create table meal (
	id int not null primary key,
	name text not null,
	meal_type text not null CHECK(meal_type IN ('breakfast', 'lunch', 'dinner'))
);

create table ingredient (
	id int not null primary key,
	name text not null
);

create table meal_ingredient (
        meal_id integer not null references meal(id),
        ingredient_id integer not null references ingredient(id),
        quantity text,
        primary key (meal_id, ingredient_id)
);

insert into meal (id, name, meal_type) values (1, 'Tacos', 'dinner');

insert into ingredient (id, name) values (1, 'Crunchy Taco Shells');
insert into ingredient (id, name) values (2, 'Ground Chicken');
insert into ingredient (id, name) values (3, 'Shredded Cheese');
insert into ingredient (id, name) values (4, 'Lettuce');
insert into ingredient (id, name) values (5, 'Sour Cream');
insert into ingredient (id, name) values (6, 'Salsa');
insert into ingredient (id, name) values (7, 'Black Olives');
insert into ingredient (id, name) values (8, 'Taco Seasoning');
insert into ingredient (id, name) values (9, 'Refried Beans');

insert into meal_ingredient (meal_id, ingredient_id, quantity) values (1, 1, '6');
insert into meal_ingredient (meal_id, ingredient_id, quantity) values (1, 2, '1 lb');
insert into meal_ingredient (meal_id, ingredient_id, quantity) values (1, 3, '1 cup');
insert into meal_ingredient (meal_id, ingredient_id, quantity) values (1, 4, '1 head');
insert into meal_ingredient (meal_id, ingredient_id, quantity) values (1, 5, '1 package');
insert into meal_ingredient (meal_id, ingredient_id, quantity) values (1, 6, '1 jar');
insert into meal_ingredient (meal_id, ingredient_id, quantity) values (1, 7, '1 can');
insert into meal_ingredient (meal_id, ingredient_id, quantity) values (1, 8, '1 packet');
insert into meal_ingredient (meal_id, ingredient_id, quantity) values (1, 9, '1 can');