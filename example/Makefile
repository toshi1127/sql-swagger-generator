# note: call scripts from /scripts

.PHONY: openapigen
openapigen:
	docker-compose -f docker-compose.apimodel.yaml run cleanup
	docker-compose -f docker-compose.apimodel.yaml run api-model-go
	docker-compose -f docker-compose.apimodel.yaml run api-model-ts
