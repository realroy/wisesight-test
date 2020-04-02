FROM python:3.7.7-alpine3.11 AS clean

COPY clean.py .
COPY requirements.txt .

RUN pip install -r requirements.txt

RUN wget "https://s3-ap-southeast-1.amazonaws.com/wisesight-public/dev-test/rawdata.csv"

RUN python clean.py


FROM mongo:4.2.3-bionic

COPY --from=clean ./data.json .

CMD mongoimport --host db --db $DB_NAME --collection $COLLECTION_NAME --file data.json --jsonArray
