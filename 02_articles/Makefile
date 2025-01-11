create_all:
	goose -dir migrations create articles sql 
	goose -dir migrations create articles_add_cat_id sql 
	goose -dir migrations create categories sql 
	goose -dir migrations create users sql 
	goose -dir migrations create roles sql 
	goose -dir migrations create users_add_role_id sql 
	goose -dir migrations create users_add_age sql 
	goose -dir migrations create users_add_password sql 

create_articles:
	goose -dir migrations create articles sql 

create_articles_cat_id:
	goose -dir migrations create articles_add_cat_id sql 

create_categories:
	goose -dir migrations create categories sql 

create_users:
	goose -dir migrations create users sql 

create_user_roles:
	goose -dir migrations create roles sql

create_users_role_id:
	goose -dir migrations create users_add_role_id sql 

create_users_age:
	goose -dir migrations create users_add_age sql 

create_user_password:
	goose -dir migrations create users_add_password sql 
	
migrate:
	goose -dir migrations postgres "postgresql://pguser:pgpassword@127.0.0.1:5432/postgres?sslmode=disable" up