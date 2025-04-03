PHONY: create

create:
	goose -dir ./migrations create -s ${NAME_MIGRATION} sql
