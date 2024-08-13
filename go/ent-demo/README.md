```sh
go run -mod=mod entgo.io/ent/cmd/ent init MODEL-NAME
go generate ./ent
go run -mod=mod ./cmd/migration/main.go MIGRATION-FILE-NAME
docker compose exec backend atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url mysql://user:password@database:3306/entdemo
```
