CURRENT_DIR=$(shell pwd)
DBURL := postgres://postgres:0412@localhost:5432/nt?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -path databases/migrations -database '${DBURL}' -verbose up

mig-down:
	migrate -path databases/migrations -database '${DBURL}' -verbose down

mig-force:
	migrate -path databases/migrations -database '${DBURL}' -verbose force 1


mig-create-table:
	migrate create -ext sql -dir databases/migrations -seq create_reservation_table

