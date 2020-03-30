FROM mongo:4.2.3-bionic

COPY ./data.json .

CMD mongoimport --host db --db $DB_NAME --collection $COLLECTION_NAME --file data.json --jsonArray
